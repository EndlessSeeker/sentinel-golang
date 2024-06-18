// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/geoip_providers/common/v3/common.proto

package commonv3

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
)

// Validate checks the field values on CommonGeoipProviderConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CommonGeoipProviderConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CommonGeoipProviderConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CommonGeoipProviderConfigMultiError, or nil if none found.
func (m *CommonGeoipProviderConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *CommonGeoipProviderConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetGeoHeadersToAdd() == nil {
		err := CommonGeoipProviderConfigValidationError{
			field:  "GeoHeadersToAdd",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetGeoHeadersToAdd()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CommonGeoipProviderConfigValidationError{
					field:  "GeoHeadersToAdd",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CommonGeoipProviderConfigValidationError{
					field:  "GeoHeadersToAdd",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetGeoHeadersToAdd()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CommonGeoipProviderConfigValidationError{
				field:  "GeoHeadersToAdd",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CommonGeoipProviderConfigMultiError(errors)
	}

	return nil
}

// CommonGeoipProviderConfigMultiError is an error wrapping multiple validation
// errors returned by CommonGeoipProviderConfig.ValidateAll() if the
// designated constraints aren't met.
type CommonGeoipProviderConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommonGeoipProviderConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommonGeoipProviderConfigMultiError) AllErrors() []error { return m }

// CommonGeoipProviderConfigValidationError is the validation error returned by
// CommonGeoipProviderConfig.Validate if the designated constraints aren't met.
type CommonGeoipProviderConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonGeoipProviderConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CommonGeoipProviderConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CommonGeoipProviderConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CommonGeoipProviderConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonGeoipProviderConfigValidationError) ErrorName() string {
	return "CommonGeoipProviderConfigValidationError"
}

// Error satisfies the builtin error interface
func (e CommonGeoipProviderConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonGeoipProviderConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonGeoipProviderConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonGeoipProviderConfigValidationError{}

// Validate checks the field values on
// CommonGeoipProviderConfig_GeolocationHeadersToAdd with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CommonGeoipProviderConfig_GeolocationHeadersToAdd) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// CommonGeoipProviderConfig_GeolocationHeadersToAdd with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in
// CommonGeoipProviderConfig_GeolocationHeadersToAddMultiError, or nil if none found.
func (m *CommonGeoipProviderConfig_GeolocationHeadersToAdd) ValidateAll() error {
	return m.validate(true)
}

func (m *CommonGeoipProviderConfig_GeolocationHeadersToAdd) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetCountry() != "" {

		if !_CommonGeoipProviderConfig_GeolocationHeadersToAdd_Country_Pattern.MatchString(m.GetCountry()) {
			err := CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError{
				field:  "Country",
				reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetCity() != "" {

		if !_CommonGeoipProviderConfig_GeolocationHeadersToAdd_City_Pattern.MatchString(m.GetCity()) {
			err := CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError{
				field:  "City",
				reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetRegion() != "" {

		if !_CommonGeoipProviderConfig_GeolocationHeadersToAdd_Region_Pattern.MatchString(m.GetRegion()) {
			err := CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError{
				field:  "Region",
				reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetAsn() != "" {

		if !_CommonGeoipProviderConfig_GeolocationHeadersToAdd_Asn_Pattern.MatchString(m.GetAsn()) {
			err := CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError{
				field:  "Asn",
				reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetIsAnon() != "" {

		if !_CommonGeoipProviderConfig_GeolocationHeadersToAdd_IsAnon_Pattern.MatchString(m.GetIsAnon()) {
			err := CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError{
				field:  "IsAnon",
				reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetAnonVpn() != "" {

		if !_CommonGeoipProviderConfig_GeolocationHeadersToAdd_AnonVpn_Pattern.MatchString(m.GetAnonVpn()) {
			err := CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError{
				field:  "AnonVpn",
				reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetAnonHosting() != "" {

		if !_CommonGeoipProviderConfig_GeolocationHeadersToAdd_AnonHosting_Pattern.MatchString(m.GetAnonHosting()) {
			err := CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError{
				field:  "AnonHosting",
				reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetAnonTor() != "" {

		if !_CommonGeoipProviderConfig_GeolocationHeadersToAdd_AnonTor_Pattern.MatchString(m.GetAnonTor()) {
			err := CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError{
				field:  "AnonTor",
				reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetAnonProxy() != "" {

		if !_CommonGeoipProviderConfig_GeolocationHeadersToAdd_AnonProxy_Pattern.MatchString(m.GetAnonProxy()) {
			err := CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError{
				field:  "AnonProxy",
				reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return CommonGeoipProviderConfig_GeolocationHeadersToAddMultiError(errors)
	}

	return nil
}

// CommonGeoipProviderConfig_GeolocationHeadersToAddMultiError is an error
// wrapping multiple validation errors returned by
// CommonGeoipProviderConfig_GeolocationHeadersToAdd.ValidateAll() if the
// designated constraints aren't met.
type CommonGeoipProviderConfig_GeolocationHeadersToAddMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CommonGeoipProviderConfig_GeolocationHeadersToAddMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CommonGeoipProviderConfig_GeolocationHeadersToAddMultiError) AllErrors() []error { return m }

// CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError is the
// validation error returned by
// CommonGeoipProviderConfig_GeolocationHeadersToAdd.Validate if the
// designated constraints aren't met.
type CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError) ErrorName() string {
	return "CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError"
}

// Error satisfies the builtin error interface
func (e CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCommonGeoipProviderConfig_GeolocationHeadersToAdd.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CommonGeoipProviderConfig_GeolocationHeadersToAddValidationError{}

var _CommonGeoipProviderConfig_GeolocationHeadersToAdd_Country_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")

var _CommonGeoipProviderConfig_GeolocationHeadersToAdd_City_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")

var _CommonGeoipProviderConfig_GeolocationHeadersToAdd_Region_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")

var _CommonGeoipProviderConfig_GeolocationHeadersToAdd_Asn_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")

var _CommonGeoipProviderConfig_GeolocationHeadersToAdd_IsAnon_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")

var _CommonGeoipProviderConfig_GeolocationHeadersToAdd_AnonVpn_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")

var _CommonGeoipProviderConfig_GeolocationHeadersToAdd_AnonHosting_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")

var _CommonGeoipProviderConfig_GeolocationHeadersToAdd_AnonTor_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")

var _CommonGeoipProviderConfig_GeolocationHeadersToAdd_AnonProxy_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")
