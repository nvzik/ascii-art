package asciiart

import (
	"fmt"
	"os"
	"strings"
)

const hash string = "81802daa5d0000076c70ba728c2deb73"

func MainWork(args []string) {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Wrong number of arguments")
		return
	}
	arg := os.Args[1]

	banner, err := os.ReadFile("standard.txt")
	Check(err)
	if CheckHash(string(banner)) != hash {
		fmt.Println("ERROR: Wrong Hash")
		return
	}
	symbols := strings.Split(strings.ReplaceAll(string(banner), "\r", ""), "\n\n")
	bargs := strings.ReplaceAll(arg, "\n", "\\n")
	text := strings.Split(bargs, "\\n")
	FinalPrint(symbols, bargs, text)
}
