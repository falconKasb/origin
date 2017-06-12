// Code generated by protoc-gen-go.
// source: examples/examplepb/a_bit_of_everything.proto
// DO NOT EDIT!

package examplepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"
import grpc_gateway_examples_sub "github.com/grpc-ecosystem/grpc-gateway/examples/sub"
import sub2 "github.com/grpc-ecosystem/grpc-gateway/examples/sub2"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// NumericEnum is one or zero.
type NumericEnum int32

const (
	// ZERO means 0
	NumericEnum_ZERO NumericEnum = 0
	// ONE means 1
	NumericEnum_ONE NumericEnum = 1
)

var NumericEnum_name = map[int32]string{
	0: "ZERO",
	1: "ONE",
}
var NumericEnum_value = map[string]int32{
	"ZERO": 0,
	"ONE":  1,
}

func (x NumericEnum) String() string {
	return proto.EnumName(NumericEnum_name, int32(x))
}
func (NumericEnum) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// DeepEnum is one or zero.
type ABitOfEverything_Nested_DeepEnum int32

const (
	// FALSE is false.
	ABitOfEverything_Nested_FALSE ABitOfEverything_Nested_DeepEnum = 0
	// TRUE is true.
	ABitOfEverything_Nested_TRUE ABitOfEverything_Nested_DeepEnum = 1
)

var ABitOfEverything_Nested_DeepEnum_name = map[int32]string{
	0: "FALSE",
	1: "TRUE",
}
var ABitOfEverything_Nested_DeepEnum_value = map[string]int32{
	"FALSE": 0,
	"TRUE":  1,
}

func (x ABitOfEverything_Nested_DeepEnum) String() string {
	return proto.EnumName(ABitOfEverything_Nested_DeepEnum_name, int32(x))
}
func (ABitOfEverything_Nested_DeepEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{0, 0, 0}
}

// Intentionaly complicated message type to cover much features of Protobuf.
// NEXT ID: 27
type ABitOfEverything struct {
	SingleNested *ABitOfEverything_Nested   `protobuf:"bytes,25,opt,name=single_nested,json=singleNested" json:"single_nested,omitempty"`
	Uuid         string                     `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Nested       []*ABitOfEverything_Nested `protobuf:"bytes,2,rep,name=nested" json:"nested,omitempty"`
	FloatValue   float32                    `protobuf:"fixed32,3,opt,name=float_value,json=floatValue" json:"float_value,omitempty"`
	DoubleValue  float64                    `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue" json:"double_value,omitempty"`
	Int64Value   int64                      `protobuf:"varint,5,opt,name=int64_value,json=int64Value" json:"int64_value,omitempty"`
	Uint64Value  uint64                     `protobuf:"varint,6,opt,name=uint64_value,json=uint64Value" json:"uint64_value,omitempty"`
	Int32Value   int32                      `protobuf:"varint,7,opt,name=int32_value,json=int32Value" json:"int32_value,omitempty"`
	Fixed64Value uint64                     `protobuf:"fixed64,8,opt,name=fixed64_value,json=fixed64Value" json:"fixed64_value,omitempty"`
	Fixed32Value uint32                     `protobuf:"fixed32,9,opt,name=fixed32_value,json=fixed32Value" json:"fixed32_value,omitempty"`
	BoolValue    bool                       `protobuf:"varint,10,opt,name=bool_value,json=boolValue" json:"bool_value,omitempty"`
	StringValue  string                     `protobuf:"bytes,11,opt,name=string_value,json=stringValue" json:"string_value,omitempty"`
	// TODO(yugui) add bytes_value
	Uint32Value         uint32      `protobuf:"varint,13,opt,name=uint32_value,json=uint32Value" json:"uint32_value,omitempty"`
	EnumValue           NumericEnum `protobuf:"varint,14,opt,name=enum_value,json=enumValue,enum=grpc.gateway.examples.examplepb.NumericEnum" json:"enum_value,omitempty"`
	Sfixed32Value       int32       `protobuf:"fixed32,15,opt,name=sfixed32_value,json=sfixed32Value" json:"sfixed32_value,omitempty"`
	Sfixed64Value       int64       `protobuf:"fixed64,16,opt,name=sfixed64_value,json=sfixed64Value" json:"sfixed64_value,omitempty"`
	Sint32Value         int32       `protobuf:"zigzag32,17,opt,name=sint32_value,json=sint32Value" json:"sint32_value,omitempty"`
	Sint64Value         int64       `protobuf:"zigzag64,18,opt,name=sint64_value,json=sint64Value" json:"sint64_value,omitempty"`
	RepeatedStringValue []string    `protobuf:"bytes,19,rep,name=repeated_string_value,json=repeatedStringValue" json:"repeated_string_value,omitempty"`
	// Types that are valid to be assigned to OneofValue:
	//	*ABitOfEverything_OneofEmpty
	//	*ABitOfEverything_OneofString
	OneofValue               isABitOfEverything_OneofValue       `protobuf_oneof:"oneof_value"`
	MapValue                 map[string]NumericEnum              `protobuf:"bytes,22,rep,name=map_value,json=mapValue" json:"map_value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value,enum=grpc.gateway.examples.examplepb.NumericEnum"`
	MappedStringValue        map[string]string                   `protobuf:"bytes,23,rep,name=mapped_string_value,json=mappedStringValue" json:"mapped_string_value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	MappedNestedValue        map[string]*ABitOfEverything_Nested `protobuf:"bytes,24,rep,name=mapped_nested_value,json=mappedNestedValue" json:"mapped_nested_value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NonConventionalNameValue string                              `protobuf:"bytes,26,opt,name=nonConventionalNameValue" json:"nonConventionalNameValue,omitempty"`
}

