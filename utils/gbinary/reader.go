package gbinary

import (
	"glotsrans/utils/gerror"
	"io"
)

type BinaryReader struct {
	src io.Reader
	buff [8]byte
}

func NewReader(src io.Reader) *BinaryReader {
	return &BinaryReader{src: src}
}

func (b *BinaryReader) readBytes(n int) {
	count, err := b.src.Read(b.buff[:n])
	if err != nil {
		panic(err)
	}

	if count < n {
		panic(gerror.CannotReadAsManyAsExpected(n, count))
	}
}

func (b *BinaryReader) ReadInt32() int32 {
	b.readBytes(4)
	return int32(b.buff[3]) << 24 | int32(b.buff[2]) << 16 | int32(b.buff[1]) << 8 | int32(b.buff[0])
}

func (b *BinaryReader) ReadUInt32() uint32 {
	b.readBytes(4)
	return uint32(b.buff[3]) << 24 | uint32(b.buff[2]) << 16 | uint32(b.buff[1]) << 8 | uint32(b.buff[0])
}

func (b *BinaryReader) ReadByteSlice(n int) []byte {
	buff := make([]byte, n)
	count, err := b.src.Read(buff)
	if err != nil {
		panic(err)
	}

	if count < n {
		panic(gerror.CannotReadAsManyAsExpected(n, count))
	}

	return buff
}
