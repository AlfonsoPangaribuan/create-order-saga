// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.30.2
// source: proto/shipping.proto

package shipping

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

type StartShippingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       int64                  `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartShippingRequest) Reset() {
	*x = StartShippingRequest{}
	mi := &file_proto_shipping_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartShippingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartShippingRequest) ProtoMessage() {}

func (x *StartShippingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartShippingRequest.ProtoReflect.Descriptor instead.
func (*StartShippingRequest) Descriptor() ([]byte, []int) {
	return file_proto_shipping_proto_rawDescGZIP(), []int{0}
}

func (x *StartShippingRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type StartShippingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StartShippingResponse) Reset() {
	*x = StartShippingResponse{}
	mi := &file_proto_shipping_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartShippingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartShippingResponse) ProtoMessage() {}

func (x *StartShippingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartShippingResponse.ProtoReflect.Descriptor instead.
func (*StartShippingResponse) Descriptor() ([]byte, []int) {
	return file_proto_shipping_proto_rawDescGZIP(), []int{1}
}

func (x *StartShippingResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type CancelShippingRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OrderId       int64                  `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelShippingRequest) Reset() {
	*x = CancelShippingRequest{}
	mi := &file_proto_shipping_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelShippingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelShippingRequest) ProtoMessage() {}

func (x *CancelShippingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelShippingRequest.ProtoReflect.Descriptor instead.
func (*CancelShippingRequest) Descriptor() ([]byte, []int) {
	return file_proto_shipping_proto_rawDescGZIP(), []int{2}
}

func (x *CancelShippingRequest) GetOrderId() int64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type CancelShippingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CancelShippingResponse) Reset() {
	*x = CancelShippingResponse{}
	mi := &file_proto_shipping_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CancelShippingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelShippingResponse) ProtoMessage() {}

func (x *CancelShippingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shipping_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelShippingResponse.ProtoReflect.Descriptor instead.
func (*CancelShippingResponse) Descriptor() ([]byte, []int) {
	return file_proto_shipping_proto_rawDescGZIP(), []int{3}
}

func (x *CancelShippingResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_proto_shipping_proto protoreflect.FileDescriptor

const file_proto_shipping_proto_rawDesc = "" +
	"\n" +
	"\x14proto/shipping.proto\x12\bshipping\"1\n" +
	"\x14StartShippingRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\x03R\aorderId\"/\n" +
	"\x15StartShippingResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status\"2\n" +
	"\x15CancelShippingRequest\x12\x19\n" +
	"\border_id\x18\x01 \x01(\x03R\aorderId\"0\n" +
	"\x16CancelShippingResponse\x12\x16\n" +
	"\x06status\x18\x01 \x01(\tR\x06status2\xb8\x01\n" +
	"\x0fShippingService\x12P\n" +
	"\rStartShipping\x12\x1e.shipping.StartShippingRequest\x1a\x1f.shipping.StartShippingResponse\x12S\n" +
	"\x0eCancelShipping\x12\x1f.shipping.CancelShippingRequest\x1a .shipping.CancelShippingResponseB+Z)create-order-saga/proto/shipping;shippingb\x06proto3"

var (
	file_proto_shipping_proto_rawDescOnce sync.Once
	file_proto_shipping_proto_rawDescData []byte
)

func file_proto_shipping_proto_rawDescGZIP() []byte {
	file_proto_shipping_proto_rawDescOnce.Do(func() {
		file_proto_shipping_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_shipping_proto_rawDesc), len(file_proto_shipping_proto_rawDesc)))
	})
	return file_proto_shipping_proto_rawDescData
}

var file_proto_shipping_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_shipping_proto_goTypes = []any{
	(*StartShippingRequest)(nil),   // 0: shipping.StartShippingRequest
	(*StartShippingResponse)(nil),  // 1: shipping.StartShippingResponse
	(*CancelShippingRequest)(nil),  // 2: shipping.CancelShippingRequest
	(*CancelShippingResponse)(nil), // 3: shipping.CancelShippingResponse
}
var file_proto_shipping_proto_depIdxs = []int32{
	0, // 0: shipping.ShippingService.StartShipping:input_type -> shipping.StartShippingRequest
	2, // 1: shipping.ShippingService.CancelShipping:input_type -> shipping.CancelShippingRequest
	1, // 2: shipping.ShippingService.StartShipping:output_type -> shipping.StartShippingResponse
	3, // 3: shipping.ShippingService.CancelShipping:output_type -> shipping.CancelShippingResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_shipping_proto_init() }
func file_proto_shipping_proto_init() {
	if File_proto_shipping_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_shipping_proto_rawDesc), len(file_proto_shipping_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_shipping_proto_goTypes,
		DependencyIndexes: file_proto_shipping_proto_depIdxs,
		MessageInfos:      file_proto_shipping_proto_msgTypes,
	}.Build()
	File_proto_shipping_proto = out.File
	file_proto_shipping_proto_goTypes = nil
	file_proto_shipping_proto_depIdxs = nil
}
