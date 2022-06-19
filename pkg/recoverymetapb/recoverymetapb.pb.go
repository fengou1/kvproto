// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: recoverymetapb.proto

package recovery

import (
	"context"
	"fmt"
	"io"
	"math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ReadRegionMetaRequest struct {
	StoreId              uint64   `protobuf:"varint,1,opt,name=store_id,json=storeId,proto3" json:"store_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRegionMetaRequest) Reset()         { *m = ReadRegionMetaRequest{} }
func (m *ReadRegionMetaRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRegionMetaRequest) ProtoMessage()    {}
func (*ReadRegionMetaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_58edb6090bdf10eb, []int{0}
}
func (m *ReadRegionMetaRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReadRegionMetaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReadRegionMetaRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReadRegionMetaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRegionMetaRequest.Merge(m, src)
}
func (m *ReadRegionMetaRequest) XXX_Size() int {
	return m.Size()
}
func (m *ReadRegionMetaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRegionMetaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRegionMetaRequest proto.InternalMessageInfo

func (m *ReadRegionMetaRequest) GetStoreId() uint64 {
	if m != nil {
		return m.StoreId
	}
	return 0
}

type RegionMeta struct {
	RegionId             uint64   `protobuf:"varint,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	AppliedIndex         uint64   `protobuf:"varint,2,opt,name=applied_index,json=appliedIndex,proto3" json:"applied_index,omitempty"`
	LastIndex            uint64   `protobuf:"varint,3,opt,name=last_index,json=lastIndex,proto3" json:"last_index,omitempty"`
	CommitIndex          uint64   `protobuf:"varint,4,opt,name=commit_index,json=commitIndex,proto3" json:"commit_index,omitempty"`
	Term                 uint64   `protobuf:"varint,5,opt,name=term,proto3" json:"term,omitempty"`
	Version              uint64   `protobuf:"varint,6,opt,name=version,proto3" json:"version,omitempty"`
	Tombstone            bool     `protobuf:"varint,7,opt,name=tombstone,proto3" json:"tombstone,omitempty"`
	StartKey             []byte   `protobuf:"bytes,8,opt,name=start_key,json=startKey,proto3" json:"start_key,omitempty"`
	EndKey               []byte   `protobuf:"bytes,9,opt,name=end_key,json=endKey,proto3" json:"end_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegionMeta) Reset()         { *m = RegionMeta{} }
func (m *RegionMeta) String() string { return proto.CompactTextString(m) }
func (*RegionMeta) ProtoMessage()    {}
func (*RegionMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_58edb6090bdf10eb, []int{1}
}
func (m *RegionMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegionMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegionMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegionMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegionMeta.Merge(m, src)
}
func (m *RegionMeta) XXX_Size() int {
	return m.Size()
}
func (m *RegionMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_RegionMeta.DiscardUnknown(m)
}

var xxx_messageInfo_RegionMeta proto.InternalMessageInfo

func (m *RegionMeta) GetRegionId() uint64 {
	if m != nil {
		return m.RegionId
	}
	return 0
}

func (m *RegionMeta) GetAppliedIndex() uint64 {
	if m != nil {
		return m.AppliedIndex
	}
	return 0
}

func (m *RegionMeta) GetLastIndex() uint64 {
	if m != nil {
		return m.LastIndex
	}
	return 0
}

func (m *RegionMeta) GetCommitIndex() uint64 {
	if m != nil {
		return m.CommitIndex
	}
	return 0
}

func (m *RegionMeta) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RegionMeta) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *RegionMeta) GetTombstone() bool {
	if m != nil {
		return m.Tombstone
	}
	return false
}

func (m *RegionMeta) GetStartKey() []byte {
	if m != nil {
		return m.StartKey
	}
	return nil
}

func (m *RegionMeta) GetEndKey() []byte {
	if m != nil {
		return m.EndKey
	}
	return nil
}

type RecoveryCmdRequest struct {
	RegionId             uint64   `protobuf:"varint,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	Term                 uint64   `protobuf:"varint,2,opt,name=term,proto3" json:"term,omitempty"`
	AsLeader             bool     `protobuf:"varint,3,opt,name=as_leader,json=asLeader,proto3" json:"as_leader,omitempty"`
	LeaderCommitIndex    bool     `protobuf:"varint,4,opt,name=leader_commit_index,json=leaderCommitIndex,proto3" json:"leader_commit_index,omitempty"`
	Tombstone            bool     `protobuf:"varint,5,opt,name=tombstone,proto3" json:"tombstone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecoveryCmdRequest) Reset()         { *m = RecoveryCmdRequest{} }
func (m *RecoveryCmdRequest) String() string { return proto.CompactTextString(m) }
func (*RecoveryCmdRequest) ProtoMessage()    {}
func (*RecoveryCmdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_58edb6090bdf10eb, []int{2}
}
func (m *RecoveryCmdRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RecoveryCmdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RecoveryCmdRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RecoveryCmdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoveryCmdRequest.Merge(m, src)
}
func (m *RecoveryCmdRequest) XXX_Size() int {
	return m.Size()
}
func (m *RecoveryCmdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoveryCmdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecoveryCmdRequest proto.InternalMessageInfo

func (m *RecoveryCmdRequest) GetRegionId() uint64 {
	if m != nil {
		return m.RegionId
	}
	return 0
}

func (m *RecoveryCmdRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RecoveryCmdRequest) GetAsLeader() bool {
	if m != nil {
		return m.AsLeader
	}
	return false
}

func (m *RecoveryCmdRequest) GetLeaderCommitIndex() bool {
	if m != nil {
		return m.LeaderCommitIndex
	}
	return false
}

func (m *RecoveryCmdRequest) GetTombstone() bool {
	if m != nil {
		return m.Tombstone
	}
	return false
}

// ensure leader assign is OK
type RecoveryCmdResponse struct {
	RegionId             uint64   `protobuf:"varint,1,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	Ok                   bool     `protobuf:"varint,2,opt,name=Ok,proto3" json:"Ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecoveryCmdResponse) Reset()         { *m = RecoveryCmdResponse{} }
