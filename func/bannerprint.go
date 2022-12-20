package asciiartfs

import (
	"fmt"
	"os"
	"strings"
)

const (
	b1 string = "81802daa5d0000076c70ba728c2deb73" // banner 1 "standard.txt"
	b2 string = "d44671e556d138171774efbababfc135" // banner 2 "shadow.txt"
	b3 string = "0021f26ad06f2f73a0cfa7b7d38d1434" // banner 3 "thinkertoy.txt"
)

func BannerPrint(s string, bannerName string) {
	arg := s
	banner, err := os.ReadFile(bannerName)
	Check(err)
	if bannerName == "standard" {
		if CheckHash(string(banner)) != b1 {
			fmt.Println("ERROR: Wrong Hash")
			return
		}
	} else if bannerName == "shadow" {
		if CheckHash(string(banner)) != b2 {
			fmt.Println("ERROR: Wrong Hash")
			return
		}
	} else if bannerName == "thinkertoy" {
		if CheckHash(string(banner)) != b3 {
			fmt.Println("ERROR: Wrong Hash")
			return
		}
	}
	symbols := strings.Split(strings.ReplaceAll(string(banner), "\r", ""), "\n\n")
	bargs := strings.ReplaceAll(arg, "\n", "\\n")
	text := strings.Split(bargs, "\\n")
	FinalPrint(symbols, bargs, text)
}
