// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/v3/account_service.proto

package account

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	v3 "github.com/kinecosystem/agora-api/genproto/common/v3"
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

type CreateAccountResponse_Result int32

const (
	CreateAccountResponse_OK     CreateAccountResponse_Result = 0
	CreateAccountResponse_EXISTS CreateAccountResponse_Result = 1
)

var CreateAccountResponse_Result_name = map[int32]string{
	0: "OK",
	1: "EXISTS",
}

var CreateAccountResponse_Result_value = map[string]int32{
	"OK":     0,
	"EXISTS": 1,
}

func (x CreateAccountResponse_Result) String() string {
	return proto.EnumName(CreateAccountResponse_Result_name, int32(x))
}

func (CreateAccountResponse_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb584e8026f725ac, []int{2, 0}
}

type GetAccountInfoResponse_Result int32

const (
	GetAccountInfoResponse_OK GetAccountInfoResponse_Result = 0
	// The specified account could not be found.
	GetAccountInfoResponse_NOT_FOUND GetAccountInfoResponse_Result = 1
)

var GetAccountInfoResponse_Result_name = map[int32]string{
	0: "OK",
	1: "NOT_FOUND",
}

var GetAccountInfoResponse_Result_value = map[string]int32{
	"OK":        0,
	"NOT_FOUND": 1,
}

func (x GetAccountInfoResponse_Result) String() string {
	return proto.EnumName(GetAccountInfoResponse_Result_name, int32(x))
}

func (GetAccountInfoResponse_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb584e8026f725ac, []int{4, 0}
}

type Events_Result int32

const (
	Events_OK Events_Result = 0
	// The specified account could not be found.
	Events_NOT_FOUND Events_Result = 1
)

var Events_Result_name = map[int32]string{
	0: "OK",
	1: "NOT_FOUND",
}

var Events_Result_value = map[string]int32{
	"OK":        0,
	"NOT_FOUND": 1,
}

func (x Events_Result) String() string {
	return proto.EnumName(Events_Result_name, int32(x))
}

func (Events_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb584e8026f725ac, []int{6, 0}
}

type AccountInfo struct {
	AccountId *v3.StellarAccountId `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// The last known sequence number of the account.
	SequenceNumber int64 `protobuf:"varint,2,opt,name=sequence_number,json=sequenceNumber,proto3" json:"sequence_number,omitempty"`
	// The last known balance, in quarks, of the account.
	Balance              int64    `protobuf:"varint,3,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountInfo) Reset()         { *m = AccountInfo{} }
func (m *AccountInfo) String() string { return proto.CompactTextString(m) }
func (*AccountInfo) ProtoMessage()    {}
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb584e8026f725ac, []int{0}
}

