package unpredictable

//go:generate go run genchacha20/main.go -- chacha20_core.go

const (
	stateSize = 16
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
	s.core(b)
	s.state[12]++
	if s.state[12] == 0 {
		s.state[13]++
	}
}
