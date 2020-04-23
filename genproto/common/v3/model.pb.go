// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/v3/model.proto

package common

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type AgoraData_TransactionType int32

const (
	AgoraData_UNKNOWN AgoraData_TransactionType = 0
	AgoraData_EARN    AgoraData_TransactionType = 1
	AgoraData_SPEND   AgoraData_TransactionType = 2
	AgoraData_P2P     AgoraData_TransactionType = 3
)

var AgoraData_TransactionType_name = map[int32]string{
	0: "UNKNOWN",
	1: "EARN",
	2: "SPEND",
	3: "P2P",
}

var AgoraData_TransactionType_value = map[string]int32{
	"UNKNOWN": 0,
	"EARN":    1,
	"SPEND":   2,
	"P2P":     3,
}

func (x AgoraData_TransactionType) String() string {
	return proto.EnumName(AgoraData_TransactionType_name, int32(x))
}

func (AgoraData_TransactionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{4, 0}
}

type StellarAccountId struct {
	// The public key of accounts always starts with a G, so we
	// ensure that the value starts with a G to prevent the secret
	// key from being used.
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StellarAccountId) Reset()         { *m = StellarAccountId{} }
func (m *StellarAccountId) String() string { return proto.CompactTextString(m) }
func (*StellarAccountId) ProtoMessage()    {}
func (*StellarAccountId) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{0}
}

func (m *StellarAccountId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StellarAccountId.Unmarshal(m, b)
}
func (m *StellarAccountId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StellarAccountId.Marshal(b, m, deterministic)
}
func (m *StellarAccountId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StellarAccountId.Merge(m, src)
}
func (m *StellarAccountId) XXX_Size() int {
	return xxx_messageInfo_StellarAccountId.Size(m)
}
func (m *StellarAccountId) XXX_DiscardUnknown() {
	xxx_messageInfo_StellarAccountId.DiscardUnknown(m)
}

var xxx_messageInfo_StellarAccountId proto.InternalMessageInfo

func (m *StellarAccountId) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type TransactionHash struct {
	// The sha256 hash of a transaction.
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionHash) Reset()         { *m = TransactionHash{} }
func (m *TransactionHash) String() string { return proto.CompactTextString(m) }
func (*TransactionHash) ProtoMessage()    {}
func (*TransactionHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{1}
}

func (m *TransactionHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionHash.Unmarshal(m, b)
}
func (m *TransactionHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionHash.Marshal(b, m, deterministic)
}
func (m *TransactionHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionHash.Merge(m, src)
}
func (m *TransactionHash) XXX_Size() int {
	return xxx_messageInfo_TransactionHash.Size(m)
}
func (m *TransactionHash) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionHash.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionHash proto.InternalMessageInfo

func (m *TransactionHash) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type BigDecimal struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BigDecimal) Reset()         { *m = BigDecimal{} }
func (m *BigDecimal) String() string { return proto.CompactTextString(m) }
func (*BigDecimal) ProtoMessage()    {}
func (*BigDecimal) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{2}
}

func (m *BigDecimal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BigDecimal.Unmarshal(m, b)
}
func (m *BigDecimal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BigDecimal.Marshal(b, m, deterministic)
}
func (m *BigDecimal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BigDecimal.Merge(m, src)
}
func (m *BigDecimal) XXX_Size() int {
	return xxx_messageInfo_BigDecimal.Size(m)
}
func (m *BigDecimal) XXX_DiscardUnknown() {
	xxx_messageInfo_BigDecimal.DiscardUnknown(m)
}

var xxx_messageInfo_BigDecimal proto.InternalMessageInfo

func (m *BigDecimal) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type AgoraDataUrl struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgoraDataUrl) Reset()         { *m = AgoraDataUrl{} }
func (m *AgoraDataUrl) String() string { return proto.CompactTextString(m) }
func (*AgoraDataUrl) ProtoMessage()    {}
func (*AgoraDataUrl) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{3}
}

func (m *AgoraDataUrl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgoraDataUrl.Unmarshal(m, b)
}
func (m *AgoraDataUrl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgoraDataUrl.Marshal(b, m, deterministic)
}
func (m *AgoraDataUrl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgoraDataUrl.Merge(m, src)
}
func (m *AgoraDataUrl) XXX_Size() int {
	return xxx_messageInfo_AgoraDataUrl.Size(m)
}
func (m *AgoraDataUrl) XXX_DiscardUnknown() {
	xxx_messageInfo_AgoraDataUrl.DiscardUnknown(m)
}

var xxx_messageInfo_AgoraDataUrl proto.InternalMessageInfo