func (m *ABitOfEverything) Reset()                    { *m = ABitOfEverything{} }
func (m *ABitOfEverything) String() string            { return proto.CompactTextString(m) }
func (*ABitOfEverything) ProtoMessage()               {}
func (*ABitOfEverything) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type isABitOfEverything_OneofValue interface {
	isABitOfEverything_OneofValue()
}

type ABitOfEverything_OneofEmpty struct {
	OneofEmpty *google_protobuf1.Empty `protobuf:"bytes,20,opt,name=oneof_empty,json=oneofEmpty,oneof"`
}
type ABitOfEverything_OneofString struct {
	OneofString string `protobuf:"bytes,21,opt,name=oneof_string,json=oneofString,oneof"`
}

func (*ABitOfEverything_OneofEmpty) isABitOfEverything_OneofValue()  {}
func (*ABitOfEverything_OneofString) isABitOfEverything_OneofValue() {}

func (m *ABitOfEverything) GetOneofValue() isABitOfEverything_OneofValue {
	if m != nil {
		return m.OneofValue
	}
	return nil
}

func (m *ABitOfEverything) GetSingleNested() *ABitOfEverything_Nested {
	if m != nil {
		return m.SingleNested
	}
	return nil
}

func (m *ABitOfEverything) GetNested() []*ABitOfEverything_Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *ABitOfEverything) GetOneofEmpty() *google_protobuf1.Empty {
	if x, ok := m.GetOneofValue().(*ABitOfEverything_OneofEmpty); ok {
		return x.OneofEmpty
	}
	return nil
}

func (m *ABitOfEverything) GetOneofString() string {
	if x, ok := m.GetOneofValue().(*ABitOfEverything_OneofString); ok {
		return x.OneofString
	}
	return ""
}

func (m *ABitOfEverything) GetMapValue() map[string]NumericEnum {
	if m != nil {
		return m.MapValue
	}
	return nil
}

func (m *ABitOfEverything) GetMappedStringValue() map[string]string {
	if m != nil {
		return m.MappedStringValue
	}
	return nil
}

