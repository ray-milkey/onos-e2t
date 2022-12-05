// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/f1ap/v1/f1ap_containers.proto

package f1apcontainersv1

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

// Validate checks the field values on ProtocolIeContainer001 with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProtocolIeContainer001) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProtocolIeContainer001 with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProtocolIeContainer001MultiError, or nil if none found.
func (m *ProtocolIeContainer001) ValidateAll() error {
	return m.validate(true)
}

func (m *ProtocolIeContainer001) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ProtocolIeContainer001MultiError(errors)
	}

	return nil
}

// ProtocolIeContainer001MultiError is an error wrapping multiple validation
// errors returned by ProtocolIeContainer001.ValidateAll() if the designated
// constraints aren't met.
type ProtocolIeContainer001MultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProtocolIeContainer001MultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProtocolIeContainer001MultiError) AllErrors() []error { return m }

// ProtocolIeContainer001ValidationError is the validation error returned by
// ProtocolIeContainer001.Validate if the designated constraints aren't met.
type ProtocolIeContainer001ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolIeContainer001ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolIeContainer001ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolIeContainer001ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolIeContainer001ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolIeContainer001ValidationError) ErrorName() string {
	return "ProtocolIeContainer001ValidationError"
}

// Error satisfies the builtin error interface
func (e ProtocolIeContainer001ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolIeContainer001.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolIeContainer001ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolIeContainer001ValidationError{}

// Validate checks the field values on ProtocolIeSingleContainer001 with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProtocolIeSingleContainer001) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProtocolIeSingleContainer001 with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProtocolIeSingleContainer001MultiError, or nil if none found.
func (m *ProtocolIeSingleContainer001) ValidateAll() error {
	return m.validate(true)
}

func (m *ProtocolIeSingleContainer001) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetValue()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProtocolIeSingleContainer001ValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProtocolIeSingleContainer001ValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProtocolIeSingleContainer001ValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ProtocolIeSingleContainer001MultiError(errors)
	}

	return nil
}

// ProtocolIeSingleContainer001MultiError is an error wrapping multiple
// validation errors returned by ProtocolIeSingleContainer001.ValidateAll() if
// the designated constraints aren't met.
type ProtocolIeSingleContainer001MultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProtocolIeSingleContainer001MultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProtocolIeSingleContainer001MultiError) AllErrors() []error { return m }

// ProtocolIeSingleContainer001ValidationError is the validation error returned
// by ProtocolIeSingleContainer001.Validate if the designated constraints
// aren't met.
type ProtocolIeSingleContainer001ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolIeSingleContainer001ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolIeSingleContainer001ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolIeSingleContainer001ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolIeSingleContainer001ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolIeSingleContainer001ValidationError) ErrorName() string {
	return "ProtocolIeSingleContainer001ValidationError"
}

// Error satisfies the builtin error interface
func (e ProtocolIeSingleContainer001ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolIeSingleContainer001.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolIeSingleContainer001ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolIeSingleContainer001ValidationError{}

// Validate checks the field values on ProtocolIeField001 with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProtocolIeField001) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProtocolIeField001 with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProtocolIeField001MultiError, or nil if none found.
func (m *ProtocolIeField001) ValidateAll() error {
	return m.validate(true)
}

func (m *ProtocolIeField001) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Criticality

	// no validation rules for Value

	if len(errors) > 0 {
		return ProtocolIeField001MultiError(errors)
	}

	return nil
}

// ProtocolIeField001MultiError is an error wrapping multiple validation errors
// returned by ProtocolIeField001.ValidateAll() if the designated constraints
// aren't met.
type ProtocolIeField001MultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProtocolIeField001MultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProtocolIeField001MultiError) AllErrors() []error { return m }

// ProtocolIeField001ValidationError is the validation error returned by
// ProtocolIeField001.Validate if the designated constraints aren't met.
type ProtocolIeField001ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolIeField001ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolIeField001ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolIeField001ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolIeField001ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolIeField001ValidationError) ErrorName() string {
	return "ProtocolIeField001ValidationError"
}