func (m *AgoraDataUrl) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type AgoraData struct {
	// A renderable title related to the transaction.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// A renderable description related to the transaction.
	Description     string                    `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	TransactionType AgoraData_TransactionType `protobuf:"varint,3,opt,name=transaction_type,json=transactionType,proto3,enum=kin.common.v3.AgoraData_TransactionType" json:"transaction_type,omitempty"`
	// The total amount, in quarks, of the 'item(s)' the transaction
	// was related to. This should be the same as the transaction
	// amount. If there is a mismatch, then the client mis-used or
	// abused the memo system, likely by circumventing the Agora
	// services.
	TotalAmount int64 `protobuf:"varint,4,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
	// The full foreign key of the off-chain data, which is
	// prefixed in the memo. This can be treated similar SKU.
	ForeignKey []byte `protobuf:"bytes,5,opt,name=foreign_key,json=foreignKey,proto3" json:"foreign_key,omitempty"`
	// The invoice related to the transaction, if the service was
	// able to resolve one using the foreign key.
	Invoice              *Invoice `protobuf:"bytes,6,opt,name=invoice,proto3" json:"invoice,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgoraData) Reset()         { *m = AgoraData{} }
func (m *AgoraData) String() string { return proto.CompactTextString(m) }
func (*AgoraData) ProtoMessage()    {}
func (*AgoraData) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{4}
}

func (m *AgoraData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgoraData.Unmarshal(m, b)
}
func (m *AgoraData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgoraData.Marshal(b, m, deterministic)
}
func (m *AgoraData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgoraData.Merge(m, src)
}
func (m *AgoraData) XXX_Size() int {
	return xxx_messageInfo_AgoraData.Size(m)
}
func (m *AgoraData) XXX_DiscardUnknown() {
	xxx_messageInfo_AgoraData.DiscardUnknown(m)
}

var xxx_messageInfo_AgoraData proto.InternalMessageInfo

func (m *AgoraData) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AgoraData) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AgoraData) GetTransactionType() AgoraData_TransactionType {
	if m != nil {
		return m.TransactionType
	}
	return AgoraData_UNKNOWN
}

func (m *AgoraData) GetTotalAmount() int64 {
	if m != nil {
		return m.TotalAmount
	}
	return 0
}

func (m *AgoraData) GetForeignKey() []byte {
	if m != nil {
		return m.ForeignKey
	}
	return nil
}

func (m *AgoraData) GetInvoice() *Invoice {
	if m != nil {
		return m.Invoice
	}
	return nil
}

type InvoiceHash struct {
	// The SHA-224 hash of the invoice.
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InvoiceHash) Reset()         { *m = InvoiceHash{} }
func (m *InvoiceHash) String() string { return proto.CompactTextString(m) }
func (*InvoiceHash) ProtoMessage()    {}
func (*InvoiceHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{5}
}

func (m *InvoiceHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InvoiceHash.Unmarshal(m, b)
}
func (m *InvoiceHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InvoiceHash.Marshal(b, m, deterministic)
}
func (m *InvoiceHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InvoiceHash.Merge(m, src)
}
func (m *InvoiceHash) XXX_Size() int {
	return xxx_messageInfo_InvoiceHash.Size(m)
}
func (m *InvoiceHash) XXX_DiscardUnknown() {
	xxx_messageInfo_InvoiceHash.DiscardUnknown(m)
}

var xxx_messageInfo_InvoiceHash proto.InternalMessageInfo

func (m *InvoiceHash) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type Invoice struct {
	Items                []*Invoice_LineItem `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Nonce                *Nonce              `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Invoice) Reset()         { *m = Invoice{} }
func (m *Invoice) String() string { return proto.CompactTextString(m) }
func (*Invoice) ProtoMessage()    {}
func (*Invoice) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{6}
}

func (m *Invoice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invoice.Unmarshal(m, b)
}
func (m *Invoice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invoice.Marshal(b, m, deterministic)
}
func (m *Invoice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invoice.Merge(m, src)
}
func (m *Invoice) XXX_Size() int {
	return xxx_messageInfo_Invoice.Size(m)
}
func (m *Invoice) XXX_DiscardUnknown() {
	xxx_messageInfo_Invoice.DiscardUnknown(m)
}

var xxx_messageInfo_Invoice proto.InternalMessageInfo

func (m *Invoice) GetItems() []*Invoice_LineItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func (m *Invoice) GetNonce() *Nonce {
	if m != nil {
		return m.Nonce
	}
	return nil
}

type Invoice_LineItem struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The amount in quarks.
	Amount int64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	// The app SKU related to this line item, if applicable.
	Sku                  []byte   `protobuf:"bytes,4,opt,name=sku,proto3" json:"sku,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Invoice_LineItem) Reset()         { *m = Invoice_LineItem{} }
func (m *Invoice_LineItem) String() string { return proto.CompactTextString(m) }
func (*Invoice_LineItem) ProtoMessage()    {}
func (*Invoice_LineItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{6, 0}
}

func (m *Invoice_LineItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Invoice_LineItem.Unmarshal(m, b)
}
func (m *Invoice_LineItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Invoice_LineItem.Marshal(b, m, deterministic)
}
func (m *Invoice_LineItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Invoice_LineItem.Merge(m, src)
}
func (m *Invoice_LineItem) XXX_Size() int {
	return xxx_messageInfo_Invoice_LineItem.Size(m)
}
func (m *Invoice_LineItem) XXX_DiscardUnknown() {
	xxx_messageInfo_Invoice_LineItem.DiscardUnknown(m)
}

var xxx_messageInfo_Invoice_LineItem proto.InternalMessageInfo

func (m *Invoice_LineItem) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Invoice_LineItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Invoice_LineItem) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Invoice_LineItem) GetSku() []byte {
	if m != nil {
		return m.Sku
	}
	return nil
}

