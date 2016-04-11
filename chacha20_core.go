package unpredictable

func (s *stream) core(output *block) {
	x00 := s.state[0]
	x01 := s.state[1]
	x02 := s.state[2]
	x03 := s.state[3]
	x04 := s.state[4]
	x05 := s.state[5]
	x06 := s.state[6]
	x07 := s.state[7]
	x08 := s.state[8]
	x09 := s.state[9]
	x10 := s.state[10]
	x11 := s.state[11]
	x12 := s.state[12]
	x13 := s.state[13]
	x14 := s.state[14]
	x15 := s.state[15]

	for i := 0; i < 20; i += 2 {
		x00 += x04
		x12 ^= x00
		x12 = (x12 << 16) | (x12 >> 16)
		x08 += x12
		x04 ^= x08
		x04 = (x04 << 12) | (x04 >> 20)
		x00 += x04
		x12 ^= x00
		x12 = (x12 << 8) | (x12 >> 24)
		x08 += x12
		x04 ^= x08
		x04 = (x04 << 7) | (x04 >> 25)

		x01 += x05
		x13 ^= x01
		x13 = (x13 << 16) | (x13 >> 16)
		x09 += x13
		x05 ^= x09
		x05 = (x05 << 12) | (x05 >> 20)
		x01 += x05
		x13 ^= x01
		x13 = (x13 << 8) | (x13 >> 24)
		x09 += x13
		x05 ^= x09
		x05 = (x05 << 7) | (x05 >> 25)

		x02 += x06
		x14 ^= x02
		x14 = (x14 << 16) | (x14 >> 16)
		x10 += x14
		x06 ^= x10
		x06 = (x06 << 12) | (x06 >> 20)
		x02 += x06
		x14 ^= x02
		x14 = (x14 << 8) | (x14 >> 24)
		x10 += x14
		x06 ^= x10
		x06 = (x06 << 7) | (x06 >> 25)

		x03 += x07
		x15 ^= x03
		x15 = (x15 << 16) | (x15 >> 16)
		x11 += x15
		x07 ^= x11
		x07 = (x07 << 12) | (x07 >> 20)
		x03 += x07
		x15 ^= x03
		x15 = (x15 << 8) | (x15 >> 24)
		x11 += x15
		x07 ^= x11
		x07 = (x07 << 7) | (x07 >> 25)

		x00 += x05
		x15 ^= x00
		x15 = (x15 << 16) | (x15 >> 16)
		x10 += x15
		x05 ^= x10
		x05 = (x05 << 12) | (x05 >> 20)
		x00 += x05
		x15 ^= x00
		x15 = (x15 << 8) | (x15 >> 24)
		x10 += x15
		x05 ^= x10
		x05 = (x05 << 7) | (x05 >> 25)

		x01 += x06
		x12 ^= x01
		x12 = (x12 << 16) | (x12 >> 16)
		x11 += x12
		x06 ^= x11
		x06 = (x06 << 12) | (x06 >> 20)
		x01 += x06
		x12 ^= x01
		x12 = (x12 << 8) | (x12 >> 24)
		x11 += x12
		x06 ^= x11
		x06 = (x06 << 7) | (x06 >> 25)

		x02 += x07
		x13 ^= x02
		x13 = (x13 << 16) | (x13 >> 16)
		x08 += x13
		x07 ^= x08
		x07 = (x07 << 12) | (x07 >> 20)
		x02 += x07
		x13 ^= x02
		x13 = (x13 << 8) | (x13 >> 24)
		x08 += x13
		x07 ^= x08
		x07 = (x07 << 7) | (x07 >> 25)

		x03 += x04
		x14 ^= x03
		x14 = (x14 << 16) | (x14 >> 16)
		x09 += x14
		x04 ^= x09
		x04 = (x04 << 12) | (x04 >> 20)
		x03 += x04
		x14 ^= x03
		x14 = (x14 << 8) | (x14 >> 24)
		x09 += x14
		x04 ^= x09
		x04 = (x04 << 7) | (x04 >> 25)

	}

	output[0] = s.state[0] + x00
	output[1] = s.state[1] + x01
	output[2] = s.state[2] + x02
	output[3] = s.state[3] + x03
	output[4] = s.state[4] + x04
	output[5] = s.state[5] + x05
	output[6] = s.state[6] + x06
	output[7] = s.state[7] + x07
	output[8] = s.state[8] + x08
	output[9] = s.state[9] + x09
	output[10] = s.state[10] + x10
	output[11] = s.state[11] + x11
	output[12] = s.state[12] + x12
	output[13] = s.state[13] + x13
	output[14] = s.state[14] + x14
	output[15] = s.state[15] + x15
	s.state[12]++
	if s.state[12] == 0 {
		s.state[13]++
	}

}
