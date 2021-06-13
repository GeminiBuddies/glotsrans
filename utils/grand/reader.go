package grand

import (
	"encoding/binary"
	"io"
	"math/rand"
)

type rngReader struct {
	source rand.Source64
}

func (r *rngReader) Read(p []byte) (n int, err error) {
	cnt := len(p)
	ptr := 0

	for ptr+8 <= cnt {
		binary.LittleEndian.PutUint64(p[ptr:ptr+8], r.source.Uint64())
		ptr += 8
	}

	if ptr < cnt {
		tail := r.source.Uint64()

		for ptr < cnt {
			p[ptr] = byte(tail)
			tail >>= 8
			ptr += 1
		}
	}

	return cnt, err
}

func Reader(source rand.Source64) io.Reader {
	return &rngReader{source: source}
}
