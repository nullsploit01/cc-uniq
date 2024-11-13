package internal

import (
	"bufio"
	"os"
)

type AdjacentUniqueLine struct {
	Line  string
	Count int
}

type Uniq struct {
	AdjacentUniqueLines []AdjacentUniqueLine
}

func NewUniq() *Uniq {
	return &Uniq{}
}

func (u *Uniq) PrintUniqueLinesFromFile(file *os.File) error {
	if err := u.ProcessFile(file); err != nil {
		return err
	}

	for _, k := range u.AdjacentUniqueLines {
		if k.Count > 1 {
			continue
		}

		println(k.Line)
	}

	return nil
}

func (u *Uniq) ProcessFile(file *os.File) error {
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

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
