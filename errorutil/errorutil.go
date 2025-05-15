package errorutil

import (
	"errors"
	"fmt"
)

// WrapError wraps an existing error with additional context
func WrapError(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}

// MustError panics if err is not nil, otherwise does nothing
func MustError(err error) {
	if err != nil {
		panic(err)
	}
}

// MustErrorf panics if err is not nil with a formatted message
func MustErrorf(err error, format string, args ...interface{}) {
	if err != nil {
		panic(fmt.Errorf(format+": %w", append(args, err)...))
	}
}

// ErrorChain creates a chain of errors
type ErrorChain struct {
	errors []error
}

// NewErrorChain creates a new error chain
func NewErrorChain() *ErrorChain {
	return &ErrorChain{
		errors: make([]error, 0),
	}
}

// Add adds an error to the chain if it's not nil
func (c *ErrorChain) Add(err error) *ErrorChain {
	if err != nil {
		c.errors = append(c.errors, err)
	}
	return c
}

// Addf adds a formatted error to the chain
func (c *ErrorChain) Addf(format string, args ...interface{}) *ErrorChain {
	c.errors = append(c.errors, fmt.Errorf(format, args...))
	return c
}

// Error returns the error chain as a string
func (c *ErrorChain) Error() string {
	if len(c.errors) == 0 {
		return ""
	}

	errStr := "Multiple errors occurred:"
	for i, err := range c.errors {
		errStr += fmt.Sprintf("\n  %d. %s", i+1, err.Error())
	}
	return errStr
}

// HasErrors returns true if the chain contains errors
func (c *ErrorChain) HasErrors() bool {
	return len(c.errors) > 0
}

// Result returns nil if there are no errors, otherwise returns the error chain
func (c *ErrorChain) Result() error {
	if c.HasErrors() {
		return c
	}
	return nil
}

// ErrNotFound is a common error for when an item isn't found
var ErrNotFound = errors.New("item not found")

// ErrInvalidInput is a common error for invalid input
var ErrInvalidInput = errors.New("invalid input")

// ErrUnauthorized is a common error for unauthorized access
var ErrUnauthorized = errors.New("unauthorized access")

// Is checks if err is a specific error. It is a wrapper around errors.Is
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As finds the first error in err's chain that matches target, and if one is found,
// sets target to that error value and returns true. It is a wrapper around errors.As
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}
