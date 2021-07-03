package ebcdic_test

import (
	"fmt"

	"github.com/intermernet/ebcdic"
)

func ExampleEncode() {
	input := "An alleged character set used on IBM dinosaurs"
	output, err := ebcdic.Encode(input)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%0X\n", output)
	// Output: C195408193938587858440838881998183A3859940A285A340A4A2858440969540C9C2D44084899596A281A499A2
}

func ExampleDecode() {
	input := []byte{0xc1, 0x95, 0x40, 0x81, 0x93, 0x93, 0x85, 0x87, 0x85, 0x84,
		0x40, 0x83, 0x88, 0x81, 0x99, 0x81, 0x83, 0xa3, 0x85, 0x99, 0x40, 0xa2,
		0x85, 0xa3, 0x40, 0xa4, 0xa2, 0x85, 0x84, 0x40, 0x96, 0x95, 0x40, 0xc9,
		0xc2, 0xd4, 0x40, 0x84, 0x89, 0x95, 0x96, 0xa2, 0x81, 0xa4, 0x99, 0xa2,
	}
	fmt.Println(string(ebcdic.Decode(input)))
	// Output: An alleged character set used on IBM dinosaurs
}
