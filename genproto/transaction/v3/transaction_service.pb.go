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

type GetHistoryRequest_Direction int32

const (
	// ASC direction returns all history items in
	// ascending (based on chain order) from the cursor.
	GetHistoryRequest_ASC GetHistoryRequest_Direction = 0
	// SESC direction returns all history items in
	// descending (based on chain order) from _before_ the cursor.
	GetHistoryRequest_DESC GetHistoryRequest_Direction = 1
)

var GetHistoryRequest_Direction_name = map[int32]string{
	0: "ASC",
	1: "DESC",
}

var GetHistoryRequest_Direction_value = map[string]int32{
	"ASC":  0,
	"DESC": 1,
}

func (x GetHistoryRequest_Direction) String() string {
	return proto.EnumName(GetHistoryRequest_Direction_name, int32(x))
}

func (GetHistoryRequest_Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67d4eb25f503063e, []int{0, 0}
}

type GetHistoryResponse_Result int32

const (
	GetHistoryResponse_OK        GetHistoryResponse_Result = 0
	GetHistoryResponse_NOT_FOUND GetHistoryResponse_Result = 1
)

var GetHistoryResponse_Result_name = map[int32]string{
	0: "OK",
	1: "NOT_FOUND",
}

var GetHistoryResponse_Result_value = map[string]int32{
	"OK":        0,
	"NOT_FOUND": 1,
}

func (x GetHistoryResponse_Result) String() string {
	return proto.EnumName(GetHistoryResponse_Result_name, int32(x))
}

func (GetHistoryResponse_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_67d4eb25f503063e, []int{1, 0}
}

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
	// Another invoice with the same invoice hash prefix exists
	// in the store. Clients should regenerate the invoice
	// with a new nonce and resubmit.
	SubmitSendResponse_INVOICE_COLLISION SubmitSendResponse_Result = 3
)

var SubmitSendResponse_Result_name = map[int32]string{
	0: "OK",
	1: "INTERNAL_ERROR",
	2: "INSUFFICIENT_BALANCE",
	3: "INVOICE_COLLISION",
}

