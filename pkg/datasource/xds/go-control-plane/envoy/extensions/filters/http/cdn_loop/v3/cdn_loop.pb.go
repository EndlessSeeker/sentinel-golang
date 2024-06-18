// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.4
// source: envoy/extensions/filters/http/cdn_loop/v3/cdn_loop.proto

package cdn_loopv3

import (
	_ "github.com/alibaba/sentinel-golang/pkg/datasource/xds/cncf/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// CDN-Loop Header filter config. See the :ref:`configuration overview
// <config_http_filters_cdn_loop>` for more information.
type CdnLoopConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The CDN identifier to use for loop checks and to append to the
	// CDN-Loop header.
	//
	// RFC 8586 calls this the cdn-id. The cdn-id can either be a
	// pseudonym or hostname the CDN is in control of.
	//
	// cdn_id must not be empty.
	CdnId string `protobuf:"bytes,1,opt,name=cdn_id,json=cdnId,proto3" json:"cdn_id,omitempty"`
	// The maximum allowed count of cdn_id in the downstream CDN-Loop
	// request header.
	//
	// The default of 0 means a request can transit the CdnLoopFilter
	// once. A value of 1 means that a request can transit the
	// CdnLoopFilter twice and so on.
	MaxAllowedOccurrences uint32 `protobuf:"varint,2,opt,name=max_allowed_occurrences,json=maxAllowedOccurrences,proto3" json:"max_allowed_occurrences,omitempty"`
}

func (x *CdnLoopConfig) Reset() {
	*x = CdnLoopConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CdnLoopConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CdnLoopConfig) ProtoMessage() {}

func (x *CdnLoopConfig) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CdnLoopConfig.ProtoReflect.Descriptor instead.
func (*CdnLoopConfig) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_rawDescGZIP(), []int{0}
}

func (x *CdnLoopConfig) GetCdnId() string {
	if x != nil {
		return x.CdnId
	}
	return ""
}

func (x *CdnLoopConfig) GetMaxAllowedOccurrences() uint32 {
	if x != nil {
		return x.MaxAllowedOccurrences
	}
	return 0
}

var File_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_rawDesc = []byte{
	0x0a, 0x38, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x63, 0x64, 0x6e, 0x5f, 0x6c, 0x6f, 0x6f, 0x70, 0x2f, 0x76, 0x33, 0x2f, 0x63, 0x64, 0x6e, 0x5f,
	0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63, 0x64, 0x6e, 0x5f, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67, 0x0a,
	0x0d, 0x43, 0x64, 0x6e, 0x4c, 0x6f, 0x6f, 0x70, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1e,
	0x0a, 0x06, 0x63, 0x64, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x05, 0x63, 0x64, 0x6e, 0x49, 0x64, 0x12, 0x36,
	0x0a, 0x17, 0x6d, 0x61, 0x78, 0x5f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x5f, 0x6f, 0x63,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x15, 0x6d, 0x61, 0x78, 0x41, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x4f, 0x63, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x42, 0xae, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10,
	0x02, 0x0a, 0x37, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x63,
	0x64, 0x6e, 0x5f, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x76, 0x33, 0x42, 0x0c, 0x43, 0x64, 0x6e, 0x4c,
	0x6f, 0x6f, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61,
	0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70,
	0x2f, 0x63, 0x64, 0x6e, 0x5f, 0x6c, 0x6f, 0x6f, 0x70, 0x2f, 0x76, 0x33, 0x3b, 0x63, 0x64, 0x6e,
	0x5f, 0x6c, 0x6f, 0x6f, 0x70, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_rawDescData = file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_rawDesc
)

func file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_rawDescData
}

var file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_goTypes = []interface{}{
	(*CdnLoopConfig)(nil), // 0: envoy.extensions.filters.http.cdn_loop.v3.CdnLoopConfig
}
var file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_init() }
func file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_init() {
	if File_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CdnLoopConfig); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto = out.File
	file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_rawDesc = nil
	file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_goTypes = nil
	file_envoy_extensions_filters_http_cdn_loop_v3_cdn_loop_proto_depIdxs = nil
}