func (m *ABitOfEverything) GetMappedNestedValue() map[string]*ABitOfEverything_Nested {
	if m != nil {
		return m.MappedNestedValue
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ABitOfEverything) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ABitOfEverything_OneofMarshaler, _ABitOfEverything_OneofUnmarshaler, _ABitOfEverything_OneofSizer, []interface{}{
		(*ABitOfEverything_OneofEmpty)(nil),
		(*ABitOfEverything_OneofString)(nil),
	}
}

func _ABitOfEverything_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ABitOfEverything)
	// oneof_value
	switch x := m.OneofValue.(type) {
	case *ABitOfEverything_OneofEmpty:
		b.EncodeVarint(20<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.OneofEmpty); err != nil {
			return err
		}
	case *ABitOfEverything_OneofString:
		b.EncodeVarint(21<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.OneofString)
	case nil:
	default:
		return fmt.Errorf("ABitOfEverything.OneofValue has unexpected type %T", x)
	}
	return nil
}

func _ABitOfEverything_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ABitOfEverything)
	switch tag {
	case 20: // oneof_value.oneof_empty
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Empty)
		err := b.DecodeMessage(msg)
		m.OneofValue = &ABitOfEverything_OneofEmpty{msg}
		return true, err
	case 21: // oneof_value.oneof_string
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OneofValue = &ABitOfEverything_OneofString{x}
		return true, err
	default:
		return false, nil
	}
}

func _ABitOfEverything_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ABitOfEverything)
	// oneof_value
	switch x := m.OneofValue.(type) {
	case *ABitOfEverything_OneofEmpty:
		s := proto.Size(x.OneofEmpty)
		n += proto.SizeVarint(20<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ABitOfEverything_OneofString:
		n += proto.SizeVarint(21<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.OneofString)))
		n += len(x.OneofString)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Nested is nested type.
type ABitOfEverything_Nested struct {
	// name is nested field.
	Name   string                           `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Amount uint32                           `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
	Ok     ABitOfEverything_Nested_DeepEnum `protobuf:"varint,3,opt,name=ok,enum=grpc.gateway.examples.examplepb.ABitOfEverything_Nested_DeepEnum" json:"ok,omitempty"`
}

func (m *ABitOfEverything_Nested) Reset()                    { *m = ABitOfEverything_Nested{} }
func (m *ABitOfEverything_Nested) String() string            { return proto.CompactTextString(m) }
func (*ABitOfEverything_Nested) ProtoMessage()               {}
func (*ABitOfEverything_Nested) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

