package upr

/*
This code is derived from github.com/codahale/chacha20 with the following LICENSE:

The MIT License (MIT)

Copyright (c) 2014 Coda Hale

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

const (
	stateSize = 16                   // the size of ChaCha20's state, in words
	rounds = 20
)

type block [stateSize]uint32

type stream struct {
	state  block
}

func (s *stream) InitKeyStream(key *[8]uint32, nonce *[2]uint32) {
	// the magic constants for 256-bit keys
	s.state[0] = 0x61707865
	s.state[1] = 0x3320646e
	s.state[2] = 0x79622d32
	s.state[3] = 0x6b206574

	s.state[4] = key[0]
	s.state[5] = key[1]
	s.state[6] = key[2]
	s.state[7] = key[3]
	s.state[8] = key[4]
	s.state[9] = key[5]
	s.state[10] = key[6]
	s.state[11] = key[7]
	s.state[12] = 0
	s.state[13] = 0
	s.state[14] = nonce[0]
	s.state[15] = nonce[1]
}

func (s *stream) KeyBlock(b *block) {
	core(&s.state, b)
	s.state[12]++
	if s.state[12] == 0 {
		s.state[13]++
	}
}

func core(input, output *block) {
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

	output[0] = input[0] + x00
	output[1] = input[1] + x01
	output[2] = input[2] + x02
	output[3] = input[3] + x03
	output[4] = input[4] + x04
	output[5] = input[5] + x05
	output[6] = input[6] + x06
	output[7] = input[7] + x07
	output[8] = input[8] + x08
	output[9] = input[9] + x09
	output[10] = input[10] + x10
	output[11] = input[11] + x11
	output[12] = input[12] + x12
	output[13] = input[13] + x13
	output[14] = input[14] + x14
	output[15] = input[15] + x15
}
