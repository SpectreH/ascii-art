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
		fmt.Println("Usage: go run . [STRING] [BANNER] [OPTION]\n\nEX: go run . something standard --color=<color>")
	} else if number == 3 {
		fmt.Println("Option not found. Availabe options are: --color; --align; --output; --reverse.")
	} else if number == 4 {
		fmt.Println("Option values not found. Option value example is: --color=blue; --align=right; --output=filename.txt; --reverse=filename.txt.")
	} else if number == 5 {
		fmt.Println("Align type not found. Availabe types are: left; right, center, justify")
	} else if number == 6 {
		fmt.Println("Invalid color optional values. Color optional values example is: [3-4]")
	} else if number == 7 {
		fmt.Println("Invalid color. Available colors are: red; yellow; orange; blue; green; purple; brown")
	}

	os.Exit(1)
}
