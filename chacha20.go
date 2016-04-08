package upr

import (
	"encoding/binary"
	"unsafe"
)

const (
	stateSize = 16                   // the size of ChaCha20's state, in words
	rounds = 20
)

type stream struct {
	state  [stateSize]uint32 // the state as an array of 16 32-bit words
}

func (s *stream) InitKeyStream(key []byte, nonce []byte) {
	// the magic constants for 256-bit keys
	s.state[0] = 0x61707865
	s.state[1] = 0x3320646e
	s.state[2] = 0x79622d32
	s.state[3] = 0x6b206574

	s.state[4] = binary.LittleEndian.Uint32(key[0:])
	s.state[5] = binary.LittleEndian.Uint32(key[4:])
	s.state[6] = binary.LittleEndian.Uint32(key[8:])
	s.state[7] = binary.LittleEndian.Uint32(key[12:])
	s.state[8] = binary.LittleEndian.Uint32(key[16:])
	s.state[9] = binary.LittleEndian.Uint32(key[20:])
	s.state[10] = binary.LittleEndian.Uint32(key[24:])
	s.state[11] = binary.LittleEndian.Uint32(key[28:])

	if len(nonce) != 8 {
		panic("invalid nonce size")
	}
	s.state[12] = 0
	s.state[13] = 0
	s.state[14] = binary.LittleEndian.Uint32(nonce[0:])
	s.state[15] = binary.LittleEndian.Uint32(nonce[4:])
}

func NewKeyStream(key []byte, nonce []byte) *stream {
	s := &stream{}
	s.InitKeyStream(key, nonce)
	return s
}

func (s *stream) KeyBlock(b []byte) {
	core2(&s.state, (*[stateSize]uint32)(unsafe.Pointer(&b[0])))
	ns := s.state[12] + 1
	s.state[12] = ns
	if ns == 0 {
		s.state[13]++
	}
}

func core2(input, output *[stateSize]uint32) {
	var (
		x00 = input[0]
		x01 = input[1]
		x02 = input[2]
		x03 = input[3]
		x04 = input[4]
		x05 = input[5]
		x06 = input[6]
		x07 = input[7]
		x08 = input[8]
		x09 = input[9]
		x10 = input[10]
		x11 = input[11]
		x12 = input[12]
		x13 = input[13]
		x14 = input[14]
		x15 = input[15]
	)

	var x uint32

	// Unrolling all 20 rounds kills performance on modern Intel processors
	// (Tested on a i5 Haswell, likely applies to Sandy Bridge+), due to uop
	// cache thrashing.  The straight forward 2 rounds per loop implementation
	// of this has double the performance of the fully unrolled version.
	for i := 0; i < rounds; i += 2 {
		x00 += x04
		x = x12 ^ x00
		x12 = (x << 16) | (x >> 16)
		x08 += x12
		x = x04 ^ x08
		x04 = (x << 12) | (x >> 20)
		x00 += x04
		x = x12 ^ x00
		x12 = (x << 8) | (x >> 24)
		x08 += x12
		x = x04 ^ x08
		x04 = (x << 7) | (x >> 25)
		x01 += x05
		x = x13 ^ x01
		x13 = (x << 16) | (x >> 16)
		x09 += x13
		x = x05 ^ x09
		x05 = (x << 12) | (x >> 20)
		x01 += x05
		x = x13 ^ x01
		x13 = (x << 8) | (x >> 24)
		x09 += x13
		x = x05 ^ x09
		x05 = (x << 7) | (x >> 25)
		x02 += x06
		x = x14 ^ x02
		x14 = (x << 16) | (x >> 16)
		x10 += x14
		x = x06 ^ x10
		x06 = (x << 12) | (x >> 20)
		x02 += x06
		x = x14 ^ x02
		x14 = (x << 8) | (x >> 24)
		x10 += x14
		x = x06 ^ x10
		x06 = (x << 7) | (x >> 25)
		x03 += x07
		x = x15 ^ x03
		x15 = (x << 16) | (x >> 16)
		x11 += x15
		x = x07 ^ x11
		x07 = (x << 12) | (x >> 20)
		x03 += x07
		x = x15 ^ x03
		x15 = (x << 8) | (x >> 24)
		x11 += x15
		x = x07 ^ x11
		x07 = (x << 7) | (x >> 25)
		x00 += x05
		x = x15 ^ x00
		x15 = (x << 16) | (x >> 16)
		x10 += x15
		x = x05 ^ x10
		x05 = (x << 12) | (x >> 20)
		x00 += x05
		x = x15 ^ x00
		x15 = (x << 8) | (x >> 24)
		x10 += x15
		x = x05 ^ x10
		x05 = (x << 7) | (x >> 25)
		x01 += x06
		x = x12 ^ x01
		x12 = (x << 16) | (x >> 16)
		x11 += x12
		x = x06 ^ x11
		x06 = (x << 12) | (x >> 20)
		x01 += x06
		x = x12 ^ x01
		x12 = (x << 8) | (x >> 24)
		x11 += x12
		x = x06 ^ x11
		x06 = (x << 7) | (x >> 25)
		x02 += x07
		x = x13 ^ x02
		x13 = (x << 16) | (x >> 16)
		x08 += x13
		x = x07 ^ x08
		x07 = (x << 12) | (x >> 20)
		x02 += x07
		x = x13 ^ x02
		x13 = (x << 8) | (x >> 24)
		x08 += x13
		x = x07 ^ x08
		x07 = (x << 7) | (x >> 25)
		x03 += x04
		x = x14 ^ x03
		x14 = (x << 16) | (x >> 16)
		x09 += x14
		x = x04 ^ x09
		x04 = (x << 12) | (x >> 20)
		x03 += x04
		x = x14 ^ x03
		x14 = (x << 8) | (x >> 24)
		x09 += x14
		x = x04 ^ x09
		x04 = (x << 7) | (x >> 25)
	}

	output[0] = x00
	output[1] = x01
	output[2] = x02
	output[3] = x03
	output[4] = x04
	output[5] = x05
	output[6] = x06
	output[7] = x07
	output[8] = x08
	output[9] = x09
	output[10] = x10
	output[11] = x11
	output[12] = x12
	output[13] = x13
	output[14] = x14
	output[15] = x15
}
