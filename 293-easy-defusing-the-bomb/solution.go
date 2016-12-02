package main

// If you cut a white cable you can't cut white or black cable.
// If you cut a red cable you have to cut a green one
// If you cut a black cable it is not allowed to cut a white, green or orange one
// If you cut a orange cable you should cut a red or black one
// If you cut a green one you have to cut a orange or white one
// If you cut a purple cable you can't cut a purple, green, orange or white cable

// white black purple red green orange

import (
	"fmt"
	"strings"
)

func main() {
	//input := "WRGW" // input 1 defused
	input := "WOGW" // input 2 boom
	valid := map[byte]string{
		'W': "PRGO",
		'R': "G",
		'B': "BPR",
		'O': "RB",
		'G': "OW",
		'P': "BR",
	}

	exploded := false
	for i := 0; i < len(input)-1; i++ {
		if !strings.ContainsRune(valid[input[i]], rune(input[i+1])) {
			exploded = true
			fmt.Println("Bomb exploded")
			break
		}
	}
	if !exploded {
		fmt.Println("Bomb was defused")
	}
}
