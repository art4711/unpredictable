package unpredictable

//go:generate go run genchacha20/main.go -- chacha20_core.go

const (
	stateSize = 16
)

type block [stateSize]uint32

type stream block

func (s *stream) InitKeyStream(key *[8]uint32, nonce *[2]uint32) {
	// the magic constants for 256-bit keys
	s[0] = 0x61707865
	s[1] = 0x3320646e
	s[2] = 0x79622d32
	s[3] = 0x6b206574

	s[4] = key[0]
	s[5] = key[1]
	s[6] = key[2]
	s[7] = key[3]
	s[8] = key[4]
	s[9] = key[5]
	s[10] = key[6]
	s[11] = key[7]
	s[12] = 0
	s[13] = 0
	s[14] = nonce[0]
	s[15] = nonce[1]
}

func (s *stream) KeyBlock(b *block) {
	s.core(b)
}

func (s *stream) KeyBlocks(b *[16]block) {
	s.core_slice(b[:])
}