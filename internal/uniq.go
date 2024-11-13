package internal

import (
	"bufio"
	"os"
)

type Uniq struct {
	AdjacentUniqueLines map[string]int
}

func NewUniq() *Uniq {
	return &Uniq{
		AdjacentUniqueLines: make(map[string]int),
	}
}

func (u *Uniq) PrintUniqueLinesFromFile(file *os.File) error {
	if err := u.ProcessFile(file); err != nil {
		return err
	}

	for k := range u.AdjacentUniqueLines {
		println(k)
	}

	return nil
}

func (u *Uniq) ProcessFile(file *os.File) error {
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}

	var lastLine string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		curr := scanner.Text()

		if curr == lastLine {
			u.AdjacentUniqueLines[curr]++
		} else {
			u.AdjacentUniqueLines[curr] = 1
			lastLine = curr
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
