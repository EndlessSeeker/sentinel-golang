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

package resources

import "strings"

type XdsListener struct {
	Name             string
	HasRds           bool
	RdsResourceNames []string
	TrafficDirection string
	// virtual inbound 15006 listener
	IsVirtualInbound bool
	//FilterChains     []XdsFilterChain
	// virtual inbound 15006 listener tls and downstream transport socket which is for mtls
	InboundTLSMode                   XdsTLSMode
	InboundDownstreamTransportSocket XdsDownstreamTransportSocket
	JwtAuthnFilter                   JwtAuthnEnovyFilter
	RBACFilter                       RBACEnvoyFilter
}

type XdsDownstreamTransportSocket struct {
	SubjectAltNamesMatch string // exact, prefix
	SubjectAltNamesValue string
	//tlsContext               sockets_tls_v3.DownstreamTlsContext
	RequireClientCertificate bool
}

type XdsFilterChain struct {
	Name             string
	FilterChainMatch XdsFilterChainMatch
	TransportSocket  XdsDownstreamTransportSocket
	Filters          []XdsFilter
}

type XdsFilterChainMatch struct {
	DestinationPort   uint32
	TransportProtocol string
}

type XdsFilter struct {
	Name                         string
	IncludeHttpConnectionManager bool
	HasRds                       bool
	RdsResourceName              string
}

type XdsHostInboundListener struct {
	MutualTLSMode   MutualTLSMode
	TransportSocket XdsDownstreamTransportSocket
	JwtAuthnFilter  JwtAuthnEnovyFilter
	RBACFilter      RBACEnvoyFilter
	// other host inbound info for protocol export here
}

func MatchSpiffe(spiffee, action, value string) bool {
	spiffee = strings.ToLower(spiffee)
	value = strings.ToLower(value)
	action = strings.ToLower(action)
	if len(action) == 0 {
		return true
	}
	if action == "exact" {
		return spiffee == value
	}
	if action == "prefix" {
		return strings.HasPrefix(spiffee, value)
	}
	if action == "contains" {
		return strings.Contains(spiffee, value)
	}
	return false
}