func (m *RecoveryCmdResponse) String() string { return proto.CompactTextString(m) }
func (*RecoveryCmdResponse) ProtoMessage()    {}
func (*RecoveryCmdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_58edb6090bdf10eb, []int{3}
}
func (m *RecoveryCmdResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RecoveryCmdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RecoveryCmdResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RecoveryCmdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoveryCmdResponse.Merge(m, src)
}
func (m *RecoveryCmdResponse) XXX_Size() int {
	return m.Size()
}
func (m *RecoveryCmdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoveryCmdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RecoveryCmdResponse proto.InternalMessageInfo

func (m *RecoveryCmdResponse) GetRegionId() uint64 {
	if m != nil {
		return m.RegionId
	}
	return 0
}

func (m *RecoveryCmdResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*ReadRegionMetaRequest)(nil), "recovery.ReadRegionMetaRequest")
	proto.RegisterType((*RegionMeta)(nil), "recovery.RegionMeta")
	proto.RegisterType((*RecoveryCmdRequest)(nil), "recovery.RecoveryCmdRequest")
	proto.RegisterType((*RecoveryCmdResponse)(nil), "recovery.RecoveryCmdResponse")
}

func init() { proto.RegisterFile("recoverymetapb.proto", fileDescriptor_58edb6090bdf10eb) }

