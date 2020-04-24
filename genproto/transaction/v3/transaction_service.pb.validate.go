// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: transaction/v3/transaction_service.proto

package transaction

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _transaction_service_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GetHistoryRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetHistoryRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetAccountId() == nil {
		return GetHistoryRequestValidationError{
			field:  "AccountId",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAccountId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetHistoryRequestValidationError{
				field:  "AccountId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetCursor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetHistoryRequestValidationError{
				field:  "Cursor",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Direction

	return nil
}

// GetHistoryRequestValidationError is the validation error returned by
// GetHistoryRequest.Validate if the designated constraints aren't met.
type GetHistoryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetHistoryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetHistoryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetHistoryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetHistoryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetHistoryRequestValidationError) ErrorName() string {
	return "GetHistoryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetHistoryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetHistoryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetHistoryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetHistoryRequestValidationError{}

// Validate checks the field values on GetHistoryResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetHistoryResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	if len(m.GetItems()) > 128 {
		return GetHistoryResponseValidationError{
			field:  "Items",
			reason: "value must contain no more than 128 item(s)",
		}
	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetHistoryResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// GetHistoryResponseValidationError is the validation error returned by
// GetHistoryResponse.Validate if the designated constraints aren't met.
type GetHistoryResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetHistoryResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetHistoryResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetHistoryResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetHistoryResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetHistoryResponseValidationError) ErrorName() string {
	return "GetHistoryResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetHistoryResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetHistoryResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetHistoryResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetHistoryResponseValidationError{}

// Validate checks the field values on SubmitTransactionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SubmitTransactionRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := len(m.GetTransactionXdr()); l < 1 || l > 10240 {
		return SubmitTransactionRequestValidationError{
			field:  "TransactionXdr",
			reason: "value length must be between 1 and 10240 bytes, inclusive",
		}
	}

	if v, ok := interface{}(m.GetInvoiceList()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubmitTransactionRequestValidationError{
				field:  "InvoiceList",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SubmitTransactionRequestValidationError is the validation error returned by
// SubmitTransactionRequest.Validate if the designated constraints aren't met.
type SubmitTransactionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubmitTransactionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubmitTransactionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubmitTransactionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubmitTransactionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubmitTransactionRequestValidationError) ErrorName() string {
	return "SubmitTransactionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SubmitTransactionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubmitTransactionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubmitTransactionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubmitTransactionRequestValidationError{}

// Validate checks the field values on SubmitTransactionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SubmitTransactionResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	for idx, item := range m.GetInvoiceErrors() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SubmitTransactionResponseValidationError{
					field:  fmt.Sprintf("InvoiceErrors[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetHash()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubmitTransactionResponseValidationError{
				field:  "Hash",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Ledger

	if l := len(m.GetResultXdr()); l < 0 || l > 10240 {
		return SubmitTransactionResponseValidationError{
			field:  "ResultXdr",
			reason: "value length must be between 0 and 10240 bytes, inclusive",
		}
	}

	return nil
}

// SubmitTransactionResponseValidationError is the validation error returned by
// SubmitTransactionResponse.Validate if the designated constraints aren't met.
type SubmitTransactionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubmitTransactionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubmitTransactionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubmitTransactionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubmitTransactionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubmitTransactionResponseValidationError) ErrorName() string {
	return "SubmitTransactionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SubmitTransactionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubmitTransactionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubmitTransactionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubmitTransactionResponseValidationError{}

// Validate checks the field values on GetTransactionRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetTransactionRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetTransactionHash() == nil {
		return GetTransactionRequestValidationError{
			field:  "TransactionHash",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetTransactionHash()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTransactionRequestValidationError{
				field:  "TransactionHash",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetTransactionRequestValidationError is the validation error returned by
// GetTransactionRequest.Validate if the designated constraints aren't met.
type GetTransactionRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTransactionRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTransactionRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTransactionRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTransactionRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTransactionRequestValidationError) ErrorName() string {
	return "GetTransactionRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetTransactionRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTransactionRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTransactionRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTransactionRequestValidationError{}

// Validate checks the field values on GetTransactionResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetTransactionResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for State

	// no validation rules for Ledger

	if v, ok := interface{}(m.GetItem()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetTransactionResponseValidationError{
				field:  "Item",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetTransactionResponseValidationError is the validation error returned by
// GetTransactionResponse.Validate if the designated constraints aren't met.
type GetTransactionResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTransactionResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTransactionResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTransactionResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTransactionResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTransactionResponseValidationError) ErrorName() string {
	return "GetTransactionResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetTransactionResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTransactionResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTransactionResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTransactionResponseValidationError{}

// Validate checks the field values on HistoryItem with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HistoryItem) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetHash() == nil {
		return HistoryItemValidationError{
			field:  "Hash",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetHash()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HistoryItemValidationError{
				field:  "Hash",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if l := len(m.GetResultXdr()); l < 1 || l > 10240 {
		return HistoryItemValidationError{
			field:  "ResultXdr",
			reason: "value length must be between 1 and 10240 bytes, inclusive",
		}
	}

	if l := len(m.GetEnvelopeXdr()); l < 1 || l > 10240 {
		return HistoryItemValidationError{
			field:  "EnvelopeXdr",
			reason: "value length must be between 1 and 10240 bytes, inclusive",
		}
	}

	if v, ok := interface{}(m.GetCursor()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HistoryItemValidationError{
				field:  "Cursor",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetAgoraDataUrl()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HistoryItemValidationError{
				field:  "AgoraDataUrl",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetOpAgoraData()) > 100 {
		return HistoryItemValidationError{
			field:  "OpAgoraData",
			reason: "value must contain no more than 100 item(s)",
		}
	}

	for idx, item := range m.GetOpAgoraData() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HistoryItemValidationError{
					field:  fmt.Sprintf("OpAgoraData[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// HistoryItemValidationError is the validation error returned by
// HistoryItem.Validate if the designated constraints aren't met.
type HistoryItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HistoryItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HistoryItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HistoryItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HistoryItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HistoryItemValidationError) ErrorName() string { return "HistoryItemValidationError" }

// Error satisfies the builtin error interface
func (e HistoryItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHistoryItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HistoryItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HistoryItemValidationError{}

// Validate checks the field values on Cursor with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Cursor) Validate() error {
	if m == nil {
		return nil
	}

	if l := len(m.GetValue()); l < 1 || l > 128 {
		return CursorValidationError{
			field:  "Value",
			reason: "value length must be between 1 and 128 bytes, inclusive",
		}
	}

	return nil
}

// CursorValidationError is the validation error returned by Cursor.Validate if
// the designated constraints aren't met.
type CursorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CursorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CursorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CursorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CursorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CursorValidationError) ErrorName() string { return "CursorValidationError" }

// Error satisfies the builtin error interface
func (e CursorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCursor.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CursorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CursorValidationError{}

// Validate checks the field values on SubmitTransactionResponse_InvoiceError
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *SubmitTransactionResponse_InvoiceError) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetOpIndex() > 100 {
		return SubmitTransactionResponse_InvoiceErrorValidationError{
			field:  "OpIndex",
			reason: "value must be less than or equal to 100",
		}
	}

	if m.GetInvoice() == nil {
		return SubmitTransactionResponse_InvoiceErrorValidationError{
			field:  "Invoice",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetInvoice()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubmitTransactionResponse_InvoiceErrorValidationError{
				field:  "Invoice",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Reason

	return nil
}

// SubmitTransactionResponse_InvoiceErrorValidationError is the validation
// error returned by SubmitTransactionResponse_InvoiceError.Validate if the
// designated constraints aren't met.
type SubmitTransactionResponse_InvoiceErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubmitTransactionResponse_InvoiceErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubmitTransactionResponse_InvoiceErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubmitTransactionResponse_InvoiceErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubmitTransactionResponse_InvoiceErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubmitTransactionResponse_InvoiceErrorValidationError) ErrorName() string {
	return "SubmitTransactionResponse_InvoiceErrorValidationError"
}

// Error satisfies the builtin error interface
func (e SubmitTransactionResponse_InvoiceErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubmitTransactionResponse_InvoiceError.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubmitTransactionResponse_InvoiceErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubmitTransactionResponse_InvoiceErrorValidationError{}
