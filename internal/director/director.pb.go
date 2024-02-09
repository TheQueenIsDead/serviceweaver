// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.15.8
// source: director.proto

package director

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

type Director struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id   int32  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Director) Reset() {
	*x = Director{}
	if protoimpl.UnsafeEnabled {
		mi := &file_director_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Director) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Director) ProtoMessage() {}

func (x *Director) ProtoReflect() protoreflect.Message {
	mi := &file_director_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Director.ProtoReflect.Descriptor instead.
func (*Director) Descriptor() ([]byte, []int) {
	return file_director_proto_rawDescGZIP(), []int{0}
}

func (x *Director) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Director) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetDirectorByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDirectorByIdRequest) Reset() {
	*x = GetDirectorByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_director_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDirectorByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDirectorByIdRequest) ProtoMessage() {}

func (x *GetDirectorByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_director_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDirectorByIdRequest.ProtoReflect.Descriptor instead.
func (*GetDirectorByIdRequest) Descriptor() ([]byte, []int) {
	return file_director_proto_rawDescGZIP(), []int{1}
}

func (x *GetDirectorByIdRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetDirectorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Director `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GetDirectorResponse) Reset() {
	*x = GetDirectorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_director_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDirectorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDirectorResponse) ProtoMessage() {}

func (x *GetDirectorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_director_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDirectorResponse.ProtoReflect.Descriptor instead.
func (*GetDirectorResponse) Descriptor() ([]byte, []int) {
	return file_director_proto_rawDescGZIP(), []int{2}
}

func (x *GetDirectorResponse) GetData() *Director {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_director_proto protoreflect.FileDescriptor

var file_director_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x2e, 0x0a, 0x08, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x44, 0x69, 0x72, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x38, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x32, 0x5d, 0x0a, 0x0f, 0x44, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x21, 0x5a, 0x1f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x77, 0x65, 0x61, 0x76, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_director_proto_rawDescOnce sync.Once
	file_director_proto_rawDescData = file_director_proto_rawDesc
)

func file_director_proto_rawDescGZIP() []byte {
	file_director_proto_rawDescOnce.Do(func() {
		file_director_proto_rawDescData = protoimpl.X.CompressGZIP(file_director_proto_rawDescData)
	})
	return file_director_proto_rawDescData
}

var file_director_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_director_proto_goTypes = []interface{}{
	(*Director)(nil),               // 0: api.Director
	(*GetDirectorByIdRequest)(nil), // 1: api.GetDirectorByIdRequest
	(*GetDirectorResponse)(nil),    // 2: api.GetDirectorResponse
}
var file_director_proto_depIdxs = []int32{
	0, // 0: api.GetDirectorResponse.Data:type_name -> api.Director
	1, // 1: api.DirectorService.GetDirectorById:input_type -> api.GetDirectorByIdRequest
	2, // 2: api.DirectorService.GetDirectorById:output_type -> api.GetDirectorResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_director_proto_init() }
func file_director_proto_init() {
	if File_director_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_director_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Director); i {
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
		file_director_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDirectorByIdRequest); i {
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
		file_director_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDirectorResponse); i {
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
			RawDescriptor: file_director_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_director_proto_goTypes,
		DependencyIndexes: file_director_proto_depIdxs,
		MessageInfos:      file_director_proto_msgTypes,
	}.Build()
	File_director_proto = out.File
	file_director_proto_rawDesc = nil
	file_director_proto_goTypes = nil
	file_director_proto_depIdxs = nil
}