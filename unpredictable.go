package upr

import (
	badrand "math/rand"
	goodrand "crypto/rand"
	"unsafe"
)

const keysz = 32
const ivsz = 8
const bufSz = 1024

type rs struct {
	c stream
	buf [bufSz]byte
	inBuf int
	count int
}

func (rs *rs) stirIfNeeded(n int) {
	if rs.count <= n {
		var kbuf [keysz + ivsz]byte
		n, err := goodrand.Read(kbuf[:])
		if err != nil {
			panic(err)
		}
		if n != keysz + ivsz {
			panic("not enough entropy")
		}
		rs.c.InitKeyStream((*[8]uint32)(unsafe.Pointer(&kbuf[0])), (*[2]uint32)(unsafe.Pointer(&kbuf[keysz])))
		rs.count = 1600000
	}
	rs.count -= n
}

func (rs *rs) rekey() {
	rs.stirIfNeeded(len(rs.buf))
	for i := 0; i < len(rs.buf); i += stateSize * 4 {
		rs.c.KeyBlock((*block)(unsafe.Pointer(&rs.buf[i])))
	}
	rs.inBuf = bufSz
}

func (rs *rs) Int63() int64 {
	if rs.inBuf < 8 {
		rs.rekey()
	}
	o := bufSz - rs.inBuf
	rs.inBuf -= 8
	i := *(*int64)(unsafe.Pointer(&rs.buf[o]))
	return i & 0x7fffffffffffffff
}

func (rs *rs) Seed(s int64) {
	panic("no")
}

func NewMathRandSource() badrand.Source {
	return &rs{}
}