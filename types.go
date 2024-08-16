package aerrors

// Errors named in line with GRPC codes and some that overlap with HTTP statuses
const (
	ErrOK                 Code = "OK"                  // HTTP: 200 GRPC: codes.OK
	ErrCanceled           Code = "CANCELED"            // HTTP: 408 GRPC: codes.Canceled
	ErrUnknown            Code = "UNKNOWN"             // HTTP: 510 GRPC: codes.Unknown
	ErrInvalidArgument    Code = "INVALID_ARGUMENT"    // HTTP: 400 GRPC: codes.InvalidArgument
	ErrDeadlineExceeded   Code = "DEADLINE_EXCEEDED"   // HTTP: 504 GRPC: codes.DeadlineExceeded
	ErrNotFound           Code = "NOT_FOUND"           // HTTP: 404 GRPC: codes.NotFound
	ErrAlreadyExists      Code = "ALREADY_EXISTS"      // HTTP: 409 GRPC: codes.AlreadyExists
	ErrPermissionDenied   Code = "PERMISSION_DENIED"   // HTTP: 403 GRPC: codes.PermissionDenied
	ErrResourceExhausted  Code = "RESOURCE_EXHAUSTED"  // HTTP: 429 GRPC: codes.ResourceExhausted
	ErrFailedPrecondition Code = "FAILED_PRECONDITION" // HTTP: 400 GRPC: codes.FailedPrecondition
	ErrAborted            Code = "ABORTED"             // HTTP: 409 GRPC: codes.Aborted
	ErrOutOfRange         Code = "OUT_OF_RANGE"        // HTTP: 422 GRPC: codes.OutOfRange
	ErrUnimplemented      Code = "UNIMPLEMENTED"       // HTTP: 501 GRPC: codes.Unimplemented
	ErrInternal           Code = "INTERNAL"            // HTTP: 500 GRPC: codes.Internal
	ErrUnavailable        Code = "UNAVAILABLE"         // HTTP: 503 GRPC: codes.Unavailable
	ErrDataLoss           Code = "DATA_LOSS"           // HTTP: 500 GRPC: codes.DataLoss
	ErrUnauthenticated    Code = "UNAUTHENTICATED"     // HTTP: 401 GRPC: codes.Unauthenticated
)

// Errors named in line with HTTP statuses
const (
	ErrBadRequest                 Code = "BAD_REQUEST"                   // HTTP: 400 GRPC: codes.InvalidArgument
	ErrUnauthorized               Code = "UNAUTHORIZED"                  // HTTP: 401 GRPC: codes.Unauthenticated
	ErrForbidden                  Code = "FORBIDDEN"                     // HTTP: 403 GRPC: codes.PermissionDenied
	ErrMethodNotAllowed           Code = "METHOD_NOT_ALLOWED"            // HTTP: 405 GRPC: codes.Unimplemented
	ErrRequestTimeout             Code = "REQUEST_TIMEOUT"               // HTTP: 408 GRPC: codes.DeadlineExceeded
	ErrConflict                   Code = "CONFLICT"                      // HTTP: 409 GRPC: codes.AlreadyExists
	ErrImATeapot                  Code = "IM_A_TEAPOT"                   // HTTP: 418 GRPC: codes.Unknown
	ErrUnprocessableEntity        Code = "UNPROCESSABLE_ENTITY"          // HTTP: 422 GRPC: codes.InvalidArgument
	ErrTooManyRequests            Code = "TOO_MANY_REQUESTS"             // HTTP: 429 GRPC: codes.ResourceExhausted
	ErrUnavailableForLegalReasons Code = "UNAVAILABLE_FOR_LEGAL_REASONS" // HTTP: 451 GRPC: codes.Unavailable
	ErrInternalServerError        Code = "INTERNAL_SERVER_ERROR"         // HTTP: 500 GRPC: codes.Internal
	ErrNotImplemented             Code = "NOT_IMPLEMENTED"               // HTTP: 501 GRPC: codes.Unimplemented
	ErrBadGateway                 Code = "BAD_GATEWAY"                   // HTTP: 502 GRPC: codes.Aborted
	ErrServiceUnavailable         Code = "SERVICE_UNAVAILABLE"           // HTTP: 503 GRPC: codes.Unavailable
	ErrGatewayTimeout             Code = "GATEWAY_TIMEOUT"               // HTTP: 504 GRPC: codes.DeadlineExceeded
)

func InvalidArgument(reason string) Builder {
	return New(ErrInvalidArgument, reason)
}

func IsValidArgument(err *AError) bool {
	return err.code == ErrInvalidArgument
}

func FailedPrecondition(reason string) Builder {
	return New(ErrFailedPrecondition, reason)
}

func IsFailedPrecondition(err *AError) bool {
	return err.code == ErrFailedPrecondition
}

func Unauthentication(reason string) Builder {
	return New(ErrUnauthenticated, reason)
}

func IsUnauthentication(err *AError) bool {
	return err.code == ErrUnauthenticated || err.code == ErrUnauthorized
}

func PermissionDenied(reason string) Builder {
	return New(ErrPermissionDenied, reason)
}

func IsPermissionDenied(err *AError) bool {
	return err.code == ErrPermissionDenied || err.code == ErrForbidden
}

func NotFound(reason string) Builder {
	return New(ErrNotFound, reason)
}

func IsNotFound(err *AError) bool {
	return err.code == ErrNotFound
}

func AlreadyExists(reason string) Builder {
	return New(ErrConflict, reason)
}

func IsAlreadyExists(err *AError) bool {
	return err.code == ErrAlreadyExists || err.code == ErrConflict || err.code == ErrAborted
}

func Internal(reason string) Builder {
	return New(ErrInternal, reason)
}

func IsInternal(err *AError) bool {
	return err.code == ErrInternal || err.code == ErrInternalServerError || err.code == ErrDataLoss
}

func Unimplemented(reason string) Builder {
	return New(ErrUnimplemented, reason)
}

func IsUnimplemented(err *AError) bool {
	return err.code == ErrUnimplemented || err.code == ErrMethodNotAllowed ||
		err.code == ErrNotImplemented
}

func Unavailable(reason string) Builder {
	return New(ErrUnavailable, reason)
}

func IsUnavailable(err *AError) bool {
	return err.code == ErrUnavailable || err.code == ErrServiceUnavailable
}

func DeadlineExceeded(reason string) Builder {
	return New(ErrGatewayTimeout, reason)
}

func IsDeadlineExceeded(err *AError) bool {
	return err.code == ErrDeadlineExceeded || err.code == ErrGatewayTimeout
}
