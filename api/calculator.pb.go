// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: api/calculator.proto

package api

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

// The request message containing the computation to be performed on the server.
type ComputeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OperandA int32  `protobuf:"varint,1,opt,name=operand_a,json=operandA,proto3" json:"operand_a,omitempty"`
	OperandB int32  `protobuf:"varint,2,opt,name=operand_b,json=operandB,proto3" json:"operand_b,omitempty"`
	Operator string `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (x *ComputeRequest) Reset() {
	*x = ComputeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeRequest) ProtoMessage() {}

func (x *ComputeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeRequest.ProtoReflect.Descriptor instead.
func (*ComputeRequest) Descriptor() ([]byte, []int) {
	return file_api_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *ComputeRequest) GetOperandA() int32 {
	if x != nil {
		return x.OperandA
	}
	return 0
}

func (x *ComputeRequest) GetOperandB() int32 {
	if x != nil {
		return x.OperandB
	}
	return 0
}

func (x *ComputeRequest) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

// The response message containing the computation result
type ComputeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ComputeResponse) Reset() {
	*x = ComputeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputeResponse) ProtoMessage() {}

func (x *ComputeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputeResponse.ProtoReflect.Descriptor instead.
func (*ComputeResponse) Descriptor() ([]byte, []int) {
	return file_api_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *ComputeResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_api_calculator_proto protoreflect.FileDescriptor

var file_api_calculator_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x22, 0x66, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x5f,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64,
	0x41, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x62, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x42, 0x12, 0x1a,
	0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x29, 0x0a, 0x0f, 0x43, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x52, 0x0a, 0x0a, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x44, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x12, 0x1a,
	0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x64, 0x75, 0x32, 0x30, 0x32, 0x30,
	0x2f, 0x6c, 0x6f, 0x76, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_calculator_proto_rawDescOnce sync.Once
	file_api_calculator_proto_rawDescData = file_api_calculator_proto_rawDesc
)

func file_api_calculator_proto_rawDescGZIP() []byte {
	file_api_calculator_proto_rawDescOnce.Do(func() {
		file_api_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_calculator_proto_rawDescData)
	})
	return file_api_calculator_proto_rawDescData
}

var file_api_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_calculator_proto_goTypes = []interface{}{
	(*ComputeRequest)(nil),  // 0: calculator.ComputeRequest
	(*ComputeResponse)(nil), // 1: calculator.ComputeResponse
}
var file_api_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.Calculator.Compute:input_type -> calculator.ComputeRequest
	1, // 1: calculator.Calculator.Compute:output_type -> calculator.ComputeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_calculator_proto_init() }
func file_api_calculator_proto_init() {
	if File_api_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeRequest); i {
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
		file_api_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputeResponse); i {
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
			RawDescriptor: file_api_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_calculator_proto_goTypes,
		DependencyIndexes: file_api_calculator_proto_depIdxs,
		MessageInfos:      file_api_calculator_proto_msgTypes,
	}.Build()
	File_api_calculator_proto = out.File
	file_api_calculator_proto_rawDesc = nil
	file_api_calculator_proto_goTypes = nil
	file_api_calculator_proto_depIdxs = nil
}