package flags

import (
	"ascii-art/constants"
	"ascii-art/errors"
	"ascii-art/structures"
	"os"
)

func DiscoverFlagType(flag string) structures.Flag {
	var result structures.Flag
	var flagInRune []rune = []rune(flag)
	var flagFound bool = false

	if flagInRune[0] == 45 && flagInRune[1] == 45 {
		for i := 0; i < len(constants.FLAG_LIST); i++ {
			wantedFlag := []rune(constants.FLAG_LIST[i])
			for k := 0; k < len(wantedFlag); k++ {
				if flagInRune[k+2] == wantedFlag[k] {
					flagFound = true
					continue
				} else {
					flagFound = false
					break
				}
			}

			if flagFound {
				result.Class = constants.FLAG_LIST[i]
				FindFlagValue(&result, flagInRune)
				break
			}
		}
	}

	if !flagFound && (len(os.Args) == 1 || len(os.Args) == 4) {
		errors.PrintErrorMessage(3)
		return result
	} else {
		return result
	}
}

func FindFlagValue(flag *structures.Flag, values []rune) {
	var valueResult string

	if 3+len(flag.Class) == len(values) || 3+len(flag.Class) > len(values) {
		errors.PrintErrorMessage(4)
	}

	if values[2+len(flag.Class)] == 61 {
		for i := 3 + len(flag.Class); i < len(values); i++ {
			valueResult = valueResult + string(values[i])
		}
	} else {
		errors.PrintErrorMessage(4)
	}

	flag.Value = valueResult
}