type Nonce struct {
	// The time the nonce was generated.
	//
	// If the generation time value is too far in the past _or_ future of the service time,
	// the nonce will be rejected.
	GenerationTime *timestamp.Timestamp `protobuf:"bytes,1,opt,name=generation_time,json=generationTime,proto3" json:"generation_time,omitempty"`
	// A random value.
	Value                int64    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nonce) Reset()         { *m = Nonce{} }
func (m *Nonce) String() string { return proto.CompactTextString(m) }
func (*Nonce) ProtoMessage()    {}
func (*Nonce) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab58ade4bb87a1ff, []int{7}
}

func (m *Nonce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nonce.Unmarshal(m, b)
}
func (m *Nonce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nonce.Marshal(b, m, deterministic)
}
func (m *Nonce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nonce.Merge(m, src)
}
func (m *Nonce) XXX_Size() int {
	return xxx_messageInfo_Nonce.Size(m)
}
func (m *Nonce) XXX_DiscardUnknown() {
	xxx_messageInfo_Nonce.DiscardUnknown(m)
}

var xxx_messageInfo_Nonce proto.InternalMessageInfo

func (m *Nonce) GetGenerationTime() *timestamp.Timestamp {
	if m != nil {
		return m.GenerationTime
	}
	return nil
}

func (m *Nonce) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterEnum("kin.common.v3.AgoraData_TransactionType", AgoraData_TransactionType_name, AgoraData_TransactionType_value)
	proto.RegisterType((*StellarAccountId)(nil), "kin.common.v3.StellarAccountId")
	proto.RegisterType((*TransactionHash)(nil), "kin.common.v3.TransactionHash")
	proto.RegisterType((*BigDecimal)(nil), "kin.common.v3.BigDecimal")
	proto.RegisterType((*AgoraDataUrl)(nil), "kin.common.v3.AgoraDataUrl")
	proto.RegisterType((*AgoraData)(nil), "kin.common.v3.AgoraData")
	proto.RegisterType((*InvoiceHash)(nil), "kin.common.v3.InvoiceHash")
	proto.RegisterType((*Invoice)(nil), "kin.common.v3.Invoice")
	proto.RegisterType((*Invoice_LineItem)(nil), "kin.common.v3.Invoice.LineItem")
	proto.RegisterType((*Nonce)(nil), "kin.common.v3.Nonce")
}

func init() { proto.RegisterFile("common/v3/model.proto", fileDescriptor_ab58ade4bb87a1ff) }

