//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: pb/http/entry.proto

package http

import (
	_ "github.com/infraboard/mcube/cmd/protoc-gen-go-ext/extension/tag"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Entry 路由条目
type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 函数名称
	FunctionName string `protobuf:"bytes,2,opt,name=function_name,json=functionName,proto3" json:"function_name" bson:"function_name"`
	// HTTP path 用于自动生成http api
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path" bson:"path"`
	// HTTP method 用于自动生成http api
	Method string `protobuf:"bytes,4,opt,name=method,proto3" json:"method" bson:"method"`
	// 资源名称
	Resource string `protobuf:"bytes,5,opt,name=resource,proto3" json:"resource" bson:"resource"`
	// 是否校验用户身份 (acccess_token)
	AuthEnable bool `protobuf:"varint,6,opt,name=auth_enable,json=authEnable,proto3" json:"auth_enable" bson:"auth_enable"`
	// 是否校验用户权限
	PermissionEnable bool `protobuf:"varint,7,opt,name=permission_enable,json=permissionEnable,proto3" json:"permission_enable" bson:"permission_enable"`
	// 允许的通过的身份标识符, 比如角色, 用户类型之类
	Allow []string `protobuf:"bytes,12,rep,name=allow,proto3" json:"allow" bson:"allow"`
	// 是否开启操作审计, 开启后这次操作将被记录
	AuditLog bool `protobuf:"varint,9,opt,name=audit_log,json=auditLog,proto3" json:"audit_log" bson:"audit_log"`
	// 名称空间不能为空
	RequiredNamespace bool `protobuf:"varint,10,opt,name=required_namespace,json=requiredNamespace,proto3" json:"required_namespace" bson:"required_namespace"`
	// 标签
	Labels map[string]string `protobuf:"bytes,8,rep,name=labels,proto3" json:"labels" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"labels"`
	// 扩展属性
	Extension map[string]string `protobuf:"bytes,11,rep,name=extension,proto3" json:"extension" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"extension"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_http_entry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_pb_http_entry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_pb_http_entry_proto_rawDescGZIP(), []int{0}
}

func (x *Entry) GetFunctionName() string {
	if x != nil {
		return x.FunctionName
	}
	return ""
}

func (x *Entry) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Entry) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Entry) GetResource() string {
	if x != nil {
		return x.Resource
	}
	return ""
}

func (x *Entry) GetAuthEnable() bool {
	if x != nil {
		return x.AuthEnable
	}
	return false
}

func (x *Entry) GetPermissionEnable() bool {
	if x != nil {
		return x.PermissionEnable
	}
	return false
}

func (x *Entry) GetAllow() []string {
	if x != nil {
		return x.Allow
	}
	return nil
}

func (x *Entry) GetAuditLog() bool {
	if x != nil {
		return x.AuditLog
	}
	return false
}

func (x *Entry) GetRequiredNamespace() bool {
	if x != nil {
		return x.RequiredNamespace
	}
	return false
}

func (x *Entry) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Entry) GetExtension() map[string]string {
	if x != nil {
		return x.Extension
	}
	return nil
}

var file_pb_http_entry_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Entry)(nil),
		Field:         20210228,
		Name:          "mcube.http.rest_api",
		Tag:           "bytes,20210228,opt,name=rest_api",
		Filename:      "pb/http/entry.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional mcube.http.Entry rest_api = 20210228;
	E_RestApi = &file_pb_http_entry_proto_extTypes[0]
)

var File_pb_http_entry_proto protoreflect.FileDescriptor

var file_pb_http_entry_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x62, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x1a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e,
	0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x63,
	0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x2d, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x74,
	0x61, 0x67, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe5,
	0x07, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x54, 0x0a, 0x0d, 0x66, 0x75, 0x6e, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2f, 0xc2, 0xde, 0x1f, 0x2b, 0x0a, 0x29, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x66, 0x75, 0x6e,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x52, 0x0c, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x31,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xc2, 0xde,
	0x1f, 0x19, 0x0a, 0x17, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x61, 0x74, 0x68, 0x22, 0x20,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x61, 0x74, 0x68, 0x22, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x39, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x21, 0xc2, 0xde, 0x1f, 0x1d, 0x0a, 0x1b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x22, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x41, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x25,
	0xc2, 0xde, 0x1f, 0x21, 0x0a, 0x1f, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x22, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x4c, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x42, 0x2b, 0xc2, 0xde, 0x1f, 0x27, 0x0a, 0x25, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x20, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x22, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x64, 0x0a,
	0x11, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x42, 0x37, 0xc2, 0xde, 0x1f, 0x33, 0x0a, 0x31,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x22, 0x52, 0x10, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x0c, 0x20, 0x03,
	0x28, 0x09, 0x42, 0x1f, 0xc2, 0xde, 0x1f, 0x1b, 0x0a, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x22, 0x52, 0x05, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x44, 0x0a, 0x09, 0x61, 0x75,
	0x64, 0x69, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x42, 0x27, 0xc2,
	0xde, 0x1f, 0x23, 0x0a, 0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x5f, 0x6c, 0x6f, 0x67, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x22, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67,
	0x12, 0x68, 0x0a, 0x12, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x42, 0x39, 0xc2, 0xde,
	0x1f, 0x35, 0x0a, 0x33, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x52, 0x11, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x63, 0x75,
	0x62, 0x65, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x21, 0xc2, 0xde, 0x1f, 0x1d, 0x0a,
	0x1b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x20, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x67, 0x0a, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a,
	0x21, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x39, 0x0a,
	0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x4f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x61,
	0x70, 0x69, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xb4, 0xc4, 0xd1, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d, 0x63,
	0x75, 0x62, 0x65, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x72, 0x65, 0x73, 0x74, 0x41, 0x70, 0x69, 0x42, 0x73, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x42, 0x0a,
	0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01, 0x5a, 0x28, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x2f, 0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x68, 0x74, 0x74,
	0x70, 0x3b, 0x68, 0x74, 0x74, 0x70, 0xf8, 0x01, 0x01, 0xa2, 0x02, 0x03, 0x47, 0x50, 0x42, 0xaa,
	0x02, 0x1a, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x52, 0x65, 0x66, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_http_entry_proto_rawDescOnce sync.Once
	file_pb_http_entry_proto_rawDescData = file_pb_http_entry_proto_rawDesc
)

func file_pb_http_entry_proto_rawDescGZIP() []byte {
	file_pb_http_entry_proto_rawDescOnce.Do(func() {
		file_pb_http_entry_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_http_entry_proto_rawDescData)
	})
	return file_pb_http_entry_proto_rawDescData
}

var file_pb_http_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_http_entry_proto_goTypes = []interface{}{
	(*Entry)(nil),                      // 0: mcube.http.Entry
	nil,                                // 1: mcube.http.Entry.LabelsEntry
	nil,                                // 2: mcube.http.Entry.ExtensionEntry
	(*descriptorpb.MethodOptions)(nil), // 3: google.protobuf.MethodOptions
}
var file_pb_http_entry_proto_depIdxs = []int32{
	1, // 0: mcube.http.Entry.labels:type_name -> mcube.http.Entry.LabelsEntry
	2, // 1: mcube.http.Entry.extension:type_name -> mcube.http.Entry.ExtensionEntry
	3, // 2: mcube.http.rest_api:extendee -> google.protobuf.MethodOptions
	0, // 3: mcube.http.rest_api:type_name -> mcube.http.Entry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	3, // [3:4] is the sub-list for extension type_name
	2, // [2:3] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_http_entry_proto_init() }
func file_pb_http_entry_proto_init() {
	if File_pb_http_entry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_http_entry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
			RawDescriptor: file_pb_http_entry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_pb_http_entry_proto_goTypes,
		DependencyIndexes: file_pb_http_entry_proto_depIdxs,
		MessageInfos:      file_pb_http_entry_proto_msgTypes,
		ExtensionInfos:    file_pb_http_entry_proto_extTypes,
	}.Build()
	File_pb_http_entry_proto = out.File
	file_pb_http_entry_proto_rawDesc = nil
	file_pb_http_entry_proto_goTypes = nil
	file_pb_http_entry_proto_depIdxs = nil
}