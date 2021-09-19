package errors

import (
	"fmt"
	"os"
)

func PrintErrorMessage(number int) {
	if number == 0 {
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\n\nEX: go run . something standard --color=<color>")
	} else if number == 1 {
		fmt.Println("Error. Missing file.")
	} else if number == 3 {
		fmt.Println("Option not found. Availabe options are: --color; --align; --output; --reverse.")
	} else if number == 4 {
		fmt.Println("Option values not found. Option value example is: --color=blue; --align=right; --output=filename.txt; --reverse=filename.txt.")
	}

	os.Exit(1)
}
