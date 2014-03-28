[![GoDoc](https://godoc.org/github.com/Intermernet/ebcdic?status.png)](https://godoc.org/github.com/Intermernet/ebcdic) [![Build Status](https://drone.io/github.com/Intermernet/ebcdic/status.png)](https://drone.io/github.com/Intermernet/ebcdic/latest)

EBCDIC / Unicode transcoding package for Go (Code page 37 only, `00..FF`).

For example will convert from the EBCDIC bytes:

    c1 95 40 81 93 93 85 87 85 84 40 83 88 81 99 81 83 a3 85 99 40 a2 85
    a3 40 a4 a2 85 84 40 96 95 40 c9 c2 d4 40 84 89 95 96 a2 81 a4 99 a2

To the Unicode string:

    An alleged character set used on IBM dinosaurs

and vice versa.

See [godoc.org](https://godoc.org/github.com/Intermernet/ebcdic) for usage.

Conversion table data from http://www.unicode.org/Public/MAPPINGS/VENDORS/MICSFT/EBCDIC/CP037.TXT

Example conversion text from http://zvon.org/comp/r/ref-Jargon_file.html#Terms~EBCDIC