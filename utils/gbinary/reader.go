package gbinary

import (
	"encoding/binary"
	"glotsrans/utils/gerror"
	"io"
)

type BinaryReader struct {
	src       io.Reader
	buff      [8]byte
	buffSlice []byte
}

func NewReader(src io.Reader) *BinaryReader {
	reader := &BinaryReader{src: src}
	reader.buffSlice = reader.buff[:]
	return reader
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
	return int32(b.ReadUInt32())
}

func (b *BinaryReader) ReadUInt32() uint32 {
	b.readBytes(4)
	return binary.LittleEndian.Uint32(b.buffSlice)
}

func (b *BinaryReader) ReadInt64() int64 {
	b.readBytes(8)
	return int64(b.ReadUInt64())
}

func (b *BinaryReader) ReadUInt64() uint64 {
	b.readBytes(8)
	return binary.LittleEndian.Uint64(b.buffSlice)
}

func (b *BinaryReader) ReadByteSlice(n int) []byte {
	buff := make([]byte, n)
	b.MustRead(buff)
	return buff
}

func (b *BinaryReader) MustRead(p []byte) {
	count, err := b.src.Read(p)
	if err != nil {
		panic(err)
	}

	if count < len(p) {
		panic(gerror.CannotReadAsManyAsExpected(len(p), count))
	}
}

func (b *BinaryReader) Read(p []byte) (n int, err error) {
	return b.src.Read(p)
}
