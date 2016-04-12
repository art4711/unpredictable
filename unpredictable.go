package unpredictable

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
	have int
	count int
	inited bool
}


func (rs *rs) willConsume(n int) int {
	if rs.count <= n {
		var kbuf [keysz + ivsz]byte
		n, err := goodrand.Read(kbuf[:])
		if err != nil {
			panic(err)
		}
		if n != keysz + ivsz {
			panic("not enough entropy")
		}
		rs.count = 1600000
		if !rs.inited {
			rs.inited = true
			rs.c.InitKeyStream((*[8]uint32)(unsafe.Pointer(&kbuf[0])), (*[2]uint32)(unsafe.Pointer(&kbuf[keysz])))
		} else {
			rs.rekey(kbuf[:])
			rs.have = 0
		}
	}
	rs.count -= n
	if rs.have < n {
		rs.rekey(nil)
	}
	r := bufSz - rs.have
	rs.have -= n
	return r
}

func (rs *rs) rekey(extra []byte) {
	rs.c.KeyBlocks((*[16]block)(unsafe.Pointer(&rs.buf[0])))
	rs.have = bufSz
	if extra != nil {
		for i := range extra {
			rs.buf[i] ^= extra[i]
		}
	}
	rs.c.InitKeyStream((*[8]uint32)(unsafe.Pointer(&rs.buf[0])), (*[2]uint32)(unsafe.Pointer(&rs.buf[keysz])))
	rs.have -= (keysz + ivsz)
}

func (rs *rs) Int63() int64 {
	o := rs.willConsume(8)
	i := *(*int64)(unsafe.Pointer(&rs.buf[o]))
	return i & 0x7fffffffffffffff
}

func (rs *rs) Seed(s int64) {
	panic("no")
}

func NewMathRandSource() badrand.Source {
	return &rs{}
}
