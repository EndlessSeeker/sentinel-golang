/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package protocol

import (
	"fmt"
	"github.com/alibaba/sentinel-golang/pkg/datasource/xds/client"
	"github.com/alibaba/sentinel-golang/pkg/datasource/xds/resources/rbac"
	"sync"

	listenerv3 "github.com/alibaba/sentinel-golang/pkg/datasource/xds/go-control-plane/envoy/config/listener/v3"
	jwtauthnv3 "github.com/alibaba/sentinel-golang/pkg/datasource/xds/go-control-plane/envoy/extensions/filters/http/jwt_authn/v3"
	envoyrbacv3 "github.com/alibaba/sentinel-golang/pkg/datasource/xds/go-control-plane/envoy/extensions/filters/http/rbac/v3"
	hcmv3 "github.com/alibaba/sentinel-golang/pkg/datasource/xds/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"
	sockets_tls_v3 "github.com/alibaba/sentinel-golang/pkg/datasource/xds/go-control-plane/envoy/extensions/transport_sockets/tls/v3"
	discoveryv3 "github.com/alibaba/sentinel-golang/pkg/datasource/xds/go-control-plane/envoy/service/discovery/v3"
	"github.com/alibaba/sentinel-golang/pkg/datasource/xds/resources"
	"github.com/alibaba/sentinel-golang/pkg/datasource/xds/utils"
	"github.com/dubbogo/gost/log/logger"
	"github.com/golang/protobuf/ptypes"
)

type LdsProtocol struct {
	xdsClientChannel *client.XdsClient
	resourcesMap     sync.Map
	stopChan         chan struct{}
	updateChan       chan resources.XdsUpdateEvent
}

func NewLdsProtocol(stopChan chan struct{}, updateChan chan resources.XdsUpdateEvent, xdsClientChannel *client.XdsClient) (*LdsProtocol, error) {
	ldsProtocol := &LdsProtocol{
		xdsClientChannel: xdsClientChannel,
		stopChan:         stopChan,
		updateChan:       updateChan,
	}
	return ldsProtocol, nil
}

func (lds *LdsProtocol) GetTypeUrl() string {
	return client.EnvoyListener
}

func (lds *LdsProtocol) SubscribeResource(resourceNames []string) error {
	return lds.xdsClientChannel.SendWithTypeUrlAndResourceNames(lds.GetTypeUrl(), resourceNames)
}

func (lds *LdsProtocol) ProcessProtocol(resp *discoveryv3.DiscoveryResponse, xdsClientChannel *client.XdsClient) error {
	if resp.GetTypeUrl() != lds.GetTypeUrl() {
		return nil
	}
	//logger.Infof("DiscoveryResponse: %s", utils.ConvertJsonString(resp))
	xdsListeners := make([]resources.XdsListener, 0)
	rdsResourceNames := make([]string, 0)
	for _, resource := range resp.GetResources() {
		ldsResource := &listenerv3.Listener{}
		if err := ptypes.UnmarshalAny(resource, ldsResource); err != nil {
			logger.Errorf("[lds Protocol] fail to extract listener: %v", err)
			continue
		}
		logger.Debugf("[lds Protocol] Listener: %s", utils.ConvertJsonString(ldsResource))
		fmt.Printf("[LdsProtocol.ProcessProtocol] listener info name: %v, listener info: %v\n", ldsResource.Name, ldsResource.String())
		xdsListener, _ := lds.parseListener(ldsResource)
		if xdsListener.HasRds {
			rdsResourceNames = append(rdsResourceNames, xdsListener.RdsResourceNames...)
		}
		xdsListeners = append(xdsListeners, xdsListener)
	}

	// notify update
	updateEvent := resources.XdsUpdateEvent{
		Type:   resources.XdsEventUpdateLDS,
		Object: xdsListeners,
	}
	lds.updateChan <- updateEvent

	info := &client.ResponseInfo{
		VersionInfo:   resp.VersionInfo,
		ResponseNonce: resp.Nonce,
		ResourceNames: []string{}, // always
	}
	lds.xdsClientChannel.ApiStore.Store(client.EnvoyListener, info)
	lds.xdsClientChannel.AckResponse(resp)

	if len(rdsResourceNames) > 0 { // Load RDS
		lds.xdsClientChannel.ApiStore.SetResourceNames(client.EnvoyRoute, rdsResourceNames)
		req := lds.xdsClientChannel.CreateRdsRequest()
		return lds.xdsClientChannel.Send(req)
	}
	return nil
}

