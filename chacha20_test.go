package unpredictable

import (
	"testing"
)

func cmp(t *testing.T, a, b block) {
	for i := range a {
		if a[i] != b[i] {
			t.Errorf("cmp failed(%d) %v != %v", i, a, b)
			return
		}
	}
}

func TestChaCha20(t *testing.T) {
	s := stream{}
	s.InitKeyStream(&[8]uint32{ 0, 1, 2, 3, 4, 5, 6, 7 }, &[2]uint32{ 17, 42 })
	var b1 block
	var b2 block
	s.KeyBlock(&b1)
	s.KeyBlock(&b2)

	e1 := block{
		2871835573, 3650430976, 856247812, 915926598,
		2377918783, 135788522, 1770096939, 3298343243,
		2918484449, 1921299346, 580272198, 646137650,
		2266901406, 1672342262, 2374457603, 3525023161,
	}
	e2 := block{
		2691987456, 1496855720, 2864756992, 683310477,
		4005061894, 989127363, 3113588261, 1295586749,
		413581430, 3188783649, 349418763, 782679946,
		431281646, 3571683293, 1617743812, 37269309,
	}

	cmp(t, b1, e1)
	cmp(t, b2, e2)
}

func TestChaCha20Blocks(t *testing.T) {
	s1 := stream{}
	s1.InitKeyStream(&[8]uint32{ 42, 17, 2, 3, 4, 5, 6, 7 }, &[2]uint32{ 17, 42 })
	s2 := stream{}
	s2.InitKeyStream(&[8]uint32{ 42, 17, 2, 3, 4, 5, 6, 7 }, &[2]uint32{ 17, 42 })

	for i := 0; i < 50000; i++ {
		var bls [16]block
		s2.KeyBlocks(&bls)
		for j := range bls {
			var bl block
			s1.KeyBlock(&bl)
			cmp(t, bl, bls[j])
		}
	}
}

func BenchmarkChaCha20(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(64))
	s := &stream{}
	s.InitKeyStream(&[8]uint32{ 0, 1, 2, 3, 4, 5, 6, 7 }, &[2]uint32{ 17, 42 })
	var bl block
	for i := 0; i < b.N; i++ {
		s.KeyBlock(&bl)
	}
}

func BenchmarkChaCha20Blocks(b *testing.B) {
	b.ReportAllocs()
	b.SetBytes(int64(1024))
	s := &stream{}
	s.InitKeyStream(&[8]uint32{ 0, 1, 2, 3, 4, 5, 6, 7 }, &[2]uint32{ 17, 42 })
	var bl [16]block
	s.KeyBlocks(&bl)
	for i := 0; i < b.N; i++ {
		s.KeyBlocks(&bl)
	}
}