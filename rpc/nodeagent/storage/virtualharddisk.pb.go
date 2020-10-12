// Code generated by protoc-gen-go. DO NOT EDIT.
// source: virtualharddisk.proto

package storage

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type VirtualHardDiskType int32

const (
	VirtualHardDiskType_OS_VIRTUALHARDDISK       VirtualHardDiskType = 0
	VirtualHardDiskType_DATADISK_VIRTUALHARDDISK VirtualHardDiskType = 1
)

var VirtualHardDiskType_name = map[int32]string{
	0: "OS_VIRTUALHARDDISK",
	1: "DATADISK_VIRTUALHARDDISK",
}

var VirtualHardDiskType_value = map[string]int32{
	"OS_VIRTUALHARDDISK":       0,
	"DATADISK_VIRTUALHARDDISK": 1,
}

func (x VirtualHardDiskType) String() string {
	return proto.EnumName(VirtualHardDiskType_name, int32(x))
}

func (VirtualHardDiskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f4b382f86170a6e5, []int{0}
}

type VirtualHardDiskRequest struct {
	VirtualHardDiskSystems []*VirtualHardDisk `protobuf:"bytes,1,rep,name=VirtualHardDiskSystems,proto3" json:"VirtualHardDiskSystems,omitempty"`
	OperationType          common.Operation   `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}           `json:"-"`
	XXX_unrecognized       []byte             `json:"-"`
	XXX_sizecache          int32              `json:"-"`
}

func (m *VirtualHardDiskRequest) Reset()         { *m = VirtualHardDiskRequest{} }
func (m *VirtualHardDiskRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDiskRequest) ProtoMessage()    {}
func (*VirtualHardDiskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b382f86170a6e5, []int{0}
}

func (m *VirtualHardDiskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHardDiskRequest.Unmarshal(m, b)
}
func (m *VirtualHardDiskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHardDiskRequest.Marshal(b, m, deterministic)
}
func (m *VirtualHardDiskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHardDiskRequest.Merge(m, src)
}
func (m *VirtualHardDiskRequest) XXX_Size() int {
	return xxx_messageInfo_VirtualHardDiskRequest.Size(m)
}
func (m *VirtualHardDiskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHardDiskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHardDiskRequest proto.InternalMessageInfo

func (m *VirtualHardDiskRequest) GetVirtualHardDiskSystems() []*VirtualHardDisk {
	if m != nil {
		return m.VirtualHardDiskSystems
	}
	return nil
}

func (m *VirtualHardDiskRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type VirtualHardDiskResponse struct {
	VirtualHardDiskSystems []*VirtualHardDisk  `protobuf:"bytes,1,rep,name=VirtualHardDiskSystems,proto3" json:"VirtualHardDiskSystems,omitempty"`
	Result                 *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                  string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}            `json:"-"`
	XXX_unrecognized       []byte              `json:"-"`
	XXX_sizecache          int32               `json:"-"`
}

func (m *VirtualHardDiskResponse) Reset()         { *m = VirtualHardDiskResponse{} }
func (m *VirtualHardDiskResponse) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDiskResponse) ProtoMessage()    {}
func (*VirtualHardDiskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b382f86170a6e5, []int{1}
}

func (m *VirtualHardDiskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHardDiskResponse.Unmarshal(m, b)
}
func (m *VirtualHardDiskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHardDiskResponse.Marshal(b, m, deterministic)
}
func (m *VirtualHardDiskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHardDiskResponse.Merge(m, src)
}
func (m *VirtualHardDiskResponse) XXX_Size() int {
	return xxx_messageInfo_VirtualHardDiskResponse.Size(m)
}
func (m *VirtualHardDiskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHardDiskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHardDiskResponse proto.InternalMessageInfo

func (m *VirtualHardDiskResponse) GetVirtualHardDiskSystems() []*VirtualHardDisk {
	if m != nil {
		return m.VirtualHardDiskSystems
	}
	return nil
}

func (m *VirtualHardDiskResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *VirtualHardDiskResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type VirtualHardDisk struct {
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Path   string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	// Storage container name to hold this vhd
	ContainerName        string              `protobuf:"bytes,5,opt,name=containerName,proto3" json:"containerName,omitempty"`
	Status               *common.Status      `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Size                 int64               `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	Dynamic              bool                `protobuf:"varint,8,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
	Blocksizebytes       int32               `protobuf:"varint,9,opt,name=blocksizebytes,proto3" json:"blocksizebytes,omitempty"`
	Logicalsectorbytes   int32               `protobuf:"varint,10,opt,name=logicalsectorbytes,proto3" json:"logicalsectorbytes,omitempty"`
	Physicalsectorbytes  int32               `protobuf:"varint,11,opt,name=physicalsectorbytes,proto3" json:"physicalsectorbytes,omitempty"`
	Controllernumber     int64               `protobuf:"varint,12,opt,name=controllernumber,proto3" json:"controllernumber,omitempty"`
	Controllerlocation   int64               `protobuf:"varint,13,opt,name=controllerlocation,proto3" json:"controllerlocation,omitempty"`
	Disknumber           int64               `protobuf:"varint,14,opt,name=disknumber,proto3" json:"disknumber,omitempty"`
	VirtualmachineName   string              `protobuf:"bytes,15,opt,name=virtualmachineName,proto3" json:"virtualmachineName,omitempty"`
	Scsipath             string              `protobuf:"bytes,16,opt,name=scsipath,proto3" json:"scsipath,omitempty"`
	Virtualharddisktype  VirtualHardDiskType `protobuf:"varint,17,opt,name=virtualharddisktype,proto3,enum=moc.nodeagent.storage.VirtualHardDiskType" json:"virtualharddisktype,omitempty"`
	Entity               *common.Entity      `protobuf:"bytes,18,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VirtualHardDisk) Reset()         { *m = VirtualHardDisk{} }
func (m *VirtualHardDisk) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDisk) ProtoMessage()    {}
func (*VirtualHardDisk) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b382f86170a6e5, []int{2}
}

func (m *VirtualHardDisk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHardDisk.Unmarshal(m, b)
}
func (m *VirtualHardDisk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHardDisk.Marshal(b, m, deterministic)
}
func (m *VirtualHardDisk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHardDisk.Merge(m, src)
}
func (m *VirtualHardDisk) XXX_Size() int {
	return xxx_messageInfo_VirtualHardDisk.Size(m)
}
func (m *VirtualHardDisk) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHardDisk.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHardDisk proto.InternalMessageInfo

func (m *VirtualHardDisk) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualHardDisk) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VirtualHardDisk) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *VirtualHardDisk) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *VirtualHardDisk) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *VirtualHardDisk) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VirtualHardDisk) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *VirtualHardDisk) GetDynamic() bool {
	if m != nil {
		return m.Dynamic
	}
	return false
}

func (m *VirtualHardDisk) GetBlocksizebytes() int32 {
	if m != nil {
		return m.Blocksizebytes
	}
	return 0
}

func (m *VirtualHardDisk) GetLogicalsectorbytes() int32 {
	if m != nil {
		return m.Logicalsectorbytes
	}
	return 0
}

func (m *VirtualHardDisk) GetPhysicalsectorbytes() int32 {
	if m != nil {
		return m.Physicalsectorbytes
	}
	return 0
}

func (m *VirtualHardDisk) GetControllernumber() int64 {
	if m != nil {
		return m.Controllernumber
	}
	return 0
}

func (m *VirtualHardDisk) GetControllerlocation() int64 {
	if m != nil {
		return m.Controllerlocation
	}
	return 0
}

func (m *VirtualHardDisk) GetDisknumber() int64 {
	if m != nil {
		return m.Disknumber
	}
	return 0
}

func (m *VirtualHardDisk) GetVirtualmachineName() string {
	if m != nil {
		return m.VirtualmachineName
	}
	return ""
}

func (m *VirtualHardDisk) GetScsipath() string {
	if m != nil {
		return m.Scsipath
	}
	return ""
}

func (m *VirtualHardDisk) GetVirtualharddisktype() VirtualHardDiskType {
	if m != nil {
		return m.Virtualharddisktype
	}
	return VirtualHardDiskType_OS_VIRTUALHARDDISK
}

func (m *VirtualHardDisk) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterEnum("moc.nodeagent.storage.VirtualHardDiskType", VirtualHardDiskType_name, VirtualHardDiskType_value)
	proto.RegisterType((*VirtualHardDiskRequest)(nil), "moc.nodeagent.storage.VirtualHardDiskRequest")
	proto.RegisterType((*VirtualHardDiskResponse)(nil), "moc.nodeagent.storage.VirtualHardDiskResponse")
	proto.RegisterType((*VirtualHardDisk)(nil), "moc.nodeagent.storage.VirtualHardDisk")
}

func init() { proto.RegisterFile("virtualharddisk.proto", fileDescriptor_f4b382f86170a6e5) }

var fileDescriptor_f4b382f86170a6e5 = []byte{
	// 681 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdd, 0x4e, 0x1b, 0x3b,
	0x10, 0x66, 0x81, 0x04, 0x98, 0x40, 0x80, 0x01, 0x72, 0xf6, 0xe4, 0x1c, 0xa1, 0x28, 0xad, 0x50,
	0x84, 0xd4, 0x0d, 0x4a, 0xfb, 0x02, 0xa1, 0x41, 0x02, 0x51, 0x81, 0xe4, 0x50, 0x2e, 0xaa, 0xaa,
	0x95, 0xe3, 0x98, 0xc4, 0xca, 0xee, 0x7a, 0x6b, 0x7b, 0xa9, 0xd2, 0x87, 0x6a, 0x1f, 0xa2, 0x97,
	0x7d, 0xa9, 0xca, 0xde, 0x25, 0x84, 0x24, 0x95, 0xb8, 0xe9, 0x9d, 0x67, 0xbe, 0x6f, 0x3e, 0x7f,
	0xfe, 0x99, 0x81, 0x83, 0x7b, 0xa1, 0x4c, 0x4a, 0xc3, 0x21, 0x55, 0xfd, 0xbe, 0xd0, 0xa3, 0x20,
	0x51, 0xd2, 0x48, 0x3c, 0x88, 0x24, 0x0b, 0x62, 0xd9, 0xe7, 0x74, 0xc0, 0x63, 0x13, 0x68, 0x23,
	0x15, 0x1d, 0xf0, 0xea, 0xe1, 0x40, 0xca, 0x41, 0xc8, 0x9b, 0x8e, 0xd4, 0x4b, 0xef, 0x9a, 0x5f,
	0x15, 0x4d, 0x12, 0xae, 0x74, 0x56, 0x56, 0xdd, 0x64, 0x32, 0x8a, 0x64, 0x9c, 0x47, 0x18, 0x4b,
	0x23, 0xee, 0x04, 0xa3, 0x46, 0x4c, 0x72, 0xff, 0xcd, 0x2a, 0xf0, 0x28, 0x31, 0xe3, 0x0c, 0xac,
	0x7f, 0xf7, 0xa0, 0x72, 0x9b, 0xf9, 0x39, 0xa7, 0xaa, 0xdf, 0x11, 0x7a, 0x44, 0xf8, 0x97, 0x94,
	0x6b, 0x83, 0x9f, 0xe6, 0x90, 0xee, 0x58, 0x1b, 0x1e, 0x69, 0xdf, 0xab, 0xad, 0x34, 0x4a, 0xad,
	0xa3, 0x60, 0xa1, 0xe3, 0x60, 0x56, 0xee, 0x0f, 0x2a, 0xf8, 0x06, 0xb6, 0xae, 0x13, 0xae, 0x9c,
	0xd5, 0x9b, 0x71, 0xc2, 0xfd, 0xe5, 0x9a, 0xd7, 0x28, 0xb7, 0xca, 0x4e, 0x76, 0x82, 0x90, 0xa7,
	0xa4, 0xfa, 0x4f, 0x0f, 0xfe, 0x99, 0x33, 0xac, 0x13, 0x19, 0x6b, 0xfe, 0xd7, 0x1d, 0xb7, 0xa0,
	0x48, 0xb8, 0x4e, 0x43, 0xe3, 0xac, 0x96, 0x5a, 0xd5, 0x20, 0xbb, 0xda, 0xe0, 0xe1, 0x6a, 0x83,
	0x53, 0x29, 0xc3, 0x5b, 0x1a, 0xa6, 0x9c, 0xe4, 0x4c, 0xdc, 0x87, 0xc2, 0x99, 0x52, 0x52, 0xf9,
	0x2b, 0x35, 0xaf, 0xb1, 0x41, 0xb2, 0xa0, 0xfe, 0xa3, 0x00, 0xdb, 0x33, 0x9b, 0x20, 0xc2, 0x6a,
	0x4c, 0x23, 0xee, 0x7b, 0x8e, 0xe8, 0xd6, 0x58, 0x86, 0x65, 0xd1, 0x77, 0xbb, 0x6d, 0x90, 0x65,
	0xd1, 0xc7, 0x0a, 0x14, 0xb5, 0x4c, 0x15, 0xe3, 0xb9, 0x5c, 0x1e, 0xd9, 0xda, 0x84, 0x9a, 0xa1,
	0xbf, 0x9a, 0xd5, 0xda, 0x35, 0xbe, 0x84, 0x2d, 0x26, 0x63, 0x43, 0x45, 0xcc, 0xd5, 0x95, 0x15,
	0x2e, 0x38, 0xf0, 0x69, 0x12, 0x5f, 0x40, 0x51, 0x1b, 0x6a, 0x52, 0xed, 0x17, 0xdd, 0x99, 0x4a,
	0xee, 0x8e, 0xba, 0x2e, 0x45, 0x72, 0xc8, 0xca, 0x6b, 0xf1, 0x8d, 0xfb, 0x6b, 0x35, 0xaf, 0xb1,
	0x42, 0xdc, 0x1a, 0x7d, 0x58, 0xeb, 0x8f, 0x63, 0x1a, 0x09, 0xe6, 0xaf, 0xd7, 0xbc, 0xc6, 0x3a,
	0x79, 0x08, 0xf1, 0x08, 0xca, 0xbd, 0x50, 0xb2, 0x91, 0xa5, 0xf5, 0xc6, 0x86, 0x6b, 0x7f, 0xa3,
	0xe6, 0x35, 0x0a, 0x64, 0x26, 0x8b, 0x01, 0x60, 0x28, 0x07, 0x82, 0xd1, 0x50, 0x73, 0x66, 0xa4,
	0xca, 0xb8, 0xe0, 0xb8, 0x0b, 0x10, 0x3c, 0x81, 0xbd, 0x64, 0x38, 0xd6, 0xb3, 0x05, 0x25, 0x57,
	0xb0, 0x08, 0xc2, 0x63, 0xd8, 0xb1, 0xa7, 0x55, 0x32, 0x0c, 0xb9, 0x8a, 0xd3, 0xa8, 0xc7, 0x95,
	0xbf, 0xe9, 0xce, 0x30, 0x97, 0xb7, 0x6e, 0x1e, 0x73, 0xa1, 0xcc, 0x5a, 0xc8, 0xdf, 0x72, 0xec,
	0x05, 0x08, 0x1e, 0x02, 0xd8, 0xee, 0xcd, 0x55, 0xcb, 0x8e, 0x37, 0x95, 0xb1, 0x7a, 0x79, 0xa3,
	0x47, 0x94, 0x0d, 0x45, 0xcc, 0xdd, 0x1b, 0x6c, 0xbb, 0x37, 0x58, 0x80, 0x60, 0x15, 0xd6, 0x35,
	0xd3, 0xc2, 0x3d, 0xe3, 0x8e, 0x63, 0x4d, 0x62, 0xfc, 0x08, 0x7b, 0x33, 0x43, 0xc3, 0xd8, 0x86,
	0xd9, 0x75, 0x0d, 0x73, 0xfc, 0xbc, 0x5f, 0x6d, 0xbb, 0x87, 0x2c, 0x92, 0xb1, 0x5f, 0x80, 0xc7,
	0x46, 0x98, 0xb1, 0x8f, 0x53, 0x5f, 0xe0, 0xcc, 0xa5, 0x48, 0x0e, 0x1d, 0x5f, 0xc2, 0xde, 0x02,
	0x41, 0xac, 0x00, 0x5e, 0x77, 0x3f, 0xdf, 0x5e, 0x90, 0x9b, 0xf7, 0xed, 0x77, 0xe7, 0x6d, 0xd2,
	0xe9, 0x5c, 0x74, 0x2f, 0x77, 0x96, 0xf0, 0x7f, 0xf0, 0x3b, 0xed, 0x9b, 0xb6, 0x8d, 0xe6, 0x50,
	0xaf, 0xf5, 0xcb, 0x83, 0xfd, 0x19, 0xb5, 0xb6, 0xf5, 0x8e, 0x02, 0x8a, 0x17, 0xf1, 0xbd, 0x1c,
	0x71, 0x7c, 0xf5, 0xcc, 0x5e, 0xcd, 0x86, 0x55, 0x35, 0x78, 0x2e, 0x3d, 0x1b, 0x15, 0xf5, 0x25,
	0x3c, 0x87, 0xdd, 0xb7, 0x43, 0xce, 0x46, 0x57, 0x53, 0x13, 0x13, 0x2b, 0x73, 0x1d, 0x7d, 0x66,
	0x87, 0x65, 0xf5, 0x5f, 0x27, 0x3f, 0x4d, 0x7d, 0x54, 0x3a, 0x3d, 0xf9, 0x10, 0x0c, 0x84, 0x19,
	0xa6, 0xbd, 0x80, 0xc9, 0xa8, 0x19, 0x09, 0xa6, 0xa4, 0x96, 0x77, 0xa6, 0x19, 0x49, 0xd6, 0x54,
	0x09, 0x6b, 0x4e, 0x5c, 0x35, 0x73, 0x57, 0xbd, 0xa2, 0x93, 0x7f, 0xfd, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0xc3, 0x94, 0xf4, 0x1e, 0x0b, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VirtualHardDiskAgentClient is the client API for VirtualHardDiskAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualHardDiskAgentClient interface {
	Invoke(ctx context.Context, in *VirtualHardDiskRequest, opts ...grpc.CallOption) (*VirtualHardDiskResponse, error)
	CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error)
}

type virtualHardDiskAgentClient struct {
	cc *grpc.ClientConn
}

func NewVirtualHardDiskAgentClient(cc *grpc.ClientConn) VirtualHardDiskAgentClient {
	return &virtualHardDiskAgentClient{cc}
}

func (c *virtualHardDiskAgentClient) Invoke(ctx context.Context, in *VirtualHardDiskRequest, opts ...grpc.CallOption) (*VirtualHardDiskResponse, error) {
	out := new(VirtualHardDiskResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.storage.VirtualHardDiskAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualHardDiskAgentClient) CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error) {
	out := new(common.NotificationResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.storage.VirtualHardDiskAgent/CheckNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualHardDiskAgentServer is the server API for VirtualHardDiskAgent service.
type VirtualHardDiskAgentServer interface {
	Invoke(context.Context, *VirtualHardDiskRequest) (*VirtualHardDiskResponse, error)
	CheckNotification(context.Context, *empty.Empty) (*common.NotificationResponse, error)
}

// UnimplementedVirtualHardDiskAgentServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualHardDiskAgentServer struct {
}

func (*UnimplementedVirtualHardDiskAgentServer) Invoke(ctx context.Context, req *VirtualHardDiskRequest) (*VirtualHardDiskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (*UnimplementedVirtualHardDiskAgentServer) CheckNotification(ctx context.Context, req *empty.Empty) (*common.NotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckNotification not implemented")
}

func RegisterVirtualHardDiskAgentServer(s *grpc.Server, srv VirtualHardDiskAgentServer) {
	s.RegisterService(&_VirtualHardDiskAgent_serviceDesc, srv)
}

func _VirtualHardDiskAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualHardDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualHardDiskAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.storage.VirtualHardDiskAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHardDiskAgentServer).Invoke(ctx, req.(*VirtualHardDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualHardDiskAgent_CheckNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualHardDiskAgentServer).CheckNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.storage.VirtualHardDiskAgent/CheckNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHardDiskAgentServer).CheckNotification(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualHardDiskAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.storage.VirtualHardDiskAgent",
	HandlerType: (*VirtualHardDiskAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _VirtualHardDiskAgent_Invoke_Handler,
		},
		{
			MethodName: "CheckNotification",
			Handler:    _VirtualHardDiskAgent_CheckNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "virtualharddisk.proto",
}