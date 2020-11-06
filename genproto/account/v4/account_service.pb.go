// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/v4/account_service.proto

package account

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

type CreateAccountResponse_Result int32

const (
	CreateAccountResponse_OK     CreateAccountResponse_Result = 0
	CreateAccountResponse_EXISTS CreateAccountResponse_Result = 1
	// Indicates that the service will not subsidize the transaction, and that
	// the caller should fund the transaction themselves.
	CreateAccountResponse_PAYER_REQUIRED CreateAccountResponse_Result = 2
	// Indicates the nonce/blockhash used in the transaction is invalid, and should
	// be refetched
	CreateAccountResponse_BAD_NONCE CreateAccountResponse_Result = 3
)

var CreateAccountResponse_Result_name = map[int32]string{
	0: "OK",
	1: "EXISTS",
	2: "PAYER_REQUIRED",
	3: "BAD_NONCE",
}

var CreateAccountResponse_Result_value = map[string]int32{
	"OK":             0,
	"EXISTS":         1,
	"PAYER_REQUIRED": 2,
	"BAD_NONCE":      3,
}

func (x CreateAccountResponse_Result) String() string {
	return proto.EnumName(CreateAccountResponse_Result_name, int32(x))
}

func (CreateAccountResponse_Result) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3f5e8e0cd9d934c3, []int{2, 0}
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
	return fileDescriptor_3f5e8e0cd9d934c3, []int{4, 0}
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
	return fileDescriptor_3f5e8e0cd9d934c3, []int{8, 0}
}

type AccountInfo struct {
	AccountId *v4.SolanaAccountId `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// The last known balance, in quarks, of the account.
	Balance              int64    `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountInfo) Reset()         { *m = AccountInfo{} }
func (m *AccountInfo) String() string { return proto.CompactTextString(m) }
func (*AccountInfo) ProtoMessage()    {}
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f5e8e0cd9d934c3, []int{0}
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

func (m *AccountInfo) GetAccountId() *v4.SolanaAccountId {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *AccountInfo) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type CreateAccountRequest struct {
	// Transaction should contain the following instructions:
	//
	//   1. SystemProgram::CreateAccount()
	//   2. SplTokenProgram::InitializeAccount()
	//   3. SplTokenProgram::SetAuthority()
	//
	// (3) only needs to be set if the service is subsidizing the
	// account creation. In that case, a SetAuthority() instruction
	// should be included that sets the CloseAuthority of the account
	// to the subsidizer. This is to prevent farming of Sol by creating
	// accounts. It should be noted that an account can only be closed
	// if there is zero kin in the account.
	//
	// If the parameters are not for the Kin token, or there are
	// any other instructions, InvalidArgument will be returned.
	Transaction          *v4.Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Commitment           v4.Commitment   `protobuf:"varint,2,opt,name=commitment,proto3,enum=kin.agora.common.v4.Commitment" json:"commitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateAccountRequest) Reset()         { *m = CreateAccountRequest{} }
func (m *CreateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAccountRequest) ProtoMessage()    {}
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f5e8e0cd9d934c3, []int{1}
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

func (m *CreateAccountRequest) GetTransaction() *v4.Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *CreateAccountRequest) GetCommitment() v4.Commitment {
	if m != nil {
		return m.Commitment
	}
	return v4.Commitment_RECENT
}

type CreateAccountResponse struct {
	Result CreateAccountResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=kin.agora.account.v4.CreateAccountResponse_Result" json:"result,omitempty"`
	// Present iff the account was created or already existed.
	AccountInfo          *AccountInfo `protobuf:"bytes,2,opt,name=account_info,json=accountInfo,proto3" json:"account_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateAccountResponse) Reset()         { *m = CreateAccountResponse{} }
func (m *CreateAccountResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAccountResponse) ProtoMessage()    {}
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f5e8e0cd9d934c3, []int{2}
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
	AccountId            *v4.SolanaAccountId `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Commitment           v4.Commitment       `protobuf:"varint,2,opt,name=commitment,proto3,enum=kin.agora.common.v4.Commitment" json:"commitment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetAccountInfoRequest) Reset()         { *m = GetAccountInfoRequest{} }
func (m *GetAccountInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetAccountInfoRequest) ProtoMessage()    {}
func (*GetAccountInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f5e8e0cd9d934c3, []int{3}
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

func (m *GetAccountInfoRequest) GetAccountId() *v4.SolanaAccountId {
	if m != nil {
		return m.AccountId
	}
	return nil
}

func (m *GetAccountInfoRequest) GetCommitment() v4.Commitment {
	if m != nil {
		return m.Commitment
	}
	return v4.Commitment_RECENT
}

type GetAccountInfoResponse struct {
	Result GetAccountInfoResponse_Result `protobuf:"varint,1,opt,name=result,proto3,enum=kin.agora.account.v4.GetAccountInfoResponse_Result" json:"result,omitempty"`
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
	return fileDescriptor_3f5e8e0cd9d934c3, []int{4}
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

type ResolveTokenAccountsRequest struct {
	AccountId            *v4.SolanaAccountId `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ResolveTokenAccountsRequest) Reset()         { *m = ResolveTokenAccountsRequest{} }
