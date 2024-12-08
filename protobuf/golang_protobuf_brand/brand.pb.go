// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: protobuf/brand.proto

package golang_protobuf_brand

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProtoBrandRepo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brands []*ProtoBrandRepo_ProtoBrand `protobuf:"bytes,1,rep,name=brands,proto3" json:"brands,omitempty"`
}

func (x *ProtoBrandRepo) Reset() {
	*x = ProtoBrandRepo{}
	mi := &file_protobuf_brand_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProtoBrandRepo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoBrandRepo) ProtoMessage() {}

func (x *ProtoBrandRepo) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_brand_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoBrandRepo.ProtoReflect.Descriptor instead.
func (*ProtoBrandRepo) Descriptor() ([]byte, []int) {
	return file_protobuf_brand_proto_rawDescGZIP(), []int{0}
}

func (x *ProtoBrandRepo) GetBrands() []*ProtoBrandRepo_ProtoBrand {
	if x != nil {
		return x.Brands
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID    *wrapperspb.Int64Value     `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Brand *ProtoBrandRepo_ProtoBrand `protobuf:"bytes,2,opt,name=Brand,proto3" json:"Brand,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	mi := &file_protobuf_brand_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_brand_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_brand_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRequest) GetID() *wrapperspb.Int64Value {
	if x != nil {
		return x.ID
	}
	return nil
}

func (x *UpdateRequest) GetBrand() *ProtoBrandRepo_ProtoBrand {
	if x != nil {
		return x.Brand
	}
	return nil
}

type ProtoBrand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Year uint32 `protobuf:"varint,3,opt,name=Year,json=year,proto3" json:"Year,omitempty"`
}

func (x *ProtoBrand) Reset() {
	*x = ProtoBrand{}
	mi := &file_protobuf_brand_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProtoBrand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoBrand) ProtoMessage() {}

func (x *ProtoBrand) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_brand_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoBrand.ProtoReflect.Descriptor instead.
func (*ProtoBrand) Descriptor() ([]byte, []int) {
	return file_protobuf_brand_proto_rawDescGZIP(), []int{2}
}

func (x *ProtoBrand) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ProtoBrand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProtoBrand) GetYear() uint32 {
	if x != nil {
		return x.Year
	}
	return 0
}

type BrandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brand *ProtoBrand `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
}

func (x *BrandResponse) Reset() {
	*x = BrandResponse{}
	mi := &file_protobuf_brand_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BrandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandResponse) ProtoMessage() {}

func (x *BrandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_brand_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrandResponse.ProtoReflect.Descriptor instead.
func (*BrandResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_brand_proto_rawDescGZIP(), []int{3}
}

func (x *BrandResponse) GetBrand() *ProtoBrand {
	if x != nil {
		return x.Brand
	}
	return nil
}

type ProtoBrandRepo_ProtoBrand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Year uint32 `protobuf:"varint,3,opt,name=Year,json=year,proto3" json:"Year,omitempty"`
}

func (x *ProtoBrandRepo_ProtoBrand) Reset() {
	*x = ProtoBrandRepo_ProtoBrand{}
	mi := &file_protobuf_brand_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProtoBrandRepo_ProtoBrand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProtoBrandRepo_ProtoBrand) ProtoMessage() {}

func (x *ProtoBrandRepo_ProtoBrand) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_brand_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProtoBrandRepo_ProtoBrand.ProtoReflect.Descriptor instead.
func (*ProtoBrandRepo_ProtoBrand) Descriptor() ([]byte, []int) {
	return file_protobuf_brand_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ProtoBrandRepo_ProtoBrand) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *ProtoBrandRepo_ProtoBrand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProtoBrandRepo_ProtoBrand) GetYear() uint32 {
	if x != nil {
		return x.Year
	}
	return 0
}

var File_protobuf_brand_proto protoreflect.FileDescriptor

var file_protobuf_brand_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x1a, 0x1b, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65,
	0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76,
	0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x01, 0x0a, 0x0e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x48, 0x0a, 0x06,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52,
	0x65, 0x70, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x06,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x1a, 0x4e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42,
	0x72, 0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x17, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a,
	0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x22, 0x84, 0x01, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x02, 0x49, 0x44, 0x12, 0x46, 0x0a, 0x05, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x05, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x22, 0x4e, 0x0a,
	0x0a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x17, 0x0a, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x22, 0x48, 0x0a,
	0x0d, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x32, 0x71, 0x0a, 0x04, 0x43, 0x72, 0x75, 0x64, 0x12,
	0x69, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x1a, 0x24, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x73, 0x42, 0x18, 0x5a, 0x16, 0x2f, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x62,
	0x72, 0x61, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_brand_proto_rawDescOnce sync.Once
	file_protobuf_brand_proto_rawDescData = file_protobuf_brand_proto_rawDesc
)

func file_protobuf_brand_proto_rawDescGZIP() []byte {
	file_protobuf_brand_proto_rawDescOnce.Do(func() {
		file_protobuf_brand_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_brand_proto_rawDescData)
	})
	return file_protobuf_brand_proto_rawDescData
}

var file_protobuf_brand_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protobuf_brand_proto_goTypes = []any{
	(*ProtoBrandRepo)(nil),            // 0: golang_protobuf_brand.ProtoBrandRepo
	(*UpdateRequest)(nil),             // 1: golang_protobuf_brand.UpdateRequest
	(*ProtoBrand)(nil),                // 2: golang_protobuf_brand.ProtoBrand
	(*BrandResponse)(nil),             // 3: golang_protobuf_brand.BrandResponse
	(*ProtoBrandRepo_ProtoBrand)(nil), // 4: golang_protobuf_brand.ProtoBrandRepo.ProtoBrand
	(*wrapperspb.Int64Value)(nil),     // 5: google.protobuf.Int64Value
}
var file_protobuf_brand_proto_depIdxs = []int32{
	4, // 0: golang_protobuf_brand.ProtoBrandRepo.brands:type_name -> golang_protobuf_brand.ProtoBrandRepo.ProtoBrand
	5, // 1: golang_protobuf_brand.UpdateRequest.ID:type_name -> google.protobuf.Int64Value
	4, // 2: golang_protobuf_brand.UpdateRequest.Brand:type_name -> golang_protobuf_brand.ProtoBrandRepo.ProtoBrand
	2, // 3: golang_protobuf_brand.BrandResponse.brand:type_name -> golang_protobuf_brand.ProtoBrand
	2, // 4: golang_protobuf_brand.Crud.Create:input_type -> golang_protobuf_brand.ProtoBrand
	3, // 5: golang_protobuf_brand.Crud.Create:output_type -> golang_protobuf_brand.BrandResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protobuf_brand_proto_init() }
func file_protobuf_brand_proto_init() {
	if File_protobuf_brand_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_brand_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_brand_proto_goTypes,
		DependencyIndexes: file_protobuf_brand_proto_depIdxs,
		MessageInfos:      file_protobuf_brand_proto_msgTypes,
	}.Build()
	File_protobuf_brand_proto = out.File
	file_protobuf_brand_proto_rawDesc = nil
	file_protobuf_brand_proto_goTypes = nil
	file_protobuf_brand_proto_depIdxs = nil
}
