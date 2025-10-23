// Package chessboard provides a functions to generate a chessboard.
package chessboard

import "fmt"

// Get returns a chessboard of the given width and length.
func Get(width int, length int) (string, error) {
	if width < 0 || width > 1000 {
		return "", fmt.Errorf("width must be > 0 and <= 1000")
	}
	if length < 0 || length > 1000 {
		return "", fmt.Errorf("length must be > 0 and <= 1000")
	}

	board := ""
	sym := ""

	for i := 0; i < length; i++ {
		if i%2 == 0 {
			sym = "#"
		} else {
			sym = " "
		}
		for j := 0; j < width; j++ {
			board = board + sym
			if sym == "#" {
				sym = " "
			} else {
				sym = "#"
			}
		}
		if i != length-1 {
			board = board + "\n"
		}
	}
	return board, nil
}
