package main

import (
	asciiart "ascii-art/func"
	"fmt"
	"os"
	"strings"
)

const hash string = "81802daa5d0000076c70ba728c2deb73"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: Wrong number of arguments")
		return
	}
	arg := os.Args[1]

	banner, err := os.ReadFile("standard.txt")
	asciiart.Check(err)
	if asciiart.CheckHash(string(banner)) != hash {
		fmt.Println("ERROR: Wrong Hash")
		return
	}
	symbols := strings.Split(strings.ReplaceAll(string(banner), "\r", ""), "\n\n")
	bargs := strings.ReplaceAll(arg, "\n", "\\n")
	text := strings.Split(bargs, "\\n")
	asciiart.FinalPrint(symbols, bargs, text)
}
