// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: envoy/extensions/filters/udp/udp_proxy/session/dynamic_forward_proxy/v3/dynamic_forward_proxy.proto

package dynamic_forward_proxyv3

import (
	_ "github.com/alibaba/sentinel-golang/pkg/datasource/xds/cncf/go/udpa/annotations"
	v3 "github.com/alibaba/sentinel-golang/pkg/datasource/xds/go-control-plane/envoy/extensions/common/dynamic_forward_proxy/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Configuration for the filter state based dynamic forward proxy filter. See the
// :ref:`architecture overview <arch_overview_http_dynamic_forward_proxy>` for
// more information. Note this filter must be used in conjunction to another filter that
// sets the 'envoy.upstream.dynamic_host' and the 'envoy.upstream.dynamic_port' filter
// state keys for the required upstream UDP session.
// [#extension: envoy.filters.udp.session.dynamic_forward_proxy]
type FilterConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The prefix to use when emitting :ref:`statistics <config_udp_session_filters_dynamic_forward_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Types that are assignable to ImplementationSpecifier:
	//
	//	*FilterConfig_DnsCacheConfig
	ImplementationSpecifier isFilterConfig_ImplementationSpecifier `protobuf_oneof:"implementation_specifier"`
	// If configured, the filter will buffer datagrams in case that it is waiting for a DNS response.
	// If this field is not configured, there will be no buffering and downstream datagrams that arrive
	// while the DNS resolution is in progress will be dropped. In case this field is set but the options
	// are not configured, the default values will be applied as described in the “BufferOptions“.
	BufferOptions *FilterConfig_BufferOptions `protobuf:"bytes,3,opt,name=buffer_options,json=bufferOptions,proto3" json:"buffer_options,omitempty"`
}

func (x *FilterConfig) Reset() {
	*x = FilterConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterConfig) ProtoMessage() {}

func (x *FilterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterConfig.ProtoReflect.Descriptor instead.
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescGZIP(), []int{0}
}

func (x *FilterConfig) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

func (m *FilterConfig) GetImplementationSpecifier() isFilterConfig_ImplementationSpecifier {
	if m != nil {
		return m.ImplementationSpecifier
	}
	return nil
}

func (x *FilterConfig) GetDnsCacheConfig() *v3.DnsCacheConfig {
	if x, ok := x.GetImplementationSpecifier().(*FilterConfig_DnsCacheConfig); ok {
		return x.DnsCacheConfig
	}
	return nil
}

func (x *FilterConfig) GetBufferOptions() *FilterConfig_BufferOptions {
	if x != nil {
		return x.BufferOptions
	}
	return nil
}

type isFilterConfig_ImplementationSpecifier interface {
	isFilterConfig_ImplementationSpecifier()
}

type FilterConfig_DnsCacheConfig struct {
	// The DNS cache configuration that the filter will attach to. Note this
	// configuration must match that of associated :ref:`dynamic forward proxy cluster configuration
	// <envoy_v3_api_field_extensions.clusters.dynamic_forward_proxy.v3.ClusterConfig.dns_cache_config>`.
	DnsCacheConfig *v3.DnsCacheConfig `protobuf:"bytes,2,opt,name=dns_cache_config,json=dnsCacheConfig,proto3,oneof"`
}

func (*FilterConfig_DnsCacheConfig) isFilterConfig_ImplementationSpecifier() {}

// Configuration for UDP datagrams buffering.
type FilterConfig_BufferOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// If set, the filter will only buffer datagrams up to the requested limit, and will drop
	// new UDP datagrams if the buffer contains the max_buffered_datagrams value at the time
	// of a new datagram arrival. If not set, the default value is 1024 datagrams.
	MaxBufferedDatagrams *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=max_buffered_datagrams,json=maxBufferedDatagrams,proto3" json:"max_buffered_datagrams,omitempty"`
	// If set, the filter will only buffer datagrams up to the requested total buffered bytes limit,
	// and will drop new UDP datagrams if the buffer contains the max_buffered_datagrams value
	// at the time of a new datagram arrival. If not set, the default value is 16,384 (16KB).
	MaxBufferedBytes *wrappers.UInt64Value `protobuf:"bytes,2,opt,name=max_buffered_bytes,json=maxBufferedBytes,proto3" json:"max_buffered_bytes,omitempty"`
}

