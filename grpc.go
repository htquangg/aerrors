package aerrors

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCCoder interface {
	GRPCCode() codes.Code
}

// nolint:gocyclo
func (err Code) GRPCCode() codes.Code {
	switch err {
	// GRPC Errors
	case ErrOK:
		return codes.OK
	case ErrCanceled:
		return codes.Canceled
	case ErrUnknown:
		return codes.Unknown
	case ErrInvalidArgument:
		return codes.InvalidArgument
	case ErrDeadlineExceeded:
		return codes.DeadlineExceeded
	case ErrNotFound:
		return codes.NotFound
	case ErrAlreadyExists:
		return codes.AlreadyExists
	case ErrPermissionDenied:
		return codes.PermissionDenied
	case ErrResourceExhausted:
		return codes.ResourceExhausted
	case ErrFailedPrecondition:
		return codes.FailedPrecondition
	case ErrAborted:
		return codes.Aborted
	case ErrOutOfRange:
		return codes.OutOfRange
	case ErrUnimplemented:
		return codes.Unimplemented
	case ErrInternal:
		return codes.Internal
	case ErrUnavailable:
		return codes.Unavailable
	case ErrDataLoss:
		return codes.DataLoss
	case ErrUnauthenticated:
		return codes.Unauthenticated

	// HTTP Errors
	case ErrBadRequest:
		return codes.InvalidArgument
	case ErrUnauthorized:
		return codes.Unauthenticated
	case ErrForbidden:
		return codes.PermissionDenied
	case ErrMethodNotAllowed:
		return codes.Unimplemented
	case ErrRequestTimeout:
		return codes.DeadlineExceeded
	case ErrConflict:
		return codes.AlreadyExists
	case ErrImATeapot:
		return codes.Unknown
	case ErrUnprocessableEntity:
		return codes.InvalidArgument
	case ErrTooManyRequests:
		return codes.ResourceExhausted
	case ErrUnavailableForLegalReasons:
		return codes.Unavailable
	case ErrInternalServerError:
		return codes.Internal
	case ErrNotImplemented:
		return codes.Unimplemented
	case ErrBadGateway:
		return codes.Aborted
	case ErrServiceUnavailable:
		return codes.Unavailable
	case ErrGatewayTimeout:
		return codes.DeadlineExceeded
	default:
		return codes.Internal
	}
}

func (err Code) GRPCStatus() *status.Status {
	return errToStatus(err)
}

func (err *Error) GRPCStatus() *status.Status {
	return errToStatus(err)
}

type grpcError struct {
	status   *status.Status
	grpcCode codes.Code
	httpCode int
	id       string
	code     string
	reason   string
	message  string
}

func (err *grpcError) Error() string {
	return fmt.Sprintf("Code=%s ID=%s Reason=%s Message=(%v)", err.code, err.id, err.reason, err.message)
}

func (err *grpcError) GRPCStatus() *status.Status {
	return err.status
}

func (err *grpcError) HTTPCode() int {
	return err.httpCode
}

func (err *grpcError) GRPCCode() codes.Code {
	return err.grpcCode
}

func (err *grpcError) TypeCode() string {
	return err.code
}

// Is returns true if any of TypeCoder, HTTPCoder, GRPCCoder are a match between the error and target
func (err *grpcError) Is(target error) bool {
	if t, ok := target.(GRPCCoder); ok && err.grpcCode == t.GRPCCode() {
		return true
	}
	if t, ok := target.(HTTPCoder); ok && err.httpCode == t.HTTPCode() {
		return true
	}
	if t, ok := target.(TypeCoder); ok && err.code == t.TypeCode() {
		return true
	}
	return false
}

// GRPCCode returns the GRPC code for the given error or codes.OK when nil or codes.Unknown otherwise
func GRPCCode(err error) codes.Code {
	if err == nil {
		return ErrOK.GRPCCode()
	}
	var e GRPCCoder
	if errors.As(err, &e) {
		return e.GRPCCode()
	}
	return ErrUnknown.GRPCCode()
}

// SendGRPCError ensures that the error being used is sent with the correct code applied
//
// Use in the server when sending errors.
// If err is nil then SendGRPCError returns nil.
func SendGRPCError(err error) error {
	if err == nil {
		return nil
	}

	// Already setup with a grpcCode
	if _, ok := status.FromError(err); ok {
		return err
	}

	s := errToStatus(err)

	return s.Err()
}