// Error satisfies the builtin error interface
func (e ProtocolIeField001ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolIeField001.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolIeField001ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolIeField001ValidationError{}

// Validate checks the field values on ProtocolIeContainerPair with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProtocolIeContainerPair) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProtocolIeContainerPair with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProtocolIeContainerPairMultiError, or nil if none found.
func (m *ProtocolIeContainerPair) ValidateAll() error {
	return m.validate(true)
}

func (m *ProtocolIeContainerPair) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ProtocolIeContainerPairMultiError(errors)
	}

	return nil
}

// ProtocolIeContainerPairMultiError is an error wrapping multiple validation
// errors returned by ProtocolIeContainerPair.ValidateAll() if the designated
// constraints aren't met.
type ProtocolIeContainerPairMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProtocolIeContainerPairMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProtocolIeContainerPairMultiError) AllErrors() []error { return m }

// ProtocolIeContainerPairValidationError is the validation error returned by
// ProtocolIeContainerPair.Validate if the designated constraints aren't met.
type ProtocolIeContainerPairValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolIeContainerPairValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolIeContainerPairValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolIeContainerPairValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolIeContainerPairValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolIeContainerPairValidationError) ErrorName() string {
	return "ProtocolIeContainerPairValidationError"
}

// Error satisfies the builtin error interface
func (e ProtocolIeContainerPairValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolIeContainerPair.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolIeContainerPairValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolIeContainerPairValidationError{}

// Validate checks the field values on ProtocolIeFieldPair with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProtocolIeFieldPair) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProtocolIeFieldPair with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProtocolIeFieldPairMultiError, or nil if none found.
func (m *ProtocolIeFieldPair) ValidateAll() error {
	return m.validate(true)
}

func (m *ProtocolIeFieldPair) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return ProtocolIeFieldPairMultiError(errors)
	}

	return nil
}

// ProtocolIeFieldPairMultiError is an error wrapping multiple validation
// errors returned by ProtocolIeFieldPair.ValidateAll() if the designated
// constraints aren't met.
type ProtocolIeFieldPairMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProtocolIeFieldPairMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProtocolIeFieldPairMultiError) AllErrors() []error { return m }

// ProtocolIeFieldPairValidationError is the validation error returned by
// ProtocolIeFieldPair.Validate if the designated constraints aren't met.
type ProtocolIeFieldPairValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolIeFieldPairValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolIeFieldPairValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolIeFieldPairValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolIeFieldPairValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolIeFieldPairValidationError) ErrorName() string {
	return "ProtocolIeFieldPairValidationError"
}

// Error satisfies the builtin error interface
func (e ProtocolIeFieldPairValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolIeFieldPair.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolIeFieldPairValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolIeFieldPairValidationError{}

// Validate checks the field values on ProtocolExtensionContainer001 with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProtocolExtensionContainer001) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProtocolExtensionContainer001 with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ProtocolExtensionContainer001MultiError, or nil if none found.
func (m *ProtocolExtensionContainer001) ValidateAll() error {
	return m.validate(true)
}

func (m *ProtocolExtensionContainer001) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ProtocolExtensionContainer001MultiError(errors)
	}

	return nil
}

// ProtocolExtensionContainer001MultiError is an error wrapping multiple
// validation errors returned by ProtocolExtensionContainer001.ValidateAll()
// if the designated constraints aren't met.
type ProtocolExtensionContainer001MultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProtocolExtensionContainer001MultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProtocolExtensionContainer001MultiError) AllErrors() []error { return m }

// ProtocolExtensionContainer001ValidationError is the validation error
// returned by ProtocolExtensionContainer001.Validate if the designated
// constraints aren't met.
type ProtocolExtensionContainer001ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolExtensionContainer001ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolExtensionContainer001ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolExtensionContainer001ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolExtensionContainer001ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolExtensionContainer001ValidationError) ErrorName() string {
	return "ProtocolExtensionContainer001ValidationError"
}

