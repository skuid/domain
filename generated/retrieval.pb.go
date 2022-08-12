// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: retrieval.proto

package generated

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

// we get a request from tides that's going to
// fire off a series of requests to ${host}/v2/metadata/retrieve/plan
// to get the plan for this work.
// POST: ${host}/v2/metadata/retrieve/plan
// headers: content-type: application/json
//					authorization: bearer ${access-token}
//					x-skuid-public-key-endpoint:
//${host}/api/v1/site/verificationkey
type RetrievalRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host     string          `protobuf:"bytes,1,opt,name=Host,proto3" json:"Host,omitempty"`
	Username string          `protobuf:"bytes,2,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string          `protobuf:"bytes,3,opt,name=Password,proto3" json:"-"` // @gotags: json:"-"
	Filter   *RetrieveFilter `protobuf:"bytes,4,opt,name=Filter,proto3" json:"Filter,omitempty"`
}

func (x *RetrievalRequest) Reset() {
	*x = RetrievalRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_retrieval_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrievalRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrievalRequest) ProtoMessage() {}

func (x *RetrievalRequest) ProtoReflect() protoreflect.Message {
	mi := &file_retrieval_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrievalRequest.ProtoReflect.Descriptor instead.
func (*RetrievalRequest) Descriptor() ([]byte, []int) {
	return file_retrieval_proto_rawDescGZIP(), []int{0}
}

func (x *RetrievalRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RetrievalRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RetrievalRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RetrievalRequest) GetFilter() *RetrieveFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type RetrieveFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName   string   `protobuf:"bytes,1,opt,name=AppName,proto3" json:"appName"`     //@gotags: json:"appName"
	PageNames []string `protobuf:"bytes,2,rep,name=PageNames,proto3" json:"pageNames"` //@gotags: json:"pageNames"
}

func (x *RetrieveFilter) Reset() {
	*x = RetrieveFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_retrieval_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveFilter) ProtoMessage() {}

func (x *RetrieveFilter) ProtoReflect() protoreflect.Message {
	mi := &file_retrieval_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveFilter.ProtoReflect.Descriptor instead.
func (*RetrieveFilter) Descriptor() ([]byte, []int) {
	return file_retrieval_proto_rawDescGZIP(), []int{1}
}

func (x *RetrieveFilter) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *RetrieveFilter) GetPageNames() []string {
	if x != nil {
		return x.PageNames
	}
	return nil
}

type RetrievalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"data"`         // @gotags: json:"data"
	PlanName string `protobuf:"bytes,2,opt,name=PlanName,proto3" json:"planName"` // @gotags: json:"planName"
	Url      string `protobuf:"bytes,3,opt,name=Url,proto3" json:"url"`           // @gotags: json:"url"
}

func (x *RetrievalResponse) Reset() {
	*x = RetrievalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_retrieval_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrievalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrievalResponse) ProtoMessage() {}

func (x *RetrievalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_retrieval_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrievalResponse.ProtoReflect.Descriptor instead.
func (*RetrievalResponse) Descriptor() ([]byte, []int) {
	return file_retrieval_proto_rawDescGZIP(), []int{2}
}

func (x *RetrievalResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RetrievalResponse) GetPlanName() string {
	if x != nil {
		return x.PlanName
	}
	return ""
}

func (x *RetrievalResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_retrieval_proto protoreflect.FileDescriptor

var file_retrieval_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x90, 0x01, 0x0a, 0x10,
	0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x48, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x30, 0x0a, 0x06,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x48,
	0x0a, 0x0e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x41, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61,
	0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x50,
	0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x55, 0x0a, 0x11, 0x52, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x42,
	0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6b,
	0x75, 0x69, 0x64, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_retrieval_proto_rawDescOnce sync.Once
	file_retrieval_proto_rawDescData = file_retrieval_proto_rawDesc
)

func file_retrieval_proto_rawDescGZIP() []byte {
	file_retrieval_proto_rawDescOnce.Do(func() {
		file_retrieval_proto_rawDescData = protoimpl.X.CompressGZIP(file_retrieval_proto_rawDescData)
	})
	return file_retrieval_proto_rawDescData
}

var file_retrieval_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_retrieval_proto_goTypes = []interface{}{
	(*RetrievalRequest)(nil),  // 0: protobuf.RetrievalRequest
	(*RetrieveFilter)(nil),    // 1: protobuf.RetrieveFilter
	(*RetrievalResponse)(nil), // 2: protobuf.RetrievalResponse
}
var file_retrieval_proto_depIdxs = []int32{
	1, // 0: protobuf.RetrievalRequest.Filter:type_name -> protobuf.RetrieveFilter
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_retrieval_proto_init() }
func file_retrieval_proto_init() {
	if File_retrieval_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_retrieval_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrievalRequest); i {
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
		file_retrieval_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveFilter); i {
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
		file_retrieval_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrievalResponse); i {
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
			RawDescriptor: file_retrieval_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_retrieval_proto_goTypes,
		DependencyIndexes: file_retrieval_proto_depIdxs,
		MessageInfos:      file_retrieval_proto_msgTypes,
	}.Build()
	File_retrieval_proto = out.File
	file_retrieval_proto_rawDesc = nil
	file_retrieval_proto_goTypes = nil
	file_retrieval_proto_depIdxs = nil
}
