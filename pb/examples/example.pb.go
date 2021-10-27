//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: pb/examples/example.proto

package examples

import (
	_ "github.com/infraboard/mcube/cmd/protoc-gen-go-ext/extension/tag"
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

type GrantType int32

const (
	GrantType_NULL      GrantType = 0
	GrantType_UNKNOWN   GrantType = 1
	GrantType_PASSWORD  GrantType = 2
	GrantType_LDAP      GrantType = 3
	GrantType_REFRESH   GrantType = 4
	GrantType_ACCESS    GrantType = 5
	GrantType_CLIENT    GrantType = 6
	GrantType_AUTH_CODE GrantType = 7
	GrantType_IMPLICIT  GrantType = 8
)

// Enum value maps for GrantType.
var (
	GrantType_name = map[int32]string{
		0: "NULL",
		1: "UNKNOWN",
		2: "PASSWORD",
		3: "LDAP",
		4: "REFRESH",
		5: "ACCESS",
		6: "CLIENT",
		7: "AUTH_CODE",
		8: "IMPLICIT",
	}
	GrantType_value = map[string]int32{
		"NULL":      0,
		"UNKNOWN":   1,
		"PASSWORD":  2,
		"LDAP":      3,
		"REFRESH":   4,
		"ACCESS":    5,
		"CLIENT":    6,
		"AUTH_CODE": 7,
		"IMPLICIT":  8,
	}
)

func (x GrantType) Enum() *GrantType {
	p := new(GrantType)
	*p = x
	return p
}

func (x GrantType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GrantType) Descriptor() protoreflect.EnumDescriptor {
	return file_pb_examples_example_proto_enumTypes[0].Descriptor()
}

func (GrantType) Type() protoreflect.EnumType {
	return &file_pb_examples_example_proto_enumTypes[0]
}

func (x GrantType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GrantType.Descriptor instead.
func (GrantType) EnumDescriptor() ([]byte, []int) {
	return file_pb_examples_example_proto_rawDescGZIP(), []int{0}
}

type Http struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol string `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty" validate:"oneof=http https"`
	// version_default is the same as version_update
	VersionDefault string    `protobuf:"bytes,2,opt,name=version_default,json=VersionDefault,proto3" json:"version_with_default" validate:"gte=0,lte=130" bson:"xxxx"`
	VersionUpdate  string    `protobuf:"bytes,3,opt,name=version_update,json=VersionUpdate,proto3" json:"version_with_update" validate:"gte=0,lte=130"`
	VersionReplace string    `protobuf:"bytes,4,opt,name=version_replace,json=VersionReplace,proto3" json:"version_with_replace" validate:"gte=0,lte=130"`
	Url            *Http_Url `protobuf:"bytes,5,opt,name=url,json=Url,proto3" json:"url_tag,omitempty"`
	VersionSkip    string    `protobuf:"bytes,6,opt,name=version_skip,json=versionSkip,proto3" json:"-" bson:"-"`
}

func (x *Http) Reset() {
	*x = Http{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_examples_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http) ProtoMessage() {}

func (x *Http) ProtoReflect() protoreflect.Message {
	mi := &file_pb_examples_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http.ProtoReflect.Descriptor instead.
func (*Http) Descriptor() ([]byte, []int) {
	return file_pb_examples_example_proto_rawDescGZIP(), []int{0}
}

func (x *Http) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *Http) GetVersionDefault() string {
	if x != nil {
		return x.VersionDefault
	}
	return ""
}

func (x *Http) GetVersionUpdate() string {
	if x != nil {
		return x.VersionUpdate
	}
	return ""
}

func (x *Http) GetVersionReplace() string {
	if x != nil {
		return x.VersionReplace
	}
	return ""
}

func (x *Http) GetUrl() *Http_Url {
	if x != nil {
		return x.Url
	}
	return nil
}

func (x *Http) GetVersionSkip() string {
	if x != nil {
		return x.VersionSkip
	}
	return ""
}

type Http_Url struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scheme string `protobuf:"bytes,1,opt,name=scheme,json=Scheme,proto3" json:"schema_tag,omitempty"` //  string scheme = 1[json_name = "Scheme"];
}

func (x *Http_Url) Reset() {
	*x = Http_Url{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_examples_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http_Url) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http_Url) ProtoMessage() {}

func (x *Http_Url) ProtoReflect() protoreflect.Message {
	mi := &file_pb_examples_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http_Url.ProtoReflect.Descriptor instead.
func (*Http_Url) Descriptor() ([]byte, []int) {
	return file_pb_examples_example_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Http_Url) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

var File_pb_examples_example_proto protoreflect.FileDescriptor

var file_pb_examples_example_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72,
	0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x63, 0x6d, 0x64,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65,
	0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67,
	0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x04, 0x0a, 0x04, 0x48,
	0x74, 0x74, 0x70, 0x12, 0x3d, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xc2, 0xde, 0x1f, 0x1d, 0x0a, 0x1b, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x3d, 0x68, 0x74, 0x74,
	0x70, 0x20, 0x68, 0x74, 0x74, 0x70, 0x73, 0x22, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x12, 0x6f, 0x0a, 0x0f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x46, 0xc2, 0xde, 0x1f,
	0x42, 0x0a, 0x40, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x67, 0x74, 0x65,
	0x3d, 0x30, 0x2c, 0x6c, 0x74, 0x65, 0x3d, 0x31, 0x33, 0x30, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x22, 0x20, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x78, 0x78,
	0x78, 0x78, 0x22, 0x52, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x12, 0x60, 0x0a, 0x0e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x39, 0xc2, 0xde, 0x1f,
	0x35, 0x0a, 0x33, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x67, 0x74, 0x65,
	0x3d, 0x30, 0x2c, 0x6c, 0x74, 0x65, 0x3d, 0x31, 0x33, 0x30, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x22, 0x52, 0x0d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x63, 0x0a, 0x0f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x3a,
	0xc2, 0xde, 0x1f, 0x36, 0x0a, 0x34, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22,
	0x67, 0x74, 0x65, 0x3d, 0x30, 0x2c, 0x6c, 0x74, 0x65, 0x3d, 0x31, 0x33, 0x30, 0x22, 0x20, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x77, 0x69, 0x74,
	0x68, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x22, 0x52, 0x0e, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x48, 0x74, 0x74,
	0x70, 0x2e, 0x55, 0x72, 0x6c, 0x42, 0x1e, 0xc2, 0xde, 0x1f, 0x1a, 0x0a, 0x18, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x75, 0x72, 0x6c, 0x5f, 0x74, 0x61, 0x67, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x3a, 0x0a, 0x0c, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x6b, 0x69, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x17, 0xc2, 0xde, 0x1f, 0x13, 0x0a, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2d, 0x22,
	0x20, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x2d, 0x22, 0x52, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x6b, 0x69, 0x70, 0x1a, 0x40, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x39, 0x0a,
	0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xc2,
	0xde, 0x1f, 0x1d, 0x0a, 0x1b, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x5f, 0x74, 0x61, 0x67, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x52, 0x06, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x2a, 0x7c, 0x0a, 0x09, 0x47, 0x72, 0x61, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x55, 0x4c, 0x4c, 0x10, 0x00, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08,
	0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x4c, 0x44,
	0x41, 0x50, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x46, 0x52, 0x45, 0x53, 0x48, 0x10,
	0x04, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x05, 0x12, 0x0a, 0x0a,
	0x06, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x55, 0x54,
	0x48, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4d, 0x50, 0x4c,
	0x49, 0x43, 0x49, 0x54, 0x10, 0x08, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f,
	0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_examples_example_proto_rawDescOnce sync.Once
	file_pb_examples_example_proto_rawDescData = file_pb_examples_example_proto_rawDesc
)

func file_pb_examples_example_proto_rawDescGZIP() []byte {
	file_pb_examples_example_proto_rawDescOnce.Do(func() {
		file_pb_examples_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_examples_example_proto_rawDescData)
	})
	return file_pb_examples_example_proto_rawDescData
}

var file_pb_examples_example_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pb_examples_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pb_examples_example_proto_goTypes = []interface{}{
	(GrantType)(0),   // 0: pb.GrantType
	(*Http)(nil),     // 1: pb.Http
	(*Http_Url)(nil), // 2: pb.Http.Url
}
var file_pb_examples_example_proto_depIdxs = []int32{
	2, // 0: pb.Http.url:type_name -> pb.Http.Url
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_examples_example_proto_init() }
func file_pb_examples_example_proto_init() {
	if File_pb_examples_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_examples_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http); i {
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
		file_pb_examples_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http_Url); i {
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
			RawDescriptor: file_pb_examples_example_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_examples_example_proto_goTypes,
		DependencyIndexes: file_pb_examples_example_proto_depIdxs,
		EnumInfos:         file_pb_examples_example_proto_enumTypes,
		MessageInfos:      file_pb_examples_example_proto_msgTypes,
	}.Build()
	File_pb_examples_example_proto = out.File
	file_pb_examples_example_proto_rawDesc = nil
	file_pb_examples_example_proto_goTypes = nil
	file_pb_examples_example_proto_depIdxs = nil
}
