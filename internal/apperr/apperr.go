// Package apperr là kiểu lỗi duy nhất mà tầng service trả lên handler:
// mang HTTP status + message cho client, và lỗi gốc (Err) để log.
package apperr

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	auth "github.com/phil/user-service/pkg/authkit"
)

type Error struct {
	Status  int
	Message string
	Err     error
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func (e *Error) Unwrap() error { return e.Err }

func New(status int, msg string) *Error { return &Error{Status: status, Message: msg} }

func BadRequest(msg string) *Error   { return New(http.StatusBadRequest, msg) }
func Unauthorized(msg string) *Error { return New(http.StatusUnauthorized, msg) }
func Forbidden(msg string) *Error    { return New(http.StatusForbidden, msg) }
func NotFound(msg string) *Error     { return New(http.StatusNotFound, msg) }
func Conflict(msg string) *Error     { return New(http.StatusConflict, msg) }

func Internal(err error) *Error {
	return &Error{Status: http.StatusInternalServerError, Message: "internal error", Err: err}
}

// Write ghi lỗi ra response dạng {"error": "..."}. Lỗi không phải *Error coi là 500.
func Write(w http.ResponseWriter, err error) {
	var e *Error
	if !errors.As(err, &e) {
		e = Internal(err)
	}
	if e.Status >= http.StatusInternalServerError {
		slog.Error("request failed", "status", e.Status, "error", e)
	}
	auth.WriteError(w, e.Status, e.Message)
}