var fileDescriptor_ab58ade4bb87a1ff = []byte{
	// 725 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6b, 0xe3, 0x46,
	0x14, 0xdd, 0x91, 0x22, 0x5b, 0x1e, 0xa5, 0x6b, 0x31, 0x6c, 0x76, 0x85, 0x59, 0x88, 0x6a, 0xfa,
	0xa0, 0x94, 0xb5, 0x64, 0xcb, 0x4b, 0xd9, 0xdd, 0x16, 0x82, 0x85, 0x97, 0x36, 0x24, 0x75, 0x83,
	0x92, 0x50, 0x88, 0x9b, 0x84, 0x89, 0x3c, 0x51, 0x06, 0x4b, 0x1a, 0x23, 0x8d, 0x0d, 0x4e, 0x1b,
	0xc8, 0x73, 0x1f, 0xfa, 0x50, 0xfa, 0x4b, 0xfa, 0xc3, 0xfa, 0x03, 0xfc, 0x54, 0x34, 0xb2, 0x1c,
	0x3b, 0xfd, 0x60, 0xdf, 0x74, 0xef, 0x39, 0xf7, 0xea, 0xcc, 0x3d, 0x97, 0x0b, 0x77, 0x02, 0x16,
	0xc7, 0x2c, 0x71, 0x66, 0x5d, 0x27, 0x66, 0x23, 0x12, 0xd9, 0x93, 0x94, 0x71, 0x86, 0x3e, 0x1b,
	0xd3, 0xc4, 0x2e, 0x20, 0x7b, 0xd6, 0x6d, 0xbc, 0x9a, 0xe1, 0x88, 0x8e, 0x30, 0x27, 0x4e, 0xf9,
	0x51, 0xf0, 0x1a, 0xbb, 0x21, 0x63, 0x61, 0x44, 0x1c, 0x11, 0x5d, 0x4f, 0x6f, 0x1c, 0x4e, 0x63,
	0x92, 0x71, 0x1c, 0x4f, 0x0a, 0x42, 0xb3, 0x0f, 0xf5, 0x13, 0x4e, 0xa2, 0x08, 0xa7, 0xbd, 0x20,
	0x60, 0xd3, 0x84, 0x1f, 0x8c, 0x50, 0x1b, 0x2a, 0x33, 0x1c, 0x4d, 0x89, 0x01, 0x4c, 0x60, 0xd5,
	0xbc, 0xc6, 0xc2, 0x7b, 0x95, 0xee, 0x98, 0xef, 0xac, 0x77, 0x6e, 0xfd, 0xf2, 0xdb, 0x21, 0x6e,
	0xdd, 0xf5, 0x5a, 0xe7, 0xed, 0xd6, 0xfb, 0x8b, 0x2f, 0xbf, 0xf0, 0x0b, 0x62, 0xd3, 0x85, 0xf5,
	0xd3, 0x14, 0x27, 0x19, 0x0e, 0x38, 0x65, 0xc9, 0x77, 0x38, 0xbb, 0x45, 0xbb, 0xeb, 0x4d, 0xb6,
	0xbd, 0xda, 0xc2, 0xab, 0xdc, 0x6d, 0xe9, 0xa6, 0x61, 0x96, 0x35, 0xe7, 0x10, 0x7a, 0x34, 0xec,
	0x93, 0x80, 0xc6, 0x38, 0x42, 0x47, 0x9b, 0xff, 0xfc, 0x6a, 0xe1, 0x75, 0xd3, 0x8e, 0xeb, 0x5c,
	0x5a, 0xfb, 0x1f, 0x86, 0x9d, 0xd6, 0xfb, 0x8b, 0x61, 0xfe, 0xc3, 0x9f, 0xdb, 0x6f, 0x3a, 0x6f,
	0xef, 0x7f, 0x69, 0xef, 0x59, 0xfb, 0x1f, 0x7e, 0xb2, 0x57, 0xa9, 0xce, 0xbd, 0xa0, 0xec, 0xed,
	0xaf, 0xf4, 0xb4, 0xe1, 0x76, 0x2f, 0x64, 0x29, 0xee, 0x63, 0x8e, 0xcf, 0xd2, 0x08, 0x99, 0x9b,
	0xdd, 0xe1, 0xc2, 0xab, 0xa6, 0x8a, 0x0e, 0x8c, 0x07, 0xbd, 0xac, 0xf8, 0x4b, 0x82, 0xb5, 0x55,
	0x49, 0xce, 0xe7, 0x94, 0x47, 0xff, 0xe4, 0x4b, 0x7e, 0x01, 0xa0, 0x37, 0x50, 0x1b, 0x91, 0x2c,
	0x48, 0xe9, 0x24, 0x7f, 0xb1, 0x21, 0x3d, 0xe1, 0xa9, 0xfe, 0x3a, 0x8c, 0x4e, 0xa0, 0xce, 0x1f,
	0xe7, 0x73, 0xc5, 0xe7, 0x13, 0x62, 0xc8, 0x26, 0xb0, 0x9e, 0xbb, 0x96, 0xbd, 0xe1, 0xa4, 0xbd,
	0xd2, 0x60, 0xaf, 0x0d, 0xf4, 0x74, 0x3e, 0x21, 0x7e, 0x9d, 0x6f, 0x26, 0xd0, 0xe7, 0x70, 0x9b,
	0x33, 0x8e, 0xa3, 0x2b, 0x1c, 0xe7, 0xc6, 0x19, 0x5b, 0x26, 0xb0, 0x64, 0x5f, 0x13, 0xb9, 0x9e,
	0x48, 0xa1, 0x3d, 0xa8, 0xdd, 0xb0, 0x94, 0xd0, 0x30, 0xb9, 0x1a, 0x93, 0xb9, 0xa1, 0x08, 0x2b,
	0xd4, 0x85, 0xa7, 0xdc, 0xc9, 0xf9, 0x5b, 0xe0, 0x12, 0x3c, 0x24, 0x73, 0xd4, 0x86, 0x55, 0x9a,
	0xcc, 0x18, 0x0d, 0x88, 0x51, 0x31, 0x81, 0xa5, 0xb9, 0x2f, 0x9f, 0x28, 0x3b, 0x28, 0x50, 0xbf,
	0xa4, 0x35, 0xbf, 0xd9, 0x30, 0x5d, 0x48, 0xd2, 0x60, 0xf5, 0x6c, 0x70, 0x38, 0xf8, 0xe1, 0xc7,
	0x81, 0xfe, 0x0c, 0xa9, 0x70, 0xeb, 0x63, 0xcf, 0x1f, 0xe8, 0x00, 0xd5, 0xa0, 0x72, 0x72, 0xfc,
	0x71, 0xd0, 0xd7, 0x25, 0x54, 0x85, 0xf2, 0xb1, 0x7b, 0xac, 0xcb, 0x4d, 0x1b, 0x6a, 0xcb, 0x8e,
	0xff, 0xb3, 0x2e, 0xaf, 0x8d, 0xd7, 0xa5, 0x41, 0x7f, 0x48, 0xb0, 0xba, 0x2c, 0x40, 0x1e, 0x54,
	0x28, 0x27, 0x71, 0x66, 0x00, 0x53, 0xb6, 0x34, 0x77, 0xf7, 0xdf, 0x95, 0xda, 0x47, 0x34, 0x21,
	0x07, 0x9c, 0xc4, 0x9e, 0xb6, 0xf0, 0xd4, 0xdf, 0x81, 0xa2, 0x02, 0xfd, 0x41, 0xf5, 0x8b, 0x52,
	0xf4, 0x16, 0x2a, 0x09, 0x4b, 0x02, 0x22, 0xac, 0xd3, 0xdc, 0x17, 0x4f, 0x7a, 0x0c, 0x72, 0x4c,
	0x8c, 0xea, 0x57, 0x20, 0xe9, 0xc0, 0x2f, 0xc8, 0x8d, 0xdf, 0x00, 0x54, 0xcb, 0xb6, 0xff, 0xbd,
	0x25, 0xe0, 0x53, 0xb6, 0xe4, 0x59, 0xee, 0xc0, 0xc6, 0x96, 0xbc, 0x84, 0x95, 0xa5, 0x95, 0xb2,
	0xb0, 0x72, 0x19, 0xa1, 0x06, 0x94, 0xb3, 0xf1, 0x54, 0xf8, 0xfb, 0xe8, 0x1e, 0xf0, 0xf3, 0x64,
	0x33, 0x82, 0x8a, 0x90, 0x8a, 0xbe, 0x87, 0xf5, 0x90, 0x24, 0x24, 0xc5, 0xc5, 0x86, 0xd1, 0xb8,
	0x90, 0xa5, 0xb9, 0x0d, 0xbb, 0xb8, 0x01, 0x76, 0x79, 0x03, 0xec, 0xd3, 0xf2, 0x06, 0x88, 0x66,
	0x7f, 0x02, 0x49, 0x05, 0xfe, 0xf3, 0xc7, 0xe2, 0x1c, 0x46, 0x2f, 0x4a, 0x3f, 0x24, 0x21, 0xa5,
	0x08, 0xbc, 0x21, 0xdc, 0x61, 0x69, 0x28, 0x46, 0x15, 0x92, 0xb5, 0x71, 0x9d, 0x7b, 0x21, 0xe5,
	0xb7, 0xd3, 0xeb, 0x3c, 0xe3, 0x8c, 0x69, 0x42, 0x02, 0x96, 0xcd, 0x33, 0x4e, 0x44, 0xd0, 0xc2,
	0x13, 0xda, 0xa2, 0x09, 0x27, 0x69, 0x82, 0x23, 0x27, 0x24, 0x89, 0x10, 0xe2, 0xac, 0x0e, 0xdb,
	0xd7, 0xc5, 0xd7, 0x75, 0x45, 0xe4, 0xbb, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xfc, 0x4a, 0xbb,
	0xb1, 0xf3, 0x04, 0x00, 0x00,
}