func init() {
	proto.RegisterType((*ABitOfEverything)(nil), "grpc.gateway.examples.examplepb.ABitOfEverything")
	proto.RegisterType((*ABitOfEverything_Nested)(nil), "grpc.gateway.examples.examplepb.ABitOfEverything.Nested")
	proto.RegisterEnum("grpc.gateway.examples.examplepb.NumericEnum", NumericEnum_name, NumericEnum_value)
	proto.RegisterEnum("grpc.gateway.examples.examplepb.ABitOfEverything_Nested_DeepEnum", ABitOfEverything_Nested_DeepEnum_name, ABitOfEverything_Nested_DeepEnum_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ABitOfEverythingService service

type ABitOfEverythingServiceClient interface {
	Create(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error)
	CreateBody(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error)
	Lookup(ctx context.Context, in *sub2.IdMessage, opts ...grpc.CallOption) (*ABitOfEverything, error)
	Update(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Delete(ctx context.Context, in *sub2.IdMessage, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	Echo(ctx context.Context, in *grpc_gateway_examples_sub.StringMessage, opts ...grpc.CallOption) (*grpc_gateway_examples_sub.StringMessage, error)
	DeepPathEcho(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error)
	Timeout(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type aBitOfEverythingServiceClient struct {
	cc *grpc.ClientConn
}

func NewABitOfEverythingServiceClient(cc *grpc.ClientConn) ABitOfEverythingServiceClient {
	return &aBitOfEverythingServiceClient{cc}
}

func (c *aBitOfEverythingServiceClient) Create(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error) {
	out := new(ABitOfEverything)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) CreateBody(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error) {
	out := new(ABitOfEverything)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/CreateBody", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) Lookup(ctx context.Context, in *sub2.IdMessage, opts ...grpc.CallOption) (*ABitOfEverything, error) {
	out := new(ABitOfEverything)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Lookup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) Update(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) Delete(ctx context.Context, in *sub2.IdMessage, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) Echo(ctx context.Context, in *grpc_gateway_examples_sub.StringMessage, opts ...grpc.CallOption) (*grpc_gateway_examples_sub.StringMessage, error) {
	out := new(grpc_gateway_examples_sub.StringMessage)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) DeepPathEcho(ctx context.Context, in *ABitOfEverything, opts ...grpc.CallOption) (*ABitOfEverything, error) {
	out := new(ABitOfEverything)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/DeepPathEcho", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aBitOfEverythingServiceClient) Timeout(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Timeout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ABitOfEverythingService service

type ABitOfEverythingServiceServer interface {
	Create(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	CreateBody(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	Lookup(context.Context, *sub2.IdMessage) (*ABitOfEverything, error)
	Update(context.Context, *ABitOfEverything) (*google_protobuf1.Empty, error)
	Delete(context.Context, *sub2.IdMessage) (*google_protobuf1.Empty, error)
	Echo(context.Context, *grpc_gateway_examples_sub.StringMessage) (*grpc_gateway_examples_sub.StringMessage, error)
	DeepPathEcho(context.Context, *ABitOfEverything) (*ABitOfEverything, error)
	Timeout(context.Context, *google_protobuf1.Empty) (*google_protobuf1.Empty, error)
}

func RegisterABitOfEverythingServiceServer(s *grpc.Server, srv ABitOfEverythingServiceServer) {
	s.RegisterService(&_ABitOfEverythingService_serviceDesc, srv)
}

func _ABitOfEverythingService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ABitOfEverything)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).Create(ctx, req.(*ABitOfEverything))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_CreateBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ABitOfEverything)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).CreateBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/CreateBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).CreateBody(ctx, req.(*ABitOfEverything))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_Lookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sub2.IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).Lookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Lookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).Lookup(ctx, req.(*sub2.IdMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ABitOfEverything)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).Update(ctx, req.(*ABitOfEverything))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(sub2.IdMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).Delete(ctx, req.(*sub2.IdMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc_gateway_examples_sub.StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).Echo(ctx, req.(*grpc_gateway_examples_sub.StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_DeepPathEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ABitOfEverything)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).DeepPathEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/DeepPathEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).DeepPathEcho(ctx, req.(*ABitOfEverything))
	}
	return interceptor(ctx, in, info, handler)
}

