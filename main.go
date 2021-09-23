package main

import (
	"ascii-art/checks"
	"ascii-art/constants"
	"ascii-art/converters"
	"ascii-art/errors"
	"ascii-art/flags"
	"ascii-art/structures"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	constants.InitConstans()

	var flagToApplyData structures.Flag
	if len(os.Args) == 2 {
		if flags.DiscoverFlagType(os.Args[1]).Class == "reverse" {
			flagToApplyData = flags.DiscoverFlagType(os.Args[1])
		}
	} else if len(os.Args) == 4 || len(os.Args) == 5 {
		flagToApplyData = flags.DiscoverFlagType(os.Args[3])
	} else if len(os.Args) == 1 {
		errors.PrintErrorMessage(0)
	}

	var bannerTemplateList []structures.Banner = LoadTemplatePack()
	var transformedInput [][]rune = TransformInput(os.Args[1])
	var bannersToPrint [][]structures.Banner = CollectNeededBanners(transformedInput, bannerTemplateList)

	ApplyFlag(flagToApplyData, bannersToPrint, bannerTemplateList)
}

// Loads all symbols from text file with ascii-characters and returns them in array
func LoadTemplatePack() []structures.Banner {
	var result []structures.Banner
	var file []byte
	var text []rune

	if len(os.Args) == 2 {
		file, _ = ioutil.ReadFile(constants.DEFAULT_PACK)
	} else {
		_, err := ioutil.ReadFile(os.Args[2] + ".txt")
		checks.CheckFile(err)
		file, _ = ioutil.ReadFile(os.Args[2] + ".txt")
	}

	text = converters.TranslateToRuneSlice(file)

	var bannerToApply structures.Banner
	var textIndex int

	if text[0] == 10 {
		textIndex = 1
	} else {
		textIndex = 2
	}

	for i := 32; i < 127; i++ {
		var tempArr [8][]rune

		for k := 0; k < 8; k++ {
			for l := 0; l < 32; l++ {

				if text[textIndex] == 13 {
					if text[textIndex+1] == 10 {
						textIndex = textIndex + 2
						break
					}
				} else if text[textIndex] == 10 {
					textIndex++
					break
				}

				tempArr[k] = append(tempArr[k], text[textIndex])
				textIndex++
			}
		}

		bannerToApply.Id = i
		bannerToApply.AsciiSymbol = tempArr
		bannerToApply.Color = constants.COLOR_LIST[0]
		result = append(result, bannerToApply)

		if i != 126 && text[textIndex] == 13 {
			textIndex = textIndex + 2
		} else {
			textIndex++
		}
	}

	return result
}

// Transform string text to 2d rune array. We separate chars by rows (If we have "/n" it means that we add new row to array)
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

// With transformed string - here we try to find needed ascii-symbol and save it in to 2d array. Here we also seperate ascii-symbol by rows like in previous function
func CollectNeededBanners(charList [][]rune, bannerList []structures.Banner) [][]structures.Banner {
	var result [][]structures.Banner

	for i := 0; i < len(charList); i++ {
		if result == nil {
			result = append(result, nil)
		}

		for k := 0; k < len(charList[i]); k++ {
			for m := 0; m < len(bannerList); m++ {
				if charList[i][k] == rune(bannerList[m].Id) {
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

// Prints all ascii-characters by our 2d banner array what we have built. Nil array means new-line
func PrintBanners(banners [][]structures.Banner) {
	for i := 0; i < len(banners); i++ {
		if banners[i] == nil {
			fmt.Println()
			continue
		}
		for k := 0; k < 8; k++ {
			for d := 0; d < len(banners[i]); d++ {
				fmt.Print(string(banners[i][d].Color.Code))
				fmt.Print(string(banners[i][d].AsciiSymbol[k]))
			}
			fmt.Println()
		}
	}
}

// Applies flag if it was given
func ApplyFlag(flagToApplyData structures.Flag, userBanners [][]structures.Banner, bannerTemplateList []structures.Banner) {
	if flagToApplyData.Class == "output" {
		flags.SaveBannerInToFile(flagToApplyData.Value, userBanners)
	} else if flagToApplyData.Class == "reverse" {
		flags.ReadBannerFromFile(flagToApplyData.Value, bannerTemplateList)
	} else if flagToApplyData.Class == "align" {
		flags.UseAlignByMode(flagToApplyData.Value, userBanners)
		PrintBanners(userBanners)
	} else if flagToApplyData.Class == "color" {
		bannersToColor := flags.FindColorOptions()
		flags.ApplyColorToBanners(bannersToColor, &userBanners, flagToApplyData)
		PrintBanners(userBanners)
	} else {
		PrintBanners(userBanners)
	}

	os.Exit(0)
}
