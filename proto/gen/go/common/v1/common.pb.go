//版本号

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: common/v1/common.proto

//import "github.com/gogo/protobuf/gogoproto/gogo.proto";
//import "github.com/envoyproxy/protoc-gen-validate/validate/validate.proto";

//指定包名

package common_v1

import (
	types "github.com/gogo/protobuf/types"
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

type ErrCode int32

const (
	ErrCode_ERR_CODE_UNSPECIFIED            ErrCode = 0
	ErrCode_ERR_CODE_SERVER_UNKNOWN         ErrCode = 999999  // 未知错误
	ErrCode_ERR_CODE_SERVER_FAILED          ErrCode = -100001 // 不重要的失败，建议重试
	ErrCode_ERR_CODE_SERVER_PARAM_INVALID   ErrCode = -100002 // 请求参数不合法
	ErrCode_ERR_CODE_SERVER_EXCEPTION       ErrCode = -100003 // 服务本身异常
	ErrCode_ERR_CODE_SERVER_REJECTED        ErrCode = -100004 // 服务忙，请求队列已满
	ErrCode_ERR_CODE_SERVER_BUSY            ErrCode = -100005 // 服务忙，被调限流
	ErrCode_ERR_CODE_CLIENT_TIMEOUT         ErrCode = -100006 // 服务调用超时
	ErrCode_ERR_CODE_CLIENT_EXCEPTION       ErrCode = -100007 // 服务调用异常
	ErrCode_ERR_CODE_CLIENT_BROKEN          ErrCode = -100008 // 服务熔断
	ErrCode_ERR_CODE_CONTEXT_NO_DEVICE      ErrCode = -200000 // context 中没有device
	ErrCode_ERR_CODE_CONTEXT_NO_INPARAM     ErrCode = -200001 // context 中没有in_param
	ErrCode_ERR_CODE_RECORD_NOT_FOUND       ErrCode = -200002 // not found
	ErrCode_ERR_CODE_CONTEXT_DEVICE_INVALID ErrCode = -200003 // device invalid
	ErrCode_ERR_CODE_CONTEXT_NO_USER        ErrCode = -200004 // context 中没有user
)

// Enum value maps for ErrCode.
var (
	ErrCode_name = map[int32]string{
		0:       "ERR_CODE_UNSPECIFIED",
		999999:  "ERR_CODE_SERVER_UNKNOWN",
		-100001: "ERR_CODE_SERVER_FAILED",
		-100002: "ERR_CODE_SERVER_PARAM_INVALID",
		-100003: "ERR_CODE_SERVER_EXCEPTION",
		-100004: "ERR_CODE_SERVER_REJECTED",
		-100005: "ERR_CODE_SERVER_BUSY",
		-100006: "ERR_CODE_CLIENT_TIMEOUT",
		-100007: "ERR_CODE_CLIENT_EXCEPTION",
		-100008: "ERR_CODE_CLIENT_BROKEN",
		-200000: "ERR_CODE_CONTEXT_NO_DEVICE",
		-200001: "ERR_CODE_CONTEXT_NO_INPARAM",
		-200002: "ERR_CODE_RECORD_NOT_FOUND",
		-200003: "ERR_CODE_CONTEXT_DEVICE_INVALID",
		-200004: "ERR_CODE_CONTEXT_NO_USER",
	}
	ErrCode_value = map[string]int32{
		"ERR_CODE_UNSPECIFIED":            0,
		"ERR_CODE_SERVER_UNKNOWN":         999999,
		"ERR_CODE_SERVER_FAILED":          -100001,
		"ERR_CODE_SERVER_PARAM_INVALID":   -100002,
		"ERR_CODE_SERVER_EXCEPTION":       -100003,
		"ERR_CODE_SERVER_REJECTED":        -100004,
		"ERR_CODE_SERVER_BUSY":            -100005,
		"ERR_CODE_CLIENT_TIMEOUT":         -100006,
		"ERR_CODE_CLIENT_EXCEPTION":       -100007,
		"ERR_CODE_CLIENT_BROKEN":          -100008,
		"ERR_CODE_CONTEXT_NO_DEVICE":      -200000,
		"ERR_CODE_CONTEXT_NO_INPARAM":     -200001,
		"ERR_CODE_RECORD_NOT_FOUND":       -200002,
		"ERR_CODE_CONTEXT_DEVICE_INVALID": -200003,
		"ERR_CODE_CONTEXT_NO_USER":        -200004,
	}
)

func (x ErrCode) Enum() *ErrCode {
	p := new(ErrCode)
	*p = x
	return p
}

func (x ErrCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_common_v1_common_proto_enumTypes[0].Descriptor()
}

func (ErrCode) Type() protoreflect.EnumType {
	return &file_common_v1_common_proto_enumTypes[0]
}

func (x ErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrCode.Descriptor instead.
func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{0}
}

type ErrSubCode int32

const (
	ErrSubCode_ERR_SUB_CODE_UNSPECIFIED  ErrSubCode = 0
	ErrSubCode_ERR_SUB_CODE_BOOK_SERVICE ErrSubCode = 101 // 内容库服务
)

// Enum value maps for ErrSubCode.
var (
	ErrSubCode_name = map[int32]string{
		0:   "ERR_SUB_CODE_UNSPECIFIED",
		101: "ERR_SUB_CODE_BOOK_SERVICE",
	}
	ErrSubCode_value = map[string]int32{
		"ERR_SUB_CODE_UNSPECIFIED":  0,
		"ERR_SUB_CODE_BOOK_SERVICE": 101,
	}
)

func (x ErrSubCode) Enum() *ErrSubCode {
	p := new(ErrSubCode)
	*p = x
	return p
}

func (x ErrSubCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrSubCode) Descriptor() protoreflect.EnumDescriptor {
	return file_common_v1_common_proto_enumTypes[1].Descriptor()
}

func (ErrSubCode) Type() protoreflect.EnumType {
	return &file_common_v1_common_proto_enumTypes[1]
}

func (x ErrSubCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrSubCode.Descriptor instead.
func (ErrSubCode) EnumDescriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{1}
}

type AppIdType int32

const (
	AppIdType_APP_ID_TYPE_UNSPECIFIED AppIdType = 0
	AppIdType_APP_ID_TYPE_LIGHTHOUSE  AppIdType = 100
)

// Enum value maps for AppIdType.
var (
	AppIdType_name = map[int32]string{
		0:   "APP_ID_TYPE_UNSPECIFIED",
		100: "APP_ID_TYPE_LIGHTHOUSE",
	}
	AppIdType_value = map[string]int32{
		"APP_ID_TYPE_UNSPECIFIED": 0,
		"APP_ID_TYPE_LIGHTHOUSE":  100,
	}
)

func (x AppIdType) Enum() *AppIdType {
	p := new(AppIdType)
	*p = x
	return p
}

func (x AppIdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AppIdType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_v1_common_proto_enumTypes[2].Descriptor()
}

func (AppIdType) Type() protoreflect.EnumType {
	return &file_common_v1_common_proto_enumTypes[2]
}

func (x AppIdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AppIdType.Descriptor instead.
func (AppIdType) EnumDescriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{2}
}

type AreaIdType int32

const (
	AreaIdType_AREA_ID_TYPE_UNSPECIFIED AreaIdType = 0
	AreaIdType_AREA_ID_TYPE_ANDROID     AreaIdType = 30
	AreaIdType_AREA_ID_TYPE_IOS         AreaIdType = 40
	AreaIdType_AREA_ID_TYPE_WEB         AreaIdType = 99
	AreaIdType_AREA_ID_TYPE_H5          AreaIdType = 98
)

// Enum value maps for AreaIdType.
var (
	AreaIdType_name = map[int32]string{
		0:  "AREA_ID_TYPE_UNSPECIFIED",
		30: "AREA_ID_TYPE_ANDROID",
		40: "AREA_ID_TYPE_IOS",
		99: "AREA_ID_TYPE_WEB",
		98: "AREA_ID_TYPE_H5",
	}
	AreaIdType_value = map[string]int32{
		"AREA_ID_TYPE_UNSPECIFIED": 0,
		"AREA_ID_TYPE_ANDROID":     30,
		"AREA_ID_TYPE_IOS":         40,
		"AREA_ID_TYPE_WEB":         99,
		"AREA_ID_TYPE_H5":          98,
	}
)

func (x AreaIdType) Enum() *AreaIdType {
	p := new(AreaIdType)
	*p = x
	return p
}

func (x AreaIdType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AreaIdType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_v1_common_proto_enumTypes[3].Descriptor()
}

func (AreaIdType) Type() protoreflect.EnumType {
	return &file_common_v1_common_proto_enumTypes[3]
}

func (x AreaIdType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AreaIdType.Descriptor instead.
func (AreaIdType) EnumDescriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{3}
}

type LanguageType int32

const (
	LanguageType_LANGUAGE_TYPE_UNSPECIFIED LanguageType = 0 // Not Set
	LanguageType_LANGUAGE_TYPE_EN          LanguageType = 1 // 英文 English
	LanguageType_LANGUAGE_TYPE_SW          LanguageType = 2 // 斯瓦希里语 Swahili
	LanguageType_LANGUAGE_TYPE_AR          LanguageType = 3 // 阿语
	LanguageType_LANGUAGE_TYPE_FR          LanguageType = 4 // 法语
)

// Enum value maps for LanguageType.
var (
	LanguageType_name = map[int32]string{
		0: "LANGUAGE_TYPE_UNSPECIFIED",
		1: "LANGUAGE_TYPE_EN",
		2: "LANGUAGE_TYPE_SW",
		3: "LANGUAGE_TYPE_AR",
		4: "LANGUAGE_TYPE_FR",
	}
	LanguageType_value = map[string]int32{
		"LANGUAGE_TYPE_UNSPECIFIED": 0,
		"LANGUAGE_TYPE_EN":          1,
		"LANGUAGE_TYPE_SW":          2,
		"LANGUAGE_TYPE_AR":          3,
		"LANGUAGE_TYPE_FR":          4,
	}
)

func (x LanguageType) Enum() *LanguageType {
	p := new(LanguageType)
	*p = x
	return p
}

func (x LanguageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LanguageType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_v1_common_proto_enumTypes[4].Descriptor()
}

func (LanguageType) Type() protoreflect.EnumType {
	return &file_common_v1_common_proto_enumTypes[4]
}

func (x LanguageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LanguageType.Descriptor instead.
func (LanguageType) EnumDescriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{4}
}

type ClientType int32

const (
	ClientType_CLIENT_TYPE_UNSPECIFIED ClientType = 0
	ClientType_CLIENT_TYPE_ANDROID     ClientType = 1
	ClientType_CLIENT_TYPE_IOS         ClientType = 5
)

// Enum value maps for ClientType.
var (
	ClientType_name = map[int32]string{
		0: "CLIENT_TYPE_UNSPECIFIED",
		1: "CLIENT_TYPE_ANDROID",
		5: "CLIENT_TYPE_IOS",
	}
	ClientType_value = map[string]int32{
		"CLIENT_TYPE_UNSPECIFIED": 0,
		"CLIENT_TYPE_ANDROID":     1,
		"CLIENT_TYPE_IOS":         5,
	}
)

func (x ClientType) Enum() *ClientType {
	p := new(ClientType)
	*p = x
	return p
}

func (x ClientType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClientType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_v1_common_proto_enumTypes[5].Descriptor()
}

func (ClientType) Type() protoreflect.EnumType {
	return &file_common_v1_common_proto_enumTypes[5]
}

func (x ClientType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClientType.Descriptor instead.
func (ClientType) EnumDescriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{5}
}

type OutParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	SubCode  int32             `protobuf:"varint,2,opt,name=sub_code,json=subCode,proto3" json:"sub_code,omitempty"`
	Message  string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Reason   string            `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
	Metadata map[string]string `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *OutParam) Reset() {
	*x = OutParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OutParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OutParam) ProtoMessage() {}

