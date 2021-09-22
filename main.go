package main

import (
	"ascii-art/converters"
	"ascii-art/errors"
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
)

// TODO refactore check step in ReadBannerFromFile()
// TODO start programm without string line-argument if we have --reverse flag
// TODO (optional) use instead of DiscoverFlagType() and FindFlagValue() - functions from package flag (standard golang package)
// TODO (optional) use default texture pack - standard.txt if user didn't give to us what textures to use

var BANNER_TEMPLATE_PACK string
var FLAG_LIST []string

// Every character will have own id by ascii-table and own form. The last one we save in 2d array
type Banner struct {
	id          int
	asciiSymbol [8][]rune
}

type Flag struct {
	class string
	value string
}

func main() {
	if len(os.Args) != 3 && len(os.Args) != 4 {
		errors.PrintErrorMessage(0)
	}

	FLAG_LIST = []string{"reverse", "color", "output", "align"}
	BANNER_TEMPLATE_PACK = os.Args[2] + ".txt"

	var flagToApplyData Flag

	if len(os.Args) == 4 {
		flagToApplyData = DiscoverFlagType(os.Args[3])
	}

	var bannerTemplateList []Banner = LoadTemplatePack()
	var transformedInput [][]rune = TransformInput(os.Args[1])
	var bannersToPrint [][]Banner = CollectNeededBanners(transformedInput, bannerTemplateList)

	ApplyFlag(flagToApplyData, bannersToPrint, bannerTemplateList)
}

// Loads all symbols from text file with ascii-characters and returns them in array
func LoadTemplatePack() []Banner {
	var result []Banner

	file, err := ioutil.ReadFile(BANNER_TEMPLATE_PACK)
	CheckFile(err)
	text := converters.TranslateToRuneSlice(file)

	var bannerToApply Banner
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

		bannerToApply.id = i
		bannerToApply.asciiSymbol = tempArr
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
func CollectNeededBanners(charList [][]rune, bannerList []Banner) [][]Banner {
	var result [][]Banner

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

// Prints all ascii-characters by our 2d banner array what we have built. Nil array means new-line
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
		errors.PrintErrorMessage(1)
		os.Exit(1)
	}
}

func DiscoverFlagType(flag string) Flag {
	var result Flag
	var flagInRune []rune = []rune(flag)
	var flagFound bool = false

	if flagInRune[0] == 45 && flagInRune[1] == 45 {
		for i := 0; i < len(FLAG_LIST); i++ {
			wantedFlag := []rune(FLAG_LIST[i])
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
				result.class = FLAG_LIST[i]
				FindFlagValue(&result, flagInRune)
				break
			}
		}
	}

	if !flagFound {
		errors.PrintErrorMessage(0)
		return result
	} else {
		return result
	}
}

func FindFlagValue(flag *Flag, values []rune) {
	var valueResult string

	if 3+len(flag.class) == len(values) || 3+len(flag.class) > len(values) {
		errors.PrintErrorMessage(4)
	}

	if values[2+len(flag.class)] == 61 {
		for i := 3 + len(flag.class); i < len(values); i++ {
			valueResult = valueResult + string(values[i])
		}
	} else {
		errors.PrintErrorMessage(4)
	}

	flag.value = valueResult
}

func ApplyFlag(flagToApplyData Flag, userBanners [][]Banner, bannerTemplateList []Banner) {
	if flagToApplyData.class == "output" {
		SaveBannerInToFile(flagToApplyData.value, userBanners)
	} else if flagToApplyData.class == "reverse" {
		ReadBannerFromFile(flagToApplyData.value, bannerTemplateList)
	} else if flagToApplyData.class == "align" {
		UseAlignByMode(flagToApplyData.value, userBanners)
	} else if flagToApplyData.class == "color" {
		return
	} else {
		PrintBanners(userBanners)
	}

	os.Exit(0)
}

func SaveBannerInToFile(fileName string, bannersToSave [][]Banner) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	CheckFile(err)

	dataWriter := bufio.NewWriter(file)
	for i := 0; i < len(bannersToSave); i++ {
		if bannersToSave[i] == nil {
			_, _ = dataWriter.WriteString("\n")
			continue
		}

		for k := 0; k < 8; k++ {
			for d := 0; d < len(bannersToSave[i]); d++ {
				_, _ = dataWriter.WriteString(string(bannersToSave[i][d].asciiSymbol[k]))
			}
			_, _ = dataWriter.WriteString("\n")
		}
	}

	_, _ = dataWriter.WriteString("\n")
	dataWriter.Flush()
	file.Close()
}

