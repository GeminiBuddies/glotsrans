package gerror

import "fmt"

func CannotReadAsManyAsExpected(expected, actual int) error {
	return fmt.Errorf("cannot read %d bytes: only %d available", expected, actual)
}

func RequestTooManyBytes(requested, limit int) error {
	return fmt.Errorf("too many (%d) bytes requested: %d maximum", requested, limit)
}
