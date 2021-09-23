package flags

import (
	"ascii-art/constants"
	"ascii-art/errors"
	"ascii-art/structures"
	"os"
	"strconv"
)

func ApplyColorToBanners(colorIndexes []int, userBanners *[][]structures.Banner, flagToApplyData structures.Flag) {
	var amountToColor int

	if len(colorIndexes) == 0 {
		amountToColor = len((*userBanners)[0]) - 1
		colorIndexes = append(colorIndexes, 0)
		colorIndexes = append(colorIndexes, len((*userBanners)[0])-1)
	} else if len(colorIndexes) == 2 && colorIndexes[1] == 0 {
		amountToColor = 1
		colorIndexes[1] = colorIndexes[0]
	} else if (colorIndexes[1] - colorIndexes[0]) < 0 {
		errors.PrintErrorMessage(6)
	} else {
		amountToColor = (colorIndexes[1] - colorIndexes[0])
	}

	if colorIndexes[1] >= len((*userBanners)[0]) {
		errors.PrintErrorMessage(6)
	}

	for i := 0; i < amountToColor; i++ {
		for k := colorIndexes[0]; k < colorIndexes[1]+1; k++ {
			for m := 0; m < len(constants.COLOR_LIST); m++ {
				if constants.COLOR_LIST[m].Name == flagToApplyData.Value {
					(*userBanners)[0][k].Color = constants.COLOR_LIST[m]
					break
				}

				if m == len(constants.COLOR_LIST)-1 {
					errors.PrintErrorMessage(7)
				}
			}
		}
	}
}

func FindColorOptions() []int {
	var indexesToColorInByte [2][]byte
	var indexesToColorInInt []int
	var currentNumber int = 0

	if len(os.Args) == 5 {
		if len(os.Args[4]) != 2 || len(os.Args[4]) != 1 {
			if os.Args[4][0] == 91 && os.Args[4][len(os.Args[4])-1] == 93 {
				for i := 1; i < len(os.Args[4])-1; i++ {
					if os.Args[4][i] == 45 {
						currentNumber++
						continue
					}
					indexesToColorInByte[currentNumber] = append(indexesToColorInByte[currentNumber], (os.Args[4][i]))
				}
			} else {
				return indexesToColorInInt
			}
		}
	} else {
		return indexesToColorInInt
	}

	for i := 0; i < len(indexesToColorInByte); i++ {
		aByteToInt, _ := strconv.Atoi(string(indexesToColorInByte[i]))
		indexesToColorInInt = append(indexesToColorInInt, aByteToInt)
	}

	return indexesToColorInInt
}
