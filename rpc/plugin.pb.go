// Code generated by protoc-gen-go. DO NOT EDIT.
// source: plugin.proto

package rpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type ChannelMessage struct {
	Channel              string   `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Source               string   `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelMessage) Reset()         { *m = ChannelMessage{} }
func (m *ChannelMessage) String() string { return proto.CompactTextString(m) }
func (*ChannelMessage) ProtoMessage()    {}
func (*ChannelMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{0}
}

func (m *ChannelMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelMessage.Unmarshal(m, b)
}
func (m *ChannelMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelMessage.Marshal(b, m, deterministic)
}
func (m *ChannelMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelMessage.Merge(m, src)
}
func (m *ChannelMessage) XXX_Size() int {
	return xxx_messageInfo_ChannelMessage.Size(m)
}
func (m *ChannelMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelMessage proto.InternalMessageInfo

func (m *ChannelMessage) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *ChannelMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ChannelMessage) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type RawMessage struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RawMessage) Reset()         { *m = RawMessage{} }
func (m *RawMessage) String() string { return proto.CompactTextString(m) }
func (*RawMessage) ProtoMessage()    {}
func (*RawMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{1}
}

func (m *RawMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawMessage.Unmarshal(m, b)
}
func (m *RawMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawMessage.Marshal(b, m, deterministic)
}
func (m *RawMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawMessage.Merge(m, src)
}
func (m *RawMessage) XXX_Size() int {
	return xxx_messageInfo_RawMessage.Size(m)
}
func (m *RawMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_RawMessage.DiscardUnknown(m)
}

var xxx_messageInfo_RawMessage proto.InternalMessageInfo

func (m *RawMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Error struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{2}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Channel struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Channel) Reset()         { *m = Channel{} }
func (m *Channel) String() string { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()    {}
func (*Channel) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{3}
}

func (m *Channel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Channel.Unmarshal(m, b)
}
func (m *Channel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Channel.Marshal(b, m, deterministic)
}
func (m *Channel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Channel.Merge(m, src)
}
func (m *Channel) XXX_Size() int {
	return xxx_messageInfo_Channel.Size(m)
}
func (m *Channel) XXX_DiscardUnknown() {
	xxx_messageInfo_Channel.DiscardUnknown(m)
}

var xxx_messageInfo_Channel proto.InternalMessageInfo

