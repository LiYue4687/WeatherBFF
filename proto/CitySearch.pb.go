// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: CitySearch.proto

package proto

import (
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

type CitySearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchValue string `protobuf:"bytes,1,opt,name=searchValue,proto3" json:"searchValue,omitempty"`
}

func (x *CitySearchRequest) Reset() {
	*x = CitySearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CitySearch_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CitySearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CitySearchRequest) ProtoMessage() {}

func (x *CitySearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_CitySearch_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CitySearchRequest.ProtoReflect.Descriptor instead.
func (*CitySearchRequest) Descriptor() ([]byte, []int) {
	return file_CitySearch_proto_rawDescGZIP(), []int{0}
}

func (x *CitySearchRequest) GetSearchValue() string {
	if x != nil {
		return x.SearchValue
	}
	return ""
}

type CitySearchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*CityItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CitySearchResponse) Reset() {
	*x = CitySearchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CitySearch_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CitySearchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CitySearchResponse) ProtoMessage() {}

func (x *CitySearchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_CitySearch_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CitySearchResponse.ProtoReflect.Descriptor instead.
func (*CitySearchResponse) Descriptor() ([]byte, []int) {
	return file_CitySearch_proto_rawDescGZIP(), []int{1}
}

func (x *CitySearchResponse) GetItems() []*CityItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type CityItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProvinceName string `protobuf:"bytes,1,opt,name=provinceName,proto3" json:"provinceName,omitempty"`
	CityName     string `protobuf:"bytes,2,opt,name=cityName,proto3" json:"cityName,omitempty"`
	CountyName   string `protobuf:"bytes,3,opt,name=countyName,proto3" json:"countyName,omitempty"`
	CityCode     string `protobuf:"bytes,4,opt,name=cityCode,proto3" json:"cityCode,omitempty"`
	AdCode       string `protobuf:"bytes,5,opt,name=adCode,proto3" json:"adCode,omitempty"`
}

func (x *CityItem) Reset() {
	*x = CityItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CitySearch_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityItem) ProtoMessage() {}

func (x *CityItem) ProtoReflect() protoreflect.Message {
	mi := &file_CitySearch_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityItem.ProtoReflect.Descriptor instead.
func (*CityItem) Descriptor() ([]byte, []int) {
	return file_CitySearch_proto_rawDescGZIP(), []int{2}
}

func (x *CityItem) GetProvinceName() string {
	if x != nil {
		return x.ProvinceName
	}
	return ""
}

func (x *CityItem) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

func (x *CityItem) GetCountyName() string {
	if x != nil {
		return x.CountyName
	}
	return ""
}

func (x *CityItem) GetCityCode() string {
	if x != nil {
		return x.CityCode
	}
	return ""
}

func (x *CityItem) GetAdCode() string {
	if x != nil {
		return x.AdCode
	}
	return ""
}

var File_CitySearch_proto protoreflect.FileDescriptor

var file_CitySearch_proto_rawDesc = []byte{
	0x0a, 0x10, 0x43, 0x69, 0x74, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x11, 0x43, 0x69, 0x74,
	0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20,
	0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x3b, 0x0a, 0x12, 0x43, 0x69, 0x74, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x69,
	0x74, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x9e, 0x01,
	0x0a, 0x08, 0x43, 0x69, 0x74, 0x79, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x69,
	0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69,
	0x74, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x64, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x64, 0x43, 0x6f, 0x64, 0x65, 0x32, 0x54,
	0x0a, 0x11, 0x43, 0x69, 0x74, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x69, 0x74, 0x79, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x1a, 0x5a, 0x18, 0x67, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CitySearch_proto_rawDescOnce sync.Once
	file_CitySearch_proto_rawDescData = file_CitySearch_proto_rawDesc
)

func file_CitySearch_proto_rawDescGZIP() []byte {
	file_CitySearch_proto_rawDescOnce.Do(func() {
		file_CitySearch_proto_rawDescData = protoimpl.X.CompressGZIP(file_CitySearch_proto_rawDescData)
	})
	return file_CitySearch_proto_rawDescData
}

var file_CitySearch_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_CitySearch_proto_goTypes = []interface{}{
	(*CitySearchRequest)(nil),  // 0: proto.CitySearchRequest
	(*CitySearchResponse)(nil), // 1: proto.CitySearchResponse
	(*CityItem)(nil),           // 2: proto.CityItem
}
var file_CitySearch_proto_depIdxs = []int32{
	2, // 0: proto.CitySearchResponse.items:type_name -> proto.CityItem
	0, // 1: proto.CitySearchService.Search:input_type -> proto.CitySearchRequest
	1, // 2: proto.CitySearchService.Search:output_type -> proto.CitySearchResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_CitySearch_proto_init() }
func file_CitySearch_proto_init() {
	if File_CitySearch_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CitySearch_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CitySearchRequest); i {
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
		file_CitySearch_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CitySearchResponse); i {
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
		file_CitySearch_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityItem); i {
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
			RawDescriptor: file_CitySearch_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_CitySearch_proto_goTypes,
		DependencyIndexes: file_CitySearch_proto_depIdxs,
		MessageInfos:      file_CitySearch_proto_msgTypes,
	}.Build()
	File_CitySearch_proto = out.File
	file_CitySearch_proto_rawDesc = nil
	file_CitySearch_proto_goTypes = nil
	file_CitySearch_proto_depIdxs = nil
}