// Error satisfies the builtin error interface
func (e ProtocolExtensionContainer001ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolExtensionContainer001.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolExtensionContainer001ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolExtensionContainer001ValidationError{}

// Validate checks the field values on ProtocolExtensionField001 with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProtocolExtensionField001) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProtocolExtensionField001 with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProtocolExtensionField001MultiError, or nil if none found.
func (m *ProtocolExtensionField001) ValidateAll() error {
	return m.validate(true)
}

func (m *ProtocolExtensionField001) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Criticality

	// no validation rules for ExtensionValue

	if len(errors) > 0 {
		return ProtocolExtensionField001MultiError(errors)
	}

	return nil
}

// ProtocolExtensionField001MultiError is an error wrapping multiple validation
// errors returned by ProtocolExtensionField001.ValidateAll() if the
// designated constraints aren't met.
type ProtocolExtensionField001MultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProtocolExtensionField001MultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProtocolExtensionField001MultiError) AllErrors() []error { return m }

// ProtocolExtensionField001ValidationError is the validation error returned by
// ProtocolExtensionField001.Validate if the designated constraints aren't met.
type ProtocolExtensionField001ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolExtensionField001ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolExtensionField001ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolExtensionField001ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolExtensionField001ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolExtensionField001ValidationError) ErrorName() string {
	return "ProtocolExtensionField001ValidationError"
}

// Error satisfies the builtin error interface
func (e ProtocolExtensionField001ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolExtensionField001.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolExtensionField001ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolExtensionField001ValidationError{}

// Validate checks the field values on PrivateIeContainer001 with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PrivateIeContainer001) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PrivateIeContainer001 with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PrivateIeContainer001MultiError, or nil if none found.
func (m *PrivateIeContainer001) ValidateAll() error {
	return m.validate(true)
}

func (m *PrivateIeContainer001) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return PrivateIeContainer001MultiError(errors)
	}

	return nil
}

// PrivateIeContainer001MultiError is an error wrapping multiple validation
// errors returned by PrivateIeContainer001.ValidateAll() if the designated
// constraints aren't met.
type PrivateIeContainer001MultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PrivateIeContainer001MultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PrivateIeContainer001MultiError) AllErrors() []error { return m }

// PrivateIeContainer001ValidationError is the validation error returned by
// PrivateIeContainer001.Validate if the designated constraints aren't met.
type PrivateIeContainer001ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PrivateIeContainer001ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PrivateIeContainer001ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PrivateIeContainer001ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PrivateIeContainer001ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PrivateIeContainer001ValidationError) ErrorName() string {
	return "PrivateIeContainer001ValidationError"
}

// Error satisfies the builtin error interface
func (e PrivateIeContainer001ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrivateIeContainer001.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PrivateIeContainer001ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PrivateIeContainer001ValidationError{}

// Validate checks the field values on PrivateIeField001 with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PrivateIeField001) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PrivateIeField001 with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PrivateIeField001MultiError, or nil if none found.
func (m *PrivateIeField001) ValidateAll() error {
	return m.validate(true)
}

func (m *PrivateIeField001) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Criticality

	// no validation rules for Value

	if len(errors) > 0 {
		return PrivateIeField001MultiError(errors)
	}

	return nil
}

// PrivateIeField001MultiError is an error wrapping multiple validation errors
// returned by PrivateIeField001.ValidateAll() if the designated constraints
// aren't met.
type PrivateIeField001MultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PrivateIeField001MultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PrivateIeField001MultiError) AllErrors() []error { return m }

// PrivateIeField001ValidationError is the validation error returned by
// PrivateIeField001.Validate if the designated constraints aren't met.
type PrivateIeField001ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PrivateIeField001ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PrivateIeField001ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PrivateIeField001ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PrivateIeField001ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PrivateIeField001ValidationError) ErrorName() string {
	return "PrivateIeField001ValidationError"
}

// Error satisfies the builtin error interface
func (e PrivateIeField001ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrivateIeField001.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PrivateIeField001ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PrivateIeField001ValidationError{}