func ReadBannerFromFile(fileName string, banners []Banner) {
	file, err := ioutil.ReadFile(fileName)
	CheckFile(err)
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
		for k := 0; k < len(banners[i].asciiSymbol[0]); k++ {
			if text[k+startIndex] == banners[i].asciiSymbol[0][k] && text[k+startIndex+charCounterInTheRow] == banners[i].asciiSymbol[1][k] &&
				text[k+startIndex+charCounterInTheRow*2] == banners[i].asciiSymbol[2][k] && text[k+startIndex+charCounterInTheRow*3] == banners[i].asciiSymbol[3][k] &&
				text[k+startIndex+charCounterInTheRow*4] == banners[i].asciiSymbol[4][k] && text[k+startIndex+charCounterInTheRow*5] == banners[i].asciiSymbol[5][k] &&
				text[k+startIndex+charCounterInTheRow*6] == banners[i].asciiSymbol[6][k] && text[k+startIndex+charCounterInTheRow*7] == banners[i].asciiSymbol[7][k] {
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
			resultString = resultString + string(banners[i].id)
			i = 0
		}
	}

	fmt.Println(resultString)
}

func GetTerminalWindowsSize() int64 {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	out = RemoveSliceElement(out, len(out)-1)
	if terminalLength, err := strconv.ParseInt(string(out[3:]), 10, 32); err == nil {
		return terminalLength
	} else {
		return -1
	}
}

func RemoveSliceElement(slice []byte, s int) []byte {
	return append(slice[:s], slice[s+1:]...)
}

func UseAlignByMode(mode string, userBanners [][]Banner) {
	var alignSymbol [8][]rune
	var alignBanner Banner
	var resultBanners [][]Banner
	var totalBannersLenght int
	var alignBannersAmount int64
	var lastSpaceIndex int

	for i := 0; i < len(alignSymbol); i++ {
		alignSymbol[i] = append(alignSymbol[i], 32)
	}

	alignBanner.asciiSymbol = alignSymbol

	for i := 0; i < len(userBanners); i++ {
		totalBannersLenght = 0

		for a := 0; a < len(userBanners[i]); a++ {
			totalBannersLenght = totalBannersLenght + len(userBanners[i][a].asciiSymbol[0])
		}

		if mode == "center" {
			alignBannersAmount = (GetTerminalWindowsSize() - int64(totalBannersLenght)) / 2
		} else if mode == "right" {
			alignBannersAmount = GetTerminalWindowsSize() - int64(totalBannersLenght)
		} else if mode == "justify" {
			var tempVar []Banner
			tempVar = userBanners[i]
			var spacePos []int

			for n := 0; n < len(userBanners[i]); n++ {
				if userBanners[i][n].id == 32 {
					spacePos = append(spacePos, n)
				}
			}

			if len(spacePos) == 0 {
				break
			}

			tempVar2 := ((GetTerminalWindowsSize() - int64(totalBannersLenght)) / int64(len(spacePos)))

			for n := 0; int64(n) < tempVar2; n++ {
				for a := 0; a < len(userBanners[i]); a++ {
					if userBanners[i][a].id == 32 {
						tempVar = insert(tempVar, a, alignBanner)
						userBanners[i] = tempVar
						a++

						if n == int(tempVar2)-1 {
							lastSpaceIndex = a
						}
					}
				}
			}

			terminalWindowsSizeRemainder := ((GetTerminalWindowsSize() - int64(totalBannersLenght)) % int64(len(spacePos)))
			for a := 0; int64(a) < terminalWindowsSizeRemainder; a++ {
				tempVar = insert(tempVar, lastSpaceIndex, alignBanner)
				userBanners[i] = tempVar
			}
		} else if mode == "left" {
			alignBannersAmount = 0
		} else {
			errors.PrintErrorMessage(5)
		}

		resultBanners = append(resultBanners, nil)

		for k := 0; int64(k) < alignBannersAmount; k++ {
			resultBanners[i] = append(resultBanners[i], alignBanner)
		}

		userBanners[i] = append(resultBanners[i], userBanners[i]...)
	}

	PrintBanners(userBanners)
}

func insert(a []Banner, index int, value Banner) []Banner {
	if len(a) == index {
		return append(a, value)
	}

	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}