func (lds *LdsProtocol) parseListener(listener *listenerv3.Listener) (resources.XdsListener, error) {
	envoyListener := resources.XdsListener{
		//FilterChains:     make([]resources.XdsFilterChain, 0),
		RdsResourceNames: make([]string, 0),
	}
	envoyListener.Name = listener.Name
	// inbound15006 tls mode and downstream transport socket
	inboundTlsMode := resources.XdsTLSMode{}
	inboundDownstreamTransportSocket := resources.XdsDownstreamTransportSocket{}
	isVirtualInbound := false

	// get 15006 and traffic direction
	if listener.GetAddress() != nil {
		if listener.GetAddress().GetSocketAddress() != nil {
			port := listener.GetAddress().GetSocketAddress().GetPortValue()
			if port == 15006 && listener.Name == "virtualInbound" {
				isVirtualInbound = true
			}
		}
	}

	envoyListener.IsVirtualInbound = isVirtualInbound
	envoyListener.TrafficDirection = listener.GetTrafficDirection().String()

	for _, filterChain := range listener.GetFilterChains() {

		envoyFilterChain := resources.XdsFilterChain{}
		envoyFilterChain.Name = filterChain.Name
		// find inboud filterChain
		foundInboundFilterChain := false
		// find catch all filterchain
		foundInboundFilterChainCatchAll := false
		foundInboundDownstreamTransportSocket := false

		if filterChain.Name == "virtualInbound" {
			foundInboundFilterChain = true
		}
		if filterChain.Name == "virtualInbound-catchall-http" {
			foundInboundFilterChainCatchAll = true
		}
		// get transport
		envoyFilterChainMatch := resources.XdsFilterChainMatch{}
		envoyDownstreamTransportSocket := resources.XdsDownstreamTransportSocket{}

		// get listener inboundTlsMode
		if filterChain.GetFilterChainMatch() != nil {
			filterChainMatch := filterChain.GetFilterChainMatch()
			if len(filterChainMatch.GetTransportProtocol()) > 0 {
				envoyFilterChainMatch.TransportProtocol = filterChainMatch.GetTransportProtocol()
				if filterChainMatch.GetTransportProtocol() == "tls" && foundInboundFilterChain && isVirtualInbound {
					inboundTlsMode.IsTls = true
				}
				if filterChainMatch.GetTransportProtocol() == "raw_buffer" && foundInboundFilterChain && isVirtualInbound {
					inboundTlsMode.IsRawBuffer = true
				}
			}
		}
		// parse downstream transport socket
		if filterChain.GetTransportSocket() != nil {
			var tlsContext sockets_tls_v3.DownstreamTlsContext
			transportSocket := filterChain.GetTransportSocket()
			typeUrl := transportSocket.GetTypedConfig().GetTypeUrl()
			if typeUrl == "type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.DownstreamTlsContext" {
				err := transportSocket.GetTypedConfig().UnmarshalTo(&tlsContext)
				if err != nil {
					logger.Errorf("can not parse to downstream tls context")
					continue
				}
				foundInboundDownstreamTransportSocket = true
				envoyDownstreamTransportSocket.RequireClientCertificate = tlsContext.RequireClientCertificate.GetValue()
				matchers := tlsContext.CommonTlsContext.GetCombinedValidationContext().DefaultValidationContext.GetMatchSubjectAltNames()
				for _, matcher := range matchers {
					if len(matcher.GetExact()) > 0 {
						envoyDownstreamTransportSocket.SubjectAltNamesMatch = "exact"
						envoyDownstreamTransportSocket.SubjectAltNamesValue = matcher.GetExact()
					}
					if len(matcher.GetPrefix()) > 0 {
						envoyDownstreamTransportSocket.SubjectAltNamesMatch = "prefix"
						envoyDownstreamTransportSocket.SubjectAltNamesValue = matcher.GetPrefix()
					}
					if len(matcher.GetContains()) > 0 {
						envoyDownstreamTransportSocket.SubjectAltNamesMatch = "contains"
						envoyDownstreamTransportSocket.SubjectAltNamesValue = matcher.GetContains()
					}
				}
			}
		}

		// found inbound TransportSocket
		if isVirtualInbound && foundInboundFilterChain && foundInboundDownstreamTransportSocket {
			inboundDownstreamTransportSocket = envoyDownstreamTransportSocket
		}

		envoyFilterChain.FilterChainMatch = envoyFilterChainMatch
		envoyFilterChain.TransportSocket = envoyDownstreamTransportSocket
		envoyFilterChain.Filters = make([]resources.XdsFilter, 0)
		// Get Rds and filters
		for _, filter := range filterChain.GetFilters() {
			envoyFilter := resources.XdsFilter{}
			envoyFilter.Name = filter.Name
			if filter.GetTypedConfig().GetTypeUrl() == "type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager" {
				// check is rds
				httpConnectionManager := &hcmv3.HttpConnectionManager{}
				if err := filter.GetTypedConfig().UnmarshalTo(httpConnectionManager); err != nil {
					logger.Errorf("can not parse to Http Connection Manager")
					continue
				}
				envoyFilter.IncludeHttpConnectionManager = true
				if httpConnectionManager.GetRds() != nil {
					envoyFilter.HasRds = true
					envoyFilter.RdsResourceName = httpConnectionManager.GetRds().RouteConfigName
					envoyListener.HasRds = true
					envoyListener.RdsResourceNames = append(envoyListener.RdsResourceNames, httpConnectionManager.GetRds().RouteConfigName)
				}

				for _, httpFilter := range httpConnectionManager.HttpFilters {
					if httpFilter.GetTypedConfig().GetTypeUrl() == "type.googleapis.com/envoy.extensions.filters.http.jwt_authn.v3.JwtAuthentication" && foundInboundFilterChainCatchAll {
						// parse jwt authn here
						jwtAuthentication := &jwtauthnv3.JwtAuthentication{}
						if err := httpFilter.GetTypedConfig().UnmarshalTo(jwtAuthentication); err != nil {
							logger.Errorf("can not parse to Http JwtAuthnEnovyFilter")
							continue
						}
						jwtAuthnFilter, err := lds.parseJwtAuthnFilter(httpFilter.Name, jwtAuthentication)
						if err != nil {
							logger.Errorf("can not convert to JwtAuthnEnovyFilter, error:%v", err)
							continue
						}
						envoyListener.JwtAuthnFilter = jwtAuthnFilter
					}

					//if httpFilter.GetTypedConfig().GetTypeUrl() == "type.googleapis.com/istio.envoy.config.filter.http.authn.v2alpha1.FilterConfig" && foundInboundFilterChainCatchAll {
					//	// parse istio authn here
					//}

					if httpFilter.GetTypedConfig().GetTypeUrl() == "type.googleapis.com/envoy.extensions.filters.http.rbac.v3.RBAC" && foundInboundFilterChainCatchAll {
						// parse rbac filter here
						rbacEnvoyFilter := &envoyrbacv3.RBAC{}
						if err := httpFilter.GetTypedConfig().UnmarshalTo(rbacEnvoyFilter); err != nil {
							logger.Errorf("can not parse to Http rbacFilter")
							continue
						}

						rbacFilter, err := lds.parseRBACFilter(httpFilter.Name, rbacEnvoyFilter)
						if err != nil {
							logger.Errorf("can not convert to RBACEnovyFilter, error: %v", err)
							continue
						}
						envoyListener.RBACFilter = rbacFilter
					}
				}
			}
			envoyFilterChain.Filters = append(envoyFilterChain.Filters, envoyFilter)
		}

		//envoyListener.FilterChains = append(envoyListener.FilterChains, envoyFilterChain)

		// Set envoy listener
		envoyListener.InboundTLSMode = inboundTlsMode
		envoyListener.InboundDownstreamTransportSocket = inboundDownstreamTransportSocket
	}

	return envoyListener, nil
}

