package errorsx

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// ========================================================
// Generic APIError Type
// ========================================================

// AppError provides the generic API and protocol agnostic error type all SDK generated exception types will implement
type AppError interface {
	error

	// Code returns the status code of the HTTP response
	Code() int

	// Unwrap returns the underlying error if one is set
	Unwrap() error
}

func WarnDuplicatedErrorWrap(err error) {
	log.Printf("FIXME: err(%T) is already wrapped as AppError, do NOT re-wrap it!", err)
}

// ========================================================
// Concreted Error Types
// ========================================================

type InvalidParamError struct {
	err error
}

func (e *InvalidParamError) Error() string {
	return fmt.Sprintf("APIGW.InvalidParamError: %v", e.err)
}

func (e *InvalidParamError) Code() int {
	return 400
}

func (e *InvalidParamError) Unwrap() error {
	return e.err
}

func NewInvalidParamError(err error) error {
	if _, ok := err.(AppError); ok {
		WarnDuplicatedErrorWrap(err)
		return err
	}

	return &InvalidParamError{
		err: err,
	}
}

type NotAllowedError struct {
	err error
}

func (e *NotAllowedError) Error() string {
	return fmt.Sprintf("APIGW.NotAllowedError: %v", e.err)
}

func (e *NotAllowedError) Code() int {
	return 503
}

func (e *NotAllowedError) Unwrap() error {
	return e.err
}

func NewNotAllowedError(err error) error {
	if _, ok := err.(AppError); ok {
		WarnDuplicatedErrorWrap(err)
		return err
	}

	return &NotAllowedError{
		err: err,
	}
}

type ResolverError struct {
	err error
}

func (e *ResolverError) Error() string {
	return fmt.Sprintf("APIGW.ResolverError: %v", e.err)
}

func (e *ResolverError) Code() int {
	return 500
}

func (e *ResolverError) Unwrap() error {
	return e.err
}

func NewResolverError(err error) error {
	if _, ok := err.(AppError); ok {
		WarnDuplicatedErrorWrap(err)
		return err
	}

	return &ResolverError{
		err: err,
	}
}

type NotFoundError struct {
	err error
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("APIGW.NotFoundError: %v", e.err)
}

func (e *NotFoundError) Code() int {
	return 404
}

func (e *NotFoundError) Unwrap() error {
	return e.err
}

func NewNotFoundError(err error) error {
	if _, ok := err.(AppError); ok {
		WarnDuplicatedErrorWrap(err)
		return err
	}

	return &NotFoundError{
		err: err,
	}
}

// ========================================================
// App Error Presenter
// ========================================================

func AppErrorPresenter(ctx context.Context, e error) *gqlerror.Error {
	err := &gqlerror.Error{}

	if gqlErr, ok := e.(*gqlerror.Error); ok {
		e = gqlErr.Unwrap()
	}

	var appErr AppError
	if errors.As(e, &appErr) {
		err.Message = appErr.Error()
		errcode.Set(err, strconv.Itoa(appErr.Code()))
		return err
	}

	err.Message = fmt.Sprintf("APIGW.InternalError: %s %+v", err.Message, e)
	errcode.Set(err, "500")
	return err
}
