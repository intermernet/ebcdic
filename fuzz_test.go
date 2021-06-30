//
//Tests for fuzzing package ebcdic
//
package ebcdic

import (
	"bytes"
	"testing"
)

func FuzzRoundTrip(f *testing.F) {
	f.Add([]byte{0xE3, 0x85, 0xA2, 0xA3, 0x89, 0x95, 0x87, 0x6B, 0x40, 0xF1, 0xF2, 0xF3})
	f.Fuzz(func(t *testing.T, data []byte) {
		if len(data) == 0 {
			t.Skip()
		}
		t.Parallel() // seed corpus tests can run in parallel
		encData, err := Encode(data)
		if err != nil {
			if err == ErrEncode {
				t.Skip()
			}
			t.Log(err)
		}
		decData := Decode(encData)
		if !bytes.Equal(decData, data) || len(decData) != len(data) {
			t.Errorf("output %v does not match the provided input %v", decData, data)
		}
	})

}
