package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var bannerTemplatePack string = "standard.txt"

type Banner struct {
	id          int
	asciiSymbol [8][]rune
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error! Must be (or at least) only one command-line argument.")
		os.Exit(0)
	}

	var bannerTemplateList []Banner = LoadTemplatePack()
	var transformedInput [][]rune = TransformInput(os.Args[1])

	var bannersToPrint [][]Banner = CollectNeededBanners(transformedInput, bannerTemplateList)

	PrintBanners(bannersToPrint)
}

func LoadTemplatePack() []Banner {
	var result []Banner

	file, err := ioutil.ReadFile(bannerTemplatePack)
	CheckFile(err)
	text := TranslateToRuneSlice(file)

	var bannerToApply Banner
	var textIndex int = 1

	for i := 32; i < 127; i++ {
		var tempArr [8][]rune

		for k := 0; k < 8; k++ {
			for l := 0; l < 32; l++ {
				if text[textIndex] == 10 {
					textIndex++
					break
				}

				tempArr[k] = append(tempArr[k], text[textIndex])
				textIndex++
			}
		}

		bannerToApply.id = i
		bannerToApply.asciiSymbol = tempArr
		result = append(result, bannerToApply)

		textIndex = textIndex + 1
	}

	return result
}

func TransformInput(text string) [][]rune {
	var result [][]rune

	textInRune := []rune(text)
	currentLine := 0

	for i := 0; i < len(textInRune); i++ {
		if textInRune[i] == 92 && i+1 < len(textInRune) {
			if textInRune[i+1] == 110 {
				result = append(result, nil)
				currentLine++
				i++
				continue
			}
		}

		if result == nil {
			result = append(result, nil)
		} else if result[0] == nil && len(result) == currentLine {
			result = append(result, nil)
		}

		result[currentLine] = append(result[currentLine], textInRune[i])
	}

	return result
}

func CollectNeededBanners(charList [][]rune, bannerList []Banner) [][]Banner {
	var result [][]Banner

	var newLineBanner Banner
	newLineBanner.id = 10

	for i := 0; i < len(charList); i++ {
		if result == nil {
			result = append(result, nil)
		}

		for k := 0; k < len(charList[i]); k++ {
			for m := 0; m < len(bannerList); m++ {
				if charList[i][k] == rune(bannerList[m].id) {
					result[i] = append(result[i], bannerList[m])
				} else {
					continue
				}
			}
		}

		if len(charList) != i+1 {
			result = append(result, nil)
		}
	}

	return result
}

func PrintBanners(banners [][]Banner) {
	for i := 0; i < len(banners); i++ {
		if banners[i] == nil {
			fmt.Println()
			continue
		}

		for k := 0; k < 8; k++ {
			for d := 0; d < len(banners[i]); d++ {
				fmt.Print(string(banners[i][d].asciiSymbol[k]))
			}
			fmt.Println()
		}
	}
}

func CheckFile(e error) {
	if e != nil {
		fmt.Println("Error. Missing file")
		os.Exit(1)
	}
}

func TranslateToRuneSlice(bytes []byte) []rune {
	var text []rune

	for i := range bytes {
		text = append(text, rune(bytes[i]))
	}

	return text
}
