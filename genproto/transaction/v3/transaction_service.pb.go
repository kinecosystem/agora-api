// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transaction/v3/transaction_service.proto

package transaction

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	v3 "github.com/kinecosystem/kin-api/genproto/common/v3"
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

type SubmitSendResponse_Result int32

const (
	SubmitSendResponse_OK SubmitSendResponse_Result = 0
	// There was an issue with submitting the transaction
	// to the underlying chain. Clients should retry with
	// a rebuilt transaction in case there is temporal
	// issues with the transaction, such as sequence number,
	// or some other chain specific errors. The detail of
	// the error is present in the result xdr.
	SubmitSendResponse_INTERNAL_ERROR SubmitSendResponse_Result = 1
	// Clients should not submit send transactions until
	// they've verified they have sufficient funds.
	SubmitSendResponse_INSUFFICIENT_BALANCE SubmitSendResponse_Result = 2
)

var SubmitSendResponse_Result_name = map[int32]string{
	0: "OK",
	1: "INTERNAL_ERROR",
	2: "INSUFFICIENT_BALANCE",
}

var SubmitSendResponse_Result_value = map[string]int32{
	"OK":                   0,
	"INTERNAL_ERROR":       1,
	"INSUFFICIENT_BALANCE": 2,
}

func (x SubmitSendResponse_Result) String() string {
	return proto.EnumName(SubmitSendResponse_Result_name, int32(x))
}

func (SubmitSendResponse_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67d4eb25f503063e, []int{3, 0}
}

type GetTransactionResponse_State int32

const (
	GetTransactionResponse_UNKNOWN            GetTransactionResponse_State = 0
	GetTransactionResponse_SUCCESS            GetTransactionResponse_State = 1
	GetTransactionResponse_INTERNAL_ERROR     GetTransactionResponse_State = 2
	GetTransactionResponse_INSUFFICIENT_FUNDS GetTransactionResponse_State = 3
)

var GetTransactionResponse_State_name = map[int32]string{
	0: "UNKNOWN",
	1: "SUCCESS",
	2: "INTERNAL_ERROR",
	3: "INSUFFICIENT_FUNDS",
}

var GetTransactionResponse_State_value = map[string]int32{
	"UNKNOWN":            0,
	"SUCCESS":            1,
	"INTERNAL_ERROR":     2,
	"INSUFFICIENT_FUNDS": 3,
}

func (x GetTransactionResponse_State) String() string {
	return proto.EnumName(GetTransactionResponse_State_name, int32(x))
}

func (GetTransactionResponse_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67d4eb25f503063e, []int{5, 0}
}

type GetHistoryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHistoryRequest) Reset()         { *m = GetHistoryRequest{} }
func (m *GetHistoryRequest) String() string { return proto.CompactTextString(m) }
func (*GetHistoryRequest) ProtoMessage()    {}
func (*GetHistoryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67d4eb25f503063e, []int{0}
}

func (m *GetHistoryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHistoryRequest.Unmarshal(m, b)
}
func (m *GetHistoryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHistoryRequest.Marshal(b, m, deterministic)
}
func (m *GetHistoryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHistoryRequest.Merge(m, src)
}
func (m *GetHistoryRequest) XXX_Size() int {
	return xxx_messageInfo_GetHistoryRequest.Size(m)
}
func (m *GetHistoryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHistoryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetHistoryRequest proto.InternalMessageInfo

type GetHistoryResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetHistoryResponse) Reset()         { *m = GetHistoryResponse{} }
func (m *GetHistoryResponse) String() string { return proto.CompactTextString(m) }
func (*GetHistoryResponse) ProtoMessage()    {}
func (*GetHistoryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67d4eb25f503063e, []int{1}
}

