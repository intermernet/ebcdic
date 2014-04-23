/*
Package ebcdic provides functions to convert Unicode to EBCDIC and vice-versa.

It uses conversion table data from http://www.unicode.org/Public/MAPPINGS/VENDORS/MICSFT/EBCDIC/CP037.TXT

Valid codepage bytes for both Unicode and EBCDIC are `0x00` .. `0xFF`

Invalid characters are replaced with `NUL` (`0x00`)

Copyright Mike Hughes 2014 (intermernet AT gmail DOT com)
*/
package ebcdic

const charMapLength = 0xFF

// Encode Unicode to EBCDIC.
// Replaces invalid characters with NUL.
func Encode(Unicode []byte) []byte {
	runes := []rune(string(Unicode)) // Convert bytes back to runes
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
func Decode(EBCDIC []byte) []byte {
	var out []byte
	for _, v := range EBCDIC {
		out = append(out, byte(unicodeMap[v]))
	}
	return out
}
