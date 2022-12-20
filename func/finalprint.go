package asciiartfs

import (
	"fmt"
	"strings"
)

func FinalPrint(symbs []string, bargs string, text []string) {
	if len(text) == 1 && text[0] == "" {
		return
	}
	OnlyNewLines := CheckNewLines(text)
	for i, word := range text {
		if OnlyNewLines && i == len(text)-1 {
			break
		}
		if word == "" {
			fmt.Println()
		} else {
			for lines := 0; lines < 8; lines++ {
				for _, rune := range word {
					fmt.Print(strings.Split(symbs[rune-32], "\n")[lines])
				}
				fmt.Println()
			}
		}
	}
}
