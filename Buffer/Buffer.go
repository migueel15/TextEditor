package buffer

import "errors"

type TextBuffer struct {
	lines []string
}

func (b *TextBuffer) Insert(line int, col int, text string) error {
	if line >= len(b.lines) || line < 0 {
		return errors.New("Line out of bounds")
	}
	if col < 0 || col > len(b.lines[line]) {
		return errors.New("Column out of bounds")
	}
	b.lines[line] = b.lines[line][:col] + text + b.lines[line][col:]
	return nil
}

func (b *TextBuffer) Delete(line int, col int, length int) error {
	if line >= len(b.lines) || line < 0 {
		return errors.New("Line out of bounds")
	}
	if col < 0 || col+length > len(b.lines[line]) {
		return errors.New("Column out of bounds")
	}

	b.lines[line] = b.lines[line][:col] + b.lines[line][col+length:]
	return nil
}
