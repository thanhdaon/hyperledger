package errors

import (
	"bytes"
	"fmt"
)

type Op string

type Error struct {
	Op   Op
	Err  error
	Kind Kind
}

func (e *Error) Error() string {
	var buf bytes.Buffer

	// Print the current operation in our stack, if any.
	if e.Op != "" {
		fmt.Fprintf(&buf, "%s -> ", e.Op)
	}

	// If wrapping an error, print its Error() message.
	// Otherwise print the error code & message.
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	}
	return buf.String()
}

func E(args ...interface{}) error {
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.Op = arg
		case error:
			e.Err = arg
		case Kind:
			e.Kind = arg
		default:
			panic("bad call to E")
		}
	}
	return e
}

func Is(err error, kinds ...Kind) bool {
	for _, k := range kinds {
		if is(err, k) {
			return true
		}
	}
	return false
}

func is(err error, kind Kind) bool {
	e, ok := err.(*Error)
	if !ok {
		return false
	}

	if e.Kind != KUnknown {
		return e.Kind.slug == kind.slug
	}

	return Is(e.Err, kind)
}

func Unwrap(err error) error {
	e, ok := err.(*Error)
	if ok {
		return Unwrap(e.Err)
	}

	return err
}
