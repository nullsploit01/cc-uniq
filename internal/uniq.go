package internal

import (
	"bufio"
	"os"

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

func (u *Uniq) PrintUniqueLinesFromFile(file *os.File) error {
	if err := u.ProcessFile(file); err != nil {
		return err
	}

	for _, k := range u.AdjacentUniqueLines {
		u.cmd.OutOrStdout().Write([]byte(k.Line + "\n"))
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
