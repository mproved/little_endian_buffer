package little_endian_buffer

import "math"

type Buffer struct {
	bytes   []byte
	pointer int
}

func BufferFromBytes(b []byte) *Buffer {
	return &Buffer{
		bytes:   b,
		pointer: 0,
	}
}

func (buf *Buffer) Bytes() []byte {
	return buf.bytes
}

func (buf *Buffer) Pointer() int {
	return buf.pointer
}

func (buf *Buffer) SetPointer(newPointer int) {
	buf.pointer = newPointer
}

func (buf *Buffer) LeftToRead() int {
	return len(buf.bytes) - buf.pointer
}

func (buf *Buffer) EnsureSize(size int) {
	neededSpace := (buf.pointer + size) - len(buf.bytes)

	if neededSpace > 0 {
		zeroSlice := make([]byte, neededSpace)
		buf.bytes = append(buf.bytes, zeroSlice...)
	}
}

func (buf *Buffer) ReadUint8() (value uint8) {
	byte1 := buf.bytes[buf.pointer+0]
	buf.pointer += 1

	value |= uint8(byte1) << 0

	return
}

func (buf *Buffer) WriteUint8(value uint8) {
	buf.EnsureSize(1)

	byte1 := byte(value >> 0)

	buf.bytes[buf.pointer+0] = byte1
	buf.pointer += 1
}

func (buf *Buffer) ReadUint16() (value uint16) {
	byte1 := buf.bytes[buf.pointer+0]
	byte2 := buf.bytes[buf.pointer+1]
	buf.pointer += 2

	value |= uint16(byte1) << 8
	value |= uint16(byte2) << 0

	return
}

func (buf *Buffer) WriteUint16(value uint16) {
	buf.EnsureSize(2)

	byte1 := byte(value >> 8)
	byte2 := byte(value >> 0)

	buf.bytes[buf.pointer+0] = byte1
	buf.bytes[buf.pointer+1] = byte2
	buf.pointer += 2
}

func (buf *Buffer) ReadUint32() (value uint32) {
	byte1 := buf.bytes[buf.pointer+0]
	byte2 := buf.bytes[buf.pointer+1]
	byte3 := buf.bytes[buf.pointer+2]
	byte4 := buf.bytes[buf.pointer+3]
	buf.pointer += 4

	value |= uint32(byte1) << 24
	value |= uint32(byte2) << 16
	value |= uint32(byte3) << 8
	value |= uint32(byte4) << 0

	return
}

func (buf *Buffer) WriteUint32(value uint32) {
	buf.EnsureSize(4)

	byte1 := byte(value >> 24)
	byte2 := byte(value >> 16)
	byte3 := byte(value >> 8)
	byte4 := byte(value >> 0)

	buf.bytes[buf.pointer+0] = byte1
	buf.bytes[buf.pointer+1] = byte2
	buf.bytes[buf.pointer+2] = byte3
	buf.bytes[buf.pointer+3] = byte4
	buf.pointer += 4
}

func (buf *Buffer) ReadUint64() (value uint64) {
	byte1 := buf.bytes[buf.pointer+0]
	byte2 := buf.bytes[buf.pointer+1]
	byte3 := buf.bytes[buf.pointer+2]
	byte4 := buf.bytes[buf.pointer+3]
	byte5 := buf.bytes[buf.pointer+4]
	byte6 := buf.bytes[buf.pointer+5]
	byte7 := buf.bytes[buf.pointer+6]
	byte8 := buf.bytes[buf.pointer+7]
	buf.pointer += 8

	value |= uint64(byte1) << 56
	value |= uint64(byte2) << 48
	value |= uint64(byte3) << 40
	value |= uint64(byte4) << 32
	value |= uint64(byte5) << 24
	value |= uint64(byte6) << 16
	value |= uint64(byte7) << 8
	value |= uint64(byte8) << 0

	return
}

func (buf *Buffer) WriteUint64(value uint64) {
	buf.EnsureSize(8)

	byte1 := byte(value >> 56)
	byte2 := byte(value >> 48)
	byte3 := byte(value >> 40)
	byte4 := byte(value >> 32)
	byte5 := byte(value >> 24)
	byte6 := byte(value >> 16)
	byte7 := byte(value >> 8)
	byte8 := byte(value >> 0)

	buf.bytes[buf.pointer+0] = byte1
	buf.bytes[buf.pointer+1] = byte2
	buf.bytes[buf.pointer+2] = byte3
	buf.bytes[buf.pointer+3] = byte4
	buf.bytes[buf.pointer+4] = byte5
	buf.bytes[buf.pointer+5] = byte6
	buf.bytes[buf.pointer+6] = byte7
	buf.bytes[buf.pointer+7] = byte8
	buf.pointer += 8
}

func (buf *Buffer) ReadInt8() int8 {
	return int8(buf.ReadUint8())
}

func (buf *Buffer) WriteInt8(value int8) {
	buf.WriteUint8(uint8(value))
}

func (buf *Buffer) ReadInt16() int16 {
	return int16(buf.ReadUint16())
}

func (buf *Buffer) WriteInt16(value int16) {
	buf.WriteUint16(uint16(value))
}

func (buf *Buffer) ReadInt32() int32 {
	return int32(buf.ReadUint32())
}

func (buf *Buffer) WriteInt32(value int32) {
	buf.WriteUint32(uint32(value))
}

func (buf *Buffer) ReadInt64() int64 {
	return int64(buf.ReadUint64())
}

func (buf *Buffer) WriteInt64(value int64) {
	buf.WriteUint64(uint64(value))
}

func (buf *Buffer) ReadBool() bool {
	if buf.ReadUint8() == 0 {
		return false
	} else {
		return true
	}
}

func (buf *Buffer) WriteBool(value bool) {
	if value {
		buf.WriteUint8(uint8(1))
	} else {
		buf.WriteUint8(uint8(0))
	}
}

func (buf *Buffer) ReadFloat32() float32 {
	return math.Float32frombits(buf.ReadUint32())
}

func (buf *Buffer) WriteFloat32(value float32) {
	buf.WriteUint32(math.Float32bits(value))
}

func (buf *Buffer) ReadFloat64() float64 {
	return math.Float64frombits(buf.ReadUint64())
}

func (buf *Buffer) WriteFloat64(value float64) {
	buf.WriteUint64(math.Float64bits(value))
}

func (buf *Buffer) ReadBytes(num int) (value []byte) {
	value = buf.bytes[buf.pointer : buf.pointer+num]
	buf.pointer += num

	return
}

func (buf *Buffer) WriteBytes(bytes []byte) {
	size := len(bytes)
	buf.EnsureSize(size)

	for i := 0; i < size; i += 1 {
		buf.bytes[buf.pointer+i] = bytes[i]
	}

	buf.pointer += size
}
