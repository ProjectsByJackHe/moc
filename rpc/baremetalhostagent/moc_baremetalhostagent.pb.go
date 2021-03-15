// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_baremetalhostagent.proto

package baremetalhostagent

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
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

type BareMetalHostOperatingSystemConfiguration struct {
	ComputerName         string   `protobuf:"bytes,1,opt,name=computerName,proto3" json:"computerName,omitempty"`
	CustomData           string   `protobuf:"bytes,2,opt,name=customData,proto3" json:"customData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BareMetalHostOperatingSystemConfiguration) Reset() {
	*m = BareMetalHostOperatingSystemConfiguration{}
}
func (m *BareMetalHostOperatingSystemConfiguration) String() string {
	return proto.CompactTextString(m)
}
func (*BareMetalHostOperatingSystemConfiguration) ProtoMessage() {}
func (*BareMetalHostOperatingSystemConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_254a3462c20ef4c7, []int{0}
}

func (m *BareMetalHostOperatingSystemConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalHostOperatingSystemConfiguration.Unmarshal(m, b)
}
func (m *BareMetalHostOperatingSystemConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalHostOperatingSystemConfiguration.Marshal(b, m, deterministic)
}
func (m *BareMetalHostOperatingSystemConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalHostOperatingSystemConfiguration.Merge(m, src)
}
func (m *BareMetalHostOperatingSystemConfiguration) XXX_Size() int {
	return xxx_messageInfo_BareMetalHostOperatingSystemConfiguration.Size(m)
}
func (m *BareMetalHostOperatingSystemConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalHostOperatingSystemConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalHostOperatingSystemConfiguration proto.InternalMessageInfo

func (m *BareMetalHostOperatingSystemConfiguration) GetComputerName() string {
	if m != nil {
		return m.ComputerName
	}
	return ""
}

func (m *BareMetalHostOperatingSystemConfiguration) GetCustomData() string {
	if m != nil {
		return m.CustomData
	}
	return ""
}

type BareMetalHost struct {
	Name                 string                                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string                                     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Os                   *BareMetalHostOperatingSystemConfiguration `protobuf:"bytes,3,opt,name=os,proto3" json:"os,omitempty"`
	Status               *common.Status                             `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Entity               *common.Entity                             `protobuf:"bytes,5,opt,name=entity,proto3" json:"entity,omitempty"`
	Tags                 *common.Tags                               `protobuf:"bytes,6,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *BareMetalHost) Reset()         { *m = BareMetalHost{} }
func (m *BareMetalHost) String() string { return proto.CompactTextString(m) }
func (*BareMetalHost) ProtoMessage()    {}
func (*BareMetalHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_254a3462c20ef4c7, []int{1}
}

func (m *BareMetalHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalHost.Unmarshal(m, b)
}
func (m *BareMetalHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalHost.Marshal(b, m, deterministic)
}
func (m *BareMetalHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalHost.Merge(m, src)
}
func (m *BareMetalHost) XXX_Size() int {
	return xxx_messageInfo_BareMetalHost.Size(m)
}
func (m *BareMetalHost) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalHost.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalHost proto.InternalMessageInfo

func (m *BareMetalHost) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BareMetalHost) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BareMetalHost) GetOs() *BareMetalHostOperatingSystemConfiguration {
	if m != nil {
		return m.Os
	}
	return nil
}

func (m *BareMetalHost) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *BareMetalHost) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *BareMetalHost) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

type BareMetalHostRequest struct {
	BareMetalHost        *BareMetalHost `protobuf:"bytes,1,opt,name=bareMetalHost,proto3" json:"bareMetalHost,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *BareMetalHostRequest) Reset()         { *m = BareMetalHostRequest{} }
func (m *BareMetalHostRequest) String() string { return proto.CompactTextString(m) }
func (*BareMetalHostRequest) ProtoMessage()    {}
func (*BareMetalHostRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_254a3462c20ef4c7, []int{2}
}

func (m *BareMetalHostRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalHostRequest.Unmarshal(m, b)
}
func (m *BareMetalHostRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalHostRequest.Marshal(b, m, deterministic)
}
func (m *BareMetalHostRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalHostRequest.Merge(m, src)
}
func (m *BareMetalHostRequest) XXX_Size() int {
	return xxx_messageInfo_BareMetalHostRequest.Size(m)
}
func (m *BareMetalHostRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalHostRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalHostRequest proto.InternalMessageInfo

func (m *BareMetalHostRequest) GetBareMetalHost() *BareMetalHost {
	if m != nil {
		return m.BareMetalHost
	}
	return nil
}

type BareMetalHostResponse struct {
	BareMetalHost        *BareMetalHost      `protobuf:"bytes,1,opt,name=bareMetalHost,proto3" json:"bareMetalHost,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *BareMetalHostResponse) Reset()         { *m = BareMetalHostResponse{} }
func (m *BareMetalHostResponse) String() string { return proto.CompactTextString(m) }
func (*BareMetalHostResponse) ProtoMessage()    {}
func (*BareMetalHostResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_254a3462c20ef4c7, []int{3}
}

func (m *BareMetalHostResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BareMetalHostResponse.Unmarshal(m, b)
}
func (m *BareMetalHostResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BareMetalHostResponse.Marshal(b, m, deterministic)
}
func (m *BareMetalHostResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BareMetalHostResponse.Merge(m, src)
}
func (m *BareMetalHostResponse) XXX_Size() int {
	return xxx_messageInfo_BareMetalHostResponse.Size(m)
}
func (m *BareMetalHostResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BareMetalHostResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BareMetalHostResponse proto.InternalMessageInfo

func (m *BareMetalHostResponse) GetBareMetalHost() *BareMetalHost {
	if m != nil {
		return m.BareMetalHost
	}
	return nil
}

func (m *BareMetalHostResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *BareMetalHostResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*BareMetalHostOperatingSystemConfiguration)(nil), "moc.baremetalhostagent.BareMetalHostOperatingSystemConfiguration")
	proto.RegisterType((*BareMetalHost)(nil), "moc.baremetalhostagent.BareMetalHost")
	proto.RegisterType((*BareMetalHostRequest)(nil), "moc.baremetalhostagent.BareMetalHostRequest")
	proto.RegisterType((*BareMetalHostResponse)(nil), "moc.baremetalhostagent.BareMetalHostResponse")
}

func init() { proto.RegisterFile("moc_baremetalhostagent.proto", fileDescriptor_254a3462c20ef4c7) }

var fileDescriptor_254a3462c20ef4c7 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xd1, 0x8a, 0xd4, 0x30,
	0x14, 0xb5, 0x75, 0xb6, 0x30, 0x77, 0x5c, 0x1f, 0xc2, 0xaa, 0x65, 0xd0, 0x61, 0xa9, 0x08, 0x2b,
	0x68, 0x8b, 0xf5, 0x0b, 0x76, 0x74, 0x41, 0x10, 0x15, 0xb3, 0xea, 0x83, 0x2f, 0x92, 0x66, 0x32,
	0xd9, 0x42, 0xd3, 0x5b, 0x93, 0x1b, 0x64, 0x1f, 0xfc, 0x25, 0x7f, 0xcd, 0x5f, 0x90, 0xa6, 0x15,
	0xa7, 0x8e, 0x0f, 0xf3, 0xe0, 0x53, 0x9b, 0x73, 0xce, 0x3d, 0x27, 0xf7, 0xe6, 0xc2, 0x7d, 0x83,
	0xf2, 0x4b, 0x25, 0xac, 0x32, 0x8a, 0x44, 0x73, 0x85, 0x8e, 0x84, 0x56, 0x2d, 0xe5, 0x9d, 0x45,
	0x42, 0x76, 0xd7, 0xa0, 0xcc, 0xf7, 0xd9, 0xe5, 0x4a, 0x23, 0xea, 0x46, 0x15, 0x41, 0x55, 0xf9,
	0x6d, 0xf1, 0xcd, 0x8a, 0xae, 0x53, 0xd6, 0x0d, 0x75, 0xcb, 0x7b, 0xbd, 0xab, 0x44, 0x63, 0xb0,
	0x1d, 0x3f, 0x23, 0xb1, 0x9a, 0x12, 0x9d, 0x27, 0xb5, 0xcb, 0x67, 0x08, 0x8f, 0xd7, 0xc2, 0xaa,
	0x37, 0x7d, 0xdc, 0x2b, 0x74, 0xf4, 0xae, 0x53, 0x56, 0x50, 0xdd, 0xea, 0xcb, 0x6b, 0x47, 0xca,
	0xbc, 0xc0, 0x76, 0x5b, 0x6b, 0xdf, 0x43, 0xd8, 0xb2, 0x0c, 0x6e, 0x8d, 0x1e, 0xf6, 0xad, 0x30,
	0x2a, 0x8d, 0x4e, 0xa3, 0xb3, 0x39, 0x9f, 0x60, 0x6c, 0x05, 0x20, 0xbd, 0x23, 0x34, 0x2f, 0x05,
	0x89, 0x34, 0x0e, 0x8a, 0x1d, 0x24, 0xfb, 0x19, 0xc1, 0xf1, 0x24, 0x91, 0x31, 0x98, 0xb5, 0x7f,
	0xdc, 0xc2, 0x3f, 0xbb, 0x0d, 0x71, 0xbd, 0x19, 0xab, 0xe3, 0x7a, 0xc3, 0xde, 0x43, 0x8c, 0x2e,
	0xbd, 0x79, 0x1a, 0x9d, 0x2d, 0xca, 0xf3, 0xfc, 0xdf, 0x43, 0xca, 0x0f, 0x6e, 0x84, 0xc7, 0xe8,
	0xd8, 0x43, 0x48, 0x1c, 0x09, 0xf2, 0x2e, 0x9d, 0x05, 0xdb, 0x45, 0xb0, 0xbd, 0x0c, 0x10, 0x1f,
	0xa9, 0x5e, 0xa4, 0x5a, 0xaa, 0xe9, 0x3a, 0x3d, 0xda, 0x11, 0x5d, 0x04, 0x88, 0x8f, 0x14, 0x7b,
	0x00, 0x33, 0x12, 0xda, 0xa5, 0x49, 0x90, 0xcc, 0x83, 0xe4, 0x83, 0xd0, 0x8e, 0x07, 0x38, 0x93,
	0x70, 0x32, 0xb9, 0x19, 0x57, 0x5f, 0xbd, 0x72, 0xc4, 0x5e, 0xc3, 0x71, 0xb5, 0x8b, 0x87, 0x01,
	0x2c, 0xca, 0x47, 0x07, 0xb5, 0xc7, 0xa7, 0xb5, 0xd9, 0x8f, 0x08, 0xee, 0xfc, 0x95, 0xe2, 0x3a,
	0x6c, 0x9d, 0xfa, 0xaf, 0x31, 0xac, 0x84, 0x84, 0x2b, 0xe7, 0x1b, 0x0a, 0x6f, 0xb3, 0x28, 0x97,
	0xf9, 0xb0, 0x98, 0xf9, 0xef, 0xc5, 0xcc, 0xd7, 0x88, 0xcd, 0x27, 0xd1, 0x78, 0xc5, 0x47, 0x25,
	0x3b, 0x81, 0xa3, 0x0b, 0x6b, 0xd1, 0x86, 0xe7, 0x9b, 0xf3, 0xe1, 0x50, 0x7e, 0x07, 0x36, 0x49,
	0x3a, 0xef, 0xc3, 0x99, 0x86, 0xe4, 0x63, 0xb7, 0x11, 0xa4, 0xd8, 0x93, 0xc3, 0xee, 0x37, 0xcc,
	0x72, 0xf9, 0xf4, 0x40, 0xf5, 0x30, 0x93, 0xec, 0xc6, 0xfa, 0xd9, 0xe7, 0x42, 0xd7, 0x74, 0xe5,
	0xab, 0x5c, 0xa2, 0x29, 0x4c, 0x2d, 0x2d, 0x3a, 0xdc, 0x52, 0x61, 0x50, 0x16, 0xb6, 0x93, 0xc5,
	0xbe, 0x55, 0x95, 0x84, 0x1e, 0x9f, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x3d, 0xae, 0x71, 0xda,
	0xc2, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BareMetalHostAgentClient is the client API for BareMetalHostAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BareMetalHostAgentClient interface {
	Update(ctx context.Context, in *BareMetalHostRequest, opts ...grpc.CallOption) (*BareMetalHostResponse, error)
}

type bareMetalHostAgentClient struct {
	cc *grpc.ClientConn
}

func NewBareMetalHostAgentClient(cc *grpc.ClientConn) BareMetalHostAgentClient {
	return &bareMetalHostAgentClient{cc}
}

func (c *bareMetalHostAgentClient) Update(ctx context.Context, in *BareMetalHostRequest, opts ...grpc.CallOption) (*BareMetalHostResponse, error) {
	out := new(BareMetalHostResponse)
	err := c.cc.Invoke(ctx, "/moc.baremetalhostagent.BareMetalHostAgent/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BareMetalHostAgentServer is the server API for BareMetalHostAgent service.
type BareMetalHostAgentServer interface {
	Update(context.Context, *BareMetalHostRequest) (*BareMetalHostResponse, error)
}

// UnimplementedBareMetalHostAgentServer can be embedded to have forward compatible implementations.
type UnimplementedBareMetalHostAgentServer struct {
}

func (*UnimplementedBareMetalHostAgentServer) Update(ctx context.Context, req *BareMetalHostRequest) (*BareMetalHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterBareMetalHostAgentServer(s *grpc.Server, srv BareMetalHostAgentServer) {
	s.RegisterService(&_BareMetalHostAgent_serviceDesc, srv)
}

func _BareMetalHostAgent_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BareMetalHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BareMetalHostAgentServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.baremetalhostagent.BareMetalHostAgent/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BareMetalHostAgentServer).Update(ctx, req.(*BareMetalHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BareMetalHostAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.baremetalhostagent.BareMetalHostAgent",
	HandlerType: (*BareMetalHostAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Update",
			Handler:    _BareMetalHostAgent_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_baremetalhostagent.proto",
}