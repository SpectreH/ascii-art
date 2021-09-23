package flags

import (
	"ascii-art/checks"
	"ascii-art/converters"
	"ascii-art/structures"
	"fmt"
	"io/ioutil"
)

// TODO refactore check step on line 34

func ReadBannerFromFile(fileName string, banners []structures.Banner) {
	file, err := ioutil.ReadFile(fileName)
	checks.CheckFile(err)
	text := converters.TranslateToRuneSlice(file)

	var charCounterInTheRow int = 1
	for i := 0; i < len(text); i++ {
		if text[i] != 10 {
			charCounterInTheRow++
		} else {
			break
		}
	}

	var symbolFound bool = false
	var resultString string
	var saveIndex int = 0
	var startIndex int = 0

	for i := 0; i < len(banners); i++ {
		symbolFound = false
		for k := 0; k < len(banners[i].AsciiSymbol[0]); k++ {
			if text[k+startIndex] == banners[i].AsciiSymbol[0][k] && text[k+startIndex+charCounterInTheRow] == banners[i].AsciiSymbol[1][k] &&
				text[k+startIndex+charCounterInTheRow*2] == banners[i].AsciiSymbol[2][k] && text[k+startIndex+charCounterInTheRow*3] == banners[i].AsciiSymbol[3][k] &&
				text[k+startIndex+charCounterInTheRow*4] == banners[i].AsciiSymbol[4][k] && text[k+startIndex+charCounterInTheRow*5] == banners[i].AsciiSymbol[5][k] &&
				text[k+startIndex+charCounterInTheRow*6] == banners[i].AsciiSymbol[6][k] && text[k+startIndex+charCounterInTheRow*7] == banners[i].AsciiSymbol[7][k] {
				symbolFound = true
				saveIndex = k + 1
			} else {
				saveIndex = startIndex
				symbolFound = false
				break
			}
		}

		if symbolFound {
			startIndex = startIndex + saveIndex
			resultString = resultString + string(banners[i].Id)
			i = -1
			continue
		}
	}

	fmt.Println(resultString)
}
