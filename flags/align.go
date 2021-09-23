package flags

import (
	"ascii-art/errors"
	"ascii-art/slices"
	"ascii-art/structures"
	"log"
	"os"
	"os/exec"
	"strconv"
)

func UseAlignByMode(mode string, userBanners [][]structures.Banner) {
	var alignSymbol [8][]rune
	var alignBanner structures.Banner
	var resultBanners [][]structures.Banner
	var totalBannersLenght int
	var alignBannersAmount int64
	var lastSpaceIndex int

	for i := 0; i < len(alignSymbol); i++ {
		alignSymbol[i] = append(alignSymbol[i], 32)
	}

	alignBanner.AsciiSymbol = alignSymbol

	for i := 0; i < len(userBanners); i++ {
		totalBannersLenght = 0

		for a := 0; a < len(userBanners[i]); a++ {
			totalBannersLenght = totalBannersLenght + len(userBanners[i][a].AsciiSymbol[0])
		}

		if mode == "center" {
			alignBannersAmount = (GetTerminalWindowsSize() - int64(totalBannersLenght)) / 2
		} else if mode == "right" {
			alignBannersAmount = GetTerminalWindowsSize() - int64(totalBannersLenght)
		} else if mode == "justify" {
			var tempVar []structures.Banner
			tempVar = userBanners[i]
			var spacePos []int

			for n := 0; n < len(userBanners[i]); n++ {
				if userBanners[i][n].Id == 32 {
					spacePos = append(spacePos, n)
				}
			}

			if len(spacePos) == 0 {
				break
			}

			tempVar2 := ((GetTerminalWindowsSize() - int64(totalBannersLenght)) / int64(len(spacePos)))

			for n := 0; int64(n) < tempVar2; n++ {
				for a := 0; a < len(userBanners[i]); a++ {
					if userBanners[i][a].Id == 32 {
						tempVar = slices.Insert(tempVar, a, alignBanner)
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
				tempVar = slices.Insert(tempVar, lastSpaceIndex, alignBanner)
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
}

func GetTerminalWindowsSize() int64 {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	out = slices.RemoveSliceElement(out, len(out)-1)
	if terminalLength, err := strconv.ParseInt(string(out[3:]), 10, 32); err == nil {
		return terminalLength
	} else {
		return -1
	}
}