var fileDescriptor_58edb6090bdf10eb = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x41, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x33, 0x26, 0x4d, 0xc6, 0xaf, 0xa1, 0x88, 0x69, 0x10, 0x26, 0x6d, 0x43, 0x30, 0x9b,
	0xac, 0x02, 0x2a, 0x37, 0x68, 0x57, 0x51, 0x80, 0x4a, 0x73, 0x01, 0x6b, 0xd2, 0x79, 0x8a, 0xac,
	0xc4, 0x1e, 0x33, 0x33, 0xad, 0xc8, 0x35, 0x58, 0x71, 0x03, 0x58, 0x72, 0x0c, 0x96, 0x2c, 0x59,
	0xa2, 0x70, 0x11, 0xe4, 0x19, 0x9b, 0x38, 0x85, 0x66, 0x37, 0xef, 0xff, 0x7e, 0x4b, 0xff, 0xfb,
	0xf5, 0x0c, 0x7d, 0x8d, 0xd7, 0xea, 0x16, 0xf5, 0x3a, 0x43, 0x2b, 0x8a, 0xf9, 0xa4, 0xd0, 0xca,
	0x2a, 0x46, 0x6b, 0x75, 0xd0, 0x5f, 0xa8, 0x85, 0x72, 0xe2, 0xab, 0xf2, 0xe5, 0xf9, 0xe0, 0x91,
	0xbe, 0x31, 0xd6, 0x3d, 0xbd, 0x10, 0x9f, 0xc3, 0x13, 0x8e, 0x42, 0x72, 0x5c, 0xa4, 0x2a, 0x7f,
	0x87, 0x56, 0x70, 0xfc, 0x70, 0x83, 0xc6, 0xb2, 0x67, 0x40, 0x8d, 0x55, 0x1a, 0x93, 0x54, 0x46,
	0x64, 0x44, 0xc6, 0x6d, 0xde, 0x75, 0xf3, 0x54, 0xc6, 0x9f, 0x02, 0x80, 0xed, 0x07, 0xec, 0x04,
	0x42, 0xed, 0xa6, 0xad, 0x95, 0x7a, 0x61, 0x2a, 0xd9, 0x4b, 0x78, 0x28, 0x8a, 0x62, 0x95, 0xa2,
	0x4c, 0xd2, 0x5c, 0xe2, 0xc7, 0x28, 0x70, 0x86, 0x5e, 0x25, 0x4e, 0x4b, 0x8d, 0x9d, 0x01, 0xac,
	0x84, 0xb1, 0x95, 0xe3, 0x81, 0x73, 0x84, 0xa5, 0xe2, 0xf1, 0x0b, 0xe8, 0x5d, 0xab, 0x2c, 0x4b,
	0x6b, 0x43, 0xdb, 0x19, 0x0e, 0xbd, 0xe6, 0x2d, 0x0c, 0xda, 0x16, 0x75, 0x16, 0x1d, 0x38, 0xe4,
	0xde, 0x2c, 0x82, 0xee, 0x2d, 0x6a, 0x93, 0xaa, 0x3c, 0xea, 0xf8, 0x05, 0xaa, 0x91, 0x9d, 0x42,
	0x68, 0x55, 0x36, 0x37, 0x56, 0xe5, 0x18, 0x75, 0x47, 0x64, 0x4c, 0xf9, 0x56, 0x28, 0xf7, 0x31,
	0x56, 0x68, 0x9b, 0x2c, 0x71, 0x1d, 0xd1, 0x11, 0x19, 0xf7, 0x38, 0x75, 0xc2, 0x0c, 0xd7, 0xec,
	0x29, 0x74, 0x31, 0x97, 0x0e, 0x85, 0x0e, 0x75, 0x30, 0x97, 0x33, 0x5c, 0xc7, 0xdf, 0x08, 0x30,
	0x5e, 0x95, 0x7f, 0x99, 0xc9, 0xba, 0xc6, 0xbd, 0xe5, 0xd4, 0xa9, 0x83, 0x46, 0xea, 0x13, 0x08,
	0x85, 0x49, 0x56, 0x28, 0x24, 0x6a, 0x57, 0x05, 0xe5, 0x54, 0x98, 0xb7, 0x6e, 0x66, 0x13, 0x38,
	0xf6, 0x24, 0xf9, 0xa7, 0x10, 0xca, 0x1f, 0x7b, 0x74, 0xd9, 0xa8, 0x65, 0x67, 0xd1, 0x83, 0x3b,
	0x8b, 0xc6, 0x17, 0x70, 0xbc, 0x93, 0xd8, 0x14, 0x2a, 0x37, 0xb8, 0x3f, 0xf2, 0x11, 0x04, 0x57,
	0x4b, 0x17, 0x98, 0xf2, 0xe0, 0x6a, 0x79, 0xfe, 0x85, 0xc0, 0xdf, 0x9b, 0x63, 0x33, 0x38, 0xda,
	0x3d, 0x26, 0xf6, 0x7c, 0x52, 0xc3, 0xc9, 0x7f, 0xcf, 0x6c, 0xd0, 0x6f, 0x1a, 0x6a, 0x18, 0xb7,
	0x5e, 0x13, 0xf6, 0x1e, 0x0e, 0x1b, 0xe9, 0xd8, 0x69, 0xd3, 0x78, 0xb7, 0xe6, 0xc1, 0xd9, 0x3d,
	0xd4, 0xaf, 0x14, 0xb7, 0xc6, 0xe4, 0xa2, 0xff, 0xf3, 0x2b, 0x25, 0xdf, 0x37, 0x43, 0xf2, 0x63,
	0x33, 0x24, 0xbf, 0x36, 0x43, 0xf2, 0xf9, 0xf7, 0xb0, 0x35, 0xef, 0xb8, 0xdf, 0xe0, 0xcd, 0x9f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x80, 0x31, 0x34, 0x38, 0x4f, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RecoveryClient is the client API for Recovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecoveryClient interface {
	// scan region meta to ready region meta
	ReadRegionMeta(ctx context.Context, in *ReadRegionMetaRequest, opts ...grpc.CallOption) (Recovery_ReadRegionMetaClient, error)
	// execute the recovery command
	RecoveryCmd(ctx context.Context, opts ...grpc.CallOption) (Recovery_RecoveryCmdClient, error)
}