// ReceiveGRPCError recreates the error with the coded Error reapplied
//
// Non-nil results can be used as both Error and *status.Status. Methods
// errors.Is()/errors.As(), and status.Convert()/status.FromError() will
// continue to work.
//
// Use in the clients when receiving errors.
// If err is nil then ReceiveGRPCError returns nil.
func ReceiveGRPCError(err error) error {
	if err == nil {
		return nil
	}

	s, ok := status.FromError(err)
	if !ok {
		return &grpcError{
			status:   s,
			grpcCode: ErrUnknown.GRPCCode(),
			httpCode: ErrUnknown.HTTPCode(),
			code:     ErrUnknown.TypeCode(),
			reason:   ErrUnknown.Error(),
			message:  err.Error(),
		}
	}

	grpcCode := s.Code()
	httpCode := ErrUnknown.HTTPCode()
	embedType := codeToError(grpcCode).TypeCode()
	id := ""
	reason := ErrUnknown.Error()

	for _, detail := range s.Details() {
		switch d := detail.(type) {
		case *ErrorDetail:
			grpcCode = codes.Code(d.GRPCCode)
			httpCode = int(d.HTTPCode)
			embedType = d.TypeCode
			id = d.ID
			reason = d.Reason
		default:
		}
	}

	return &grpcError{
		status:   s,
		grpcCode: grpcCode,
		httpCode: httpCode,
		code:     embedType,
		id:       id,
		reason:   reason,
		message:  s.Message(),
	}
}

// convert a code to a known Error type;
func codeToError(code codes.Code) Code {
	switch code {
	case codes.OK:
		return ErrOK
	case codes.Canceled:
		return ErrCanceled
	case codes.Unknown:
		return ErrUnknown
	case codes.InvalidArgument:
		return ErrInvalidArgument
	case codes.DeadlineExceeded:
		return ErrDeadlineExceeded
	case codes.NotFound:
		return ErrNotFound
	case codes.AlreadyExists:
		return ErrAlreadyExists
	case codes.PermissionDenied:
		return ErrPermissionDenied
	case codes.ResourceExhausted:
		return ErrResourceExhausted
	case codes.FailedPrecondition:
		return ErrFailedPrecondition
	case codes.Aborted:
		return ErrAborted
	case codes.OutOfRange:
		return ErrOutOfRange
	case codes.Unimplemented:
		return ErrUnimplemented
	case codes.Internal:
		return ErrInternal
	case codes.Unavailable:
		return ErrUnavailable
	case codes.DataLoss:
		return ErrDataLoss
	case codes.Unauthenticated:
		return ErrUnauthenticated
	default:
		return ErrInternal
	}
}

// convert an error into a gRPC *status.Status
func errToStatus(err error) *status.Status {
	grpcCode := ErrUnknown.GRPCCode()
	httpCode := ErrUnknown.HTTPCode()
	typeCode := ErrUnknown.TypeCode()

	// Set the grpcCode based on GRPCCoder output; otherwise leave as Unknown
	var grpcCoder GRPCCoder
	if errors.As(err, &grpcCoder) {
		grpcCode = grpcCoder.GRPCCode()
	}

	// short circuit building detailed errors if the code is OK
	if grpcCode == codes.OK {
		return status.New(codes.OK, "")
	}

	// Set the httpCode based on HTTPCoder output; otherwise leave as Unknown
	var httpCoder HTTPCoder
	if errors.As(err, &httpCoder) {
		httpCode = httpCoder.HTTPCode()
	}

	// Embed the specific error "type"; otherwise leave as "UNKNOWN"
	var typeCoder TypeCoder
	if errors.As(err, &typeCoder) {
		typeCode = typeCoder.TypeCode()
	}

	errInfo := &ErrorDetail{
		TypeCode: typeCode,
		GRPCCode: int64(grpcCode),
		HTTPCode: int64(httpCode),
	}

	var e Error
	if ok := errors.As(err, &e); ok {
		errInfo.Reason = e.reason
		errInfo.Message = e.message
	}

	s, _ := status.New(grpcCode, err.Error()).WithDetails(errInfo)

	return s
}