func (x *OutParam) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OutParam.ProtoReflect.Descriptor instead.
func (*OutParam) Descriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{0}
}

func (x *OutParam) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OutParam) GetSubCode() int32 {
	if x != nil {
		return x.SubCode
	}
	return 0
}

func (x *OutParam) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *OutParam) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *OutParam) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	SubCode  int32             `protobuf:"varint,2,opt,name=sub_code,json=subCode,proto3" json:"sub_code,omitempty"`
	Message  string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Reason   string            `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
	Metadata map[string]string `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data     *types.Any        `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_v1_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_v1_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_common_v1_common_proto_rawDescGZIP(), []int{1}
}

func (x *CommonResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CommonResponse) GetSubCode() int32 {
	if x != nil {
		return x.SubCode
	}
	return 0
}

func (x *CommonResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CommonResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *CommonResponse) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CommonResponse) GetData() *types.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_common_v1_common_proto protoreflect.FileDescriptor

var file_common_v1_common_proto_rawDesc = []byte{
	0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x0e, 0x67, 0x6f, 0x67, 0x6f, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe7, 0x01, 0x0a, 0x08, 0x4f, 0x75, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x75, 0x62, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x3d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x4f, 0x75, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x9d, 0x02,
	0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x75, 0x62, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x43, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0xc8, 0x04,
	0x0a, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x52, 0x52,
	0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0xbf,
	0x84, 0x3d, 0x12, 0x23, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0xdf, 0xf2, 0xf9,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x2a, 0x0a, 0x1d, 0x45, 0x52, 0x52, 0x5f, 0x43,
	0x4f, 0x44, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0xde, 0xf2, 0xf9, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0x01, 0x12, 0x26, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x58, 0x43, 0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e,
	0x10, 0xdd, 0xf2, 0xf9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x25, 0x0a, 0x18, 0x45,
	0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x52,
	0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0xdc, 0xf2, 0xf9, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0x01, 0x12, 0x21, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x53,
	0x45, 0x52, 0x56, 0x45, 0x52, 0x5f, 0x42, 0x55, 0x53, 0x59, 0x10, 0xdb, 0xf2, 0xf9, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x24, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54,
	0x10, 0xda, 0xf2, 0xf9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x26, 0x0a, 0x19, 0x45,
	0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x45,
	0x58, 0x43, 0x45, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0xd9, 0xf2, 0xf9, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0x01, 0x12, 0x23, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x42, 0x52, 0x4f, 0x4b, 0x45, 0x4e, 0x10, 0xd8, 0xf2,
	0xf9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x27, 0x0a, 0x1a, 0x45, 0x52, 0x52, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x4e, 0x4f, 0x5f,
	0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x10, 0xc0, 0xe5, 0xf3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x01, 0x12, 0x28, 0x0a, 0x1b, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x4f,
	0x4e, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x4e, 0x4f, 0x5f, 0x49, 0x4e, 0x50, 0x41, 0x52, 0x41, 0x4d,
	0x10, 0xbf, 0xe5, 0xf3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x12, 0x26, 0x0a, 0x19, 0x45,
	0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0xbe, 0xe5, 0xf3, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0x01, 0x12, 0x2c, 0x0a, 0x1f, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f,
	0x43, 0x4f, 0x4e, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0xbd, 0xe5, 0xf3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x01, 0x12, 0x25, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x4f,
	0x4e, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x4e, 0x4f, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x10, 0xbc, 0xe5,
	0xf3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01, 0x2a, 0x49, 0x0a, 0x0a, 0x45, 0x72, 0x72, 0x53,
	0x75, 0x62, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x5f, 0x53, 0x55,
	0x42, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x5f, 0x53, 0x55, 0x42, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43,
	0x45, 0x10, 0x65, 0x2a, 0x44, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1b, 0x0a, 0x17, 0x41, 0x50, 0x50, 0x5f, 0x49, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a,
	0x16, 0x41, 0x50, 0x50, 0x5f, 0x49, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x49, 0x47,
	0x48, 0x54, 0x48, 0x4f, 0x55, 0x53, 0x45, 0x10, 0x64, 0x2a, 0x85, 0x01, 0x0a, 0x0a, 0x41, 0x72,
	0x65, 0x61, 0x49, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x41, 0x52, 0x45, 0x41,
	0x5f, 0x49, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x52, 0x45, 0x41, 0x5f, 0x49,
	0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x1e,
	0x12, 0x14, 0x0a, 0x10, 0x41, 0x52, 0x45, 0x41, 0x5f, 0x49, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x49, 0x4f, 0x53, 0x10, 0x28, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x52, 0x45, 0x41, 0x5f, 0x49,
	0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x45, 0x42, 0x10, 0x63, 0x12, 0x13, 0x0a, 0x0f,
	0x41, 0x52, 0x45, 0x41, 0x5f, 0x49, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x35, 0x10,
	0x62, 0x2a, 0x85, 0x01, 0x0a, 0x0c, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x45, 0x4e, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x41, 0x4e, 0x47, 0x55,
	0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x57, 0x10, 0x02, 0x12, 0x14, 0x0a,
	0x10, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41,
	0x52, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x41, 0x4e, 0x47, 0x55, 0x41, 0x47, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x52, 0x10, 0x04, 0x2a, 0x57, 0x0a, 0x0a, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4c, 0x49, 0x45, 0x4e,
	0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x41, 0x4e, 0x44, 0x52, 0x4f, 0x49, 0x44, 0x10, 0x01, 0x12, 0x13, 0x0a,
	0x0f, 0x43, 0x4c, 0x49, 0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4f, 0x53,
	0x10, 0x05, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x7a, 0x64, 0x62, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x2f, 0x75, 0x74, 0x69, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_v1_common_proto_rawDescOnce sync.Once
	file_common_v1_common_proto_rawDescData = file_common_v1_common_proto_rawDesc
)

func file_common_v1_common_proto_rawDescGZIP() []byte {
	file_common_v1_common_proto_rawDescOnce.Do(func() {
		file_common_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_v1_common_proto_rawDescData)
	})
	return file_common_v1_common_proto_rawDescData
}

var file_common_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_common_v1_common_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_common_v1_common_proto_goTypes = []interface{}{
	(ErrCode)(0),           // 0: common.v1.ErrCode
	(ErrSubCode)(0),        // 1: common.v1.ErrSubCode
	(AppIdType)(0),         // 2: common.v1.AppIdType
	(AreaIdType)(0),        // 3: common.v1.AreaIdType
	(LanguageType)(0),      // 4: common.v1.LanguageType
	(ClientType)(0),        // 5: common.v1.ClientType
	(*OutParam)(nil),       // 6: common.v1.OutParam
	(*CommonResponse)(nil), // 7: common.v1.CommonResponse
	nil,                    // 8: common.v1.OutParam.MetadataEntry
	nil,                    // 9: common.v1.CommonResponse.MetadataEntry
	(*types.Any)(nil),      // 10: google.protobuf.Any
}
var file_common_v1_common_proto_depIdxs = []int32{
	8,  // 0: common.v1.OutParam.metadata:type_name -> common.v1.OutParam.MetadataEntry
	9,  // 1: common.v1.CommonResponse.metadata:type_name -> common.v1.CommonResponse.MetadataEntry
	10, // 2: common.v1.CommonResponse.data:type_name -> google.protobuf.Any
	3,  // [3:3] is the sub-list for method output_type
	3,  // [3:3] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_common_v1_common_proto_init() }
func file_common_v1_common_proto_init() {
	if File_common_v1_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_v1_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OutParam); i {
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
		file_common_v1_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
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
			RawDescriptor: file_common_v1_common_proto_rawDesc,
			NumEnums:      6,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_v1_common_proto_goTypes,
		DependencyIndexes: file_common_v1_common_proto_depIdxs,
		EnumInfos:         file_common_v1_common_proto_enumTypes,
		MessageInfos:      file_common_v1_common_proto_msgTypes,
	}.Build()
	File_common_v1_common_proto = out.File
	file_common_v1_common_proto_rawDesc = nil
	file_common_v1_common_proto_goTypes = nil
	file_common_v1_common_proto_depIdxs = nil
}
