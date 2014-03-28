/*
Package ebcdic provides functions to convert Unicode to EBCDIC and vice-versa.

It uses conversion table data from http://www.unicode.org/Public/MAPPINGS/VENDORS/MICSFT/EBCDIC/CP037.TXT

Copyright Mike Hughes 2014 (intermernet AT gmail DOT com)
*/
package ebcdic

const charMapLength = 0xFF

// Encode Unicode to EBCDIC.
// Replaces invalid characters with NUL.
func Encode(in []byte) []byte {
	runes := []rune(string(in)) // Convert bytes back to runes
	var out []byte
	for _, v := range runes {
		if v <= charMapLength {
			out = append(out, byte(ebcdicMap[v]))
		} else {
			out = append(out, 0) // NUL if out of range
		}
	}
	return out
}

// Decode EBCDIC to Unicode.
// Replaces invalid characters with NUL.
func Decode(in []byte) []byte {
	var out []byte
	for _, v := range in {
		if v <= charMapLength {
			out = append(out, byte(unicodeMap[v]))
		} else {
			out = append(out, 0) // NUL if out of range
		}
	}
	return out
}