func (m *Channel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{4}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Route struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{5}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

type HttpRequest struct {
	Header               []*HttpHeader `protobuf:"bytes,1,rep,name=header,proto3" json:"header,omitempty"`
	Body                 []byte        `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Path                 string        `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Method               string        `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *HttpRequest) Reset()         { *m = HttpRequest{} }
func (m *HttpRequest) String() string { return proto.CompactTextString(m) }
func (*HttpRequest) ProtoMessage()    {}
func (*HttpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{6}
}

func (m *HttpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpRequest.Unmarshal(m, b)
}
func (m *HttpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpRequest.Marshal(b, m, deterministic)
}
func (m *HttpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpRequest.Merge(m, src)
}
func (m *HttpRequest) XXX_Size() int {
	return xxx_messageInfo_HttpRequest.Size(m)
}
func (m *HttpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HttpRequest proto.InternalMessageInfo

func (m *HttpRequest) GetHeader() []*HttpHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HttpRequest) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *HttpRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HttpRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

type HttpResponse struct {
	Header               []*HttpHeader `protobuf:"bytes,1,rep,name=header,proto3" json:"header,omitempty"`
	Body                 []byte        `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Status               int32         `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *HttpResponse) Reset()         { *m = HttpResponse{} }
func (m *HttpResponse) String() string { return proto.CompactTextString(m) }
func (*HttpResponse) ProtoMessage()    {}
func (*HttpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{7}
}

func (m *HttpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpResponse.Unmarshal(m, b)
}
func (m *HttpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpResponse.Marshal(b, m, deterministic)
}
func (m *HttpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpResponse.Merge(m, src)
}
func (m *HttpResponse) XXX_Size() int {
	return xxx_messageInfo_HttpResponse.Size(m)
}
func (m *HttpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HttpResponse proto.InternalMessageInfo

func (m *HttpResponse) GetHeader() []*HttpHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *HttpResponse) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *HttpResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type HttpHeader struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpHeader) Reset()         { *m = HttpHeader{} }
func (m *HttpHeader) String() string { return proto.CompactTextString(m) }
func (*HttpHeader) ProtoMessage()    {}
func (*HttpHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_22a625af4bc1cc87, []int{8}
}

func (m *HttpHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpHeader.Unmarshal(m, b)
}
func (m *HttpHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpHeader.Marshal(b, m, deterministic)
}
func (m *HttpHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpHeader.Merge(m, src)
}
func (m *HttpHeader) XXX_Size() int {
	return xxx_messageInfo_HttpHeader.Size(m)
}
func (m *HttpHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpHeader.DiscardUnknown(m)
}

var xxx_messageInfo_HttpHeader proto.InternalMessageInfo

func (m *HttpHeader) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *HttpHeader) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*ChannelMessage)(nil), "rpc.ChannelMessage")
	proto.RegisterType((*RawMessage)(nil), "rpc.RawMessage")
	proto.RegisterType((*Error)(nil), "rpc.Error")
	proto.RegisterType((*Channel)(nil), "rpc.Channel")
	proto.RegisterType((*Empty)(nil), "rpc.Empty")
	proto.RegisterType((*Route)(nil), "rpc.Route")
	proto.RegisterType((*HttpRequest)(nil), "rpc.HttpRequest")
	proto.RegisterType((*HttpResponse)(nil), "rpc.HttpResponse")
	proto.RegisterType((*HttpHeader)(nil), "rpc.HttpHeader")
}

func init() {
	proto.RegisterFile("plugin.proto", fileDescriptor_22a625af4bc1cc87)
}

var fileDescriptor_22a625af4bc1cc87 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x8d, 0x49, 0xd3, 0x6a, 0xa7, 0xd1, 0xb2, 0x18, 0x84, 0xa2, 0x4a, 0x88, 0xe2, 0x03, 0xf4,
	0x54, 0x56, 0x65, 0xa5, 0xbd, 0xb3, 0xaa, 0x54, 0x0e, 0x48, 0x95, 0xd5, 0x13, 0xe2, 0xe2, 0x26,
	0x26, 0xa9, 0x68, 0x62, 0x63, 0x3b, 0x85, 0xfe, 0x48, 0xfe, 0x13, 0xf2, 0x47, 0x68, 0x5a, 0x21,
	0x2e, 0x7b, 0x9b, 0xf7, 0xe6, 0x8d, 0x9f, 0xc7, 0x33, 0x86, 0x54, 0xee, 0xdb, 0x72, 0xd7, 0xcc,
	0xa5, 0x12, 0x46, 0xe0, 0x58, 0xc9, 0x9c, 0x7c, 0x85, 0xeb, 0x87, 0x8a, 0x35, 0x0d, 0xdf, 0x7f,
	0xe6, 0x5a, 0xb3, 0x92, 0xe3, 0x0c, 0x46, 0xb9, 0x67, 0x32, 0x34, 0x45, 0xb3, 0x2b, 0xda, 0x41,
	0x9b, 0xa9, 0xbd, 0x28, 0x7b, 0xe2, 0x33, 0x01, 0xe2, 0x97, 0x30, 0xd4, 0xa2, 0x55, 0x39, 0xcf,
	0x62, 0x97, 0x08, 0x88, 0xbc, 0x05, 0xa0, 0xec, 0x67, 0xef, 0xe4, 0xae, 0x1e, 0x9d, 0xd5, 0x93,
	0x37, 0x90, 0x2c, 0x95, 0x12, 0xea, 0x3f, 0x92, 0x57, 0x30, 0x0a, 0x17, 0xc5, 0x18, 0x06, 0x0d,
	0xab, 0x3b, 0x85, 0x8b, 0xc9, 0x08, 0x92, 0x65, 0x2d, 0xcd, 0x91, 0xbc, 0x86, 0x84, 0x8a, 0xd6,
	0xb8, 0x3b, 0x49, 0xc5, 0xbf, 0xed, 0x7e, 0x05, 0x5d, 0x40, 0xe4, 0x00, 0xe3, 0x95, 0x31, 0x92,
	0xf2, 0x1f, 0x2d, 0xd7, 0x06, 0xbf, 0x83, 0x61, 0xc5, 0x59, 0xc1, 0x55, 0x86, 0xa6, 0xf1, 0x6c,
	0xbc, 0x78, 0x3a, 0x57, 0x32, 0x9f, 0x5b, 0xc5, 0xca, 0xd1, 0x34, 0xa4, 0xad, 0xeb, 0x56, 0x14,
	0x47, 0xd7, 0x7a, 0x4a, 0x5d, 0x6c, 0x39, 0xc9, 0x4c, 0x15, 0xba, 0x76, 0xb1, 0xf5, 0xad, 0xb9,
	0xa9, 0x44, 0x91, 0x0d, 0xbc, 0xaf, 0x47, 0x24, 0x87, 0xd4, 0xfb, 0x6a, 0x29, 0x1a, 0xcd, 0x1f,
	0x67, 0x6c, 0x1f, 0xdc, 0x30, 0xd3, 0x6a, 0x67, 0x9d, 0xd0, 0x80, 0xc8, 0x1d, 0xc0, 0xe9, 0x04,
	0x7c, 0x03, 0xf1, 0x77, 0x7e, 0x0c, 0xfd, 0xdb, 0x10, 0xbf, 0x80, 0xe4, 0xc0, 0xf6, 0x6d, 0x37,
	0x40, 0x0f, 0x16, 0xbf, 0x11, 0x5c, 0x7d, 0xa2, 0x0f, 0x6b, 0xb7, 0x1d, 0x78, 0x0a, 0x03, 0xb9,
	0x6b, 0x4a, 0x0c, 0xee, 0x42, 0xee, 0x55, 0x27, 0xbd, 0x98, 0x44, 0xf8, 0x1e, 0xb0, 0xe6, 0x4d,
	0x71, 0xb1, 0x38, 0xcf, 0x9d, 0xe6, 0x9c, 0xec, 0x0a, 0xed, 0x70, 0x49, 0x84, 0xdf, 0xc3, 0xb5,
	0x2d, 0xec, 0xed, 0x84, 0xef, 0xfa, 0x44, 0x5c, 0x14, 0xdc, 0xc1, 0xb8, 0xe4, 0x26, 0xe4, 0x34,
	0x4e, 0xfb, 0x16, 0x93, 0x7f, 0x19, 0x92, 0xe8, 0x16, 0x2d, 0x96, 0x00, 0xab, 0xcd, 0x66, 0x1d,
	0xfa, 0xb9, 0x07, 0x28, 0xb9, 0xe9, 0xe6, 0xfd, 0xec, 0xef, 0x33, 0x77, 0x93, 0x98, 0xdc, 0xf4,
	0x28, 0x27, 0x22, 0xd1, 0x0c, 0xdd, 0xa2, 0x8f, 0xc9, 0x17, 0xfb, 0x45, 0xb6, 0x43, 0xf7, 0x5d,
	0x3e, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x60, 0xa7, 0xbf, 0x3e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// IRCPluginClient is the client API for IRCPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IRCPluginClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	SendChannelMessage(ctx context.Context, in *ChannelMessage, opts ...grpc.CallOption) (*Error, error)
	SendRawMessage(ctx context.Context, in *RawMessage, opts ...grpc.CallOption) (*Error, error)
	GetMessages(ctx context.Context, in *Channel, opts ...grpc.CallOption) (IRCPlugin_GetMessagesClient, error)
}

type iRCPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewIRCPluginClient(cc grpc.ClientConnInterface) IRCPluginClient {
	return &iRCPluginClient{cc}
}

func (c *iRCPluginClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/rpc.IRCPlugin/ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iRCPluginClient) SendChannelMessage(ctx context.Context, in *ChannelMessage, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/rpc.IRCPlugin/sendChannelMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iRCPluginClient) SendRawMessage(ctx context.Context, in *RawMessage, opts ...grpc.CallOption) (*Error, error) {
	out := new(Error)
	err := c.cc.Invoke(ctx, "/rpc.IRCPlugin/sendRawMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iRCPluginClient) GetMessages(ctx context.Context, in *Channel, opts ...grpc.CallOption) (IRCPlugin_GetMessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_IRCPlugin_serviceDesc.Streams[0], "/rpc.IRCPlugin/getMessages", opts...)
	if err != nil {
		return nil, err
	}
	x := &iRCPluginGetMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type IRCPlugin_GetMessagesClient interface {
	Recv() (*ChannelMessage, error)
	grpc.ClientStream
}

type iRCPluginGetMessagesClient struct {
	grpc.ClientStream
}

func (x *iRCPluginGetMessagesClient) Recv() (*ChannelMessage, error) {
	m := new(ChannelMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// IRCPluginServer is the server API for IRCPlugin service.
type IRCPluginServer interface {
	Ping(context.Context, *Empty) (*Empty, error)
	SendChannelMessage(context.Context, *ChannelMessage) (*Error, error)
	SendRawMessage(context.Context, *RawMessage) (*Error, error)
	GetMessages(*Channel, IRCPlugin_GetMessagesServer) error
}

// UnimplementedIRCPluginServer can be embedded to have forward compatible implementations.
type UnimplementedIRCPluginServer struct {
}

func (*UnimplementedIRCPluginServer) Ping(ctx context.Context, req *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedIRCPluginServer) SendChannelMessage(ctx context.Context, req *ChannelMessage) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendChannelMessage not implemented")
}
func (*UnimplementedIRCPluginServer) SendRawMessage(ctx context.Context, req *RawMessage) (*Error, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRawMessage not implemented")
}
func (*UnimplementedIRCPluginServer) GetMessages(req *Channel, srv IRCPlugin_GetMessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMessages not implemented")
}

func RegisterIRCPluginServer(s *grpc.Server, srv IRCPluginServer) {
	s.RegisterService(&_IRCPlugin_serviceDesc, srv)
}

func _IRCPlugin_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IRCPluginServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.IRCPlugin/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IRCPluginServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _IRCPlugin_SendChannelMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IRCPluginServer).SendChannelMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.IRCPlugin/SendChannelMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IRCPluginServer).SendChannelMessage(ctx, req.(*ChannelMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _IRCPlugin_SendRawMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IRCPluginServer).SendRawMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.IRCPlugin/SendRawMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IRCPluginServer).SendRawMessage(ctx, req.(*RawMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _IRCPlugin_GetMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Channel)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(IRCPluginServer).GetMessages(m, &iRCPluginGetMessagesServer{stream})
}

type IRCPlugin_GetMessagesServer interface {
	Send(*ChannelMessage) error
	grpc.ServerStream
}

type iRCPluginGetMessagesServer struct {
	grpc.ServerStream
}

func (x *iRCPluginGetMessagesServer) Send(m *ChannelMessage) error {
	return x.ServerStream.SendMsg(m)
}

var _IRCPlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.IRCPlugin",
	HandlerType: (*IRCPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ping",
			Handler:    _IRCPlugin_Ping_Handler,
		},
		{
			MethodName: "sendChannelMessage",
			Handler:    _IRCPlugin_SendChannelMessage_Handler,
		},
		{
			MethodName: "sendRawMessage",
			Handler:    _IRCPlugin_SendRawMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getMessages",
			Handler:       _IRCPlugin_GetMessages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "plugin.proto",
}

// HTTPPluginClient is the client API for HTTPPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HTTPPluginClient interface {
	GetRequest(ctx context.Context, opts ...grpc.CallOption) (HTTPPlugin_GetRequestClient, error)
}

type hTTPPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewHTTPPluginClient(cc grpc.ClientConnInterface) HTTPPluginClient {
	return &hTTPPluginClient{cc}
}

func (c *hTTPPluginClient) GetRequest(ctx context.Context, opts ...grpc.CallOption) (HTTPPlugin_GetRequestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_HTTPPlugin_serviceDesc.Streams[0], "/rpc.HTTPPlugin/getRequest", opts...)
	if err != nil {
		return nil, err
	}
	x := &hTTPPluginGetRequestClient{stream}
	return x, nil
}

type HTTPPlugin_GetRequestClient interface {
	Send(*HttpResponse) error
	Recv() (*HttpRequest, error)
	grpc.ClientStream
}

type hTTPPluginGetRequestClient struct {
	grpc.ClientStream
}

func (x *hTTPPluginGetRequestClient) Send(m *HttpResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *hTTPPluginGetRequestClient) Recv() (*HttpRequest, error) {
	m := new(HttpRequest)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HTTPPluginServer is the server API for HTTPPlugin service.
type HTTPPluginServer interface {
	GetRequest(HTTPPlugin_GetRequestServer) error
}

// UnimplementedHTTPPluginServer can be embedded to have forward compatible implementations.
type UnimplementedHTTPPluginServer struct {
}

func (*UnimplementedHTTPPluginServer) GetRequest(srv HTTPPlugin_GetRequestServer) error {
	return status.Errorf(codes.Unimplemented, "method GetRequest not implemented")
}

func RegisterHTTPPluginServer(s *grpc.Server, srv HTTPPluginServer) {
	s.RegisterService(&_HTTPPlugin_serviceDesc, srv)
}

func _HTTPPlugin_GetRequest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HTTPPluginServer).GetRequest(&hTTPPluginGetRequestServer{stream})
}

type HTTPPlugin_GetRequestServer interface {
	Send(*HttpRequest) error
	Recv() (*HttpResponse, error)
	grpc.ServerStream
}

type hTTPPluginGetRequestServer struct {
	grpc.ServerStream
}

func (x *hTTPPluginGetRequestServer) Send(m *HttpRequest) error {
	return x.ServerStream.SendMsg(m)
}

func (x *hTTPPluginGetRequestServer) Recv() (*HttpResponse, error) {
	m := new(HttpResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _HTTPPlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.HTTPPlugin",
	HandlerType: (*HTTPPluginServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getRequest",
			Handler:       _HTTPPlugin_GetRequest_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "plugin.proto",
}