func (x *FilterConfig_BufferOptions) Reset() {
	*x = FilterConfig_BufferOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilterConfig_BufferOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilterConfig_BufferOptions) ProtoMessage() {}

func (x *FilterConfig_BufferOptions) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilterConfig_BufferOptions.ProtoReflect.Descriptor instead.
func (*FilterConfig_BufferOptions) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FilterConfig_BufferOptions) GetMaxBufferedDatagrams() *wrappers.UInt32Value {
	if x != nil {
		return x.MaxBufferedDatagrams
	}
	return nil
}

func (x *FilterConfig_BufferOptions) GetMaxBufferedBytes() *wrappers.UInt64Value {
	if x != nil {
		return x.MaxBufferedBytes
	}
	return nil
}

var File_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDesc = []byte{
	0x0a, 0x63, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x64, 0x70, 0x2f, 0x75,
	0x64, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x47, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x75, 0x64, 0x70, 0x2e, 0x75, 0x64, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x1a, 0x40,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x76, 0x33,
	0x2f, 0x64, 0x6e, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x04, 0x0a, 0x0c, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61,
	0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x76, 0x0a, 0x10, 0x64, 0x6e, 0x73, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33,
	0x2e, 0x44, 0x6e, 0x73, 0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x0e, 0x64, 0x6e, 0x73,
	0x43, 0x61, 0x63, 0x68, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x8a, 0x01, 0x0a, 0x0e,
	0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x63, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e,
	0x75, 0x64, 0x70, 0x2e, 0x75, 0x64, 0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x2e, 0x46,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x42, 0x75, 0x66, 0x66,
	0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0d, 0x62, 0x75, 0x66, 0x66, 0x65,
	0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0xaf, 0x01, 0x0a, 0x0d, 0x42, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x52, 0x0a, 0x16, 0x6d, 0x61,
	0x78, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x67,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x42, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x4a,
	0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x42, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x42, 0x1a, 0x0a, 0x18, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x65, 0x72, 0x42, 0x84, 0x02, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x02, 0x0a, 0x55, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x75, 0x64, 0x70, 0x2e, 0x75, 0x64,
	0x70, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x76, 0x33, 0x42, 0x18, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69,
	0x63, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x86, 0x01, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x64, 0x70, 0x2f, 0x75, 0x64, 0x70, 0x5f, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x76, 0x33, 0x3b, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x76, 0x33, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescData = file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDesc
)

func file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescData)
	})
	return file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDescData
}

var file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_goTypes = []interface{}{
	(*FilterConfig)(nil),               // 0: envoy.extensions.filters.udp.udp_proxy.session.dynamic_forward_proxy.v3.FilterConfig
	(*FilterConfig_BufferOptions)(nil), // 1: envoy.extensions.filters.udp.udp_proxy.session.dynamic_forward_proxy.v3.FilterConfig.BufferOptions
	(*v3.DnsCacheConfig)(nil),          // 2: envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig
	(*wrappers.UInt32Value)(nil),       // 3: google.protobuf.UInt32Value
	(*wrappers.UInt64Value)(nil),       // 4: google.protobuf.UInt64Value
}
var file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.udp.udp_proxy.session.dynamic_forward_proxy.v3.FilterConfig.dns_cache_config:type_name -> envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig
	1, // 1: envoy.extensions.filters.udp.udp_proxy.session.dynamic_forward_proxy.v3.FilterConfig.buffer_options:type_name -> envoy.extensions.filters.udp.udp_proxy.session.dynamic_forward_proxy.v3.FilterConfig.BufferOptions
	3, // 2: envoy.extensions.filters.udp.udp_proxy.session.dynamic_forward_proxy.v3.FilterConfig.BufferOptions.max_buffered_datagrams:type_name -> google.protobuf.UInt32Value
	4, // 3: envoy.extensions.filters.udp.udp_proxy.session.dynamic_forward_proxy.v3.FilterConfig.BufferOptions.max_buffered_bytes:type_name -> google.protobuf.UInt64Value
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_init()
}
func file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_init() {
	if File_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilterConfig_BufferOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*FilterConfig_DnsCacheConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto = out.File
	file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_rawDesc = nil
	file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_goTypes = nil
	file_envoy_extensions_filters_udp_udp_proxy_session_dynamic_forward_proxy_v3_dynamic_forward_proxy_proto_depIdxs = nil
}