type recoveryClient struct {
	cc *grpc.ClientConn
}

func NewRecoveryClient(cc *grpc.ClientConn) RecoveryClient {
	return &recoveryClient{cc}
}

func (c *recoveryClient) ReadRegionMeta(ctx context.Context, in *ReadRegionMetaRequest, opts ...grpc.CallOption) (Recovery_ReadRegionMetaClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Recovery_serviceDesc.Streams[0], "/recovery.recovery/ReadRegionMeta", opts...)
	if err != nil {
		return nil, err
	}
	x := &recoveryReadRegionMetaClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Recovery_ReadRegionMetaClient interface {
	Recv() (*RegionMeta, error)
	grpc.ClientStream
}

type recoveryReadRegionMetaClient struct {
	grpc.ClientStream
}

func (x *recoveryReadRegionMetaClient) Recv() (*RegionMeta, error) {
	m := new(RegionMeta)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *recoveryClient) RecoveryCmd(ctx context.Context, opts ...grpc.CallOption) (Recovery_RecoveryCmdClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Recovery_serviceDesc.Streams[1], "/recovery.recovery/RecoveryCmd", opts...)
	if err != nil {
		return nil, err
	}
	x := &recoveryRecoveryCmdClient{stream}
	return x, nil
}

type Recovery_RecoveryCmdClient interface {
	Send(*RecoveryCmdRequest) error
	CloseAndRecv() (*RecoveryCmdResponse, error)
	grpc.ClientStream
}

type recoveryRecoveryCmdClient struct {
	grpc.ClientStream
}

func (x *recoveryRecoveryCmdClient) Send(m *RecoveryCmdRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *recoveryRecoveryCmdClient) CloseAndRecv() (*RecoveryCmdResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RecoveryCmdResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RecoveryServer is the server API for Recovery service.
type RecoveryServer interface {
	// scan region meta to ready region meta
	ReadRegionMeta(*ReadRegionMetaRequest, Recovery_ReadRegionMetaServer) error
	// execute the recovery command
	RecoveryCmd(Recovery_RecoveryCmdServer) error
}

// UnimplementedRecoveryServer can be embedded to have forward compatible implementations.
type UnimplementedRecoveryServer struct {
}

func (*UnimplementedRecoveryServer) ReadRegionMeta(req *ReadRegionMetaRequest, srv Recovery_ReadRegionMetaServer) error {
	return status.Errorf(codes.Unimplemented, "method ReadRegionMeta not implemented")
}
func (*UnimplementedRecoveryServer) RecoveryCmd(srv Recovery_RecoveryCmdServer) error {
	return status.Errorf(codes.Unimplemented, "method RecoveryCmd not implemented")
}

func RegisterRecoveryServer(s *grpc.Server, srv RecoveryServer) {
	s.RegisterService(&_Recovery_serviceDesc, srv)
}

func _Recovery_ReadRegionMeta_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadRegionMetaRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RecoveryServer).ReadRegionMeta(m, &recoveryReadRegionMetaServer{stream})
}

