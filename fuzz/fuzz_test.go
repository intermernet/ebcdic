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
	// f.Add([]byte{0xE3, 0x85, 0xA2, 0xA3, 0x89, 0x95, 0x87, 0x6B, 0x40, 0xF1, 0xF2, 0xF3})
	// f.Add([]byte("1日2本3語"))
	f.Fuzz(func(t *testing.T, data []byte) {
		if len(data) == 0 {
			t.Skip("Skipping zero length input")
		}
		t.Parallel() // seed corpus tests can run in parallel
		encData, err := ebcdic.Encode(string(data))
		if err != nil {
			if err == ebcdic.ErrEncode {
				t.Skipf("Skipping round trip test for %#v: %v\n", data, err)
			}
			t.Log(err)
		}
		decData := ebcdic.Decode(encData)
		if !bytes.Equal([]byte(decData), data) || len(decData) != len(data) {
			t.Errorf("output %#v does not match the provided input %#v", decData, data)
		}
	})

}
