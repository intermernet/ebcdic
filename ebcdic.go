/*
ebcdic provides functions to convert Unicode to EBCDIC and vice-versa.

Copyright Mike Hughes 2014 (intermernet AT gmail DOT com)
*/
package ebcdic

import "unicode"

// Encode Unicode to EBCDIC
// Replaces invalid characters with NUL
func Encode(in []byte) []byte {
	runes := []rune(string(in)) // Convert bytes back to runes
	var out []byte
	for _, v := range runes {
		if v <= unicode.MaxLatin1 {
			out = append(out, byte(ebcdicMap[v]))
		} else {
			out = append(out, 0) // NUL if out of range
		}
	}
	return out
}

// Decode EBCDIC to Unicode
// Replaces invalid characters with NUL
func Decode(in []byte) []byte {
	var out []byte
	for _, v := range in {
		if v <= unicode.MaxLatin1 {
			out = append(out, byte(unicodeMap[v]))
		} else {
			out = append(out, 0) // NUL if out of range
		}
	}
	return out
}
