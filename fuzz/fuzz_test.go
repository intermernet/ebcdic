//
//Tests for fuzzing package ebcdic
//

// +build gofuzzbeta
package ebcdic_fuzz

import (
	"bytes"
	"testing"

	"github.com/intermernet/ebcdic"
)

func FuzzRoundTrip(f *testing.F) {
	f.Add([]byte{0xE3, 0x85, 0xA2, 0xA3, 0x89, 0x95, 0x87, 0x6B, 0x40, 0xF1, 0xF2, 0xF3})
	f.Fuzz(func(t *testing.T, data []byte) {
		if len(data) == 0 {
			t.Skip()
		}
		t.Parallel() // seed corpus tests can run in parallel
		encData, err := ebcdic.Encode(data)
		if err != nil {
			if err == ebcdic.ErrEncode {
				t.Skip()
			}
			t.Log(err)
		}
		decData := ebcdic.Decode(encData)
		if !bytes.Equal(decData, data) || len(decData) != len(data) {
			t.Errorf("output %v does not match the provided input %v", decData, data)
		}
	})

}
