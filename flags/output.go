package flags

import (
	"ascii-art/structures"
	"bufio"
	"os"
)

func SaveBannerInToFile(fileName string, bannersToSave [][]structures.Banner) {
	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	dataWriter := bufio.NewWriter(file)
	for i := 0; i < len(bannersToSave); i++ {
		if bannersToSave[i] == nil {
			_, _ = dataWriter.WriteString("\n")
			continue
		}

		for k := 0; k < 8; k++ {
			for d := 0; d < len(bannersToSave[i]); d++ {
				_, _ = dataWriter.WriteString(string(bannersToSave[i][d].AsciiSymbol[k]))
			}
			_, _ = dataWriter.WriteString("\n")
		}
	}

	_, _ = dataWriter.WriteString("\n")
	dataWriter.Flush()
	file.Close()
}
