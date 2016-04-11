package unpredictable

func (s stream) core(output *block) {
	var (
		x00 = s.state[0]
		x01 = s.state[1]
		x02 = s.state[2]
		x03 = s.state[3]
		x04 = s.state[4]
		x05 = s.state[5]
		x06 = s.state[6]
		x07 = s.state[7]
		x08 = s.state[8]
		x09 = s.state[9]
		x10 = s.state[10]
		x11 = s.state[11]
		x12 = s.state[12]
		x13 = s.state[13]
		x14 = s.state[14]
		x15 = s.state[15]
	)

	var x uint32

	for i := 0; i < 20; i += 2 {
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
}
