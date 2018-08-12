package uuid

import (
	"fmt"
	"testing"
)

func TestGenerateV4(t *testing.T) {
	var data [16]byte
	GenerateV4(&data)
	t.Log(data)
}

func TestFormat(t *testing.T) {
	var data [16]byte
	GenerateV4(&data)
	s1 := Format(&data)
	t.Log("s1", s1)
	s2 := fmt.Sprintf("%02x-%02x-%02x-%02x-%02x", data[0:4], data[4:6], data[6:8], data[8:10], data[10:16])
	t.Log("s2", s2)
	if s1 != s2 {
		t.Fatal()
	}
}

var benchid [16]byte

func init() {
	GenerateV4(&benchid)
}

func BenchmarkFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Format(&benchid)
	}
}
