package gformat

import "strings"

var lower = []byte("0123456789abcdef")
var upper = []byte("0123456789ABCDEF")

func ToHexString(data []byte, upperCase bool) string {
	l := len(data)

	sb := strings.Builder{}
	sb.Grow(l * 2)

	digits := lower
	if upperCase {
		digits = upper
	}

	for i := 0; i < l; i++ {
		b := data[i]
		sb.WriteByte(digits[b>>4])
		sb.WriteByte(digits[b&15])
	}

	return sb.String()
}
