//
// Package ebcdic provides functions to convert Unicode to EBCDIC and vice-versa.
//
// It uses conversion table data from http://www.unicode.org/Public/MAPPINGS/VENDORS/MICSFT/EBCDIC/CP037.TXT
//
// Valid codepage bytes for both Unicode and EBCDIC are `0x00` .. `0xFF`
//
// Invalid characters are replaced with `NUL` (`0x00`)
//
// Copyright Mike Hughes 2014 - 2021 (intermernet AT gmail DOT com)
//
package ebcdic

import (
	"errors"
)

const charMapLength = 0xFF

var ErrEncode = errors.New("could not encode characters")

// Encode Unicode to EBCDIC.
// Replaces Unicode runes > codepoint FF with NUL.
func Encode(Unicode string) ([]byte, error) {
	var (
		out []byte
		err error
	)
	for _, v := range Unicode {
		if v <= charMapLength { // Unicode <= FF, in valid translation range
			out = append(out, ebcdicMap[v])
		} else {
			out = append(out, 0) // Replace with NUL if out of range
			err = ErrEncode
		}
		//fmt.Printf("char: %X\terr: %v\n", v, err)
	}
	return out, err
}

// Decode EBCDIC to Unicode.
func Decode(EBCDIC []byte) string {
	var out []rune
	for _, v := range EBCDIC {
		out = append(out, unicodeMap[v])
	}
	return string(out)
}
