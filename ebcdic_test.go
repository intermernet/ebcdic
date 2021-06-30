//
// Tests for package ebcdic
//
package ebcdic

import (
	"bytes"
	"testing"
)

var (
	// Valid string, can translate both ways
	unicodeString = "Testing, 123"
	ebcdicBytes   = []byte{0xE3, 0x85, 0xA2, 0xA3, 0x89, 0x95, 0x87, 0x6B, 0x40, 0xF1, 0xF2, 0xF3}

	// Invalid string, contains Unicode code-points higher than 255. Invalid characters replaced with NUL
	failString = "Testing, 1日2本3語"
	failBytes  = []byte{0xE3, 0x85, 0xA2, 0xA3, 0x89, 0x95, 0x87, 0x6B, 0x40, 0xF1, 0x00, 0xF2, 0x00, 0xF3, 0x00}
	failResult = "Testing, 1\x002\x003\x00"
)

// Test Encoding of valid string
func TestEncode(t *testing.T) {
	out, err := Encode([]byte(unicodeString))
	if err != nil {
		t.Error(err)
	}
	if bytes.Compare(out, ebcdicBytes) != 0 {
		t.Error("encountered Encoding error.")
	}
}

// Test Decoding of valid string
func TestDecode(t *testing.T) {
	if string(Decode(ebcdicBytes)) != unicodeString {
		t.Error("encountered Decoding error.")
	}
}

// Test Encoding of invalid string
func TestEncodeFail(t *testing.T) {
	out, err := Encode([]byte(failString))
	if err != nil && err != ErrEncode {
		t.Error("encountered Encoding Failure error.")
	}
	if bytes.Compare(out, failBytes) != 0 {
		t.Error("encountered Encoding Failure error.")
	}
}

// Test Decoding of invalid string
func TestDecodeFail(t *testing.T) {
	if string(Decode(failBytes)) != failResult {
		t.Error("encountered Decoding Failure error.")
	}
}

// Test Encoding of entire Unicode <= 0xFF character map
func TestEncodeCharMap(t *testing.T) {
	out, err := Encode([]byte(string(unicodeMap)))
	if err != nil {
		t.Error(err)
	}
	if bytes.Compare(out, ordered()) != 0 {
		t.Error("encountered Encoding CharMap error.")
	}
}

// Test Decoding of entire EBCDIC character map
func TestDecodeCharMap(t *testing.T) {
	if bytes.Compare(Decode(ebcdicMap), ordered()) != 0 {
		t.Error("encountered Decoding CharMap error.")
	}
}

// Output an ordered byte-slice, 0x00..0x100
func ordered() []byte {
	out := make([]byte, charMapLength+1)
	for i := 0; i <= charMapLength; i++ {
		out[i] = byte(i)
	}
	return out
}
