package aerrors

// Errors named in line with GRPC codes and some that overlap with HTTP statuses
const (
	ErrOK                 ErrorCode = "OK"                  // HTTP: 200 GRPC: codes.OK
	ErrCanceled           ErrorCode = "CANCELED"            // HTTP: 408 GRPC: codes.Canceled
	ErrUnknown            ErrorCode = "UNKNOWN"             // HTTP: 510 GRPC: codes.Unknown
	ErrInvalidArgument    ErrorCode = "INVALID_ARGUMENT"    // HTTP: 400 GRPC: codes.InvalidArgument
	ErrDeadlineExceeded   ErrorCode = "DEADLINE_EXCEEDED"   // HTTP: 504 GRPC: codes.DeadlineExceeded
	ErrNotFound           ErrorCode = "NOT_FOUND"           // HTTP: 404 GRPC: codes.NotFound
	ErrAlreadyExists      ErrorCode = "ALREADY_EXISTS"      // HTTP: 409 GRPC: codes.AlreadyExists
	ErrPermissionDenied   ErrorCode = "PERMISSION_DENIED"   // HTTP: 403 GRPC: codes.PermissionDenied
	ErrResourceExhausted  ErrorCode = "RESOURCE_EXHAUSTED"  // HTTP: 429 GRPC: codes.ResourceExhausted
	ErrFailedPrecondition ErrorCode = "FAILED_PRECONDITION" // HTTP: 400 GRPC: codes.FailedPrecondition
	ErrAborted            ErrorCode = "ABORTED"             // HTTP: 409 GRPC: codes.Aborted
	ErrOutOfRange         ErrorCode = "OUT_OF_RANGE"        // HTTP: 422 GRPC: codes.OutOfRange
	ErrUnimplemented      ErrorCode = "UNIMPLEMENTED"       // HTTP: 501 GRPC: codes.Unimplemented
	ErrInternal           ErrorCode = "INTERNAL"            // HTTP: 500 GRPC: codes.Internal
	ErrUnavailable        ErrorCode = "UNAVAILABLE"         // HTTP: 503 GRPC: codes.Unavailable
	ErrDataLoss           ErrorCode = "DATA_LOSS"           // HTTP: 500 GRPC: codes.DataLoss
	ErrUnauthenticated    ErrorCode = "UNAUTHENTICATED"     // HTTP: 401 GRPC: codes.Unauthenticated
)

// Errors named in line with HTTP statuses
const (
	ErrBadRequest                 ErrorCode = "BAD_REQUEST"                   // HTTP: 400 GRPC: codes.InvalidArgument
	ErrUnauthorized               ErrorCode = "UNAUTHORIZED"                  // HTTP: 401 GRPC: codes.Unauthenticated
	ErrForbidden                  ErrorCode = "FORBIDDEN"                     // HTTP: 403 GRPC: codes.PermissionDenied
	ErrMethodNotAllowed           ErrorCode = "METHOD_NOT_ALLOWED"            // HTTP: 405 GRPC: codes.Unimplemented
	ErrRequestTimeout             ErrorCode = "REQUEST_TIMEOUT"               // HTTP: 408 GRPC: codes.DeadlineExceeded
	ErrConflict                   ErrorCode = "CONFLICT"                      // HTTP: 409 GRPC: codes.AlreadyExists
	ErrImATeapot                  ErrorCode = "IM_A_TEAPOT"                   // HTTP: 418 GRPC: codes.Unknown
	ErrUnprocessableEntity        ErrorCode = "UNPROCESSABLE_ENTITY"          // HTTP: 422 GRPC: codes.InvalidArgument
	ErrTooManyRequests            ErrorCode = "TOO_MANY_REQUESTS"             // HTTP: 429 GRPC: codes.ResourceExhausted
	ErrUnavailableForLegalReasons ErrorCode = "UNAVAILABLE_FOR_LEGAL_REASONS" // HTTP: 451 GRPC: codes.Unavailable
	ErrInternalServerError        ErrorCode = "INTERNAL_SERVER_ERROR"         // HTTP: 500 GRPC: codes.Internal
	ErrNotImplemented             ErrorCode = "NOT_IMPLEMENTED"               // HTTP: 501 GRPC: codes.Unimplemented
	ErrBadGateway                 ErrorCode = "BAD_GATEWAY"                   // HTTP: 502 GRPC: codes.Aborted
	ErrServiceUnavailable         ErrorCode = "SERVICE_UNAVAILABLE"           // HTTP: 503 GRPC: codes.Unavailable
	ErrGatewayTimeout             ErrorCode = "GATEWAY_TIMEOUT"               // HTTP: 504 GRPC: codes.DeadlineExceeded
)

func InvalidArgument(reason string, opts ...AErrorOption) *AError {
	return New(ErrInvalidArgument, reason, opts...)
}

func IsValidArgument(err *AError) bool {
	return err.code == ErrInvalidArgument
}

func FailedPrecondition(reason string, opts ...AErrorOption) *AError {
	return New(ErrFailedPrecondition, reason, opts...)
}

func IsFailedPrecondition(err *AError) bool {
	return err.code == ErrFailedPrecondition
}

func Unauthentication(reason string, opts ...AErrorOption) *AError {
	return New(ErrUnauthenticated, reason, opts...)
}

func IsUnauthentication(err *AError) bool {
	return err.code == ErrUnauthenticated || err.code == ErrUnauthorized
}

func PermissionDenied(reason string, opts ...AErrorOption) *AError {
	return New(ErrPermissionDenied, reason, opts...)
}

func IsPermissionDenied(err *AError) bool {
	return err.code == ErrPermissionDenied || err.code == ErrForbidden
}

func NotFound(reason string, opts ...AErrorOption) *AError {
	return New(ErrNotFound, reason, opts...)
}

func IsNotFound(err *AError) bool {
	return err.code == ErrNotFound
}

func AlreadyExists(reason string, opts ...AErrorOption) *AError {
	return New(ErrConflict, reason, opts...)
}

func IsAlreadyExists(err *AError) bool {
	return err.code == ErrAlreadyExists || err.code == ErrConflict || err.code == ErrAborted
}

func Internal(reason string, opts ...AErrorOption) *AError {
	return New(ErrInternal, reason, opts...)
}

func IsInternal(err *AError) bool {
	return err.code == ErrInternal || err.code == ErrInternalServerError || err.code == ErrDataLoss
}

func Unimplemented(reason string, opts ...AErrorOption) *AError {
	return New(ErrUnimplemented, reason, opts...)
}

func IsUnimplemented(err *AError) bool {
	return err.code == ErrUnimplemented || err.code == ErrMethodNotAllowed || err.code == ErrNotImplemented
}

func Unavailable(reason string, opts ...AErrorOption) *AError {
	return New(ErrUnavailable, reason, opts...)
}

func IsUnavailable(err *AError) bool {
	return err.code == ErrUnavailable || err.code == ErrServiceUnavailable
}

func DeadlineExceeded(reason string, opts ...AErrorOption) *AError {
	return New(ErrGatewayTimeout, reason, opts...)
}

func IsDeadlineExceeded(err *AError) bool {
	return err.code == ErrDeadlineExceeded || err.code == ErrGatewayTimeout
}
