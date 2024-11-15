package internal

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "testwrite")
	if err != nil {
		t.Fatalf("Unable to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "testfile.txt")
	testData := "Hello, world!"

	err = WriteToFile(testFile, testData)
	if err != nil {
		t.Errorf("Failed to write to file: %v", err)
	}

	content, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Fatalf("Unable to read back the file: %v", err)
	}

	if string(content) != testData {
		t.Errorf("File content mismatch: got %s, want %s", string(content), testData)
	}

	invalidFile := "/invalid_path/testfile.txt"
	err = WriteToFile(invalidFile, testData)
	if err == nil {
		t.Errorf("Expected an error when writing to an invalid path, but got nil")
	}
}
