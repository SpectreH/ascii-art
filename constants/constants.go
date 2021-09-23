package constants

import (
	"ascii-art/structures"
)

var DEFAULT_PACK string
var FLAG_LIST []string
var COLOR_LIST []structures.Color

func InitConstans() {
	DEFAULT_PACK = "standard.txt"
	FLAG_LIST = []string{"reverse", "color", "output", "align"}
	COLOR_LIST = LoadColors()
}

func LoadColors() []structures.Color {
	var result []structures.Color
	var colorToAppend structures.Color
	colorNames := []string{"white", "red", "yellow", "orange", "blue", "green", "purple", "brown"}
	colorCodes := [][]byte{{27, 91, 48, 109},
		{27, 91, 51, 56, 59, 53, 59, 49, 109},
		{27, 91, 51, 56, 59, 53, 59, 49, 49, 109},
		{27, 91, 51, 56, 59, 53, 59, 50, 48, 56, 109},
		{27, 91, 51, 56, 59, 53, 59, 52, 109},
		{27, 91, 51, 56, 59, 53, 59, 56, 50, 109},
		{27, 91, 51, 56, 59, 53, 59, 49, 50, 57, 109},
		{27, 91, 51, 56, 59, 53, 59, 57, 52, 109}}

	for i := 0; i < len(colorNames); i++ {
		colorToAppend.Name = colorNames[i]
		colorToAppend.Code = colorCodes[i]
		result = append(result, colorToAppend)
	}

	return result
}