type Recovery_ReadRegionMetaServer interface {
	Send(*RegionMeta) error
	grpc.ServerStream
}

type recoveryReadRegionMetaServer struct {
	grpc.ServerStream
}

func (x *recoveryReadRegionMetaServer) Send(m *RegionMeta) error {
	return x.ServerStream.SendMsg(m)
}

func _Recovery_RecoveryCmd_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RecoveryServer).RecoveryCmd(&recoveryRecoveryCmdServer{stream})
}

type Recovery_RecoveryCmdServer interface {
	SendAndClose(*RecoveryCmdResponse) error
	Recv() (*RecoveryCmdRequest, error)
	grpc.ServerStream
}

type recoveryRecoveryCmdServer struct {
	grpc.ServerStream
}

func (x *recoveryRecoveryCmdServer) SendAndClose(m *RecoveryCmdResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *recoveryRecoveryCmdServer) Recv() (*RecoveryCmdRequest, error) {
	m := new(RecoveryCmdRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Recovery_serviceDesc = grpc.ServiceDesc{
	ServiceName: "recovery.recovery",
	HandlerType: (*RecoveryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadRegionMeta",
			Handler:       _Recovery_ReadRegionMeta_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RecoveryCmd",
			Handler:       _Recovery_RecoveryCmd_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "recoverymetapb.proto",
}

func (m *ReadRegionMetaRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReadRegionMetaRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReadRegionMetaRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.StoreId != 0 {
		i = encodeVarintRecoverymetapb(dAtA, i, uint64(m.StoreId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RegionMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegionMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegionMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.EndKey) > 0 {
		i -= len(m.EndKey)
		copy(dAtA[i:], m.EndKey)
		i = encodeVarintRecoverymetapb(dAtA, i, uint64(len(m.EndKey)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.StartKey) > 0 {
		i -= len(m.StartKey)
		copy(dAtA[i:], m.StartKey)
		i = encodeVarintRecoverymetapb(dAtA, i, uint64(len(m.StartKey)))
		i--
		dAtA[i] = 0x42
	}
	if m.Tombstone {
		i--
		if m.Tombstone {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if m.Version != 0 {
		i = encodeVarintRecoverymetapb(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x30
	}
	if m.Term != 0 {
		i = encodeVarintRecoverymetapb(dAtA, i, uint64(m.Term))
		i--
		dAtA[i] = 0x28
	}
	if m.CommitIndex != 0 {
		i = encodeVarintRecoverymetapb(dAtA, i, uint64(m.CommitIndex))
		i--
		dAtA[i] = 0x20
	}
	if m.LastIndex != 0 {
		i = encodeVarintRecoverymetapb(dAtA, i, uint64(m.LastIndex))
		i--
		dAtA[i] = 0x18
	}
	if m.AppliedIndex != 0 {
		i = encodeVarintRecoverymetapb(dAtA, i, uint64(m.AppliedIndex))
		i--
		dAtA[i] = 0x10
	}
	if m.RegionId != 0 {
		i = encodeVarintRecoverymetapb(dAtA, i, uint64(m.RegionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RecoveryCmdRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RecoveryCmdRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RecoveryCmdRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Tombstone {
		i--
		if m.Tombstone {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	if m.LeaderCommitIndex {
		i--
		if m.LeaderCommitIndex {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if m.AsLeader {
		i--
		if m.AsLeader {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.Term != 0 {
		i = encodeVarintRecoverymetapb(dAtA, i, uint64(m.Term))
		i--
		dAtA[i] = 0x10
	}
	if m.RegionId != 0 {
		i = encodeVarintRecoverymetapb(dAtA, i, uint64(m.RegionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RecoveryCmdResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RecoveryCmdResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RecoveryCmdResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Ok {
		i--
		if m.Ok {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.RegionId != 0 {
		i = encodeVarintRecoverymetapb(dAtA, i, uint64(m.RegionId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRecoverymetapb(dAtA []byte, offset int, v uint64) int {
	offset -= sovRecoverymetapb(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReadRegionMetaRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StoreId != 0 {
		n += 1 + sovRecoverymetapb(uint64(m.StoreId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RegionMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RegionId != 0 {
		n += 1 + sovRecoverymetapb(uint64(m.RegionId))
	}
	if m.AppliedIndex != 0 {
		n += 1 + sovRecoverymetapb(uint64(m.AppliedIndex))
	}
	if m.LastIndex != 0 {
		n += 1 + sovRecoverymetapb(uint64(m.LastIndex))
	}
	if m.CommitIndex != 0 {
		n += 1 + sovRecoverymetapb(uint64(m.CommitIndex))
	}
	if m.Term != 0 {
		n += 1 + sovRecoverymetapb(uint64(m.Term))
	}
	if m.Version != 0 {
		n += 1 + sovRecoverymetapb(uint64(m.Version))
	}
	if m.Tombstone {
		n += 2
	}
	l = len(m.StartKey)
	if l > 0 {
		n += 1 + l + sovRecoverymetapb(uint64(l))
	}
	l = len(m.EndKey)
	if l > 0 {
		n += 1 + l + sovRecoverymetapb(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RecoveryCmdRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RegionId != 0 {
		n += 1 + sovRecoverymetapb(uint64(m.RegionId))
	}
	if m.Term != 0 {
		n += 1 + sovRecoverymetapb(uint64(m.Term))
	}
	if m.AsLeader {
		n += 2
	}
	if m.LeaderCommitIndex {
		n += 2
	}
	if m.Tombstone {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RecoveryCmdResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RegionId != 0 {
		n += 1 + sovRecoverymetapb(uint64(m.RegionId))
	}
	if m.Ok {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRecoverymetapb(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRecoverymetapb(x uint64) (n int) {
	return sovRecoverymetapb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReadRegionMetaRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecoverymetapb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReadRegionMetaRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReadRegionMetaRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreId", wireType)
			}
			m.StoreId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StoreId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRecoverymetapb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecoverymetapb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RegionMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecoverymetapb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RegionMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegionMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionId", wireType)
			}
			m.RegionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppliedIndex", wireType)
			}
			m.AppliedIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppliedIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastIndex", wireType)
			}
			m.LastIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CommitIndex", wireType)
			}
			m.CommitIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CommitIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tombstone", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Tombstone = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRecoverymetapb
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRecoverymetapb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartKey = append(m.StartKey[:0], dAtA[iNdEx:postIndex]...)
			if m.StartKey == nil {
				m.StartKey = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRecoverymetapb
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRecoverymetapb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndKey = append(m.EndKey[:0], dAtA[iNdEx:postIndex]...)
			if m.EndKey == nil {
				m.EndKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecoverymetapb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecoverymetapb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RecoveryCmdRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecoverymetapb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RecoveryCmdRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RecoveryCmdRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionId", wireType)
			}
			m.RegionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Term", wireType)
			}
			m.Term = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Term |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AsLeader", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AsLeader = bool(v != 0)
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaderCommitIndex", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.LeaderCommitIndex = bool(v != 0)
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tombstone", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Tombstone = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRecoverymetapb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecoverymetapb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RecoveryCmdResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecoverymetapb
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RecoveryCmdResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RecoveryCmdResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegionId", wireType)
			}
			m.RegionId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegionId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ok", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Ok = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRecoverymetapb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecoverymetapb
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRecoverymetapb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRecoverymetapb
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRecoverymetapb
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRecoverymetapb
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRecoverymetapb
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRecoverymetapb
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRecoverymetapb        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRecoverymetapb          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRecoverymetapb = fmt.Errorf("proto: unexpected end of group")
)
