//版本号

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: pagination/pagination.proto

//指定包名

package pagination

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
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

type SortOrder int32

const (
	SortOrder_SORT_ORDER_UNSPECIFIED SortOrder = 0
	SortOrder_SORT_ORDER_DESC        SortOrder = 1
	SortOrder_SORT_ORDER_ASC         SortOrder = 2
)

// Enum value maps for SortOrder.
var (
	SortOrder_name = map[int32]string{
		0: "SORT_ORDER_UNSPECIFIED",
		1: "SORT_ORDER_DESC",
		2: "SORT_ORDER_ASC",
	}
	SortOrder_value = map[string]int32{
		"SORT_ORDER_UNSPECIFIED": 0,
		"SORT_ORDER_DESC":        1,
		"SORT_ORDER_ASC":         2,
	}
)

func (x SortOrder) Enum() *SortOrder {
	p := new(SortOrder)
	*p = x
	return p
}

func (x SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_pagination_pagination_proto_enumTypes[0].Descriptor()
}

func (SortOrder) Type() protoreflect.EnumType {
	return &file_pagination_pagination_proto_enumTypes[0]
}

func (x SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortOrder.Descriptor instead.
func (SortOrder) EnumDescriptor() ([]byte, []int) {
	return file_pagination_pagination_proto_rawDescGZIP(), []int{0}
}

type PaginationReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32                `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int32                `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sort     map[string]SortOrder `protobuf:"bytes,3,rep,name=sort,proto3" json:"sort,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3,enum=pagination.SortOrder"`
}

func (x *PaginationReq) Reset() {
	*x = PaginationReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pagination_pagination_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationReq) ProtoMessage() {}

func (x *PaginationReq) ProtoReflect() protoreflect.Message {
	mi := &file_pagination_pagination_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationReq.ProtoReflect.Descriptor instead.
func (*PaginationReq) Descriptor() ([]byte, []int) {
	return file_pagination_pagination_proto_rawDescGZIP(), []int{0}
}

func (x *PaginationReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationReq) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PaginationReq) GetSort() map[string]SortOrder {
	if x != nil {
		return x.Sort
	}
	return nil
}

type TimeRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start int64 `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End   int64 `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *TimeRange) Reset() {
	*x = TimeRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pagination_pagination_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeRange) ProtoMessage() {}

func (x *TimeRange) ProtoReflect() protoreflect.Message {
	mi := &file_pagination_pagination_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeRange.ProtoReflect.Descriptor instead.
func (*TimeRange) Descriptor() ([]byte, []int) {
	return file_pagination_pagination_proto_rawDescGZIP(), []int{1}
}

func (x *TimeRange) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *TimeRange) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

var File_pagination_pagination_proto protoreflect.FileDescriptor

var file_pagination_pagination_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0f, 0x67, 0x6f, 0x67, 0x6f, 0x2f,
	0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x0d, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x27, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x1a, 0x05, 0x18, 0x88, 0x27, 0x20,
	0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3d, 0x0a, 0x04, 0x73,
	0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x04,
	0xc8, 0xde, 0x1f, 0x00, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x1a, 0x4e, 0x0a, 0x09, 0x53, 0x6f,
	0x72, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x45, 0x0a, 0x09, 0x54, 0x69,
	0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x03, 0x65, 0x6e,
	0x64, 0x2a, 0x50, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x16, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x4f,
	0x52, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x44, 0x45, 0x53, 0x43, 0x10, 0x01, 0x12,
	0x12, 0x0a, 0x0e, 0x53, 0x4f, 0x52, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x41, 0x53,
	0x43, 0x10, 0x02, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x61, 0x7a, 0x64, 0x62, 0x61, 0x61, 0x61, 0x61, 0x61, 0x61, 0x2f, 0x75, 0x74, 0x69,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x3b, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pagination_pagination_proto_rawDescOnce sync.Once
	file_pagination_pagination_proto_rawDescData = file_pagination_pagination_proto_rawDesc
)

func file_pagination_pagination_proto_rawDescGZIP() []byte {
	file_pagination_pagination_proto_rawDescOnce.Do(func() {
		file_pagination_pagination_proto_rawDescData = protoimpl.X.CompressGZIP(file_pagination_pagination_proto_rawDescData)
	})
	return file_pagination_pagination_proto_rawDescData
}

var file_pagination_pagination_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pagination_pagination_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pagination_pagination_proto_goTypes = []interface{}{
	(SortOrder)(0),        // 0: pagination.SortOrder
	(*PaginationReq)(nil), // 1: pagination.PaginationReq
	(*TimeRange)(nil),     // 2: pagination.TimeRange
	nil,                   // 3: pagination.PaginationReq.SortEntry
}
var file_pagination_pagination_proto_depIdxs = []int32{
	3, // 0: pagination.PaginationReq.sort:type_name -> pagination.PaginationReq.SortEntry
	0, // 1: pagination.PaginationReq.SortEntry.value:type_name -> pagination.SortOrder
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pagination_pagination_proto_init() }
func file_pagination_pagination_proto_init() {
	if File_pagination_pagination_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pagination_pagination_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationReq); i {
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
		file_pagination_pagination_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeRange); i {
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
			RawDescriptor: file_pagination_pagination_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pagination_pagination_proto_goTypes,
		DependencyIndexes: file_pagination_pagination_proto_depIdxs,
		EnumInfos:         file_pagination_pagination_proto_enumTypes,
		MessageInfos:      file_pagination_pagination_proto_msgTypes,
	}.Build()
	File_pagination_pagination_proto = out.File
	file_pagination_pagination_proto_rawDesc = nil
	file_pagination_pagination_proto_goTypes = nil
	file_pagination_pagination_proto_depIdxs = nil
}
