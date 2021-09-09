// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common.proto

//指定包名

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ErrCode int32

const (
	ErrCode_success              ErrCode = 0
	ErrCode_unknown_error        ErrCode = 999999
	ErrCode_server_fail          ErrCode = -100001
	ErrCode_server_param_invalid ErrCode = -100002
	ErrCode_server_exception     ErrCode = -100003
	ErrCode_server_reject        ErrCode = -100004
	ErrCode_server_busy          ErrCode = -100005
	ErrCode_client_timeout       ErrCode = -100006
	ErrCode_client_exception     ErrCode = -100007
	ErrCode_client_broken        ErrCode = -100008
	ErrCode_no_device_error      ErrCode = -200000
	ErrCode_no_in_param_error    ErrCode = -200001
	ErrCode_not_found            ErrCode = -200002
)

var ErrCode_name = map[int32]string{
	0:       "success",
	999999:  "unknown_error",
	-100001: "server_fail",
	-100002: "server_param_invalid",
	-100003: "server_exception",
	-100004: "server_reject",
	-100005: "server_busy",
	-100006: "client_timeout",
	-100007: "client_exception",
	-100008: "client_broken",
	-200000: "no_device_error",
	-200001: "no_in_param_error",
	-200002: "not_found",
}

var ErrCode_value = map[string]int32{
	"success":              0,
	"unknown_error":        999999,
	"server_fail":          -100001,
	"server_param_invalid": -100002,
	"server_exception":     -100003,
	"server_reject":        -100004,
	"server_busy":          -100005,
	"client_timeout":       -100006,
	"client_exception":     -100007,
	"client_broken":        -100008,
	"no_device_error":      -200000,
	"no_in_param_error":    -200001,
	"not_found":            -200002,
}

func (x ErrCode) String() string {
	return proto.EnumName(ErrCode_name, int32(x))
}

func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

type ErrSubCode int32

const (
	ErrSubCode_err_sub_code_unknown                 ErrSubCode = 0
	ErrSubCode_err_sub_code_book_service            ErrSubCode = 101
	ErrSubCode_err_sub_code_go_book_service         ErrSubCode = 200
	ErrSubCode_err_sub_code_agreement_service       ErrSubCode = 201
	ErrSubCode_err_sub_code_book_store_service      ErrSubCode = 202
	ErrSubCode_err_sub_code_go_device_service       ErrSubCode = 203
	ErrSubCode_err_sub_code_task_service            ErrSubCode = 204
	ErrSubCode_err_sub_code_account_service         ErrSubCode = 205
	ErrSubCode_err_sub_code_writing_contest_service ErrSubCode = 206
	ErrSubCode_err_sub_code_go_user_service         ErrSubCode = 207
	ErrSubCode_err_sub_code_balance_service         ErrSubCode = 208
	ErrSubCode_err_sub_code_track_service           ErrSubCode = 209
)

var ErrSubCode_name = map[int32]string{
	0:   "err_sub_code_unknown",
	101: "err_sub_code_book_service",
	200: "err_sub_code_go_book_service",
	201: "err_sub_code_agreement_service",
	202: "err_sub_code_book_store_service",
	203: "err_sub_code_go_device_service",
	204: "err_sub_code_task_service",
	205: "err_sub_code_account_service",
	206: "err_sub_code_writing_contest_service",
	207: "err_sub_code_go_user_service",
	208: "err_sub_code_balance_service",
	209: "err_sub_code_track_service",
}

var ErrSubCode_value = map[string]int32{
	"err_sub_code_unknown":                 0,
	"err_sub_code_book_service":            101,
	"err_sub_code_go_book_service":         200,
	"err_sub_code_agreement_service":       201,
	"err_sub_code_book_store_service":      202,
	"err_sub_code_go_device_service":       203,
	"err_sub_code_task_service":            204,
	"err_sub_code_account_service":         205,
	"err_sub_code_writing_contest_service": 206,
	"err_sub_code_go_user_service":         207,
	"err_sub_code_balance_service":         208,
	"err_sub_code_track_service":           209,
}