func (m *ResolveTokenAccountsRequest) String() string { return proto.CompactTextString(m) }
func (*ResolveTokenAccountsRequest) ProtoMessage()    {}
func (*ResolveTokenAccountsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f5e8e0cd9d934c3, []int{5}
}

func (m *ResolveTokenAccountsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveTokenAccountsRequest.Unmarshal(m, b)
}
func (m *ResolveTokenAccountsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveTokenAccountsRequest.Marshal(b, m, deterministic)
}
func (m *ResolveTokenAccountsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveTokenAccountsRequest.Merge(m, src)
}
func (m *ResolveTokenAccountsRequest) XXX_Size() int {
	return xxx_messageInfo_ResolveTokenAccountsRequest.Size(m)
}
func (m *ResolveTokenAccountsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveTokenAccountsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveTokenAccountsRequest proto.InternalMessageInfo

func (m *ResolveTokenAccountsRequest) GetAccountId() *v4.SolanaAccountId {
	if m != nil {
		return m.AccountId
	}
	return nil
}

type ResolveTokenAccountsResponse struct {
	// Zero or more accounts that are owned by the provided account id.
	//
	// If the provided account is also a token account, it will be first in the list.
	// Otherwise, the list order should not be depended on, as there is no reliable way
	// to sort accounts based on creation time.
	TokenAccounts        []*v4.SolanaAccountId `protobuf:"bytes,1,rep,name=token_accounts,json=tokenAccounts,proto3" json:"token_accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ResolveTokenAccountsResponse) Reset()         { *m = ResolveTokenAccountsResponse{} }
func (m *ResolveTokenAccountsResponse) String() string { return proto.CompactTextString(m) }
func (*ResolveTokenAccountsResponse) ProtoMessage()    {}
func (*ResolveTokenAccountsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f5e8e0cd9d934c3, []int{6}
}

func (m *ResolveTokenAccountsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveTokenAccountsResponse.Unmarshal(m, b)
}
func (m *ResolveTokenAccountsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveTokenAccountsResponse.Marshal(b, m, deterministic)
}
func (m *ResolveTokenAccountsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveTokenAccountsResponse.Merge(m, src)
}
func (m *ResolveTokenAccountsResponse) XXX_Size() int {
	return xxx_messageInfo_ResolveTokenAccountsResponse.Size(m)
}
func (m *ResolveTokenAccountsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveTokenAccountsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveTokenAccountsResponse proto.InternalMessageInfo

func (m *ResolveTokenAccountsResponse) GetTokenAccounts() []*v4.SolanaAccountId {
	if m != nil {
		return m.TokenAccounts
	}
	return nil
}

type GetEventsRequest struct {
	AccountId            *v4.SolanaAccountId `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetEventsRequest) Reset()         { *m = GetEventsRequest{} }
func (m *GetEventsRequest) String() string { return proto.CompactTextString(m) }
func (*GetEventsRequest) ProtoMessage()    {}
func (*GetEventsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f5e8e0cd9d934c3, []int{7}
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

func (m *GetEventsRequest) GetAccountId() *v4.SolanaAccountId {
	if m != nil {
		return m.AccountId
	}
	return nil
}

type Events struct {
	Result               Events_Result `protobuf:"varint,1,opt,name=result,proto3,enum=kin.agora.account.v4.Events_Result" json:"result,omitempty"`
	Events               []*Event      `protobuf:"bytes,2,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Events) Reset()         { *m = Events{} }
func (m *Events) String() string { return proto.CompactTextString(m) }
func (*Events) ProtoMessage()    {}
func (*Events) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f5e8e0cd9d934c3, []int{8}
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
	return fileDescriptor_3f5e8e0cd9d934c3, []int{9}
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
	return fileDescriptor_3f5e8e0cd9d934c3, []int{10}
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
	Transaction          *v4.Transaction      `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	TransactionError     *v4.TransactionError `protobuf:"bytes,2,opt,name=transaction_error,json=transactionError,proto3" json:"transaction_error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *TransactionEvent) Reset()         { *m = TransactionEvent{} }
func (m *TransactionEvent) String() string { return proto.CompactTextString(m) }
func (*TransactionEvent) ProtoMessage()    {}
func (*TransactionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f5e8e0cd9d934c3, []int{11}
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

func (m *TransactionEvent) GetTransaction() *v4.Transaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *TransactionEvent) GetTransactionError() *v4.TransactionError {
	if m != nil {
		return m.TransactionError
	}
	return nil
}

func init() {
	proto.RegisterEnum("kin.agora.account.v4.CreateAccountResponse_Result", CreateAccountResponse_Result_name, CreateAccountResponse_Result_value)
	proto.RegisterEnum("kin.agora.account.v4.GetAccountInfoResponse_Result", GetAccountInfoResponse_Result_name, GetAccountInfoResponse_Result_value)
	proto.RegisterEnum("kin.agora.account.v4.Events_Result", Events_Result_name, Events_Result_value)
	proto.RegisterType((*AccountInfo)(nil), "kin.agora.account.v4.AccountInfo")
	proto.RegisterType((*CreateAccountRequest)(nil), "kin.agora.account.v4.CreateAccountRequest")
	proto.RegisterType((*CreateAccountResponse)(nil), "kin.agora.account.v4.CreateAccountResponse")
	proto.RegisterType((*GetAccountInfoRequest)(nil), "kin.agora.account.v4.GetAccountInfoRequest")
	proto.RegisterType((*GetAccountInfoResponse)(nil), "kin.agora.account.v4.GetAccountInfoResponse")
	proto.RegisterType((*ResolveTokenAccountsRequest)(nil), "kin.agora.account.v4.ResolveTokenAccountsRequest")
	proto.RegisterType((*ResolveTokenAccountsResponse)(nil), "kin.agora.account.v4.ResolveTokenAccountsResponse")
	proto.RegisterType((*GetEventsRequest)(nil), "kin.agora.account.v4.GetEventsRequest")
	proto.RegisterType((*Events)(nil), "kin.agora.account.v4.Events")
	proto.RegisterType((*Event)(nil), "kin.agora.account.v4.Event")
	proto.RegisterType((*AccountUpdateEvent)(nil), "kin.agora.account.v4.AccountUpdateEvent")
	proto.RegisterType((*TransactionEvent)(nil), "kin.agora.account.v4.TransactionEvent")
}

func init() { proto.RegisterFile("account/v4/account_service.proto", fileDescriptor_3f5e8e0cd9d934c3) }

var fileDescriptor_3f5e8e0cd9d934c3 = []byte{
	// 791 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xd1, 0x6e, 0xe2, 0x46,
	0x14, 0xcd, 0x40, 0xe3, 0x94, 0x4b, 0x40, 0xee, 0x88, 0xb4, 0x88, 0x20, 0x85, 0xba, 0x6d, 0x14,
	0x35, 0xad, 0x49, 0x09, 0x6f, 0x51, 0x1b, 0xe1, 0x40, 0xd3, 0x34, 0x2d, 0xa4, 0x06, 0xaa, 0xb6,
	0xaa, 0x84, 0x26, 0x30, 0x21, 0x16, 0xc6, 0x43, 0xed, 0x01, 0x29, 0x95, 0x56, 0xda, 0xe7, 0x7d,
	0xdc, 0x4f, 0x58, 0x69, 0x77, 0x7f, 0x20, 0x9f, 0xb0, 0x8f, 0xfb, 0x25, 0xfb, 0x09, 0x79, 0x5a,
	0x61, 0x0f, 0x60, 0x1c, 0x2f, 0x21, 0xda, 0xe4, 0x6d, 0xb8, 0xdc, 0x73, 0xee, 0x9d, 0x7b, 0x66,
	0xce, 0x18, 0x72, 0xa4, 0xdd, 0x66, 0x43, 0x8b, 0xe7, 0x47, 0xc5, 0xbc, 0x58, 0xb6, 0x1c, 0x6a,
	0x8f, 0x8c, 0x36, 0x55, 0x07, 0x36, 0xe3, 0x0c, 0xa7, 0x7a, 0x86, 0xa5, 0x92, 0x2e, 0xb3, 0x89,
	0x2a, 0x12, 0xd4, 0x51, 0x31, 0xf3, 0xc5, 0x88, 0x98, 0x46, 0x87, 0x70, 0x9a, 0x9f, 0x2c, 0xbc,
	0xf4, 0xcc, 0x46, 0x9b, 0xf5, 0xfb, 0xcc, 0x1a, 0xf3, 0xf5, 0x59, 0x87, 0x9a, 0x5e, 0x58, 0xf9,
	0x1f, 0xe2, 0x25, 0x0f, 0x7d, 0x62, 0x5d, 0x30, 0xfc, 0x3b, 0xc0, 0xa4, 0x9a, 0xd1, 0x49, 0xa3,
	0x1c, 0xda, 0x89, 0x17, 0xbe, 0x56, 0x67, 0x95, 0x3c, 0x12, 0x75, 0x54, 0x54, 0xeb, 0xcc, 0x24,
	0x16, 0x99, 0x60, 0x3b, 0xda, 0xa7, 0x37, 0xda, 0xea, 0x33, 0x14, 0x91, 0x91, 0x1e, 0x23, 0x93,
	0x20, 0xce, 0xc2, 0xda, 0x39, 0x31, 0x89, 0xd5, 0xa6, 0xe9, 0x48, 0x0e, 0xed, 0x44, 0xb5, 0xc8,
	0x1e, 0xd2, 0x27, 0x21, 0xe5, 0x25, 0x82, 0xd4, 0x91, 0x4d, 0x09, 0xa7, 0x82, 0x46, 0xa7, 0xff,
	0x0d, 0xa9, 0xc3, 0xf1, 0x6f, 0x10, 0xe7, 0x36, 0xb1, 0x1c, 0xd2, 0xe6, 0x06, 0xb3, 0x44, 0x1b,
	0xb9, 0xd0, 0x36, 0x1a, 0xb3, 0x3c, 0x5f, 0x0b, 0x7e, 0x38, 0x3e, 0x04, 0x18, 0xe7, 0x1b, 0xbc,
	0x4f, 0x2d, 0xee, 0xf6, 0x91, 0x2c, 0x6c, 0x85, 0x92, 0x1d, 0x4d, 0xd3, 0x74, 0x1f, 0x44, 0x79,
	0x87, 0x60, 0x23, 0xd0, 0xa7, 0x33, 0x60, 0x96, 0x43, 0xf1, 0xaf, 0x20, 0xd9, 0xd4, 0x19, 0x9a,
	0xdc, 0xed, 0x31, 0x59, 0x28, 0xa8, 0x61, 0xa2, 0xa8, 0xa1, 0x60, 0x55, 0x77, 0x91, 0xba, 0x60,
	0xc0, 0x65, 0x58, 0x9f, 0x8e, 0xde, 0xba, 0x60, 0x6e, 0xa3, 0xf1, 0xc2, 0x97, 0xe1, 0x8c, 0x3e,
	0xcd, 0xf4, 0x38, 0x99, 0xfd, 0x50, 0x0e, 0x41, 0xf2, 0x78, 0xb1, 0x04, 0x91, 0xda, 0xa9, 0xbc,
	0x82, 0x01, 0xa4, 0xca, 0x5f, 0x27, 0xf5, 0x46, 0x5d, 0x46, 0x18, 0x43, 0xf2, 0xac, 0xf4, 0x77,
	0x45, 0x6f, 0xe9, 0x95, 0x3f, 0x9a, 0x27, 0x7a, 0xa5, 0x2c, 0x47, 0x70, 0x02, 0x62, 0x5a, 0xa9,
	0xdc, 0xaa, 0xd6, 0xaa, 0x47, 0x15, 0x39, 0xaa, 0xbc, 0x46, 0xb0, 0x71, 0x4c, 0xb9, 0xbf, 0x80,
	0x50, 0xe5, 0x81, 0xcf, 0xc6, 0x47, 0xcb, 0xf2, 0x16, 0xc1, 0xe7, 0xc1, 0x4e, 0x85, 0x2e, 0xa7,
	0x01, 0x5d, 0xf6, 0xc3, 0xa7, 0x18, 0x8e, 0x7e, 0x1c, 0x61, 0xb6, 0x6e, 0x09, 0x93, 0x80, 0x58,
	0xb5, 0xd6, 0x68, 0xfd, 0x5c, 0x6b, 0x56, 0xcb, 0x32, 0x52, 0x4c, 0xd8, 0xd4, 0xa9, 0xc3, 0xcc,
	0x11, 0x6d, 0xb0, 0x1e, 0xb5, 0x04, 0x91, 0xf3, 0x38, 0xd3, 0x57, 0x7a, 0x90, 0x0d, 0xaf, 0x36,
	0x9d, 0x60, 0x92, 0x8f, 0xff, 0x68, 0x09, 0x88, 0x93, 0x46, 0xb9, 0xe8, 0xb2, 0x25, 0xf5, 0x04,
	0xf7, 0x93, 0x2a, 0x04, 0xe4, 0x63, 0xca, 0x2b, 0x23, 0xfa, 0x78, 0xfb, 0x79, 0x85, 0x40, 0xf2,
	0x0a, 0xe0, 0x83, 0x80, 0xf8, 0x5f, 0x85, 0x2b, 0xe5, 0x65, 0x07, 0xc5, 0xfe, 0x09, 0x24, 0xea,
	0xfe, 0x91, 0x8e, 0xb8, 0xfb, 0xdd, 0x5c, 0x00, 0xd6, 0x62, 0x37, 0x9a, 0xf4, 0x1c, 0x45, 0xe5,
	0xa7, 0x48, 0x17, 0xa8, 0xbb, 0x65, 0x7e, 0x83, 0x60, 0xd5, 0x45, 0xe3, 0x7f, 0x21, 0x35, 0x99,
	0xc0, 0x70, 0x30, 0x76, 0xea, 0x96, 0xcb, 0x21, 0x66, 0xb1, 0xb3, 0xf0, 0x7c, 0x35, 0x5d, 0x80,
	0xcb, 0xf3, 0xcb, 0x8a, 0x8e, 0xc9, 0xad, 0x28, 0x6e, 0xc2, 0x67, 0x3e, 0x13, 0x14, 0xd4, 0xde,
	0xd1, 0xdd, 0x0e, 0xa7, 0xf6, 0x59, 0xe9, 0x84, 0x58, 0xe6, 0x81, 0x98, 0x26, 0xc1, 0x27, 0xfc,
	0x6a, 0x40, 0x95, 0x0e, 0xe0, 0xdb, 0xad, 0xe0, 0x6a, 0xe0, 0xaa, 0xa0, 0x25, 0xaf, 0x8a, 0xdf,
	0xba, 0xfd, 0x97, 0xe6, 0x1a, 0x81, 0x1c, 0x6c, 0xeb, 0x81, 0x5f, 0x07, 0x3d, 0x30, 0x27, 0xdb,
	0x66, 0xb6, 0x98, 0xd3, 0x37, 0x77, 0x71, 0x56, 0xc6, 0xc9, 0xf3, 0x43, 0x1a, 0x47, 0x0a, 0xd7,
	0x51, 0x58, 0x13, 0xbb, 0xc3, 0x97, 0x90, 0x98, 0xb3, 0x7f, 0xfc, 0xed, 0x52, 0x6f, 0x84, 0x7b,
	0x49, 0x32, 0xbb, 0xf7, 0x78, 0x4f, 0x70, 0x0f, 0x92, 0xf3, 0x86, 0x86, 0x77, 0x97, 0xb3, 0x3d,
	0xaf, 0xd6, 0x77, 0xf7, 0xf1, 0x48, 0xfc, 0x04, 0x52, 0x61, 0xfe, 0x81, 0x7f, 0x08, 0x67, 0x59,
	0xe0, 0x6c, 0x99, 0xc2, 0x7d, 0x20, 0xa2, 0x7c, 0x1d, 0x62, 0x53, 0x47, 0xc1, 0xdb, 0x1f, 0xec,
	0x7c, 0xce, 0x72, 0x32, 0xd9, 0x45, 0x46, 0xb0, 0x87, 0x34, 0x13, 0xb2, 0xcc, 0xee, 0xfa, 0x92,
	0xba, 0xd4, 0xf2, 0x25, 0xfe, 0xf3, 0x63, 0xd7, 0xe0, 0x97, 0xc3, 0xf3, 0xf1, 0x59, 0xc8, 0xf7,
	0x0c, 0x8b, 0xb6, 0x99, 0x73, 0xe5, 0x70, 0xda, 0xcf, 0xbb, 0xd9, 0xdf, 0x93, 0x81, 0x91, 0xef,
	0x52, 0xcb, 0xfd, 0xae, 0xca, 0xcf, 0x3e, 0xdf, 0x0e, 0xc4, 0xf2, 0x45, 0x64, 0xbd, 0x74, 0xa6,
	0x89, 0x9d, 0xfc, 0x59, 0x3c, 0x97, 0xdc, 0xc4, 0xfd, 0xf7, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0b,
	0xfd, 0xf2, 0xfe, 0xea, 0x09, 0x00, 0x00,
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
	// CreateAccount creates a kin token account.
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	// GetAccountInfo returns the information of a specified account.
	GetAccountInfo(ctx context.Context, in *GetAccountInfoRequest, opts ...grpc.CallOption) (*GetAccountInfoResponse, error)
	// ResolveTokenAccounts resolves a set of Token Accounts owned by the specified account ID.
	ResolveTokenAccounts(ctx context.Context, in *ResolveTokenAccountsRequest, opts ...grpc.CallOption) (*ResolveTokenAccountsResponse, error)
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
	err := c.cc.Invoke(ctx, "/kin.agora.account.v4.Account/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetAccountInfo(ctx context.Context, in *GetAccountInfoRequest, opts ...grpc.CallOption) (*GetAccountInfoResponse, error) {
	out := new(GetAccountInfoResponse)
	err := c.cc.Invoke(ctx, "/kin.agora.account.v4.Account/GetAccountInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) ResolveTokenAccounts(ctx context.Context, in *ResolveTokenAccountsRequest, opts ...grpc.CallOption) (*ResolveTokenAccountsResponse, error) {
	out := new(ResolveTokenAccountsResponse)
	err := c.cc.Invoke(ctx, "/kin.agora.account.v4.Account/ResolveTokenAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (Account_GetEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Account_serviceDesc.Streams[0], "/kin.agora.account.v4.Account/GetEvents", opts...)
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
	// CreateAccount creates a kin token account.
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	// GetAccountInfo returns the information of a specified account.
	GetAccountInfo(context.Context, *GetAccountInfoRequest) (*GetAccountInfoResponse, error)
	// ResolveTokenAccounts resolves a set of Token Accounts owned by the specified account ID.
	ResolveTokenAccounts(context.Context, *ResolveTokenAccountsRequest) (*ResolveTokenAccountsResponse, error)
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
func (*UnimplementedAccountServer) ResolveTokenAccounts(ctx context.Context, req *ResolveTokenAccountsRequest) (*ResolveTokenAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveTokenAccounts not implemented")
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
		FullMethod: "/kin.agora.account.v4.Account/CreateAccount",
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
		FullMethod: "/kin.agora.account.v4.Account/GetAccountInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).GetAccountInfo(ctx, req.(*GetAccountInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_ResolveTokenAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveTokenAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).ResolveTokenAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kin.agora.account.v4.Account/ResolveTokenAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).ResolveTokenAccounts(ctx, req.(*ResolveTokenAccountsRequest))
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
	ServiceName: "kin.agora.account.v4.Account",
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
		{
			MethodName: "ResolveTokenAccounts",
			Handler:    _Account_ResolveTokenAccounts_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEvents",
			Handler:       _Account_GetEvents_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "account/v4/account_service.proto",
}
