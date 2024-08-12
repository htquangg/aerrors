/*
 * Copyright © 2024 Alex Huynh.
 */

// Package aerrors builds on Go 1.13 errors adding HTTP and GRPC code to your errors.
//
// Wrapping any error other than an Error will return an error with the message formatted
// as "<message>: <error>".
//
// Wrapping an Error will return an error with an unaltered error message.
//
// # Transmitting errors over GRPC
//
// The aerrors produced with wrap, that have also been wrapped first with an Err* can be
// send with SendGRPCError() and received with ReceiveGRPCError().
//
// You may want to create and use GRPC server and client interceptors to avoid having to
// call the Send/Receive methods in every handler.
//
// The Err* constants are errors and can be used directly is desired.
package aerrors
