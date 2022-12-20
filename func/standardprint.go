package asciiartfs

import (
	"fmt"
	"os"
	"strings"
)

func StandardPrint(s string) {
	arg := s
	banner, err := os.ReadFile("standard.txt")
	Check(err)
	if CheckHash(string(banner)) != b1 {
		fmt.Println("ERROR: Wrong Hash")
		return
	}
	symbols := strings.Split(strings.ReplaceAll(string(banner), "\r", ""), "\n\n")
	bargs := strings.ReplaceAll(arg, "\n", "\\n")
	text := strings.Split(bargs, "\\n")
	FinalPrint(symbols, bargs, text)
}