func (lds *LdsProtocol) parseRBACFilter(name string, envoyRBAC *envoyrbacv3.RBAC) (resources.RBACEnvoyFilter, error) {
	logger.Infof("[lds Protocol] RBAC Filter: %s", utils.ConvertJsonString(envoyRBAC))
	rbacEnvoyFilter := resources.RBACEnvoyFilter{
		Name: name,
	}
	rbac, err := rbac.NewRBAC(envoyRBAC)
	if err != nil {
		return rbacEnvoyFilter, err
	}
	rbacEnvoyFilter.RBAC = rbac
	return rbacEnvoyFilter, nil
}

func (lds *LdsProtocol) parseJwtAuthnFilter(name string, envoyAuthentication *jwtauthnv3.JwtAuthentication) (resources.JwtAuthnEnovyFilter, error) {
	jwtAuthnFilter := resources.JwtAuthnEnovyFilter{
		Name: name,
	}
	jwtAuthentication := &resources.JwtAuthentication{
		Providers: make(map[string]*resources.JwtProvider),
		Rules:     make([]*resources.JwtRequirementRule, 0),
	}

	// parse providers
	for providerName, envoyProvider := range envoyAuthentication.Providers {
		jwtProvider := &resources.JwtProvider{
			ProviderName:         providerName,
			Issuer:               envoyProvider.Issuer,
			Audiences:            envoyProvider.Audiences,
			Forward:              envoyProvider.Forward,
			ForwardPayloadHeader: envoyProvider.ForwardPayloadHeader,
		}
		// parse from_headers
		if len(envoyProvider.FromHeaders) > 0 {
			jwtProvider.FromHeaders = make([]resources.JwtHeader, len(envoyProvider.FromHeaders))
			for i, header := range envoyProvider.FromHeaders {
				jwtProvider.FromHeaders[i] = resources.JwtHeader{
					Name:        header.Name,
					ValuePrefix: header.ValuePrefix,
				}
			}
		} else {
			// Add default from header
			jwtProvider.FromHeaders = make([]resources.JwtHeader, 1)
			jwtProvider.FromHeaders[0] = resources.JwtHeader{
				Name:        "Authorization",
				ValuePrefix: "Bearer ",
			}
		}
		// parse localjwks
		if localJwks := envoyProvider.GetLocalJwks(); localJwks != nil {
			// only get from linlineString or inlinebytes
			inLineString := ""
			if len(localJwks.GetInlineString()) > 0 {
				inLineString = localJwks.GetInlineString()
			}
			if localJwks.GetInlineBytes() != nil {
				inLineString = string(localJwks.GetInlineBytes())
			}
			jwkSet, err := resources.UnmarshalJwks(inLineString)
			if err != nil {
				return jwtAuthnFilter, err
			}
			jwtProvider.LocalJwks = &resources.LocalJwks{
				InlineString: inLineString,
				Keys:         jwkSet,
			}
		}
		jwtAuthentication.Providers[jwtProvider.ProviderName] = jwtProvider
	}
	// parse rules
	for _, rule := range envoyAuthentication.Rules {
		jwtRequirementRule := &resources.JwtRequirementRule{}
		// get routematch
		routeMatch := &resources.JwtRouteMatch{}
		if len(rule.GetMatch().GetPrefix()) > 0 {
			routeMatch.Action = "prefix"
			routeMatch.Value = rule.GetMatch().GetPrefix()
		}
		if len(rule.GetMatch().GetPath()) > 0 {
			routeMatch.Action = "path"
			routeMatch.Value = rule.GetMatch().GetPath()
		}
		if rule.GetMatch().GetSafeRegex() != nil {
			routeMatch.Action = "saferegex"
			routeMatch.Value = rule.GetMatch().GetSafeRegex().GetRegex()
		}
		if rule.GetMatch().GetCaseSensitive() != nil {
			routeMatch.CaseSensitive = rule.GetMatch().GetCaseSensitive().Value
		}
		jwtRequirementRule.Match = routeMatch
		// get requirement
		if rule.GetRequirementName() != "" {
			jwtRequirementRule.RequirementName = rule.GetRequirementName()
		}
		// parse requires
		if rule.GetRequires() != nil {
			jwtRequirement := &resources.SimpleJwtRequirement{}
			requires := rule.GetRequires()
			nextNames, m, f := getRequiresProviderNames(requires)
			jwtRequirement.ProviderNames = nextNames
			jwtRequirement.AllowMissing = m
			jwtRequirement.AllowMissingOrFailed = f
			if requires.GetAllowMissing() != nil {
				jwtRequirement.AllowMissing = true
			}
			if requires.GetAllowMissingOrFailed() != nil {
				jwtRequirement.AllowMissingOrFailed = true
			}
			jwtRequirementRule.Requires = jwtRequirement
		}

		jwtAuthentication.Rules = append(jwtAuthentication.Rules, jwtRequirementRule)

	}
	jwtAuthnFilter.JwtAuthentication = jwtAuthentication
	return jwtAuthnFilter, nil
}

func getRequiresProviderNames(requires *jwtauthnv3.JwtRequirement) ([]string, bool, bool) {
	providerNames := make([]string, 0)
	AllowMissing := false
	AllowMissingOrFailed := false
	if requires.GetProviderName() != "" {
		providerNames = append(providerNames, requires.GetProviderName())
	}
	if requires.GetRequiresAny() != nil {
		for _, anyRequires := range requires.GetRequiresAny().GetRequirements() {
			nextNames, m, f := getRequiresProviderNames(anyRequires)
			providerNames = append(providerNames, nextNames...)
			AllowMissing = m
			AllowMissingOrFailed = f
		}
	}
	if requires.GetRequiresAll() != nil {
		for _, anyRequires := range requires.GetRequiresAll().GetRequirements() {
			nextNames, m, f := getRequiresProviderNames(anyRequires)
			providerNames = append(providerNames, nextNames...)
			AllowMissing = m
			AllowMissingOrFailed = f
		}
	}
	if requires.GetAllowMissing() != nil {
		AllowMissing = true
	}
	if requires.GetAllowMissingOrFailed() != nil {
		AllowMissingOrFailed = true
	}
	return providerNames, AllowMissing, AllowMissingOrFailed
}