func (m *AccountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountInfo.Unmarshal(m, b)
}
func (m *AccountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountInfo.Marshal(b, m, deterministic)
}
func (m *AccountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountInfo.Merge(m, src)
}
func (m *AccountInfo) XXX_Size() int {
	return xxx_messageInfo_AccountInfo.Size(m)
}
func (m *AccountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AccountInfo proto.InternalMessageInfo

func (m *AccountInfo) GetAccountId() *v3.StellarAccountId {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *AccountInfo) GetSequenceNumber() int64 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *AccountInfo) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type CreateAccountRequest struct {
	// Account to be created
	AccountId            *v3.StellarAccountId `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreateAccountRequest) Reset()         { *m = CreateAccountRequest{} }
func (m *CreateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAccountRequest) ProtoMessage()    {}
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb584e8026f725ac, []int{1}
}

func (m *CreateAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountRequest.Unmarshal(m, b)
}
func (m *CreateAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountRequest.Marshal(b, m, deterministic)
}
func (m *CreateAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountRequest.Merge(m, src)
}
func (m *CreateAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAccountRequest.Size(m)
}
func (m *CreateAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountRequest proto.InternalMessageInfo

func (m *CreateAccountRequest) GetAccountId() *v3.StellarAccountId {
	if m != nil {
		return m.AccountId
	}
	return nil
}

type CreateAccountResponse struct {
	Result CreateAccountResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=kin.agora.account.v3.CreateAccountResponse_Result" json:"result,omitempty"`
	// Will be present if the account was created or already existed.
	AccountInfo          *AccountInfo `protobuf:"bytes,2,opt,name=account_info,json=accountInfo,proto3" json:"account_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateAccountResponse) Reset()         { *m = CreateAccountResponse{} }
func (m *CreateAccountResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAccountResponse) ProtoMessage()    {}
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb584e8026f725ac, []int{2}
}

func (m *CreateAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountResponse.Unmarshal(m, b)
}
func (m *CreateAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountResponse.Marshal(b, m, deterministic)
}
func (m *CreateAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountResponse.Merge(m, src)
}
func (m *CreateAccountResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAccountResponse.Size(m)
}
func (m *CreateAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountResponse proto.InternalMessageInfo

func (m *CreateAccountResponse) GetResult() CreateAccountResponse_Result {
	if m != nil {
		return m.Result
	}
	return CreateAccountResponse_OK
}

func (m *CreateAccountResponse) GetAccountInfo() *AccountInfo {
	if m != nil {
		return m.AccountInfo
	}
	return nil
}

type GetAccountInfoRequest struct {
	AccountId            *v3.StellarAccountId `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetAccountInfoRequest) Reset()         { *m = GetAccountInfoRequest{} }
func (m *GetAccountInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountInfoRequest) ProtoMessage()    {}
func (*GetAccountInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb584e8026f725ac, []int{3}
}

func (m *GetAccountInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountInfoRequest.Unmarshal(m, b)
}
func (m *GetAccountInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetAccountInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountInfoRequest.Merge(m, src)
}
func (m *GetAccountInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetAccountInfoRequest.Size(m)
}
func (m *GetAccountInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountInfoRequest proto.InternalMessageInfo

func (m *GetAccountInfoRequest) GetAccountId() *v3.StellarAccountId {
	if m != nil {
		return m.AccountId
	}
	return nil
}

type GetAccountInfoResponse struct {
	Result GetAccountInfoResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=kin.agora.account.v3.GetAccountInfoResponse_Result" json:"result,omitempty"`
	// Present iff result == Result::OK
	AccountInfo          *AccountInfo `protobuf:"bytes,2,opt,name=account_info,json=accountInfo,proto3" json:"account_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetAccountInfoResponse) Reset()         { *m = GetAccountInfoResponse{} }
func (m *GetAccountInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetAccountInfoResponse) ProtoMessage()    {}
func (*GetAccountInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb584e8026f725ac, []int{4}
}

func (m *GetAccountInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountInfoResponse.Unmarshal(m, b)
}
func (m *GetAccountInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetAccountInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountInfoResponse.Merge(m, src)
}
func (m *GetAccountInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetAccountInfoResponse.Size(m)
}
func (m *GetAccountInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountInfoResponse proto.InternalMessageInfo

func (m *GetAccountInfoResponse) GetResult() GetAccountInfoResponse_Result {
	if m != nil {
		return m.Result
	}
	return GetAccountInfoResponse_OK
}

func (m *GetAccountInfoResponse) GetAccountInfo() *AccountInfo {
	if m != nil {
		return m.AccountInfo
	}
	return nil
}

type GetEventsRequest struct {
	// The id of the account to stream events for
	AccountId            *v3.StellarAccountId `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetEventsRequest) Reset()         { *m = GetEventsRequest{} }
func (m *GetEventsRequest) String() string { return proto.CompactTextString(m) }
func (*GetEventsRequest) ProtoMessage()    {}
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb584e8026f725ac, []int{5}
}