func (m *GetHistoryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHistoryResponse.Unmarshal(m, b)
}
func (m *GetHistoryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHistoryResponse.Marshal(b, m, deterministic)
}
func (m *GetHistoryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHistoryResponse.Merge(m, src)
}
func (m *GetHistoryResponse) XXX_Size() int {
	return xxx_messageInfo_GetHistoryResponse.Size(m)
}
func (m *GetHistoryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHistoryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetHistoryResponse proto.InternalMessageInfo

type SubmitSendRequest struct {
	// The raw XDR bytes (not base64 encoded) of the transaction envelope.
	TransactionXdr       []byte   `protobuf:"bytes,1,opt,name=transaction_xdr,json=transactionXdr,proto3" json:"transaction_xdr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitSendRequest) Reset()         { *m = SubmitSendRequest{} }
func (m *SubmitSendRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitSendRequest) ProtoMessage()    {}
func (*SubmitSendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67d4eb25f503063e, []int{2}
}

func (m *SubmitSendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitSendRequest.Unmarshal(m, b)
}
func (m *SubmitSendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitSendRequest.Marshal(b, m, deterministic)
}
func (m *SubmitSendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitSendRequest.Merge(m, src)
}
func (m *SubmitSendRequest) XXX_Size() int {
	return xxx_messageInfo_SubmitSendRequest.Size(m)
}
func (m *SubmitSendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitSendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitSendRequest proto.InternalMessageInfo

func (m *SubmitSendRequest) GetTransactionXdr() []byte {
	if m != nil {
		return m.TransactionXdr
	}
	return nil
}

type SubmitSendResponse struct {
	Result SubmitSendResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=kin.transaction.v3.SubmitSendResponse_Result" json:"result,omitempty"`
	// The hash of the transaction, which may be used for other RPCs.
	Hash *v3.TransactionHash `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// The ledger in which the submitted transaction was included in.
	//
	// Non-zero on success.
	Ledger int64 `protobuf:"varint,3,opt,name=ledger,proto3" json:"ledger,omitempty"`
	// The transaction result XDR.
	ResultXdr            []byte   `protobuf:"bytes,4,opt,name=result_xdr,json=resultXdr,proto3" json:"result_xdr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubmitSendResponse) Reset()         { *m = SubmitSendResponse{} }
func (m *SubmitSendResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitSendResponse) ProtoMessage()    {}
func (*SubmitSendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67d4eb25f503063e, []int{3}
}

func (m *SubmitSendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubmitSendResponse.Unmarshal(m, b)
}
func (m *SubmitSendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubmitSendResponse.Marshal(b, m, deterministic)
}
func (m *SubmitSendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubmitSendResponse.Merge(m, src)
}
func (m *SubmitSendResponse) XXX_Size() int {
	return xxx_messageInfo_SubmitSendResponse.Size(m)
}
func (m *SubmitSendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubmitSendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubmitSendResponse proto.InternalMessageInfo

func (m *SubmitSendResponse) GetResult() SubmitSendResponse_Result {
	if m != nil {
		return m.Result
	}
	return SubmitSendResponse_OK
}

func (m *SubmitSendResponse) GetHash() *v3.TransactionHash {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *SubmitSendResponse) GetLedger() int64 {
	if m != nil {
		return m.Ledger
	}
	return 0
}

func (m *SubmitSendResponse) GetResultXdr() []byte {
	if m != nil {
		return m.ResultXdr
	}
	return nil
}

type GetTransactionRequest struct {
	TransactionHash      *v3.TransactionHash `protobuf:"bytes,1,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetTransactionRequest) Reset()         { *m = GetTransactionRequest{} }
func (m *GetTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*GetTransactionRequest) ProtoMessage()    {}
func (*GetTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_67d4eb25f503063e, []int{4}
}

func (m *GetTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionRequest.Unmarshal(m, b)
}
func (m *GetTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionRequest.Marshal(b, m, deterministic)
}
func (m *GetTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionRequest.Merge(m, src)
}
func (m *GetTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_GetTransactionRequest.Size(m)
}
func (m *GetTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionRequest proto.InternalMessageInfo

func (m *GetTransactionRequest) GetTransactionHash() *v3.TransactionHash {
	if m != nil {
		return m.TransactionHash
	}
	return nil
}

type GetTransactionResponse struct {
	// The state of the transaction. The states are the same as
	// SubmitSend, with the exception of UNKNOWN, which indicates
	// that the system is not yet aware of the transaction. This
	// cano occur if the transaction is still pending, or has been
	// dropped.
	//
	// If the transaction state is UNKNOWN for an extended period of
	// time, it is likely that it was dropped. As a result, clients
	// should limit the total times GetTransaction is called for a
	// an UNKNOWN transaction.
	State GetTransactionResponse_State `protobuf:"varint,1,opt,name=state,proto3,enum=kin.transaction.v3.GetTransactionResponse_State" json:"state,omitempty"`
	// Non-zero when state == State::SUCCESS
	Ledger int64 `protobuf:"varint,2,opt,name=ledger,proto3" json:"ledger,omitempty"`
	// Present when state != STATE::UNKNOWN
	ResultXdr            []byte   `protobuf:"bytes,3,opt,name=result_xdr,json=resultXdr,proto3" json:"result_xdr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTransactionResponse) Reset()         { *m = GetTransactionResponse{} }
func (m *GetTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*GetTransactionResponse) ProtoMessage()    {}
func (*GetTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_67d4eb25f503063e, []int{5}
}

func (m *GetTransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTransactionResponse.Unmarshal(m, b)
}
func (m *GetTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTransactionResponse.Marshal(b, m, deterministic)
}
func (m *GetTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTransactionResponse.Merge(m, src)
}
func (m *GetTransactionResponse) XXX_Size() int {
	return xxx_messageInfo_GetTransactionResponse.Size(m)
}
func (m *GetTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTransactionResponse proto.InternalMessageInfo

func (m *GetTransactionResponse) GetState() GetTransactionResponse_State {
	if m != nil {
		return m.State
	}
	return GetTransactionResponse_UNKNOWN
}

func (m *GetTransactionResponse) GetLedger() int64 {
	if m != nil {
		return m.Ledger
	}
	return 0
}

func (m *GetTransactionResponse) GetResultXdr() []byte {
	if m != nil {
		return m.ResultXdr
	}
	return nil
}

func init() {
	proto.RegisterEnum("kin.transaction.v3.SubmitSendResponse_Result", SubmitSendResponse_Result_name, SubmitSendResponse_Result_value)
	proto.RegisterEnum("kin.transaction.v3.GetTransactionResponse_State", GetTransactionResponse_State_name, GetTransactionResponse_State_value)
	proto.RegisterType((*GetHistoryRequest)(nil), "kin.transaction.v3.GetHistoryRequest")
	proto.RegisterType((*GetHistoryResponse)(nil), "kin.transaction.v3.GetHistoryResponse")
	proto.RegisterType((*SubmitSendRequest)(nil), "kin.transaction.v3.SubmitSendRequest")
	proto.RegisterType((*SubmitSendResponse)(nil), "kin.transaction.v3.SubmitSendResponse")
	proto.RegisterType((*GetTransactionRequest)(nil), "kin.transaction.v3.GetTransactionRequest")
	proto.RegisterType((*GetTransactionResponse)(nil), "kin.transaction.v3.GetTransactionResponse")
}

func init() {
	proto.RegisterFile("transaction/v3/transaction_service.proto", fileDescriptor_67d4eb25f503063e)
}

var fileDescriptor_67d4eb25f503063e = []byte{
	// 569 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x5d, 0xd2, 0xb5, 0x83, 0x6f, 0xa8, 0x64, 0x66, 0x1b, 0x55, 0x0f, 0xa8, 0x8a, 0x04, 0x74,
	0x48, 0x4b, 0x50, 0x7b, 0x04, 0x21, 0x35, 0x25, 0x5d, 0xab, 0x8d, 0x14, 0x39, 0xad, 0x40, 0xec,
	0x50, 0xa5, 0x8d, 0x95, 0x46, 0x6b, 0x93, 0x12, 0xbb, 0x15, 0xdb, 0x89, 0x33, 0xbf, 0x81, 0x1f,
	0xc1, 0x4f, 0xe3, 0xbc, 0x13, 0x8a, 0x93, 0xa9, 0x4e, 0x1b, 0xc4, 0xb8, 0xe5, 0x8b, 0xbf, 0xf7,
	0xfc, 0xbd, 0xf7, 0x6c, 0x43, 0x9d, 0x45, 0x4e, 0x40, 0x9d, 0x09, 0xf3, 0xc3, 0x40, 0x5f, 0x35,
	0x75, 0xa1, 0x1c, 0x51, 0x12, 0xad, 0xfc, 0x09, 0xd1, 0x16, 0x51, 0xc8, 0x42, 0x84, 0xae, 0xfc,
	0x40, 0x13, 0x96, 0xb5, 0x55, 0xb3, 0xfa, 0x74, 0xe5, 0xcc, 0x7c, 0xd7, 0x61, 0x44, 0xbf, 0xfb,
	0x48, 0x9a, 0xab, 0x47, 0x93, 0x70, 0x3e, 0x4f, 0x18, 0xe7, 0xa1, 0x4b, 0x66, 0xc9, 0x6f, 0xf5,
	0x09, 0x1c, 0x9c, 0x11, 0xd6, 0xf5, 0x29, 0x0b, 0xa3, 0x6b, 0x4c, 0xbe, 0x2e, 0x09, 0x65, 0xea,
	0x21, 0x20, 0xf1, 0x27, 0x5d, 0x84, 0x01, 0x25, 0x6a, 0x17, 0x0e, 0xec, 0xe5, 0x78, 0xee, 0x33,
	0x9b, 0x04, 0x6e, 0xda, 0x8a, 0x9a, 0xf0, 0x58, 0x1c, 0xf0, 0x9b, 0x1b, 0x55, 0xa4, 0x9a, 0x54,
	0x7f, 0x64, 0xc0, 0xad, 0xb1, 0x77, 0x53, 0x54, 0xa4, 0xca, 0xf7, 0x8f, 0xb8, 0x2c, 0xb4, 0x7c,
	0x76, 0x23, 0xf5, 0xa7, 0x0c, 0x48, 0xa4, 0x4a, 0x36, 0x40, 0x26, 0x94, 0x22, 0x42, 0x97, 0x33,
	0xc6, 0x29, 0xca, 0x8d, 0x53, 0x6d, 0x5b, 0xa0, 0xb6, 0x8d, 0xd3, 0x30, 0x07, 0xe1, 0x14, 0x8c,
	0xde, 0xc2, 0xee, 0xd4, 0xa1, 0xd3, 0x8a, 0x5c, 0x93, 0xea, 0xfb, 0x8d, 0x67, 0x9c, 0x24, 0x11,
	0x1f, 0xe3, 0x07, 0x6b, 0xba, 0xae, 0x43, 0xa7, 0xc6, 0x83, 0x5b, 0xa3, 0xf8, 0x43, 0x92, 0x15,
	0x09, 0x73, 0x14, 0x3a, 0x86, 0xd2, 0x8c, 0xb8, 0x1e, 0x89, 0x2a, 0x85, 0x9a, 0x54, 0x2f, 0xe0,
	0xb4, 0x42, 0x27, 0x00, 0x09, 0x3f, 0xd7, 0xb8, 0xbb, 0xa5, 0xf1, 0x61, 0xb2, 0x1a, 0xcb, 0x7b,
	0x07, 0xa5, 0x64, 0x24, 0x54, 0x02, 0xb9, 0x7f, 0xae, 0xec, 0x20, 0x04, 0xe5, 0x9e, 0x35, 0x30,
	0xb1, 0xd5, 0xba, 0x18, 0x99, 0x18, 0xf7, 0xb1, 0x22, 0xa1, 0x0a, 0x1c, 0xf6, 0x2c, 0x7b, 0xd8,
	0xe9, 0xf4, 0xda, 0x3d, 0xd3, 0x1a, 0x8c, 0x8c, 0xd6, 0x45, 0xcb, 0x6a, 0x9b, 0x8a, 0xac, 0xce,
	0xe0, 0xe8, 0x8c, 0x30, 0x61, 0xd0, 0x3b, 0xb3, 0x6d, 0x50, 0x44, 0xb3, 0xb9, 0x4a, 0xe9, 0x3f,
	0x55, 0x8a, 0x71, 0xc5, 0x4b, 0xea, 0x6f, 0x09, 0x8e, 0x37, 0xb7, 0x4b, 0x03, 0xe9, 0x40, 0x91,
	0x32, 0x87, 0x91, 0x34, 0x8f, 0xd7, 0x79, 0x79, 0xe4, 0x43, 0x35, 0x3b, 0xc6, 0xe1, 0x04, 0x2e,
	0x78, 0x2a, 0x67, 0x3c, 0x7d, 0x99, 0xf1, 0xb4, 0xc0, 0x3d, 0x8d, 0x27, 0xbd, 0x29, 0x6c, 0x38,
	0xfa, 0x01, 0x8a, 0x9c, 0x10, 0xed, 0xc3, 0xde, 0xd0, 0x3a, 0xb7, 0xfa, 0x9f, 0x2c, 0x65, 0x27,
	0x2e, 0xec, 0x61, 0xbb, 0x6d, 0xda, 0xb6, 0x22, 0xe5, 0x58, 0x2c, 0xa3, 0x63, 0x40, 0x19, 0x8b,
	0x3b, 0x43, 0xeb, 0xbd, 0xad, 0x14, 0x1a, 0xbf, 0x64, 0xd8, 0x17, 0x86, 0x46, 0x97, 0x00, 0xeb,
	0xf3, 0x8e, 0x9e, 0xff, 0x45, 0x66, 0xf6, 0x92, 0x54, 0x5f, 0xfc, 0xab, 0x2d, 0x35, 0xf1, 0x12,
	0x60, 0x7d, 0x66, 0xf3, 0xc9, 0xb7, 0xae, 0x55, 0x3e, 0x79, 0xce, 0x95, 0xf1, 0xa0, 0x9c, 0x0d,
	0x00, 0x9d, 0xdc, 0x27, 0xa4, 0x64, 0x93, 0x57, 0xf7, 0xcf, 0xd3, 0x70, 0xa0, 0x1a, 0x46, 0x1e,
	0x07, 0x78, 0x64, 0x13, 0xf4, 0xa5, 0xed, 0xf9, 0x6c, 0xba, 0x1c, 0xc7, 0x07, 0x50, 0xbf, 0xf2,
	0x03, 0x32, 0x09, 0xe9, 0x35, 0x65, 0x84, 0x17, 0xa7, 0xce, 0xc2, 0xd7, 0x3d, 0x12, 0xf0, 0xf7,
	0x46, 0xcf, 0x3e, 0x6e, 0x6f, 0x84, 0x72, 0x5c, 0xe2, 0x1d, 0xcd, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x55, 0xcc, 0x60, 0x54, 0x01, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionClient interface {
	GetHistory(ctx context.Context, in *GetHistoryRequest, opts ...grpc.CallOption) (*GetHistoryResponse, error)
	// SubmitSend submits a send transaction, which _may_ be whitelisted.
	//
	// If the memo does not conform to the 'memo standard' (todo: name),
	// the transaction is not eligible for whitelisting.
	SubmitSend(ctx context.Context, in *SubmitSendRequest, opts ...grpc.CallOption) (*SubmitSendResponse, error)
	// GetTransaction blocks until the specified transaction is resolved,
	// or until the RPC is cancelled by client / server, or fails.
	//
	// A transaction is resolved if it has succeeded for failed.
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error)
}

type transactionClient struct {
	cc *grpc.ClientConn
}

func NewTransactionClient(cc *grpc.ClientConn) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) GetHistory(ctx context.Context, in *GetHistoryRequest, opts ...grpc.CallOption) (*GetHistoryResponse, error) {
	out := new(GetHistoryResponse)
	err := c.cc.Invoke(ctx, "/kin.transaction.v3.Transaction/GetHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) SubmitSend(ctx context.Context, in *SubmitSendRequest, opts ...grpc.CallOption) (*SubmitSendResponse, error) {
	out := new(SubmitSendResponse)
	err := c.cc.Invoke(ctx, "/kin.transaction.v3.Transaction/SubmitSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transactionClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error) {
	out := new(GetTransactionResponse)
	err := c.cc.Invoke(ctx, "/kin.transaction.v3.Transaction/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
type TransactionServer interface {
	GetHistory(context.Context, *GetHistoryRequest) (*GetHistoryResponse, error)
	// SubmitSend submits a send transaction, which _may_ be whitelisted.
	//
	// If the memo does not conform to the 'memo standard' (todo: name),
	// the transaction is not eligible for whitelisting.
	SubmitSend(context.Context, *SubmitSendRequest) (*SubmitSendResponse, error)
	// GetTransaction blocks until the specified transaction is resolved,
	// or until the RPC is cancelled by client / server, or fails.
	//
	// A transaction is resolved if it has succeeded for failed.
	GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error)
}

// UnimplementedTransactionServer can be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (*UnimplementedTransactionServer) GetHistory(ctx context.Context, req *GetHistoryRequest) (*GetHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistory not implemented")
}
func (*UnimplementedTransactionServer) SubmitSend(ctx context.Context, req *SubmitSendRequest) (*SubmitSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitSend not implemented")
}
func (*UnimplementedTransactionServer) GetTransaction(ctx context.Context, req *GetTransactionRequest) (*GetTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}

func RegisterTransactionServer(s *grpc.Server, srv TransactionServer) {
	s.RegisterService(&_Transaction_serviceDesc, srv)
}

func _Transaction_GetHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).GetHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kin.transaction.v3.Transaction/GetHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).GetHistory(ctx, req.(*GetHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_SubmitSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitSendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).SubmitSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kin.transaction.v3.Transaction/SubmitSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).SubmitSend(ctx, req.(*SubmitSendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Transaction_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kin.transaction.v3.Transaction/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kin.transaction.v3.Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHistory",
			Handler:    _Transaction_GetHistory_Handler,
		},
		{
			MethodName: "SubmitSend",
			Handler:    _Transaction_SubmitSend_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _Transaction_GetTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transaction/v3/transaction_service.proto",
}
