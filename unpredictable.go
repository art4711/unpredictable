package upr

import (
	badrand "math/rand"
	goodrand "crypto/rand"
	"unsafe"
)

type rs struct {
	buf [4096]byte
	window []byte
	c stream
	count int
}

const keysz = 32
const ivsz = 8

func (rs *rs) cinit() {
	var kbuf [keysz + ivsz]byte

	n, err := goodrand.Read(kbuf[:])
	if err != nil {
		panic(err)
	}
	if n != keysz + ivsz {
		panic("not enough entropy")
	}
	rs.c.InitKeyStream(kbuf[:keysz], kbuf[keysz:])
	rs.count = 1600000
}

func (rs *rs) stirIfNeeded(n int) {
	if rs.count <= n {
		rs.cinit()
	}
	if rs.count <= n {
		rs.count = 0
	} else {
		rs.count -= n
	}
}

func (rs *rs) rekey() {
	rs.stirIfNeeded(len(rs.buf))
	for i := 0; i < len(rs.buf); i += stateSize * 4 {
		rs.c.KeyBlock(rs.buf[i:])
	}
	rs.window = rs.buf[:]
}

func (rs *rs) Int63() int64 {
	if len(rs.window) < 8 {
		rs.rekey()
	}
	i := *(*int64)(unsafe.Pointer(&rs.window[0]))
	rs.window = rs.window[8:]
	return i & 0x7fffffffffffffff
}

func (rs *rs) Seed(s int64) {
	panic("no")
}

func NewMathRandSource() badrand.Source {
	return &rs{}
}