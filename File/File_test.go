package file

import (
	"os"
	"testing"
)

func TestNewFile(t *testing.T) {
	path := "testfile.txt"
	defer os.Remove(path) // Clean up test file after test

	file := NewFile(path)
	if file == nil {
		t.Errorf("NewFile returned nil")
	}

	// Check if TextBuffer is initialized
	if file.Buffer == nil {
		t.Errorf("TextBuffer not initialized")
	}

	// Check if path matches
	if file.path != path {
		t.Errorf("Path mismatch, expected %s but got %s", path, file.path)
	}
}

func TestOpenFile(t *testing.T) {
	// Prepare a test file
	testContent := "Line 1\nLine 2\nLine 3"
	path := "testfile.txt"
	defer os.Remove(path) // Clean up test file after test
	err := os.WriteFile(path, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to prepare test file: %v", err)
	}

	file, err := NewFromFile(path)
	if err != nil {
		t.Fatalf("OpenFile returned error: %v", err)
	}

	// Check if path matches
	if file.path != path {
		t.Errorf("Path mismatch, expected %s but got %s", path, file.path)
	}

	// Check if lines are correctly read into TextBuffer
	expectedLines := []string{"Line 1", "Line 2", "Line 3"}
	lines := file.Buffer.GetLines()
	if len(lines) != len(expectedLines) {
		t.Errorf("Number of lines mismatch, expected %d but got %d", len(expectedLines), len(lines))
	}

	// Check each line
	for i := 0; i < len(expectedLines); i++ {
		if lines[i] != expectedLines[i] {
			t.Errorf("Line mismatch at index %d, expected '%s' but got '%s'", i, expectedLines[i], lines[i])
		}
	}
}

func TestFile_Save(t *testing.T) {
	// Prepare a test file
	path := "testfile.txt"
	defer os.Remove(path) // Clean up test file after test

	file := NewFile(path)
	file.Buffer.Append("Line 1")
	file.Buffer.Append("Line 2")
	file.Buffer.Append("Line 3")

	err := file.Save()
	if err != nil {
		t.Fatalf("Save returned error: %v", err)
	}

	// Read saved file to verify content
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read saved file: %v", err)
	}

	// Check content
	expectedContent := "Line 1\nLine 2\nLine 3\n"
	if string(data) != expectedContent {
		t.Errorf("Saved file content mismatch, expected '%s' but got '%s'", expectedContent, string(data))
	}
}
