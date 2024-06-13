package file

import (
	"bufio"
	"os"

	buffer "github.com/migueel15/TextEditor/Buffer"
)

type File struct {
	path   string
	Buffer *buffer.TextBuffer
}

func NewFile(path string) *File {
	return &File{path, buffer.NewTextBuffer()}
}

func NewFromFile(path string) (*File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	TextBuffer := buffer.NewTextBuffer()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		TextBuffer.Append(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &File{path, TextBuffer}, nil
}

func (f *File) Save() error {
	file, err := os.Create(f.path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range f.Buffer.GetLines() {
		writer.WriteString(line + "\n")
	}
	writer.Flush()
	return nil
}

// func (f *File) ReadBuffer() error {
// }
