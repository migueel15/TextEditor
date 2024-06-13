package buffer

import (
	"testing"
)

func TestInsertIntoBuffer(t *testing.T) {
	cases := []struct {
		before      []string
		position    [2]int
		text        string
		expect      []string
		expectError bool
	}{
		{
			before: []string{
				"Esta es la primera linea",
				"Esta es la segunda linea",
			},
			text:     "inserto esto",
			position: [2]int{0, 7},
			expect: []string{
				"Esta esinserto esto la primera linea",
				"Esta es la segunda linea",
			},
			expectError: false,
		},
		{
			before: []string{
				"Esta es la primera linea",
				"Esta es la segunda linea",
			},
			text:     "nuevo texto",
			position: [2]int{1, 20},
			expect: []string{
				"Esta esinsertp esto la primera linea",
				"Esta es la segunda lnuevo textoinea",
			},
			expectError: false,
		},
		{
			before: []string{
				"Esta es la primera linea",
				"Esta es la segunda linea",
			},
			text:     "nuevo",
			position: [2]int{0, 25},
			expect: []string{
				"Esta es la primera linea",
				"Esta es la segunda linea",
			},
			expectError: true,
		},
		{
			before: []string{
				"Esta es la primera linea",
				"Esta es la segunda linea",
			},
			text:     "nuevo",
			position: [2]int{2, 10},
			expect: []string{
				"Esta es la primera linea",
				"Esta es la segunda linea",
			},
			expectError: true,
		},
	}

	for idx, c := range cases {
		b := TextBuffer{c.before}
		err := b.Insert(c.position[0], c.position[1], c.text)
		if err != nil && !c.expectError {
			t.Errorf("Error inserting(%d): Return was: %v ; Wanted an error: %v", idx, err, c.expectError)

		} else if err != nil {
			continue
		} else if b.lines[c.position[0]] != c.expect[c.position[0]] {
			t.Errorf("Error inserting(%d):\nWas:%v\nWants:%v", idx, b.lines, c.expect)
		}
	}
}

func TestDeleteFromBuffer(t *testing.T) {
	cases := []struct {
		before      []string
		position    [2]int
		length      int
		expect      []string
		expectError bool
	}{
		{
			before: []string{
				"Esta es la primera linea",
				"Esta es la segunda linea",
			},
			position: [2]int{0, 7},
			length:   3,
			expect: []string{
				"Esta es primera linea",
				"Esta es la segunda linea",
			},
			expectError: false,
		},
		{
			before: []string{
				"Esta es la primera linea",
				"Esta es la segunda linea",
			},
			position: [2]int{1, 19},
			length:   5,
			expect: []string{
				"Esta es la primera linea",
				"Esta es la segunda ",
			},
			expectError: false,
		},
		{
			before: []string{
				"Esta es la primera linea",
				"Esta es la segunda linea",
			},
			position: [2]int{0, 25},
			length:   2,
			expect: []string{
				"Esta es la primera linea",
				"Esta es la segunda linea",
			},
			expectError: true,
		},
		{
			before: []string{
				"Esta es la primera linea",
				"Esta es la segunda linea",
			},
			position: [2]int{2, 10},
			length:   3,
			expect: []string{
				"Esta es la primera linea",
				"Esta es la segunda linea",
			},
			expectError: true,
		},
	}

	for idx, c := range cases {
		b := TextBuffer{c.before}
		err := b.Delete(c.position[0], c.position[1], c.length)
		if err != nil && !c.expectError {
			t.Errorf("Error deleting(%d): Return was: %v ; Wanted an error: %v", idx, err, c.expectError)

		} else if err != nil {
			continue
		} else if b.lines[c.position[0]] != c.expect[c.position[0]] {
			t.Errorf("Error deleting(%d):\nWas:%v ; Wanted:%v", idx, b.lines, c.expect)
		}
	}
}
