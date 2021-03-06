// Code generated by protoc-gen-go. DO NOT EDIT.
// source: airdrop/v4/airdrop_service.proto

package airdrop

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	v4 "github.com/kinecosystem/agora-api/genproto/common/v4"
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

type RequestAirdropResponse_Result int32

const (
	RequestAirdropResponse_OK RequestAirdropResponse_Result = 0
	// The target account does not exist.
	RequestAirdropResponse_NOT_FOUND RequestAirdropResponse_Result = 1
	// The airdrop service does not have enough kin.
	RequestAirdropResponse_INSUFFICIENT_KIN RequestAirdropResponse_Result = 2
)

var RequestAirdropResponse_Result_name = map[int32]string{
	0: "OK",
	1: "NOT_FOUND",
	2: "INSUFFICIENT_KIN",
}

var RequestAirdropResponse_Result_value = map[string]int32{
	"OK":               0,
	"NOT_FOUND":        1,
	"INSUFFICIENT_KIN": 2,
}

func (x RequestAirdropResponse_Result) String() string {
	return proto.EnumName(RequestAirdropResponse_Result_name, int32(x))
}

func (RequestAirdropResponse_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f20258ebb4c54c8d, []int{1, 0}
}

type RequestAirdropRequest struct {
	AccountId            *v4.SolanaAccountId `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Quarks               uint64              `protobuf:"varint,2,opt,name=quarks,proto3" json:"quarks,omitempty"`
	Commitment           v4.Commitment       `protobuf:"varint,3,opt,name=commitment,proto3,enum=kin.agora.common.v4.Commitment" json:"commitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RequestAirdropRequest) Reset()         { *m = RequestAirdropRequest{} }
func (m *RequestAirdropRequest) String() string { return proto.CompactTextString(m) }
func (*RequestAirdropRequest) ProtoMessage()    {}
func (*RequestAirdropRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f20258ebb4c54c8d, []int{0}
}

func (m *RequestAirdropRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestAirdropRequest.Unmarshal(m, b)
}
func (m *RequestAirdropRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestAirdropRequest.Marshal(b, m, deterministic)
}
func (m *RequestAirdropRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestAirdropRequest.Merge(m, src)
}
func (m *RequestAirdropRequest) XXX_Size() int {
	return xxx_messageInfo_RequestAirdropRequest.Size(m)
}
func (m *RequestAirdropRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestAirdropRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RequestAirdropRequest proto.InternalMessageInfo

func (m *RequestAirdropRequest) GetAccountId() *v4.SolanaAccountId {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *RequestAirdropRequest) GetQuarks() uint64 {
	if m != nil {
		return m.Quarks
	}
	return 0
}

func (m *RequestAirdropRequest) GetCommitment() v4.Commitment {
	if m != nil {
		return m.Commitment
	}
	return v4.Commitment_RECENT
}

type RequestAirdropResponse struct {
	Result RequestAirdropResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=kin.agora.airdrop.v4.RequestAirdropResponse_Result" json:"result,omitempty"`
	// The signature of the transaction, if result == Result::OK.
	Signature            *v4.TransactionSignature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *RequestAirdropResponse) Reset()         { *m = RequestAirdropResponse{} }
func (m *RequestAirdropResponse) String() string { return proto.CompactTextString(m) }
func (*RequestAirdropResponse) ProtoMessage()    {}
func (*RequestAirdropResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f20258ebb4c54c8d, []int{1}
}

func (m *RequestAirdropResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestAirdropResponse.Unmarshal(m, b)
}
func (m *RequestAirdropResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestAirdropResponse.Marshal(b, m, deterministic)
}
func (m *RequestAirdropResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestAirdropResponse.Merge(m, src)
}
func (m *RequestAirdropResponse) XXX_Size() int {
	return xxx_messageInfo_RequestAirdropResponse.Size(m)
}
func (m *RequestAirdropResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestAirdropResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RequestAirdropResponse proto.InternalMessageInfo

func (m *RequestAirdropResponse) GetResult() RequestAirdropResponse_Result {
	if m != nil {
		return m.Result
	}
	return RequestAirdropResponse_OK
}

func (m *RequestAirdropResponse) GetSignature() *v4.TransactionSignature {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterEnum("kin.agora.airdrop.v4.RequestAirdropResponse_Result", RequestAirdropResponse_Result_name, RequestAirdropResponse_Result_value)
	proto.RegisterType((*RequestAirdropRequest)(nil), "kin.agora.airdrop.v4.RequestAirdropRequest")
	proto.RegisterType((*RequestAirdropResponse)(nil), "kin.agora.airdrop.v4.RequestAirdropResponse")
}

func init() { proto.RegisterFile("airdrop/v4/airdrop_service.proto", fileDescriptor_f20258ebb4c54c8d) }

var fileDescriptor_f20258ebb4c54c8d = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6e, 0xd3, 0x4e,
	0x10, 0xc6, 0xff, 0xeb, 0x7f, 0x71, 0xc9, 0x14, 0x22, 0x6b, 0xd5, 0x42, 0x14, 0x21, 0xd5, 0x8a,
	0x38, 0x04, 0x41, 0xd7, 0x92, 0x1b, 0x4e, 0x1c, 0x50, 0x5c, 0x08, 0xb2, 0x22, 0x1c, 0xe4, 0xa4,
	0x1c, 0xb8, 0x44, 0x5b, 0x7b, 0x65, 0x56, 0xb1, 0x77, 0xd3, 0xdd, 0xb5, 0x25, 0x5e, 0x81, 0xc7,
	0xe0, 0x69, 0x78, 0x11, 0x5e, 0xa2, 0x27, 0x14, 0xdb, 0x21, 0x01, 0xe5, 0x00, 0xb7, 0xd9, 0xd1,
	0xf7, 0xfb, 0x66, 0xbe, 0xd1, 0x82, 0x4b, 0xb9, 0x4a, 0x95, 0x5c, 0x7b, 0xd5, 0xc8, 0x6b, 0xcb,
	0xa5, 0x66, 0xaa, 0xe2, 0x09, 0x23, 0x6b, 0x25, 0x8d, 0xc4, 0xa7, 0x2b, 0x2e, 0x08, 0xcd, 0xa4,
	0xa2, 0xa4, 0x15, 0x90, 0x6a, 0xd4, 0x7f, 0x5c, 0xd1, 0x9c, 0xa7, 0xd4, 0x30, 0x6f, 0x5b, 0x34,
	0xf2, 0xfe, 0x59, 0x22, 0x8b, 0x42, 0x8a, 0x8d, 0x5f, 0x21, 0x53, 0x96, 0x37, 0xed, 0xc1, 0x77,
	0x04, 0x67, 0x31, 0xbb, 0x2d, 0x99, 0x36, 0xe3, 0xc6, 0xa5, 0x7d, 0xe1, 0xf7, 0x00, 0x34, 0x49,
	0x64, 0x29, 0xcc, 0x92, 0xa7, 0x3d, 0xe4, 0xa2, 0xe1, 0x89, 0xff, 0x94, 0xec, 0x86, 0x36, 0x7e,
	0xa4, 0x1a, 0x91, 0xb9, 0xcc, 0xa9, 0xa0, 0xe3, 0x46, 0x1c, 0xa6, 0xc1, 0xfd, 0xbb, 0xe0, 0xde,
	0x57, 0x64, 0x39, 0x28, 0xee, 0xd0, 0x6d, 0x13, 0x9f, 0x83, 0x7d, 0x5b, 0x52, 0xb5, 0xd2, 0x3d,
	0xcb, 0x45, 0xc3, 0xa3, 0xe0, 0xf8, 0x2e, 0x38, 0xf2, 0x2d, 0x17, 0xc5, 0x6d, 0x1b, 0xbf, 0x06,
	0xd8, 0x58, 0x72, 0x53, 0x30, 0x61, 0x7a, 0xff, 0xbb, 0x68, 0xd8, 0xf5, 0xcf, 0x0f, 0xce, 0xbb,
	0xfa, 0x25, 0x8b, 0xf7, 0x90, 0xc1, 0x0f, 0x04, 0x8f, 0xfe, 0x8c, 0xa2, 0xd7, 0x52, 0x68, 0x86,
	0xa7, 0x60, 0x2b, 0xa6, 0xcb, 0xdc, 0xd4, 0x39, 0xba, 0xfe, 0x25, 0x39, 0x74, 0x3c, 0x72, 0x98,
	0x26, 0x71, 0x8d, 0xc6, 0xad, 0x05, 0x7e, 0x07, 0x1d, 0xcd, 0x33, 0x41, 0x4d, 0xa9, 0x58, 0x1d,
	0xe6, 0xc4, 0x7f, 0x76, 0x70, 0xcf, 0x85, 0xa2, 0x42, 0xd3, 0xc4, 0x70, 0x29, 0xe6, 0x5b, 0x20,
	0xde, 0xb1, 0x83, 0x97, 0x60, 0x37, 0xd6, 0xd8, 0x06, 0x6b, 0x36, 0x75, 0xfe, 0xc3, 0x0f, 0xa1,
	0x13, 0xcd, 0x16, 0xcb, 0xc9, 0xec, 0x3a, 0x7a, 0xe3, 0x20, 0x7c, 0x0a, 0x4e, 0x18, 0xcd, 0xaf,
	0x27, 0x93, 0xf0, 0x2a, 0x7c, 0x1b, 0x2d, 0x96, 0xd3, 0x30, 0x72, 0x2c, 0xbf, 0x82, 0xe3, 0x76,
	0x43, 0xbc, 0x82, 0xee, 0xef, 0x3b, 0xe3, 0xe7, 0x7f, 0x97, 0xac, 0x7e, 0xf5, 0x5f, 0xfc, 0xcb,
	0x19, 0x82, 0x12, 0x9e, 0x48, 0x95, 0xed, 0x21, 0x19, 0x13, 0x7b, 0xd8, 0xa7, 0x49, 0xc6, 0xcd,
	0xe7, 0xf2, 0x66, 0x73, 0x00, 0x6f, 0xc5, 0x05, 0x4b, 0xa4, 0xfe, 0xa2, 0x0d, 0x2b, 0xbc, 0x5a,
	0x7d, 0x41, 0xd7, 0xfc, 0x82, 0x0b, 0xc3, 0x94, 0xa0, 0xb9, 0x97, 0x31, 0x51, 0xff, 0x3f, 0x6f,
	0xf7, 0xcd, 0x5f, 0xb5, 0xe5, 0x37, 0xeb, 0xc1, 0xf8, 0x43, 0xd0, 0x0e, 0xff, 0x38, 0xba, 0xb1,
	0x6b, 0xe1, 0xe5, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x18, 0x1f, 0x76, 0xe7, 0x12, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AirdropClient is the client API for Airdrop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AirdropClient interface {
	// RequestAirdrop requests an air drop of kin to the target account.
	RequestAirdrop(ctx context.Context, in *RequestAirdropRequest, opts ...grpc.CallOption) (*RequestAirdropResponse, error)
}

type airdropClient struct {
	cc *grpc.ClientConn
}

func NewAirdropClient(cc *grpc.ClientConn) AirdropClient {
	return &airdropClient{cc}
}

func (c *airdropClient) RequestAirdrop(ctx context.Context, in *RequestAirdropRequest, opts ...grpc.CallOption) (*RequestAirdropResponse, error) {
	out := new(RequestAirdropResponse)
	err := c.cc.Invoke(ctx, "/kin.agora.airdrop.v4.Airdrop/RequestAirdrop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AirdropServer is the server API for Airdrop service.
type AirdropServer interface {
	// RequestAirdrop requests an air drop of kin to the target account.
	RequestAirdrop(context.Context, *RequestAirdropRequest) (*RequestAirdropResponse, error)
}

// UnimplementedAirdropServer can be embedded to have forward compatible implementations.
type UnimplementedAirdropServer struct {
}

func (*UnimplementedAirdropServer) RequestAirdrop(ctx context.Context, req *RequestAirdropRequest) (*RequestAirdropResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestAirdrop not implemented")
}

func RegisterAirdropServer(s *grpc.Server, srv AirdropServer) {
	s.RegisterService(&_Airdrop_serviceDesc, srv)
}

func _Airdrop_RequestAirdrop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAirdropRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AirdropServer).RequestAirdrop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kin.agora.airdrop.v4.Airdrop/RequestAirdrop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AirdropServer).RequestAirdrop(ctx, req.(*RequestAirdropRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Airdrop_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kin.agora.airdrop.v4.Airdrop",
	HandlerType: (*AirdropServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestAirdrop",
			Handler:    _Airdrop_RequestAirdrop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "airdrop/v4/airdrop_service.proto",
}