func (x ErrSubCode) String() string {
	return proto.EnumName(ErrSubCode_name, int32(x))
}

func (ErrSubCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

type AppIdType int32

const (
	AppIdType_APP_ID_UNKNOWN    AppIdType = 0
	AppIdType_APP_ID_LIGHTHOUSE AppIdType = 100
)

var AppIdType_name = map[int32]string{
	0:   "APP_ID_UNKNOWN",
	100: "APP_ID_LIGHTHOUSE",
}

var AppIdType_value = map[string]int32{
	"APP_ID_UNKNOWN":    0,
	"APP_ID_LIGHTHOUSE": 100,
}

func (x AppIdType) String() string {
	return proto.EnumName(AppIdType_name, int32(x))
}

func (AppIdType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

type AreaIdType int32

const (
	AreaIdType_AREA_ID_UNKNOWN AreaIdType = 0
	AreaIdType_AREA_ID_ANDROID AreaIdType = 30
	AreaIdType_AREA_ID_IOS     AreaIdType = 40
	AreaIdType_AREA_ID_WEB     AreaIdType = 99
	AreaIdType_AREA_ID_H5      AreaIdType = 98
)

var AreaIdType_name = map[int32]string{
	0:  "AREA_ID_UNKNOWN",
	30: "AREA_ID_ANDROID",
	40: "AREA_ID_IOS",
	99: "AREA_ID_WEB",
	98: "AREA_ID_H5",
}

var AreaIdType_value = map[string]int32{
	"AREA_ID_UNKNOWN": 0,
	"AREA_ID_ANDROID": 30,
	"AREA_ID_IOS":     40,
	"AREA_ID_WEB":     99,
	"AREA_ID_H5":      98,
}

func (x AreaIdType) String() string {
	return proto.EnumName(AreaIdType_name, int32(x))
}

func (AreaIdType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

type LanguageType int32

const (
	LanguageType_LAN_ZH LanguageType = 0
	LanguageType_LAN_EN LanguageType = 1
	LanguageType_LAN_ES LanguageType = 2
	LanguageType_LAN_IN LanguageType = 3
	LanguageType_LAN_TL LanguageType = 4
	LanguageType_LAN_MS LanguageType = 5
	LanguageType_LAN_KO LanguageType = 6
)

var LanguageType_name = map[int32]string{
	0: "LAN_ZH",
	1: "LAN_EN",
	2: "LAN_ES",
	3: "LAN_IN",
	4: "LAN_TL",
	5: "LAN_MS",
	6: "LAN_KO",
}

var LanguageType_value = map[string]int32{
	"LAN_ZH": 0,
	"LAN_EN": 1,
	"LAN_ES": 2,
	"LAN_IN": 3,
	"LAN_TL": 4,
	"LAN_MS": 5,
	"LAN_KO": 6,
}

func (x LanguageType) String() string {
	return proto.EnumName(LanguageType_name, int32(x))
}

func (LanguageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

type ClientType int32

const (
	ClientType_CLIENT_TYPE_UNKNOWN ClientType = 0
	ClientType_CLIENT_TYPE_ANDROID ClientType = 1
	ClientType_CLIENT_TYPE_IOS     ClientType = 5
)

var ClientType_name = map[int32]string{
	0: "CLIENT_TYPE_UNKNOWN",
	1: "CLIENT_TYPE_ANDROID",
	5: "CLIENT_TYPE_IOS",
}

var ClientType_value = map[string]int32{
	"CLIENT_TYPE_UNKNOWN": 0,
	"CLIENT_TYPE_ANDROID": 1,
	"CLIENT_TYPE_IOS":     5,
}

func (x ClientType) String() string {
	return proto.EnumName(ClientType_name, int32(x))
}

func (ClientType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{5}
}

type ChannelType int32

const (
	ChannelType_ChannelType_Unknown ChannelType = 0
	ChannelType_sdk_phoenix         ChannelType = 2
)

var ChannelType_name = map[int32]string{
	0: "ChannelType_Unknown",
	2: "sdk_phoenix",
}

var ChannelType_value = map[string]int32{
	"ChannelType_Unknown": 0,
	"sdk_phoenix":         2,
}

func (x ChannelType) String() string {
	return proto.EnumName(ChannelType_name, int32(x))
}

func (ChannelType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{6}
}

type SDKType int32

const (
	SDKType_SDK_UNKNOWN SDKType = 0
	SDKType_SDK_FICOOL  SDKType = 1
	SDKType_SDK_PHOENIX SDKType = 2
)

var SDKType_name = map[int32]string{
	0: "SDK_UNKNOWN",
	1: "SDK_FICOOL",
	2: "SDK_PHOENIX",
}

var SDKType_value = map[string]int32{
	"SDK_UNKNOWN": 0,
	"SDK_FICOOL":  1,
	"SDK_PHOENIX": 2,
}

func (x SDKType) String() string {
	return proto.EnumName(SDKType_name, int32(x))
}

func (SDKType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{7}
}

type OutParam struct {
	Code                 int32             `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	SubCode              int32             `protobuf:"varint,2,opt,name=sub_code,json=subCode,proto3" json:"sub_code,omitempty"`
	Message              string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Reason               string            `protobuf:"bytes,4,opt,name=reason,proto3" json:"reason,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,5,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OutParam) Reset()         { *m = OutParam{} }
func (m *OutParam) String() string { return proto.CompactTextString(m) }
func (*OutParam) ProtoMessage()    {}
func (*OutParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}
func (m *OutParam) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OutParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OutParam.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OutParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutParam.Merge(m, src)
}
func (m *OutParam) XXX_Size() int {
	return m.Size()
}
func (m *OutParam) XXX_DiscardUnknown() {
	xxx_messageInfo_OutParam.DiscardUnknown(m)
}

var xxx_messageInfo_OutParam proto.InternalMessageInfo

func (m *OutParam) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *OutParam) GetSubCode() int32 {
	if m != nil {
		return m.SubCode
	}
	return 0
}

func (m *OutParam) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *OutParam) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *OutParam) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterEnum("common.ErrCode", ErrCode_name, ErrCode_value)
	proto.RegisterEnum("common.ErrSubCode", ErrSubCode_name, ErrSubCode_value)
	proto.RegisterEnum("common.AppIdType", AppIdType_name, AppIdType_value)
	proto.RegisterEnum("common.AreaIdType", AreaIdType_name, AreaIdType_value)
	proto.RegisterEnum("common.LanguageType", LanguageType_name, LanguageType_value)
	proto.RegisterEnum("common.ClientType", ClientType_name, ClientType_value)
	proto.RegisterEnum("common.ChannelType", ChannelType_name, ChannelType_value)
	proto.RegisterEnum("common.SDKType", SDKType_name, SDKType_value)
	proto.RegisterType((*OutParam)(nil), "common.OutParam")
	proto.RegisterMapType((map[string]string)(nil), "common.OutParam.MetadataEntry")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 863 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x55, 0xcd, 0x73, 0xdb, 0xc4,
	0x1b, 0xf6, 0xda, 0xb1, 0x1d, 0xbf, 0x6e, 0xe2, 0xed, 0x26, 0xbf, 0xfe, 0xdc, 0xd0, 0xba, 0x29,
	0xf4, 0x10, 0x7c, 0x48, 0x66, 0x60, 0xf8, 0x98, 0x66, 0x98, 0xc1, 0xb5, 0x05, 0xf6, 0xc4, 0x95,
	0x3c, 0xb6, 0x33, 0x85, 0x5e, 0x76, 0x56, 0xd2, 0xd6, 0x11, 0xb6, 0x77, 0x3d, 0x2b, 0x29, 0x6d,
	0x38, 0xf3, 0x47, 0xf0, 0x4f, 0xf0, 0x2f, 0x00, 0xc7, 0xf2, 0x5d, 0x6e, 0x7c, 0xc3, 0x04, 0x38,
	0x71, 0x83, 0x13, 0x27, 0x18, 0xc9, 0xfa, 0xb0, 0x48, 0x75, 0xf1, 0xf3, 0x3e, 0xcf, 0xb3, 0x7a,
	0x9f, 0xf7, 0x1d, 0x49, 0x86, 0x4b, 0x96, 0x9c, 0xcf, 0xa5, 0xd8, 0x5f, 0x28, 0xe9, 0x49, 0x52,
	0x5a, 0x56, 0xcf, 0xfe, 0x86, 0x60, 0xdd, 0xf0, 0xbd, 0x01, 0x53, 0x6c, 0x4e, 0x08, 0xac, 0x59,
	0xd2, 0xe6, 0x75, 0xb4, 0x8b, 0xf6, 0x8a, 0xc3, 0x10, 0x93, 0xab, 0xb0, 0xee, 0xfa, 0x26, 0x0d,
	0xf9, 0x7c, 0xc8, 0x97, 0x5d, 0xdf, 0x6c, 0x07, 0x52, 0x1d, 0xca, 0x73, 0xee, 0xba, 0x6c, 0xc2,
	0xeb, 0x85, 0x5d, 0xb4, 0x57, 0x19, 0xc6, 0x25, 0xb9, 0x02, 0x25, 0xc5, 0x99, 0x2b, 0x45, 0x7d,
	0x2d, 0x14, 0xa2, 0x8a, 0xdc, 0x86, 0xf5, 0x39, 0xf7, 0x98, 0xcd, 0x3c, 0x56, 0x2f, 0xee, 0x16,
	0xf6, 0xaa, 0x2f, 0x34, 0xf6, 0xa3, 0x58, 0x71, 0x88, 0xfd, 0xbb, 0x91, 0x41, 0x13, 0x9e, 0x3a,
	0x1b, 0x26, 0xfe, 0x9d, 0x43, 0xd8, 0xc8, 0x48, 0x04, 0x43, 0x61, 0xca, 0xcf, 0xc2, 0xb0, 0x95,
	0x61, 0x00, 0xc9, 0x36, 0x14, 0x4f, 0xd9, 0xcc, 0x5f, 0x06, 0xad, 0x0c, 0x97, 0xc5, 0xed, 0xfc,
	0xab, 0xa8, 0xf9, 0x47, 0x1e, 0xca, 0x9a, 0x52, 0x61, 0xec, 0x2a, 0x94, 0x5d, 0xdf, 0xb2, 0xb8,
	0xeb, 0xe2, 0x1c, 0xd9, 0x82, 0x0d, 0x5f, 0x4c, 0x85, 0x7c, 0x28, 0x28, 0x57, 0x4a, 0x2a, 0xfc,
	0xd1, 0x7b, 0xaf, 0x91, 0x3a, 0x54, 0x5d, 0xae, 0x4e, 0xb9, 0xa2, 0x0f, 0x98, 0x33, 0xc3, 0x3f,
	0xff, 0xf9, 0xf7, 0x3f, 0xe1, 0x85, 0xc8, 0x4d, 0xd8, 0x8e, 0x94, 0x45, 0x10, 0x96, 0x3a, 0xe2,
	0x94, 0xcd, 0x1c, 0x1b, 0xff, 0x94, 0x5a, 0xae, 0x03, 0x8e, 0x2c, 0xfc, 0x91, 0xc5, 0x17, 0x9e,
	0x23, 0x05, 0xfe, 0x31, 0x95, 0x77, 0x60, 0x23, 0x92, 0x15, 0x7f, 0x87, 0x5b, 0x1e, 0xfe, 0x21,
	0xd5, 0xd2, 0xbe, 0xa6, 0xef, 0x9e, 0xe1, 0xef, 0x53, 0xe5, 0x19, 0xd8, 0xb4, 0x66, 0x0e, 0x17,
	0x1e, 0xf5, 0x9c, 0x39, 0x97, 0xbe, 0x87, 0xbf, 0xcb, 0x74, 0x8c, 0xc4, 0xb4, 0xe3, 0xb7, 0x99,
	0x8e, 0x91, 0x6c, 0x2a, 0x39, 0xe5, 0x02, 0x7f, 0x93, 0x6a, 0xd7, 0xa0, 0x26, 0x24, 0xb5, 0xf9,
	0xa9, 0x63, 0xf1, 0x68, 0x01, 0x1f, 0xff, 0xfe, 0x57, 0xac, 0x36, 0xe0, 0xb2, 0x90, 0xd4, 0x11,
	0xd1, 0xb0, 0xd1, 0x82, 0x52, 0xfd, 0x0a, 0x54, 0x84, 0xf4, 0xe8, 0x03, 0xe9, 0x0b, 0x1b, 0x7f,
	0x98, 0xf0, 0xcd, 0x0f, 0x0a, 0x00, 0x9a, 0x52, 0xa3, 0xe4, 0x39, 0xd9, 0xe6, 0x4a, 0xd1, 0xf8,
	0x31, 0xa2, 0xd1, 0xc2, 0x71, 0x8e, 0x5c, 0x87, 0xab, 0x19, 0xc5, 0x94, 0x72, 0x4a, 0x83, 0x15,
	0x38, 0x16, 0xc7, 0x9c, 0xdc, 0x84, 0x6b, 0x19, 0x79, 0x22, 0xb3, 0x8e, 0xc7, 0x88, 0x3c, 0x07,
	0x8d, 0x8c, 0x85, 0x4d, 0x14, 0xe7, 0xf3, 0x60, 0xd8, 0xd8, 0xf4, 0x09, 0x22, 0xb7, 0xe0, 0xc6,
	0x53, 0xda, 0x78, 0x52, 0xf1, 0xc4, 0xf5, 0xe9, 0xc5, 0x5b, 0x4d, 0x92, 0xc5, 0xc4, 0xa6, 0xcf,
	0x82, 0x95, 0x64, 0x13, 0x7b, 0xcc, 0x4d, 0xf3, 0x7c, 0x8e, 0x2e, 0x44, 0x66, 0x96, 0x25, 0xfd,
	0x95, 0x34, 0x5f, 0x20, 0xf2, 0x3c, 0xdc, 0xca, 0x58, 0x1e, 0x2a, 0xc7, 0x73, 0xc4, 0x84, 0x5a,
	0x52, 0x78, 0xdc, 0x4d, 0xad, 0x5f, 0xa2, 0xa7, 0x2d, 0xc0, 0x77, 0xb9, 0x4a, 0x2c, 0x5f, 0x5d,
	0xb4, 0x98, 0x6c, 0xc6, 0xc4, 0x4a, 0xe6, 0x27, 0x88, 0xdc, 0x80, 0x9d, 0x6c, 0x66, 0xc5, 0xac,
	0x34, 0xf4, 0xd7, 0xa8, 0xf9, 0x32, 0x54, 0x5a, 0x8b, 0x45, 0xcf, 0x1e, 0x9f, 0x2d, 0x38, 0x21,
	0xb0, 0xd9, 0x1a, 0x0c, 0x68, 0xaf, 0x43, 0x8f, 0xf5, 0x23, 0xdd, 0xb8, 0xa7, 0xe3, 0x1c, 0xf9,
	0x1f, 0x5c, 0x8e, 0xb8, 0x7e, 0xef, 0xcd, 0xee, 0xb8, 0x6b, 0x1c, 0x8f, 0x34, 0x6c, 0x37, 0x4f,
	0x00, 0x5a, 0x8a, 0xb3, 0xe8, 0xe0, 0x16, 0xd4, 0x5a, 0x43, 0xad, 0x95, 0x3d, 0xb9, 0x42, 0xb6,
	0xf4, 0xce, 0xd0, 0xe8, 0x75, 0x70, 0x83, 0xd4, 0xa0, 0x1a, 0x93, 0x3d, 0x63, 0x84, 0xf7, 0x56,
	0x89, 0x7b, 0xda, 0x1d, 0x6c, 0x91, 0x4d, 0x80, 0x98, 0xe8, 0xbe, 0x84, 0xcd, 0xa6, 0x09, 0x97,
	0xfa, 0x4c, 0x4c, 0x7c, 0x36, 0xe1, 0x61, 0x2f, 0x80, 0x52, 0xbf, 0xa5, 0xd3, 0xfb, 0x5d, 0x9c,
	0x8b, 0xb1, 0xa6, 0x63, 0x94, 0xe0, 0x11, 0xce, 0xc7, 0xb8, 0xa7, 0xe3, 0x42, 0x8c, 0xc7, 0x7d,
	0xbc, 0x16, 0xe3, 0xbb, 0x23, 0x5c, 0x8c, 0xf1, 0x91, 0x81, 0x4b, 0xcd, 0x11, 0x40, 0x3b, 0x7c,
	0x4f, 0xc2, 0x0e, 0xff, 0x87, 0xad, 0x76, 0xbf, 0xa7, 0xe9, 0x63, 0x3a, 0x7e, 0x7b, 0xa0, 0xad,
	0x4c, 0xf4, 0x1f, 0x21, 0x9e, 0x0a, 0x05, 0xa3, 0xae, 0x0a, 0xc1, 0x64, 0xc5, 0xe6, 0x2b, 0x50,
	0x6d, 0x9f, 0x30, 0x21, 0xf8, 0x2c, 0xb9, 0x6b, 0x5a, 0xd2, 0xe3, 0xe4, 0x4d, 0xa8, 0x41, 0xd5,
	0xb5, 0xa7, 0x74, 0x71, 0x22, 0xb9, 0x70, 0x1e, 0xe1, 0x7c, 0xf3, 0x10, 0xca, 0xa3, 0xce, 0x51,
	0x78, 0xa8, 0x06, 0xd5, 0x51, 0xe7, 0x68, 0x25, 0xc2, 0x26, 0x40, 0x40, 0xbc, 0xd1, 0x6b, 0x1b,
	0x46, 0x1f, 0xa3, 0xd8, 0x30, 0xe8, 0x1a, 0x9a, 0xde, 0x7b, 0x0b, 0xe7, 0xef, 0xbc, 0xfe, 0xf8,
	0xbc, 0x81, 0x9e, 0x9c, 0x37, 0xd0, 0x2f, 0xe7, 0x0d, 0xf4, 0xfe, 0xaf, 0x8d, 0xdc, 0xfd, 0xfd,
	0x89, 0xe3, 0x9d, 0xf8, 0x66, 0xf0, 0xb5, 0x3d, 0x60, 0xef, 0xda, 0x26, 0x0b, 0xaf, 0x03, 0xdf,
	0x73, 0x66, 0x07, 0xe1, 0x3f, 0xc2, 0xc1, 0xf2, 0x3b, 0x7c, 0xb8, 0xfc, 0x31, 0x4b, 0x21, 0xf9,
	0xe2, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xce, 0xef, 0x66, 0x12, 0x36, 0x06, 0x00, 0x00,
}

func (m *OutParam) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OutParam) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OutParam) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Metadata) > 0 {
		for k := range m.Metadata {
			v := m.Metadata[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintCommon(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintCommon(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintCommon(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Reason) > 0 {
		i -= len(m.Reason)
		copy(dAtA[i:], m.Reason)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Reason)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SubCode != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.SubCode))
		i--
		dAtA[i] = 0x10
	}
	if m.Code != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *OutParam) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovCommon(uint64(m.Code))
	}
	if m.SubCode != 0 {
		n += 1 + sovCommon(uint64(m.SubCode))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if len(m.Metadata) > 0 {
		for k, v := range m.Metadata {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovCommon(uint64(len(k))) + 1 + len(v) + sovCommon(uint64(len(v)))
			n += mapEntrySize + 1 + sovCommon(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *OutParam) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OutParam: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OutParam: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubCode", wireType)
			}
			m.SubCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubCode |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCommon
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCommon
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthCommon
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthCommon
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCommon
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthCommon
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthCommon
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipCommon(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthCommon
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Metadata[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCommon = fmt.Errorf("proto: unexpected end of group")
)
