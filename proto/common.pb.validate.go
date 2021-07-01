// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common.proto

package proto

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

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
)

// Validate checks the field values on CommonResp with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *CommonResp) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Result

	// no validation rules for Message

	// no validation rules for Reason

	// no validation rules for Rid

	return nil
}

// CommonRespValidationError is the validation error returned by
// CommonResp.Validate if the designated constraints aren't met.
type CommonRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonRespValidationError) ErrorName() string { return "CommonRespValidationError" }

// Error satisfies the builtin error interface
func (e CommonRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonRespValidationError{}

// Validate checks the field values on InParam with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *InParam) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AppId

	// no validation rules for AreaId

	// no validation rules for Version

	// no validation rules for UserId

	// no validation rules for DeviceId

	// no validation rules for ClientIp

	// no validation rules for Channel

	// no validation rules for Source

	// no validation rules for UserIp

	// no validation rules for Language

	// no validation rules for Country

	// no validation rules for AuthorId

	return nil
}

// InParamValidationError is the validation error returned by InParam.Validate
// if the designated constraints aren't met.
type InParamValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InParamValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InParamValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InParamValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InParamValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InParamValidationError) ErrorName() string { return "InParamValidationError" }

// Error satisfies the builtin error interface
func (e InParamValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInParam.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InParamValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InParamValidationError{}

// Validate checks the field values on OutParam with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OutParam) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for SubCode

	// no validation rules for Message

	return nil
}

// OutParamValidationError is the validation error returned by
// OutParam.Validate if the designated constraints aren't met.
type OutParamValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OutParamValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OutParamValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OutParamValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OutParamValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OutParamValidationError) ErrorName() string { return "OutParamValidationError" }

// Error satisfies the builtin error interface
func (e OutParamValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOutParam.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OutParamValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OutParamValidationError{}