var SubmitSendResponse_Result_value = map[string]int32{
	"OK":                   0,
	"INTERNAL_ERROR":       1,
	"INSUFFICIENT_BALANCE": 2,
	"INVOICE_COLLISION":    3,
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
	// Account to get history for.
	AccountId *v3.StellarAccountId `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// An optional history cursor indicating where in
	// the history to 'resume' from.
	Cursor *Cursor `protobuf:"bytes,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	// The order in which to return history items from
	// the cursor.
	Direction            GetHistoryRequest_Direction `protobuf:"varint,3,opt,name=direction,proto3,enum=kin.transaction.v3.GetHistoryRequest_Direction" json:"direction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
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

func (m *GetHistoryRequest) GetAccountId() *v3.StellarAccountId {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *GetHistoryRequest) GetCursor() *Cursor {
	if m != nil {
		return m.Cursor
	}
	return nil
}

func (m *GetHistoryRequest) GetDirection() GetHistoryRequest_Direction {
	if m != nil {
		return m.Direction
	}
	return GetHistoryRequest_ASC
}

type GetHistoryResponse struct {
	Result               GetHistoryResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=kin.transaction.v3.GetHistoryResponse_Result" json:"result,omitempty"`
	Items                []*HistoryItem            `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
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

func (m *GetHistoryResponse) GetResult() GetHistoryResponse_Result {
	if m != nil {
		return m.Result
	}
	return GetHistoryResponse_OK
}

func (m *GetHistoryResponse) GetItems() []*HistoryItem {
	if m != nil {
		return m.Items
	}
	return nil
}

type SubmitSendRequest struct {
	// The raw XDR bytes (not base64 encoded) of the transaction envelope.
	TransactionXdr []byte `protobuf:"bytes,1,opt,name=transaction_xdr,json=transactionXdr,proto3" json:"transaction_xdr,omitempty"`
	// An invoice indicating what the submitted transaction is for.
	//
	// If an invoice is included in this request, it is expected that the foreign key in
	// the transaction memo is the 230-bit prefix of the invoice hash. The submitted
	// invoice will be included in any calls made to the configured webhook(s) of the
	// third-party app the transaction pertains to.
	//
	// The submitted invoice data will only be available for retrieval from the service it
	// was submitted to and not directly from the blockchain nor any other deployments of
	// the service.
	Invoice              *v3.Invoice `protobuf:"bytes,2,opt,name=invoice,proto3" json:"invoice,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
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

func (m *SubmitSendRequest) GetInvoice() *v3.Invoice {
	if m != nil {
		return m.Invoice
	}
	return nil
}

type SubmitSendResponse struct {
	Result SubmitSendResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=kin.transaction.v3.SubmitSendResponse_Result" json:"result,omitempty"`
	// The hash of the transaction, if it was submitted.
	// May be used for other RPCs.
	Hash *v3.TransactionHash `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	// The ledger in which the submitted transaction was included in.
	//
	// Non-zero on success.
	Ledger int64 `protobuf:"varint,3,opt,name=ledger,proto3" json:"ledger,omitempty"`
	// The transaction result XDR, if a transaction was submitted.
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
	// Present when state != State::UNKNOWN
	Item                 *HistoryItem `protobuf:"bytes,3,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
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

func (m *GetTransactionResponse) GetItem() *HistoryItem {
	if m != nil {
		return m.Item
	}
	return nil
}

type HistoryItem struct {
	// The hash of the transaction.
	Hash *v3.TransactionHash `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// Contains the transaction result from when the transaction was submitted
	ResultXdr []byte `protobuf:"bytes,2,opt,name=result_xdr,json=resultXdr,proto3" json:"result_xdr,omitempty"`
	// Contains the transaction envelope for the transaction.
	EnvelopeXdr []byte `protobuf:"bytes,3,opt,name=envelope_xdr,json=envelopeXdr,proto3" json:"envelope_xdr,omitempty"`
	// The cursor position of this item.
	Cursor *Cursor `protobuf:"bytes,4,opt,name=cursor,proto3" json:"cursor,omitempty"`
	// If the transaction contained a valid Agora generated memo, then
	// agora_data_url will be present. Clients may use this URL to obtain
	// off-chain data.
	//
	// todo: standardize agora_data_url protocol / spec. likely just a GET
	// with the transaction hash as a parameter.
	AgoraDataUrl *v3.AgoraDataUrl `protobuf:"bytes,5,opt,name=agora_data_url,json=agoraDataUrl,proto3" json:"agora_data_url,omitempty"`
	// If the service was able to resolve, or manages, the off-chain chain
	// data associated with the memo, AgoraData will be populated.
	AgoraData            *v3.AgoraData `protobuf:"bytes,6,opt,name=agora_data,json=agoraData,proto3" json:"agora_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *HistoryItem) Reset()         { *m = HistoryItem{} }
func (m *HistoryItem) String() string { return proto.CompactTextString(m) }
func (*HistoryItem) ProtoMessage()    {}
func (*HistoryItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_67d4eb25f503063e, []int{6}
}

func (m *HistoryItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HistoryItem.Unmarshal(m, b)
}
func (m *HistoryItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HistoryItem.Marshal(b, m, deterministic)
}
func (m *HistoryItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HistoryItem.Merge(m, src)
}
func (m *HistoryItem) XXX_Size() int {
	return xxx_messageInfo_HistoryItem.Size(m)
}
func (m *HistoryItem) XXX_DiscardUnknown() {
	xxx_messageInfo_HistoryItem.DiscardUnknown(m)
}

var xxx_messageInfo_HistoryItem proto.InternalMessageInfo

func (m *HistoryItem) GetHash() *v3.TransactionHash {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *HistoryItem) GetResultXdr() []byte {
	if m != nil {
		return m.ResultXdr
	}
	return nil
}

func (m *HistoryItem) GetEnvelopeXdr() []byte {
	if m != nil {
		return m.EnvelopeXdr
	}
	return nil
}

func (m *HistoryItem) GetCursor() *Cursor {
	if m != nil {
		return m.Cursor
	}
	return nil
}

func (m *HistoryItem) GetAgoraDataUrl() *v3.AgoraDataUrl {
	if m != nil {
		return m.AgoraDataUrl
	}
	return nil
}

func (m *HistoryItem) GetAgoraData() *v3.AgoraData {
	if m != nil {
		return m.AgoraData
	}
	return nil
}

type Cursor struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cursor) Reset()         { *m = Cursor{} }
func (m *Cursor) String() string { return proto.CompactTextString(m) }
func (*Cursor) ProtoMessage()    {}
func (*Cursor) Descriptor() ([]byte, []int) {
	return fileDescriptor_67d4eb25f503063e, []int{7}
}

func (m *Cursor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cursor.Unmarshal(m, b)
}
func (m *Cursor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cursor.Marshal(b, m, deterministic)
}
func (m *Cursor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cursor.Merge(m, src)
}
func (m *Cursor) XXX_Size() int {
	return xxx_messageInfo_Cursor.Size(m)
}
func (m *Cursor) XXX_DiscardUnknown() {
	xxx_messageInfo_Cursor.DiscardUnknown(m)
}

var xxx_messageInfo_Cursor proto.InternalMessageInfo

func (m *Cursor) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterEnum("kin.transaction.v3.GetHistoryRequest_Direction", GetHistoryRequest_Direction_name, GetHistoryRequest_Direction_value)
	proto.RegisterEnum("kin.transaction.v3.GetHistoryResponse_Result", GetHistoryResponse_Result_name, GetHistoryResponse_Result_value)
	proto.RegisterEnum("kin.transaction.v3.SubmitSendResponse_Result", SubmitSendResponse_Result_name, SubmitSendResponse_Result_value)
	proto.RegisterEnum("kin.transaction.v3.GetTransactionResponse_State", GetTransactionResponse_State_name, GetTransactionResponse_State_value)
	proto.RegisterType((*GetHistoryRequest)(nil), "kin.transaction.v3.GetHistoryRequest")
	proto.RegisterType((*GetHistoryResponse)(nil), "kin.transaction.v3.GetHistoryResponse")
	proto.RegisterType((*SubmitSendRequest)(nil), "kin.transaction.v3.SubmitSendRequest")
	proto.RegisterType((*SubmitSendResponse)(nil), "kin.transaction.v3.SubmitSendResponse")
	proto.RegisterType((*GetTransactionRequest)(nil), "kin.transaction.v3.GetTransactionRequest")
	proto.RegisterType((*GetTransactionResponse)(nil), "kin.transaction.v3.GetTransactionResponse")
	proto.RegisterType((*HistoryItem)(nil), "kin.transaction.v3.HistoryItem")
	proto.RegisterType((*Cursor)(nil), "kin.transaction.v3.Cursor")
}

func init() {
	proto.RegisterFile("transaction/v3/transaction_service.proto", fileDescriptor_67d4eb25f503063e)
}

var fileDescriptor_67d4eb25f503063e = []byte{
	// 880 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x16, 0xa9, 0x1f, 0x5b, 0x23, 0x57, 0xa5, 0x17, 0xb5, 0x4b, 0xa8, 0x40, 0x62, 0x10, 0x68,
	0xe1, 0x04, 0x30, 0x15, 0x48, 0x87, 0x1e, 0xda, 0x8b, 0x48, 0x49, 0x31, 0x11, 0x87, 0x2a, 0x96,
	0x56, 0x1b, 0xb4, 0x07, 0x61, 0x2d, 0x2e, 0x64, 0x22, 0x14, 0xe9, 0x2e, 0x57, 0x42, 0x93, 0x93,
	0xcf, 0x3d, 0xf6, 0x2d, 0x7a, 0xeb, 0x23, 0xf4, 0x25, 0xfa, 0x32, 0x46, 0x0f, 0x01, 0x97, 0x94,
	0xb5, 0x12, 0x15, 0x44, 0xb9, 0x79, 0x38, 0xdf, 0xf7, 0xed, 0xec, 0x7c, 0xb3, 0x63, 0xc1, 0x39,
	0x67, 0x24, 0x4a, 0xc8, 0x94, 0x07, 0x71, 0xd4, 0x5e, 0x76, 0xdb, 0x52, 0x38, 0x49, 0x28, 0x5b,
	0x06, 0x53, 0x6a, 0xde, 0xb1, 0x98, 0xc7, 0x08, 0xbd, 0x0d, 0x22, 0x53, 0x4a, 0x9b, 0xcb, 0x6e,
	0xeb, 0xeb, 0x25, 0x09, 0x03, 0x9f, 0x70, 0xda, 0x5e, 0xfd, 0x91, 0x81, 0x5b, 0x27, 0xd3, 0x78,
	0x3e, 0xcf, 0x14, 0xe7, 0xb1, 0x4f, 0xc3, 0xec, 0xb3, 0x71, 0xaf, 0xc2, 0xf1, 0x4b, 0xca, 0x2f,
	0x83, 0x84, 0xc7, 0xec, 0x1d, 0xa6, 0xbf, 0x2f, 0x68, 0xc2, 0xd1, 0x25, 0x00, 0x99, 0x4e, 0xe3,
	0x45, 0xc4, 0x27, 0x81, 0xaf, 0x2b, 0x67, 0xca, 0x79, 0xa3, 0xf3, 0xd4, 0x4c, 0x8f, 0xcb, 0x54,
	0xcc, 0x65, 0xd7, 0xf4, 0x38, 0x0d, 0x43, 0xc2, 0x7a, 0x19, 0xce, 0xf1, 0xad, 0xc3, 0x07, 0xab,
	0xfa, 0xa7, 0xa2, 0x6a, 0x0a, 0xae, 0x93, 0xd5, 0x47, 0xd4, 0x81, 0xda, 0x74, 0xc1, 0x92, 0x98,
	0xe9, 0xaa, 0x50, 0x69, 0x99, 0xc5, 0xa2, 0x4d, 0x5b, 0x20, 0x70, 0x8e, 0x44, 0xaf, 0xa1, 0xee,
	0x07, 0x8c, 0x8a, 0xb4, 0x5e, 0x3e, 0x53, 0xce, 0x9b, 0x9d, 0xf6, 0x2e, 0x5a, 0xa1, 0x6e, 0xb3,
	0xbf, 0xa2, 0xe1, 0xb5, 0x82, 0xf1, 0x04, 0xea, 0x8f, 0xdf, 0xd1, 0x01, 0x94, 0x7b, 0x9e, 0xad,
	0x95, 0xd0, 0x21, 0x54, 0xfa, 0x03, 0xcf, 0xd6, 0x14, 0xe3, 0x5f, 0x05, 0x90, 0x2c, 0x95, 0xdc,
	0xc5, 0x51, 0x42, 0xd1, 0x00, 0x6a, 0x8c, 0x26, 0x8b, 0x90, 0x8b, 0xfb, 0x37, 0x3b, 0x17, 0x9f,
	0x2a, 0x21, 0xe3, 0x99, 0x58, 0x90, 0x70, 0x4e, 0x46, 0x16, 0x54, 0x03, 0x4e, 0xe7, 0x89, 0xae,
	0x9e, 0x95, 0x1f, 0xbb, 0xb8, 0xa5, 0x92, 0x4b, 0x38, 0x9c, 0xce, 0xad, 0xc6, 0x83, 0x75, 0xf8,
	0x97, 0x52, 0x3d, 0x2c, 0x69, 0xf7, 0x0a, 0xce, 0xa8, 0xc6, 0x53, 0xa8, 0x65, 0xaa, 0xa8, 0x06,
	0xea, 0xe8, 0x95, 0x56, 0x42, 0x5f, 0x40, 0xdd, 0x1d, 0x5d, 0x4f, 0x86, 0xa3, 0xb1, 0xdb, 0xd7,
	0x14, 0xe3, 0x3d, 0x1c, 0x7b, 0x8b, 0x9b, 0x79, 0xc0, 0x3d, 0x1a, 0xf9, 0x2b, 0x13, 0xbb, 0xf0,
	0xa5, 0x3c, 0x3b, 0x7f, 0xf8, 0x4c, 0xdc, 0xe4, 0xc8, 0x82, 0x07, 0xeb, 0xe0, 0x7d, 0x55, 0x53,
	0xf4, 0xfb, 0x9f, 0x70, 0x53, 0x82, 0xbc, 0xf1, 0x19, 0x7a, 0x01, 0x07, 0x41, 0xb4, 0x8c, 0x83,
	0x29, 0xcd, 0x0d, 0x3b, 0xdd, 0xb2, 0xdd, 0xc9, 0xb2, 0x78, 0x05, 0x33, 0xfe, 0x56, 0x01, 0xc9,
	0x87, 0x7f, 0x4e, 0xfb, 0x8a, 0xbc, 0xed, 0xf6, 0x75, 0xa0, 0x72, 0x4b, 0x92, 0xdb, 0xbc, 0x98,
	0x27, 0x5b, 0xc5, 0x5c, 0xaf, 0xe5, 0x2e, 0x49, 0x72, 0x8b, 0x05, 0x16, 0x9d, 0x42, 0x2d, 0xa4,
	0xfe, 0x8c, 0x32, 0x31, 0x3c, 0x65, 0x9c, 0x47, 0xe8, 0x19, 0x40, 0xa6, 0x2a, 0x7a, 0x51, 0x91,
	0x7b, 0x51, 0x4a, 0x7b, 0x51, 0xcf, 0xb2, 0x6f, 0x7c, 0x66, 0x8c, 0x0b, 0x1d, 0x47, 0xd0, 0x74,
	0xdc, 0xeb, 0x01, 0x76, 0x7b, 0x57, 0x93, 0x01, 0xc6, 0x23, 0xac, 0x29, 0x48, 0x87, 0xaf, 0x1c,
	0xd7, 0x1b, 0x0f, 0x87, 0x8e, 0xed, 0x0c, 0xdc, 0xeb, 0x89, 0xd5, 0xbb, 0xea, 0xb9, 0xf6, 0x40,
	0x53, 0xd1, 0x09, 0x1c, 0x3b, 0xee, 0xcf, 0x23, 0xc7, 0x1e, 0x4c, 0xec, 0xd1, 0xd5, 0x95, 0xe3,
	0x39, 0x23, 0x57, 0x2b, 0x1b, 0x21, 0x9c, 0xbc, 0xa4, 0x5c, 0xaa, 0x7a, 0xe5, 0x95, 0x07, 0x9a,
	0xec, 0x95, 0xb8, 0xb2, 0xb2, 0xcf, 0x95, 0xa5, 0x57, 0x27, 0xbb, 0x9d, 0xa6, 0x8c, 0xff, 0x15,
	0x38, 0xdd, 0x3e, 0x2e, 0x77, 0x67, 0x08, 0xd5, 0x84, 0x13, 0x4e, 0x73, 0x73, 0x5e, 0x7c, 0x64,
	0xb6, 0x77, 0x50, 0x4d, 0x2f, 0xe5, 0xe1, 0x8c, 0x2e, 0xb5, 0x5a, 0xdd, 0x68, 0x75, 0x17, 0x2a,
	0xe9, 0xe8, 0x0a, 0x03, 0x3e, 0x3d, 0xf4, 0x58, 0x80, 0x8d, 0xd7, 0x50, 0x15, 0xe2, 0xa8, 0x01,
	0x07, 0x63, 0xf7, 0x95, 0x3b, 0xfa, 0xc5, 0xd5, 0x4a, 0x69, 0xe0, 0x8d, 0x6d, 0x7b, 0xe0, 0x79,
	0x9a, 0xb2, 0xc3, 0x05, 0x15, 0x9d, 0x02, 0xda, 0x70, 0x61, 0x38, 0x76, 0xfb, 0x9e, 0x56, 0x36,
	0xfe, 0x53, 0xa1, 0x21, 0x1d, 0x82, 0x7e, 0xcc, 0x47, 0xe9, 0x73, 0xfb, 0x9a, 0x0d, 0xd5, 0xe6,
	0xf0, 0xa8, 0x85, 0x87, 0xb4, 0x1e, 0x1e, 0x74, 0x01, 0x47, 0x34, 0x5a, 0xd2, 0x30, 0xbe, 0xa3,
	0x02, 0x5c, 0x2e, 0x80, 0x1b, 0xab, 0x7c, 0x0a, 0x5f, 0xaf, 0xc8, 0xca, 0xde, 0x2b, 0xb2, 0x07,
	0x4d, 0x32, 0x8b, 0x19, 0x99, 0xf8, 0x84, 0x93, 0xc9, 0x82, 0x85, 0x7a, 0x55, 0x70, 0xbf, 0xd9,
	0xba, 0x55, 0x2f, 0x05, 0xf5, 0x09, 0x27, 0x63, 0x16, 0xe2, 0x23, 0x22, 0x45, 0xe8, 0x7b, 0x80,
	0xb5, 0x84, 0x5e, 0x13, 0x74, 0xfd, 0x63, 0x74, 0x5c, 0x7f, 0xe4, 0x1a, 0xcf, 0xa1, 0x96, 0x55,
	0x83, 0xce, 0xa0, 0xba, 0x24, 0xe1, 0x82, 0x16, 0xf6, 0x8a, 0x82, 0xb3, 0x44, 0xe7, 0x1f, 0x15,
	0x1a, 0x52, 0x67, 0xd1, 0x6f, 0x00, 0xeb, 0x95, 0x89, 0xbe, 0xdd, 0x6b, 0xab, 0xb7, 0xbe, 0xdb,
	0x6f, 0xf3, 0xa6, 0xe2, 0xeb, 0x85, 0xb2, 0x5b, 0xbc, 0xb0, 0x25, 0x77, 0x8b, 0xef, 0xd8, 0x67,
	0x33, 0x68, 0x6e, 0x3e, 0x08, 0xf4, 0x6c, 0x9f, 0x47, 0x93, 0x1d, 0xf2, 0x7c, 0xff, 0xf7, 0x65,
	0x11, 0x68, 0xc5, 0x6c, 0x26, 0x08, 0x33, 0xba, 0x4d, 0xfa, 0xd5, 0x9e, 0x05, 0xfc, 0x76, 0x71,
	0x93, 0x7a, 0xd4, 0x7e, 0x1b, 0x44, 0x74, 0x1a, 0x27, 0xef, 0x12, 0x4e, 0x45, 0x70, 0x41, 0xee,
	0x82, 0xf6, 0x8c, 0x46, 0xe2, 0x3f, 0x7b, 0x7b, 0xf3, 0x67, 0xc4, 0x0f, 0x52, 0x78, 0x53, 0x13,
	0x88, 0xee, 0x87, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x3c, 0x71, 0x11, 0x6b, 0x08, 0x00, 0x00,
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
	// GetHistory returns the transaction history for an account,
	// with additional off-chain data if available.
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
	// GetHistory returns the transaction history for an account,
	// with additional off-chain data if available.
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
