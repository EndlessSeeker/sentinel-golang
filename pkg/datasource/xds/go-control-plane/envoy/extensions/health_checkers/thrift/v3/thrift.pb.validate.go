// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/health_checkers/thrift/v3/thrift.proto

package thriftv3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"

	v3 "github.com/alibaba/sentinel-golang/pkg/datasource/xds/go-control-plane/envoy/extensions/filters/network/thrift_proxy/v3"
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
	_ = sort.Sort

	_ = v3.TransportType(0)
)

// Validate checks the field values on Thrift with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Thrift) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Thrift with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ThriftMultiError, or nil if none found.
func (m *Thrift) ValidateAll() error {
	return m.validate(true)
}

func (m *Thrift) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetMethodName()) < 1 {
		err := ThriftValidationError{
			field:  "MethodName",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := v3.TransportType_name[int32(m.GetTransport())]; !ok {
		err := ThriftValidationError{
			field:  "Transport",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := v3.ProtocolType_name[int32(m.GetProtocol())]; !ok {
		err := ThriftValidationError{
			field:  "Protocol",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ThriftMultiError(errors)
	}

	return nil
}

// ThriftMultiError is an error wrapping multiple validation errors returned by
// Thrift.ValidateAll() if the designated constraints aren't met.
type ThriftMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ThriftMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ThriftMultiError) AllErrors() []error { return m }

// ThriftValidationError is the validation error returned by Thrift.Validate if
// the designated constraints aren't met.
type ThriftValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ThriftValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ThriftValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ThriftValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ThriftValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ThriftValidationError) ErrorName() string { return "ThriftValidationError" }

// Error satisfies the builtin error interface
func (e ThriftValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sThrift.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ThriftValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ThriftValidationError{}