func _ABitOfEverythingService_Timeout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ABitOfEverythingServiceServer).Timeout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.gateway.examples.examplepb.ABitOfEverythingService/Timeout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ABitOfEverythingServiceServer).Timeout(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ABitOfEverythingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.gateway.examples.examplepb.ABitOfEverythingService",
	HandlerType: (*ABitOfEverythingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ABitOfEverythingService_Create_Handler,
		},
		{
			MethodName: "CreateBody",
			Handler:    _ABitOfEverythingService_CreateBody_Handler,
		},
		{
			MethodName: "Lookup",
			Handler:    _ABitOfEverythingService_Lookup_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ABitOfEverythingService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ABitOfEverythingService_Delete_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _ABitOfEverythingService_Echo_Handler,
		},
		{
			MethodName: "DeepPathEcho",
			Handler:    _ABitOfEverythingService_DeepPathEcho_Handler,
		},
		{
			MethodName: "Timeout",
			Handler:    _ABitOfEverythingService_Timeout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("examples/examplepb/a_bit_of_everything.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x57, 0x4d, 0x4f, 0x1b, 0x47,
	0x18, 0xce, 0xda, 0x60, 0xf0, 0xbb, 0x18, 0xcc, 0x10, 0x88, 0xe3, 0xb4, 0x82, 0xba, 0x1f, 0x42,
	0x14, 0xed, 0x0a, 0x13, 0x55, 0x09, 0x52, 0x15, 0x41, 0x70, 0x95, 0xaa, 0x09, 0xa4, 0x4b, 0x92,
	0x03, 0x6a, 0x64, 0xad, 0xf1, 0x60, 0x56, 0x78, 0x77, 0x56, 0xfb, 0xe1, 0x62, 0x51, 0x7a, 0xe8,
	0xa1, 0x7f, 0xa0, 0x87, 0x9e, 0x9a, 0x4b, 0xa5, 0xaa, 0x97, 0x1e, 0x7b, 0xee, 0x8f, 0xe8, 0x5f,
	0xe8, 0x0f, 0xe9, 0xbb, 0x33, 0xbb, 0xcb, 0xac, 0xc1, 0x32, 0x1f, 0x52, 0x4e, 0xde, 0x99, 0x79,
	0xde, 0xe7, 0x79, 0x3f, 0xf6, 0x7d, 0x67, 0x0d, 0xab, 0xf4, 0xc4, 0xb4, 0xdd, 0x2e, 0xf5, 0xf5,
	0xf8, 0xc1, 0x6d, 0xe9, 0x66, 0xb3, 0x65, 0x05, 0x4d, 0x76, 0xd8, 0xa4, 0x3d, 0xea, 0xf5, 0x83,
	0x23, 0xcb, 0xe9, 0x68, 0xae, 0xc7, 0x02, 0x46, 0x16, 0x3b, 0x9e, 0x7b, 0xa0, 0x75, 0xcc, 0x80,
	0x7e, 0x6f, 0xf6, 0xb5, 0xc4, 0x54, 0x4b, 0x4d, 0xab, 0x1f, 0x74, 0x18, 0xeb, 0x74, 0xa9, 0x6e,
	0xba, 0x96, 0x6e, 0x3a, 0x0e, 0x0b, 0xcc, 0xc0, 0x62, 0x8e, 0x2f, 0xcc, 0xab, 0x0f, 0xe2, 0x53,
	0xbe, 0x6a, 0x85, 0x87, 0x3a, 0xb5, 0xdd, 0xa0, 0x1f, 0x1f, 0x56, 0x53, 0x4f, 0xfc, 0xb0, 0xa5,
	0xdb, 0xd4, 0xf7, 0xcd, 0x0e, 0x4d, 0x0c, 0xe5, 0xb3, 0x7a, 0xf6, 0xb0, 0xf6, 0xdb, 0x34, 0x94,
	0x37, 0xb7, 0xac, 0x60, 0xf7, 0xb0, 0x91, 0xfa, 0x4b, 0xde, 0x42, 0xc9, 0xc7, 0xdf, 0x2e, 0x6d,
	0x3a, 0xd4, 0x0f, 0x68, 0xbb, 0x72, 0x7f, 0x49, 0x59, 0x56, 0xeb, 0x8f, 0xb4, 0x11, 0x11, 0x68,
	0x83, 0x4c, 0xda, 0x0e, 0xb7, 0x37, 0xa6, 0x04, 0x9d, 0x58, 0x11, 0x02, 0x63, 0x61, 0x68, 0xb5,
	0x2b, 0x0a, 0xb2, 0x16, 0x0d, 0xfe, 0x4c, 0x5e, 0x42, 0x21, 0xd6, 0xca, 0x2d, 0xe5, 0x6f, 0xa5,
	0x15, 0xf3, 0x90, 0x45, 0x50, 0x0f, 0xbb, 0xcc, 0x0c, 0x9a, 0x3d, 0xb3, 0x1b, 0xd2, 0x4a, 0x1e,
	0xc5, 0x72, 0x06, 0xf0, 0xad, 0x37, 0xd1, 0x0e, 0xf9, 0x08, 0xa6, 0xda, 0x2c, 0x6c, 0x61, 0x94,
	0x02, 0x31, 0x86, 0x08, 0xc5, 0x50, 0xc5, 0x9e, 0x80, 0x20, 0x87, 0xe5, 0x04, 0x5f, 0x3c, 0x8c,
	0x11, 0xe3, 0x88, 0xc8, 0x1b, 0xc0, 0xb7, 0x52, 0x8e, 0x50, 0x46, 0x14, 0x10, 0x31, 0x66, 0xa8,
	0xa1, 0x04, 0x11, 0x1c, 0xeb, 0xf5, 0x18, 0x31, 0x81, 0x88, 0x71, 0xce, 0xb1, 0x5e, 0x17, 0x80,
	0x8f, 0xa1, 0x74, 0x68, 0x9d, 0xd0, 0x76, 0x4a, 0x32, 0x89, 0x90, 0x82, 0x31, 0x15, 0x6f, 0x66,
	0x41, 0x29, 0x4f, 0x11, 0x41, 0x13, 0x31, 0x28, 0x61, 0xfa, 0x10, 0xa0, 0xc5, 0x58, 0x37, 0x46,
	0x00, 0x22, 0x26, 0x8d, 0x62, 0xb4, 0x93, 0x3a, 0xeb, 0x07, 0x1e, 0xa6, 0x2a, 0x06, 0xa8, 0x3c,
	0xff, 0xaa, 0xd8, 0xcb, 0xc4, 0x93, 0xaa, 0x94, 0x10, 0x52, 0x12, 0xf1, 0x24, 0x22, 0xdf, 0x00,
	0x50, 0x27, 0xb4, 0x63, 0xc0, 0x34, 0x02, 0xa6, 0xeb, 0xab, 0x23, 0xab, 0xb5, 0x13, 0xda, 0xd4,
	0xb3, 0x0e, 0x1a, 0x68, 0x69, 0x14, 0x23, 0x7b, 0x41, 0xf6, 0x29, 0x4c, 0xfb, 0xd9, 0xb8, 0x66,
	0x90, 0x70, 0xc6, 0x28, 0xf9, 0x99, 0xc0, 0x52, 0x58, 0x9a, 0xa3, 0x32, 0xc2, 0xca, 0x09, 0x4c,
	0xaa, 0x86, 0x2f, 0x7b, 0x3f, 0x8b, 0xa0, 0x59, 0x0c, 0x50, 0xf2, 0x3e, 0x86, 0xa4, 0x3c, 0x04,
	0x21, 0x44, 0x40, 0x12, 0x96, 0x3a, 0xcc, 0x7b, 0xd4, 0xa5, 0x18, 0x4b, 0xbb, 0x99, 0xc9, 0xd7,
	0x1c, 0xbe, 0x99, 0x45, 0x63, 0x2e, 0x39, 0xdc, 0x93, 0xf2, 0xf6, 0x18, 0x54, 0xe6, 0xd0, 0xa8,
	0xeb, 0xa3, 0xa6, 0xac, 0xdc, 0xe5, 0xfd, 0xb2, 0xa0, 0x89, 0x96, 0xd5, 0x92, 0x96, 0xd5, 0x1a,
	0xd1, 0xe9, 0xb3, 0x3b, 0x06, 0x70, 0x30, 0x5f, 0x61, 0x65, 0xa7, 0x84, 0xa9, 0xd0, 0xaa, 0xcc,
	0x47, 0x55, 0x41, 0x8c, 0x20, 0x14, 0x22, 0xe4, 0x3b, 0x28, 0xda, 0xa6, 0x1b, 0xfb, 0xb1, 0xc0,
	0x3b, 0xe4, 0xc9, 0xf5, 0x3b, 0xe4, 0x85, 0xe9, 0x72, 0x77, 0x1b, 0x4e, 0xe0, 0xf5, 0x8d, 0x49,
	0x3b, 0x5e, 0x92, 0x13, 0x98, 0xc3, 0x67, 0x77, 0x30, 0xde, 0x7b, 0x5c, 0xe7, 0xd9, 0x8d, 0x74,
	0xdc, 0x4c, 0x7e, 0x84, 0xe0, 0xac, 0x3d, 0xb8, 0x2f, 0x29, 0x8b, 0xae, 0x8d, 0x95, 0x2b, 0xb7,
	0x53, 0x16, 0x93, 0xe0, 0xa2, 0xb2, 0xb4, 0x4f, 0x36, 0xa0, 0xe2, 0x30, 0xe7, 0x29, 0x73, 0x7a,
	0xd4, 0x89, 0xc6, 0xac, 0xd9, 0xdd, 0x31, 0x6d, 0xd1, 0xf6, 0x95, 0x2a, 0x6f, 0x8c, 0xa1, 0xe7,
	0xd5, 0x3f, 0x15, 0x28, 0x9c, 0xcf, 0x32, 0x07, 0xf7, 0x93, 0x59, 0x16, 0x3d, 0x93, 0x05, 0x28,
	0x98, 0x36, 0x0b, 0x9d, 0x00, 0x67, 0x59, 0xd4, 0x3e, 0xf1, 0x8a, 0x7c, 0x0b, 0x39, 0x76, 0xcc,
	0x07, 0xd1, 0x74, 0x7d, 0xf3, 0xa6, 0xf3, 0x4d, 0xdb, 0xa6, 0xd4, 0xe5, 0x6d, 0x84, 0x64, 0xb5,
	0x45, 0x98, 0x4c, 0xd6, 0xa4, 0x08, 0xe3, 0x5f, 0x6d, 0x3e, 0xdf, 0x6b, 0x94, 0xef, 0x90, 0x49,
	0x18, 0x7b, 0x65, 0xbc, 0x6e, 0x94, 0x95, 0xaa, 0x05, 0xa5, 0x4c, 0xd5, 0x49, 0x19, 0xf2, 0xc7,
	0xb4, 0x1f, 0xfb, 0x1b, 0x3d, 0x92, 0x2d, 0x18, 0x17, 0x59, 0xcf, 0xdd, 0xa0, 0x97, 0x85, 0xe9,
	0x46, 0xee, 0x91, 0x52, 0xdd, 0x86, 0x85, 0xcb, 0x0b, 0x7f, 0x89, 0xe6, 0x5d, 0x59, 0xb3, 0x28,
	0xb3, 0xfc, 0x98, 0xb0, 0x0c, 0x16, 0xf1, 0x12, 0x96, 0x1d, 0x99, 0xe5, 0x36, 0x77, 0xc6, 0xb9,
	0xfe, 0x56, 0x29, 0xe9, 0x64, 0xbe, 0xb5, 0xb2, 0x04, 0xaa, 0x14, 0x6e, 0x94, 0xd8, 0xfd, 0x86,
	0xb1, 0x8b, 0x29, 0x9e, 0x80, 0xfc, 0xee, 0x0e, 0x66, 0xb8, 0xfe, 0xab, 0x0a, 0xf7, 0x06, 0x79,
	0xf7, 0xa8, 0xd7, 0xb3, 0x0e, 0x28, 0x79, 0x97, 0x87, 0xc2, 0x53, 0x2f, 0x9a, 0x16, 0x64, 0xed,
	0xda, 0xce, 0x55, 0xaf, 0x6f, 0x52, 0xfb, 0x2b, 0xf7, 0xd3, 0xbf, 0xff, 0xfd, 0x92, 0xfb, 0x23,
	0x57, 0xfb, 0x3d, 0xa7, 0xf7, 0xd6, 0x92, 0xef, 0x92, 0xcb, 0xbe, 0x4a, 0xf4, 0x53, 0xe9, 0x7a,
	0x3c, 0xd3, 0x4f, 0xe5, 0xbb, 0x10, 0x97, 0xd2, 0x90, 0x3c, 0xd3, 0x7d, 0xea, 0x9a, 0x9e, 0x19,
	0x30, 0x4f, 0x3f, 0x0d, 0x33, 0x07, 0xa7, 0xd2, 0xb8, 0xc5, 0x55, 0x66, 0x46, 0x27, 0x6b, 0xe9,
	0xfc, 0xfc, 0x76, 0xc2, 0x85, 0x3c, 0x6b, 0xbe, 0xc4, 0x85, 0xeb, 0x51, 0xc4, 0xeb, 0x2b, 0x67,
	0x42, 0x44, 0x32, 0xf3, 0x07, 0x79, 0xfc, 0x41, 0x21, 0x7f, 0xc0, 0x20, 0xeb, 0xe4, 0xb0, 0x46,
	0x3e, 0x23, 0xef, 0x14, 0x00, 0x51, 0xa0, 0x2d, 0xd6, 0xee, 0xbf, 0xa7, 0x22, 0xad, 0xf0, 0x1a,
	0x7d, 0x52, 0x5b, 0x1c, 0x51, 0xa1, 0x0d, 0x65, 0x85, 0xfc, 0x00, 0x85, 0xe7, 0x8c, 0x1d, 0x87,
	0x2e, 0x99, 0xd1, 0xa2, 0xcf, 0x37, 0xed, 0xeb, 0xf6, 0x0b, 0xf1, 0x01, 0x77, 0x13, 0x65, 0x8d,
	0x2b, 0x2f, 0x93, 0xcf, 0x46, 0xbe, 0x1b, 0xd1, 0x47, 0xd9, 0x19, 0xf9, 0x19, 0x07, 0xdd, 0x6b,
	0xb7, 0x7d, 0xc3, 0xf7, 0x77, 0xc8, 0xfd, 0x57, 0x5b, 0xe3, 0x5e, 0x7c, 0x5e, 0xbd, 0xa2, 0x17,
	0x51, 0x1a, 0x4c, 0x28, 0x6c, 0xd3, 0x2e, 0x45, 0x3f, 0x2e, 0xa4, 0x61, 0x98, 0x4a, 0x1c, 0xeb,
	0xca, 0x55, 0x63, 0xfd, 0x47, 0x81, 0xb1, 0xc6, 0xc1, 0x11, 0x23, 0xcb, 0x43, 0x22, 0x45, 0x5d,
	0x4d, 0x8c, 0xb6, 0x44, 0xfa, 0xca, 0xc8, 0xda, 0x01, 0x77, 0xe6, 0x2d, 0x59, 0x1d, 0xe5, 0x0c,
	0x45, 0x0f, 0xf4, 0x53, 0xf1, 0xe2, 0xee, 0xdf, 0xaf, 0x95, 0xf5, 0x5e, 0x3d, 0xc5, 0x47, 0x67,
	0x1b, 0x62, 0x54, 0xed, 0x13, 0x72, 0xe1, 0x88, 0xfc, 0xad, 0xc0, 0x54, 0x74, 0x1b, 0xbc, 0x34,
	0x83, 0x23, 0x1e, 0xc9, 0xfb, 0x79, 0x9d, 0x9f, 0xf0, 0xd8, 0x1e, 0xd7, 0x1e, 0x8e, 0x4c, 0x74,
	0xe6, 0x4f, 0x85, 0x16, 0xdd, 0x95, 0xbc, 0xb8, 0x6f, 0x60, 0xe2, 0x95, 0x65, 0x53, 0x16, 0x06,
	0x64, 0x48, 0x31, 0x87, 0x16, 0xf9, 0x01, 0xd7, 0x9e, 0x27, 0x73, 0x72, 0x32, 0x02, 0x41, 0xb6,
	0xa5, 0xee, 0x17, 0x53, 0xb7, 0x5b, 0x05, 0x6e, 0xb9, 0xfe, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x7f, 0xc8, 0x3e, 0x80, 0xb4, 0x0d, 0x00, 0x00,
}