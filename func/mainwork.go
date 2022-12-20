package asciiartfs

import "fmt"

func MainWork(args []string) {
	if len(args) != 2 && len(args) != 3 {
		fmt.Println("Error: wrong number of arguments")
		return
	} else if len(args) == 3 {
		if args[2] == "shadow" {
			BannerPrint(args[1], "shadow.txt")
		} else if args[2] == "thinkertoy" {
			BannerPrint(args[1], "thinkertoy.txt")
		} else if args[2] == "standard" {
			BannerPrint(args[1], "standard.txt")
		} else {
			fmt.Println("Error: wrong name of banner")
		}
	} else if len(args) == 2 {
		StandardPrint(args[1])
	}
}
