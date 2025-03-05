// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: parser.proto

package parser

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetParserElementsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SiteUrl       string                 `protobuf:"bytes,1,opt,name=site_url,json=siteUrl,proto3" json:"site_url,omitempty"`
	Selection     string                 `protobuf:"bytes,2,opt,name=selection,proto3" json:"selection,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetParserElementsRequest) Reset() {
	*x = GetParserElementsRequest{}
	mi := &file_parser_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetParserElementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetParserElementsRequest) ProtoMessage() {}

func (x *GetParserElementsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parser_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetParserElementsRequest.ProtoReflect.Descriptor instead.
func (*GetParserElementsRequest) Descriptor() ([]byte, []int) {
	return file_parser_proto_rawDescGZIP(), []int{0}
}

func (x *GetParserElementsRequest) GetSiteUrl() string {
	if x != nil {
		return x.SiteUrl
	}
	return ""
}

func (x *GetParserElementsRequest) GetSelection() string {
	if x != nil {
		return x.Selection
	}
	return ""
}

type SelectionElement struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	HrefImage     string                 `protobuf:"bytes,2,opt,name=href_image,json=hrefImage,proto3" json:"href_image,omitempty"`
	AgeLimit      int32                  `protobuf:"varint,3,opt,name=age_limit,json=ageLimit,proto3" json:"age_limit,omitempty"`
	Price         int32                  `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	TimeSeans     string                 `protobuf:"bytes,5,opt,name=time_seans,json=timeSeans,proto3" json:"time_seans,omitempty"`
	HrefSeans     string                 `protobuf:"bytes,6,opt,name=href_seans,json=hrefSeans,proto3" json:"href_seans,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SelectionElement) Reset() {
	*x = SelectionElement{}
	mi := &file_parser_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SelectionElement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectionElement) ProtoMessage() {}

func (x *SelectionElement) ProtoReflect() protoreflect.Message {
	mi := &file_parser_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectionElement.ProtoReflect.Descriptor instead.
func (*SelectionElement) Descriptor() ([]byte, []int) {
	return file_parser_proto_rawDescGZIP(), []int{1}
}

func (x *SelectionElement) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SelectionElement) GetHrefImage() string {
	if x != nil {
		return x.HrefImage
	}
	return ""
}

func (x *SelectionElement) GetAgeLimit() int32 {
	if x != nil {
		return x.AgeLimit
	}
	return 0
}

func (x *SelectionElement) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *SelectionElement) GetTimeSeans() string {
	if x != nil {
		return x.TimeSeans
	}
	return ""
}

func (x *SelectionElement) GetHrefSeans() string {
	if x != nil {
		return x.HrefSeans
	}
	return ""
}

type GetParserElementsResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	Lenght           int32                  `protobuf:"varint,1,opt,name=lenght,proto3" json:"lenght,omitempty"`
	Selectionelement []*SelectionElement    `protobuf:"bytes,2,rep,name=selectionelement,proto3" json:"selectionelement,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *GetParserElementsResponse) Reset() {
	*x = GetParserElementsResponse{}
	mi := &file_parser_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetParserElementsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetParserElementsResponse) ProtoMessage() {}

func (x *GetParserElementsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parser_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetParserElementsResponse.ProtoReflect.Descriptor instead.
func (*GetParserElementsResponse) Descriptor() ([]byte, []int) {
	return file_parser_proto_rawDescGZIP(), []int{2}
}

func (x *GetParserElementsResponse) GetLenght() int32 {
	if x != nil {
		return x.Lenght
	}
	return 0
}

func (x *GetParserElementsResponse) GetSelectionelement() []*SelectionElement {
	if x != nil {
		return x.Selectionelement
	}
	return nil
}

var File_parser_proto protoreflect.FileDescriptor

var file_parser_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03,
	0x61, 0x70, 0x69, 0x22, 0x53, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x73, 0x69, 0x74, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x69, 0x74, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xb8, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x72, 0x65, 0x66, 0x5f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x72, 0x65, 0x66, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x65,
	0x61, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x53,
	0x65, 0x61, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x72, 0x65, 0x66, 0x5f, 0x73, 0x65, 0x61,
	0x6e, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x68, 0x72, 0x65, 0x66, 0x53, 0x65,
	0x61, 0x6e, 0x73, 0x22, 0x76, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72,
	0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x68, 0x74, 0x12, 0x41, 0x0a, 0x10, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x10, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0x5c, 0x0a, 0x06, 0x50,
	0x61, 0x72, 0x73, 0x65, 0x72, 0x12, 0x52, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x72, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x73, 0x65, 0x72, 0x45, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x14, 0x5a, 0x12, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_parser_proto_rawDescOnce sync.Once
	file_parser_proto_rawDescData []byte
)

func file_parser_proto_rawDescGZIP() []byte {
	file_parser_proto_rawDescOnce.Do(func() {
		file_parser_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_parser_proto_rawDesc), len(file_parser_proto_rawDesc)))
	})
	return file_parser_proto_rawDescData
}

var file_parser_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_parser_proto_goTypes = []any{
	(*GetParserElementsRequest)(nil),  // 0: api.GetParserElementsRequest
	(*SelectionElement)(nil),          // 1: api.SelectionElement
	(*GetParserElementsResponse)(nil), // 2: api.GetParserElementsResponse
}
var file_parser_proto_depIdxs = []int32{
	1, // 0: api.GetParserElementsResponse.selectionelement:type_name -> api.SelectionElement
	0, // 1: api.Parser.GetParserElements:input_type -> api.GetParserElementsRequest
	2, // 2: api.Parser.GetParserElements:output_type -> api.GetParserElementsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_parser_proto_init() }
func file_parser_proto_init() {
	if File_parser_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_parser_proto_rawDesc), len(file_parser_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_parser_proto_goTypes,
		DependencyIndexes: file_parser_proto_depIdxs,
		MessageInfos:      file_parser_proto_msgTypes,
	}.Build()
	File_parser_proto = out.File
	file_parser_proto_goTypes = nil
	file_parser_proto_depIdxs = nil
}