func (m *GetEventsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEventsRequest.Unmarshal(m, b)
}
func (m *GetEventsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEventsRequest.Marshal(b, m, deterministic)
}
func (m *GetEventsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEventsRequest.Merge(m, src)
}
func (m *GetEventsRequest) XXX_Size() int {
	return xxx_messageInfo_GetEventsRequest.Size(m)
}
func (m *GetEventsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEventsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEventsRequest proto.InternalMessageInfo

func (m *GetEventsRequest) GetAccountId() *v3.StellarAccountId {
	if m != nil {
		return m.AccountId
	}
	return nil
}

type Events struct {
	Result               Events_Result `protobuf:"varint,1,opt,name=result,proto3,enum=kin.agora.account.v3.Events_Result" json:"result,omitempty"`
	Events               []*Event      `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Events) Reset()         { *m = Events{} }
func (m *Events) String() string { return proto.CompactTextString(m) }
func (*Events) ProtoMessage()    {}
func (*Events) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb584e8026f725ac, []int{6}
}

func (m *Events) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Events.Unmarshal(m, b)
}
func (m *Events) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Events.Marshal(b, m, deterministic)
}
func (m *Events) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Events.Merge(m, src)
}
func (m *Events) XXX_Size() int {
	return xxx_messageInfo_Events.Size(m)
}
func (m *Events) XXX_DiscardUnknown() {
	xxx_messageInfo_Events.DiscardUnknown(m)
}

var xxx_messageInfo_Events proto.InternalMessageInfo

func (m *Events) GetResult() Events_Result {
	if m != nil {
		return m.Result
	}
	return Events_OK
}

func (m *Events) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type Event struct {
	// Types that are valid to be assigned to Type:
	//	*Event_AccountUpdateEvent
	//	*Event_TransactionEvent
	Type                 isEvent_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb584e8026f725ac, []int{7}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

type isEvent_Type interface {
	isEvent_Type()
}

type Event_AccountUpdateEvent struct {
	AccountUpdateEvent *AccountUpdateEvent `protobuf:"bytes,1,opt,name=account_update_event,json=accountUpdateEvent,proto3,oneof"`
}

type Event_TransactionEvent struct {
	TransactionEvent *TransactionEvent `protobuf:"bytes,2,opt,name=transaction_event,json=transactionEvent,proto3,oneof"`
}

func (*Event_AccountUpdateEvent) isEvent_Type() {}

func (*Event_TransactionEvent) isEvent_Type() {}

func (m *Event) GetType() isEvent_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Event) GetAccountUpdateEvent() *AccountUpdateEvent {
	if x, ok := m.GetType().(*Event_AccountUpdateEvent); ok {
		return x.AccountUpdateEvent
	}
	return nil
}

func (m *Event) GetTransactionEvent() *TransactionEvent {
	if x, ok := m.GetType().(*Event_TransactionEvent); ok {
		return x.TransactionEvent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Event_AccountUpdateEvent)(nil),
		(*Event_TransactionEvent)(nil),
	}
}

// An event that gets sent when an account's information has changed.
type AccountUpdateEvent struct {
	// The account information most recently obtained by the service.
	AccountInfo          *AccountInfo `protobuf:"bytes,1,opt,name=account_info,json=accountInfo,proto3" json:"account_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *AccountUpdateEvent) Reset()         { *m = AccountUpdateEvent{} }
func (m *AccountUpdateEvent) String() string { return proto.CompactTextString(m) }
func (*AccountUpdateEvent) ProtoMessage()    {}
func (*AccountUpdateEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb584e8026f725ac, []int{8}
}

func (m *AccountUpdateEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountUpdateEvent.Unmarshal(m, b)
}
func (m *AccountUpdateEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountUpdateEvent.Marshal(b, m, deterministic)
}
func (m *AccountUpdateEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountUpdateEvent.Merge(m, src)
}
func (m *AccountUpdateEvent) XXX_Size() int {
	return xxx_messageInfo_AccountUpdateEvent.Size(m)
}
func (m *AccountUpdateEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountUpdateEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AccountUpdateEvent proto.InternalMessageInfo

func (m *AccountUpdateEvent) GetAccountInfo() *AccountInfo {
	if m != nil {
		return m.AccountInfo
	}
	return nil
}

// An event that gets sent when a transaction related to an account has been
// successfully submitted to the blockchain.
type TransactionEvent struct {
	// The transaction envelope XDR.
	EnvelopeXdr []byte `protobuf:"bytes,1,opt,name=envelope_xdr,json=envelopeXdr,proto3" json:"envelope_xdr,omitempty"`
	// The transaction result XDR.
	ResultXdr            []byte   `protobuf:"bytes,2,opt,name=result_xdr,json=resultXdr,proto3" json:"result_xdr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionEvent) Reset()         { *m = TransactionEvent{} }
func (m *TransactionEvent) String() string { return proto.CompactTextString(m) }
func (*TransactionEvent) ProtoMessage()    {}
func (*TransactionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb584e8026f725ac, []int{9}
}

func (m *TransactionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionEvent.Unmarshal(m, b)
}
func (m *TransactionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionEvent.Marshal(b, m, deterministic)
}
func (m *TransactionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionEvent.Merge(m, src)
}
func (m *TransactionEvent) XXX_Size() int {
	return xxx_messageInfo_TransactionEvent.Size(m)
}
func (m *TransactionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionEvent proto.InternalMessageInfo

func (m *TransactionEvent) GetEnvelopeXdr() []byte {
	if m != nil {
		return m.EnvelopeXdr
	}
	return nil
}

func (m *TransactionEvent) GetResultXdr() []byte {
	if m != nil {
		return m.ResultXdr
	}
	return nil
}

func init() {
	proto.RegisterEnum("kin.agora.account.v3.CreateAccountResponse_Result", CreateAccountResponse_Result_name, CreateAccountResponse_Result_value)
	proto.RegisterEnum("kin.agora.account.v3.GetAccountInfoResponse_Result", GetAccountInfoResponse_Result_name, GetAccountInfoResponse_Result_value)
	proto.RegisterEnum("kin.agora.account.v3.Events_Result", Events_Result_name, Events_Result_value)
	proto.RegisterType((*AccountInfo)(nil), "kin.agora.account.v3.AccountInfo")
	proto.RegisterType((*CreateAccountRequest)(nil), "kin.agora.account.v3.CreateAccountRequest")
	proto.RegisterType((*CreateAccountResponse)(nil), "kin.agora.account.v3.CreateAccountResponse")
	proto.RegisterType((*GetAccountInfoRequest)(nil), "kin.agora.account.v3.GetAccountInfoRequest")
	proto.RegisterType((*GetAccountInfoResponse)(nil), "kin.agora.account.v3.GetAccountInfoResponse")
	proto.RegisterType((*GetEventsRequest)(nil), "kin.agora.account.v3.GetEventsRequest")
	proto.RegisterType((*Events)(nil), "kin.agora.account.v3.Events")
	proto.RegisterType((*Event)(nil), "kin.agora.account.v3.Event")
	proto.RegisterType((*AccountUpdateEvent)(nil), "kin.agora.account.v3.AccountUpdateEvent")
	proto.RegisterType((*TransactionEvent)(nil), "kin.agora.account.v3.TransactionEvent")
}

func init() { proto.RegisterFile("account/v3/account_service.proto", fileDescriptor_eb584e8026f725ac) }

var fileDescriptor_eb584e8026f725ac = []byte{
	// 711 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xed, 0x38, 0x5f, 0xdd, 0x2f, 0x37, 0x69, 0x31, 0xa3, 0x16, 0xa2, 0x10, 0x89, 0x10, 0x44,
	0x15, 0x28, 0x75, 0xaa, 0x64, 0x59, 0x09, 0xa9, 0xa6, 0x3f, 0x94, 0x4a, 0x69, 0xe5, 0xb4, 0xa8,
	0x42, 0x48, 0xd1, 0xc4, 0x9e, 0xa6, 0x56, 0x9d, 0x99, 0x60, 0x4f, 0x2c, 0xca, 0xaa, 0x6b, 0x96,
	0x3c, 0x02, 0x0b, 0xc4, 0x0b, 0xf0, 0x06, 0x48, 0x6c, 0x78, 0xa2, 0xae, 0x50, 0xc6, 0x76, 0x93,
	0xb8, 0xa6, 0xb4, 0x12, 0xdd, 0x8d, 0x67, 0xce, 0x39, 0xf7, 0xce, 0xb9, 0x73, 0xaf, 0xa1, 0x4c,
	0x2c, 0x8b, 0x0f, 0x98, 0xa8, 0x05, 0x8d, 0x5a, 0xb4, 0x6c, 0xfb, 0xd4, 0x0b, 0x1c, 0x8b, 0xea,
	0x7d, 0x8f, 0x0b, 0x8e, 0xe7, 0x4f, 0x1c, 0xa6, 0x93, 0x2e, 0xf7, 0x88, 0x1e, 0x01, 0xf4, 0xa0,
	0x51, 0xbc, 0x1f, 0x10, 0xd7, 0xb1, 0x89, 0xa0, 0xb5, 0x78, 0x11, 0xc2, 0x8b, 0x0b, 0x16, 0xef,
	0xf5, 0x38, 0x1b, 0xea, 0xf5, 0xb8, 0x4d, 0xdd, 0x70, 0xbb, 0xf2, 0x0d, 0x41, 0x6e, 0x2d, 0xa4,
	0x6f, 0xb3, 0x23, 0x8e, 0x9b, 0x00, 0x71, 0x38, 0xc7, 0x2e, 0xa0, 0x32, 0xaa, 0xe6, 0xea, 0x4f,
	0xf4, 0x51, 0xa8, 0x50, 0x45, 0x0f, 0x1a, 0x7a, 0x4b, 0x50, 0xd7, 0x25, 0x5e, 0x4c, 0xb6, 0x8d,
	0xff, 0xcf, 0x8d, 0xe9, 0x4f, 0x48, 0xd1, 0x90, 0x99, 0x25, 0xf1, 0x26, 0x5e, 0x82, 0x3b, 0x3e,
	0x7d, 0x3f, 0xa0, 0xcc, 0xa2, 0x6d, 0x36, 0xe8, 0x75, 0xa8, 0x57, 0x50, 0xca, 0xa8, 0x9a, 0x31,
	0x94, 0x15, 0x64, 0xce, 0xc5, 0x47, 0x4d, 0x79, 0x82, 0x4b, 0x30, 0xd3, 0x21, 0x2e, 0x61, 0x16,
	0x2d, 0x64, 0x2e, 0x40, 0xf1, 0x56, 0xe5, 0x08, 0xe6, 0x5f, 0x7a, 0x94, 0x08, 0x1a, 0x85, 0x34,
	0x87, 0x64, 0x5f, 0xfc, 0xeb, 0x94, 0x2b, 0x3f, 0x11, 0x2c, 0x24, 0x02, 0xf9, 0x7d, 0xce, 0x7c,
	0x8a, 0x5f, 0x83, 0xea, 0x51, 0x7f, 0xe0, 0x0a, 0x19, 0x65, 0xae, 0x5e, 0xd7, 0xd3, 0x6a, 0xa0,
	0xa7, 0x92, 0x75, 0x53, 0x32, 0xcd, 0x48, 0x01, 0xaf, 0x43, 0xfe, 0x22, 0x6b, 0x76, 0xc4, 0xa5,
	0x2b, 0xb9, 0xfa, 0xa3, 0x74, 0xc5, 0xb1, 0x0a, 0x99, 0x39, 0x32, 0xfa, 0xa8, 0x94, 0x40, 0x0d,
	0x75, 0xb1, 0x0a, 0xca, 0xee, 0x8e, 0x36, 0x85, 0x01, 0xd4, 0x8d, 0xc3, 0xed, 0xd6, 0x7e, 0x4b,
	0x43, 0x95, 0x2e, 0x2c, 0x6c, 0x51, 0x31, 0x4e, 0xbe, 0x25, 0xcb, 0x7e, 0x21, 0xb8, 0x97, 0x8c,
	0x14, 0x79, 0xb6, 0x93, 0xf0, 0xac, 0x91, 0x7e, 0xc3, 0x74, 0xf6, 0xed, 0x98, 0xf6, 0xf0, 0x92,
	0x69, 0xb3, 0x90, 0x6d, 0xee, 0xee, 0xb7, 0x37, 0x77, 0x0f, 0x9a, 0xeb, 0x1a, 0xaa, 0x74, 0x40,
	0xdb, 0xa2, 0x62, 0x23, 0xa0, 0x4c, 0xf8, 0xb7, 0x65, 0xd9, 0x57, 0x04, 0x6a, 0x18, 0x01, 0xaf,
	0x26, 0x2c, 0x7a, 0x9c, 0x7e, 0x9f, 0x10, 0x9d, 0xb4, 0xe4, 0x05, 0xa8, 0x54, 0x1e, 0x14, 0x94,
	0x72, 0xa6, 0x9a, 0xab, 0x3f, 0xb8, 0x82, 0x6c, 0x64, 0xcf, 0x0d, 0xf5, 0x33, 0xca, 0x68, 0x67,
	0xc8, 0x8c, 0x58, 0x7f, 0x37, 0xe3, 0x07, 0x82, 0x69, 0xc9, 0xc6, 0xef, 0x60, 0x3e, 0xb6, 0x60,
	0xd0, 0x1f, 0x8e, 0x96, 0xb6, 0xd4, 0x88, 0xcc, 0xa8, 0x5e, 0x59, 0x85, 0x03, 0x49, 0x90, 0x3a,
	0xaf, 0xa6, 0x4c, 0x4c, 0x2e, 0xed, 0xe2, 0x03, 0xb8, 0x2b, 0x3c, 0xc2, 0x7c, 0x62, 0x09, 0x87,
	0xb3, 0x48, 0x3a, 0x2c, 0xf0, 0x62, 0xba, 0xf4, 0xfe, 0x08, 0x1e, 0x0b, 0x6b, 0x22, 0xb1, 0x67,
	0xa8, 0xf0, 0x9f, 0x38, 0xed, 0xd3, 0x8a, 0x0d, 0xf8, 0x72, 0x2a, 0xb8, 0x99, 0x78, 0x50, 0xe8,
	0x9a, 0x0f, 0x6a, 0xac, 0xa6, 0x13, 0x4f, 0xcb, 0x05, 0x2d, 0x99, 0x15, 0x5e, 0x86, 0x3c, 0x65,
	0x01, 0x75, 0x79, 0x9f, 0xb6, 0x3f, 0xd8, 0x9e, 0x8c, 0x91, 0x37, 0xe0, 0xdc, 0x98, 0xf9, 0x38,
	0xad, 0xa1, 0xc2, 0xd9, 0x9e, 0x99, 0x8b, 0xcf, 0x0f, 0x6d, 0x0f, 0x3f, 0x05, 0x08, 0x4b, 0x2b,
	0xc1, 0xca, 0x38, 0x78, 0x6a, 0x08, 0xce, 0x86, 0xa7, 0x87, 0xb6, 0x57, 0xff, 0xae, 0xc0, 0x4c,
	0x94, 0x14, 0x3e, 0x86, 0xd9, 0x89, 0xb9, 0x83, 0x9f, 0x5d, 0x6b, 0x38, 0xc9, 0xc7, 0x5d, 0x5c,
	0xba, 0xc1, 0x20, 0xc3, 0x27, 0x30, 0x37, 0xd9, 0xad, 0x78, 0xe9, 0x7a, 0x3d, 0x1d, 0xc6, 0x7a,
	0x7e, 0x93, 0x01, 0x80, 0x5b, 0x90, 0xbd, 0x68, 0x45, 0xbc, 0xf8, 0x47, 0xea, 0x44, 0xaf, 0x16,
	0x4b, 0x57, 0x35, 0xd0, 0x0a, 0x32, 0x06, 0x50, 0xe2, 0x5e, 0x77, 0x0c, 0xd4, 0xa5, 0x6c, 0x0c,
	0xf8, 0x76, 0xb3, 0xeb, 0x88, 0xe3, 0x41, 0x67, 0xd8, 0xd3, 0xb5, 0x13, 0x87, 0x51, 0x8b, 0xfb,
	0xa7, 0xbe, 0xa0, 0xbd, 0x9a, 0x44, 0x2f, 0x93, 0xbe, 0xb3, 0xec, 0x30, 0x41, 0x3d, 0x46, 0xdc,
	0x5a, 0x97, 0x32, 0xf9, 0x27, 0xad, 0x8d, 0x7e, 0xd8, 0xab, 0xd1, 0xf2, 0x8b, 0x92, 0x5f, 0xdb,
	0x33, 0xa2, 0x5b, 0xbd, 0x69, 0x74, 0x54, 0x09, 0x6c, 0xfc, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x41,
	0xe5, 0x5a, 0xa1, 0xdc, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	// CreateAccount creates an account using a the service's configured seed
	// account.
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	// GetAccountInfo returns the information of a specified account.
	GetAccountInfo(ctx context.Context, in *GetAccountInfoRequest, opts ...grpc.CallOption) (*GetAccountInfoResponse, error)
	// GetEvents returns a stream of events related to the specified account.
	//
	// Note: Only events occurring after stream initiation will be returned.
	GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (Account_GetEventsClient, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/kin.agora.account.v3.Account/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetAccountInfo(ctx context.Context, in *GetAccountInfoRequest, opts ...grpc.CallOption) (*GetAccountInfoResponse, error) {
	out := new(GetAccountInfoResponse)
	err := c.cc.Invoke(ctx, "/kin.agora.account.v3.Account/GetAccountInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (Account_GetEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Account_serviceDesc.Streams[0], "/kin.agora.account.v3.Account/GetEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &accountGetEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Account_GetEventsClient interface {
	Recv() (*Events, error)
	grpc.ClientStream
}

type accountGetEventsClient struct {
	grpc.ClientStream
}

func (x *accountGetEventsClient) Recv() (*Events, error) {
	m := new(Events)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	// CreateAccount creates an account using a the service's configured seed
	// account.
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	// GetAccountInfo returns the information of a specified account.
	GetAccountInfo(context.Context, *GetAccountInfoRequest) (*GetAccountInfoResponse, error)
	// GetEvents returns a stream of events related to the specified account.
	//
	// Note: Only events occurring after stream initiation will be returned.
	GetEvents(*GetEventsRequest, Account_GetEventsServer) error
}

// UnimplementedAccountServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (*UnimplementedAccountServer) CreateAccount(ctx context.Context, req *CreateAccountRequest) (*CreateAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (*UnimplementedAccountServer) GetAccountInfo(ctx context.Context, req *GetAccountInfoRequest) (*GetAccountInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccountInfo not implemented")
}
func (*UnimplementedAccountServer) GetEvents(req *GetEventsRequest, srv Account_GetEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kin.agora.account.v3.Account/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetAccountInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).GetAccountInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kin.agora.account.v3.Account/GetAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetAccountInfo(ctx, req.(*GetAccountInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_GetEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetEventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AccountServer).GetEvents(m, &accountGetEventsServer{stream})
}

type Account_GetEventsServer interface {
	Send(*Events) error
	grpc.ServerStream
}

type accountGetEventsServer struct {
	grpc.ServerStream
}

func (x *accountGetEventsServer) Send(m *Events) error {
	return x.ServerStream.SendMsg(m)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kin.agora.account.v3.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _Account_CreateAccount_Handler,
		},
		{
			MethodName: "GetAccountInfo",
			Handler:    _Account_GetAccountInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEvents",
			Handler:       _Account_GetEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "account/v3/account_service.proto",
}
