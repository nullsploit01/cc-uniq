package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type AdjacentUniqueLine struct {
	Line  string
	Count int
}

type Uniq struct {
	cmd                 *cobra.Command
	AdjacentUniqueLines []AdjacentUniqueLine
}

func NewUniq(cmd *cobra.Command) *Uniq {
	return &Uniq{cmd: cmd}
}

func (u *Uniq) PrintUniqueLinesFromFile(file *os.File, outputFileName string, withCount bool) error {
	if err := u.ProcessFile(file); err != nil {
		return err
	}

	var outputData []string

	for i, k := range u.AdjacentUniqueLines {
		dataToAppend := k.Line

		if withCount {
			dataToAppend = fmt.Sprint(u.AdjacentUniqueLines[i].Count) + " " + dataToAppend
		}

		outputData = append(outputData, dataToAppend)
	}

	if outputFileName != "" {
		data := strings.Join(outputData, "\n")
		WriteToFile(outputFileName, data)
		return nil
	}

	for _, data := range outputData {
		u.cmd.OutOrStdout().Write([]byte(data + "\n"))
	}

	return nil
}

func (u *Uniq) ProcessFile(file *os.File) error {
	var lastLine string
	var lastAdjacentLine *AdjacentUniqueLine

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		curr := scanner.Text()
		if lastAdjacentLine != nil && curr == lastLine {
			lastAdjacentLine.Count++
		} else {
			adjacentLine := AdjacentUniqueLine{Line: curr, Count: 1}
			u.AdjacentUniqueLines = append(u.AdjacentUniqueLines, adjacentLine)

			lastLine = curr
			lastAdjacentLine = &u.AdjacentUniqueLines[len(u.AdjacentUniqueLines)-1]
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
