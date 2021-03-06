// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: carrotmember/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgAddMember struct {
	Creator       string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	MemberAddress string     `protobuf:"bytes,2,opt,name=memberAddress,proto3" json:"memberAddress,omitempty"`
	TotalCnt      uint64     `protobuf:"varint,3,opt,name=totalCnt,proto3" json:"totalCnt,omitempty"`
	Amount        types.Coin `protobuf:"bytes,4,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgAddMember) Reset()         { *m = MsgAddMember{} }
func (m *MsgAddMember) String() string { return proto.CompactTextString(m) }
func (*MsgAddMember) ProtoMessage()    {}
func (*MsgAddMember) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d5db3a02cdc0bca, []int{0}
}
func (m *MsgAddMember) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddMember) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddMember.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddMember) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddMember.Merge(m, src)
}
func (m *MsgAddMember) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddMember) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddMember.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddMember proto.InternalMessageInfo

func (m *MsgAddMember) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgAddMember) GetMemberAddress() string {
	if m != nil {
		return m.MemberAddress
	}
	return ""
}

func (m *MsgAddMember) GetTotalCnt() uint64 {
	if m != nil {
		return m.TotalCnt
	}
	return 0
}

func (m *MsgAddMember) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

type MsgAddMemberResponse struct {
}

func (m *MsgAddMemberResponse) Reset()         { *m = MsgAddMemberResponse{} }
func (m *MsgAddMemberResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddMemberResponse) ProtoMessage()    {}
func (*MsgAddMemberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d5db3a02cdc0bca, []int{1}
}
func (m *MsgAddMemberResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddMemberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddMemberResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddMemberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddMemberResponse.Merge(m, src)
}
func (m *MsgAddMemberResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddMemberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddMemberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddMemberResponse proto.InternalMessageInfo

type MsgSendReward struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *MsgSendReward) Reset()         { *m = MsgSendReward{} }
func (m *MsgSendReward) String() string { return proto.CompactTextString(m) }
func (*MsgSendReward) ProtoMessage()    {}
func (*MsgSendReward) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d5db3a02cdc0bca, []int{2}
}
func (m *MsgSendReward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendReward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendReward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendReward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendReward.Merge(m, src)
}
func (m *MsgSendReward) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendReward) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendReward.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendReward proto.InternalMessageInfo

func (m *MsgSendReward) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

type MsgSendRewardResponse struct {
}

func (m *MsgSendRewardResponse) Reset()         { *m = MsgSendRewardResponse{} }
func (m *MsgSendRewardResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendRewardResponse) ProtoMessage()    {}
func (*MsgSendRewardResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d5db3a02cdc0bca, []int{3}
}
func (m *MsgSendRewardResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendRewardResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendRewardResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendRewardResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendRewardResponse.Merge(m, src)
}
func (m *MsgSendRewardResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendRewardResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendRewardResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendRewardResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAddMember)(nil), "hupayxcom.carrotmember.carrotmember.MsgAddMember")
	proto.RegisterType((*MsgAddMemberResponse)(nil), "hupayxcom.carrotmember.carrotmember.MsgAddMemberResponse")
	proto.RegisterType((*MsgSendReward)(nil), "hupayxcom.carrotmember.carrotmember.MsgSendReward")
	proto.RegisterType((*MsgSendRewardResponse)(nil), "hupayxcom.carrotmember.carrotmember.MsgSendRewardResponse")
}

func init() { proto.RegisterFile("carrotmember/tx.proto", fileDescriptor_2d5db3a02cdc0bca) }

var fileDescriptor_2d5db3a02cdc0bca = []byte{
	// 366 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4e, 0xf2, 0x40,
	0x10, 0xc7, 0xbb, 0x1f, 0x84, 0xef, 0x63, 0x3f, 0xb9, 0x34, 0xa0, 0xb5, 0x87, 0x4a, 0xaa, 0x07,
	0x3c, 0xb8, 0x1b, 0xf0, 0x40, 0xf4, 0x06, 0x9c, 0x7b, 0xb0, 0xde, 0xbc, 0x6d, 0xdb, 0x4d, 0x21,
	0xb1, 0x9d, 0x66, 0x77, 0x51, 0x48, 0x7c, 0x08, 0xdf, 0xc1, 0x97, 0xe1, 0xc8, 0xd1, 0x93, 0x31,
	0xf0, 0x1e, 0xc6, 0xd0, 0x52, 0x69, 0x0f, 0x26, 0x70, 0x9b, 0xd9, 0x9d, 0xff, 0x7f, 0x7e, 0x33,
	0x19, 0xdc, 0xf2, 0x99, 0x10, 0xa0, 0x22, 0x1e, 0x79, 0x5c, 0x50, 0x35, 0x23, 0x89, 0x00, 0x05,
	0xfa, 0xf9, 0x78, 0x9a, 0xb0, 0xf9, 0xcc, 0x87, 0x88, 0x14, 0x0b, 0x4a, 0x89, 0x69, 0xf9, 0x20,
	0x23, 0x90, 0xd4, 0x63, 0x92, 0xd3, 0xa7, 0xae, 0xc7, 0x15, 0xeb, 0x52, 0x1f, 0x26, 0x71, 0x66,
	0x62, 0x36, 0x43, 0x08, 0x21, 0x0d, 0xe9, 0x26, 0xca, 0x5e, 0xed, 0x37, 0x84, 0x8f, 0x1c, 0x19,
	0x0e, 0x82, 0xc0, 0x49, 0x6d, 0x74, 0x03, 0xff, 0xf5, 0x05, 0x67, 0x0a, 0x84, 0x81, 0xda, 0xa8,
	0x53, 0x77, 0xf3, 0x54, 0xbf, 0xc0, 0x8d, 0xac, 0xd5, 0x20, 0x08, 0x04, 0x97, 0xd2, 0xf8, 0x93,
	0xfe, 0x97, 0x1f, 0x75, 0x13, 0xff, 0x53, 0xa0, 0xd8, 0xe3, 0x28, 0x56, 0x46, 0xa5, 0x8d, 0x3a,
	0x55, 0xf7, 0x27, 0xd7, 0xfb, 0xb8, 0xc6, 0x22, 0x98, 0xc6, 0xca, 0xa8, 0xb6, 0x51, 0xe7, 0x7f,
	0xef, 0x94, 0x64, 0xcc, 0x64, 0xc3, 0x4c, 0xb6, 0xcc, 0x64, 0x04, 0x93, 0x78, 0x58, 0x5d, 0x7c,
	0x9c, 0x69, 0xee, 0xb6, 0xdc, 0x3e, 0xc6, 0xcd, 0x22, 0xa4, 0xcb, 0x65, 0x02, 0xb1, 0xe4, 0xf6,
	0x25, 0x6e, 0x38, 0x32, 0xbc, 0xe7, 0x71, 0xe0, 0xf2, 0x67, 0x26, 0x82, 0xdf, 0xe9, 0xed, 0x13,
	0xdc, 0x2a, 0x95, 0xe6, 0x1e, 0xbd, 0x2f, 0x84, 0x2b, 0x8e, 0x0c, 0xf5, 0x39, 0xae, 0xef, 0xb6,
	0xd0, 0x25, 0x7b, 0xac, 0x9c, 0x14, 0x99, 0xcc, 0x9b, 0x83, 0x25, 0x39, 0x82, 0xfe, 0x82, 0x71,
	0x61, 0x86, 0xde, 0xbe, 0x46, 0x3b, 0x8d, 0x79, 0x7b, 0xb8, 0x26, 0xef, 0x3e, 0xbc, 0x5b, 0xac,
	0x2c, 0xb4, 0x5c, 0x59, 0xe8, 0x73, 0x65, 0xa1, 0xd7, 0xb5, 0xa5, 0x2d, 0xd7, 0x96, 0xf6, 0xbe,
	0xb6, 0xb4, 0x87, 0x7e, 0x38, 0x51, 0xe3, 0xa9, 0x47, 0x7c, 0x88, 0x68, 0xe6, 0x7f, 0xb5, 0x09,
	0x33, 0xcf, 0x6c, 0x06, 0x3a, 0xa3, 0xe5, 0x9b, 0x9d, 0x27, 0x5c, 0x7a, 0xb5, 0xf4, 0xb8, 0xae,
	0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x99, 0x2c, 0x0d, 0x9e, 0xd0, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	AddMember(ctx context.Context, in *MsgAddMember, opts ...grpc.CallOption) (*MsgAddMemberResponse, error)
	SendReward(ctx context.Context, in *MsgSendReward, opts ...grpc.CallOption) (*MsgSendRewardResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddMember(ctx context.Context, in *MsgAddMember, opts ...grpc.CallOption) (*MsgAddMemberResponse, error) {
	out := new(MsgAddMemberResponse)
	err := c.cc.Invoke(ctx, "/hupayxcom.carrotmember.carrotmember.Msg/AddMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SendReward(ctx context.Context, in *MsgSendReward, opts ...grpc.CallOption) (*MsgSendRewardResponse, error) {
	out := new(MsgSendRewardResponse)
	err := c.cc.Invoke(ctx, "/hupayxcom.carrotmember.carrotmember.Msg/SendReward", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	AddMember(context.Context, *MsgAddMember) (*MsgAddMemberResponse, error)
	SendReward(context.Context, *MsgSendReward) (*MsgSendRewardResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AddMember(ctx context.Context, req *MsgAddMember) (*MsgAddMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMember not implemented")
}
func (*UnimplementedMsgServer) SendReward(ctx context.Context, req *MsgSendReward) (*MsgSendRewardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendReward not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AddMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddMember)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hupayxcom.carrotmember.carrotmember.Msg/AddMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddMember(ctx, req.(*MsgAddMember))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SendReward_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSendReward)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SendReward(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hupayxcom.carrotmember.carrotmember.Msg/SendReward",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SendReward(ctx, req.(*MsgSendReward))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hupayxcom.carrotmember.carrotmember.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMember",
			Handler:    _Msg_AddMember_Handler,
		},
		{
			MethodName: "SendReward",
			Handler:    _Msg_SendReward_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "carrotmember/tx.proto",
}

func (m *MsgAddMember) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddMember) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddMember) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.TotalCnt != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TotalCnt))
		i--
		dAtA[i] = 0x18
	}
	if len(m.MemberAddress) > 0 {
		i -= len(m.MemberAddress)
		copy(dAtA[i:], m.MemberAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MemberAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddMemberResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddMemberResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddMemberResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSendReward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendReward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendReward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendRewardResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendRewardResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendRewardResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgAddMember) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.MemberAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TotalCnt != 0 {
		n += 1 + sovTx(uint64(m.TotalCnt))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgAddMemberResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSendReward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSendRewardResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAddMember) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddMember: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddMember: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemberAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MemberAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalCnt", wireType)
			}
			m.TotalCnt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalCnt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgAddMemberResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgAddMemberResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddMemberResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSendReward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSendReward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendReward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSendRewardResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSendRewardResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendRewardResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
