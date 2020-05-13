// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/core/v3/grpc_method_list.proto

package envoy_config_core_v3

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

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _grpc_method_list_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on GrpcMethodList with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GrpcMethodList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetServices() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return GrpcMethodListValidationError{
						field:  fmt.Sprintf("Services[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// GrpcMethodListValidationError is the validation error returned by
// GrpcMethodList.Validate if the designated constraints aren't met.
type GrpcMethodListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrpcMethodListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrpcMethodListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrpcMethodListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrpcMethodListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrpcMethodListValidationError) ErrorName() string { return "GrpcMethodListValidationError" }

// Error satisfies the builtin error interface
func (e GrpcMethodListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcMethodList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrpcMethodListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrpcMethodListValidationError{}

// Validate checks the field values on GrpcMethodList_Service with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GrpcMethodList_Service) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return GrpcMethodList_ServiceValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	if len(m.GetMethodNames()) < 1 {
		return GrpcMethodList_ServiceValidationError{
			field:  "MethodNames",
			reason: "value must contain at least 1 item(s)",
		}
	}

	return nil
}

// GrpcMethodList_ServiceValidationError is the validation error returned by
// GrpcMethodList_Service.Validate if the designated constraints aren't met.
type GrpcMethodList_ServiceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GrpcMethodList_ServiceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GrpcMethodList_ServiceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GrpcMethodList_ServiceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GrpcMethodList_ServiceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GrpcMethodList_ServiceValidationError) ErrorName() string {
	return "GrpcMethodList_ServiceValidationError"
}

// Error satisfies the builtin error interface
func (e GrpcMethodList_ServiceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGrpcMethodList_Service.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GrpcMethodList_ServiceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GrpcMethodList_ServiceValidationError{}
