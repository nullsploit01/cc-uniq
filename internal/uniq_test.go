package internal

import (
	"bytes"
	"os"
	"testing"

	"github.com/spf13/cobra"
)

func TestProcessFile(t *testing.T) {
	content := "hello\nworld\nhello\nhello\nworld"
	file := createMockFile(content)
	defer file.Close()

	uniq := NewUniq(nil) // nil as no Cobra command is needed for testing
	err := uniq.ProcessFile(file)
	if err != nil {
		t.Errorf("ProcessFile returned an error: %v", err)
	}

	if len(uniq.AdjacentUniqueLines) != 4 {
		println(uniq.AdjacentUniqueLines)
		t.Errorf("Expected 3 unique lines, got %d", len(uniq.AdjacentUniqueLines))
	}
}

func TestPrintUniqueLinesFromFile(t *testing.T) {
	content := "hello\nworld\nhello\nhello\nworld\n"
	file := createMockFile(content)
	defer file.Close()

	uniq := NewUniq(nil)

	// Redirect output to capture it for testing
	var output bytes.Buffer
	uniq.cmd = &cobra.Command{}
	uniq.cmd.SetOut(&output)

	err := uniq.PrintUniqueLinesFromFile(file, "", false, false, true)
	if err != nil {
		t.Errorf("PrintUniqueLinesFromFile returned an error: %v", err)
	}

	expectedOutput := "hello\nworld\nworld\n"
	if output.String() != expectedOutput {
		t.Errorf("Expected output %q, got %q", expectedOutput, output.String())
	}
}

func createMockFile(content string) *os.File {
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		panic(err)
	}
	tmpfile.WriteString(content)
	tmpfile.Seek(0, os.SEEK_SET)
	return tmpfile
}
