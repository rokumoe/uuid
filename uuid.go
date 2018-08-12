package uuid

import (
	"crypto/rand"
	"unsafe"
)

func GenerateV4(u *[16]byte) {
	rand.Read(u[:])
	u[7] = (u[7] & 0x0f) | 0x40
	u[8] = (u[8] & 0x3f) | 0x80
}

const digits = "0123456789abcdef"

func hexOctet(s *[36]byte, i int, u *[16]byte, j int) {
	l := u[j]
	s[i+0] = digits[l>>4]
	s[i+1] = digits[l&0xf]
	h := u[j+1]
	s[i+2] = digits[h>>4]
	s[i+3] = digits[h&0xf]
}

func Format(u *[16]byte) string {
	s := new([36]byte)
	hexOctet(s, 0, u, 0)
	hexOctet(s, 4, u, 2)
	s[8] = '-'
	hexOctet(s, 9, u, 4)
	s[13] = '-'
	hexOctet(s, 14, u, 6)
	s[18] = '-'
	hexOctet(s, 19, u, 8)
	s[23] = '-'
	hexOctet(s, 24, u, 10)
	hexOctet(s, 28, u, 12)
	hexOctet(s, 32, u, 14)
	var hack [2]uintptr
	hack[0] = uintptr(unsafe.Pointer(s))
	hack[1] = 36
	return *(*string)(unsafe.Pointer(&hack))
}
