package gerror

import "fmt"

func CannotReadAsManyAsExpected(expected, actual int) error {
	return fmt.Errorf("cannot read %d bytes: only %d available", expected, actual)
}
