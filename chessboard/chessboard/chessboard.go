// Package chessboard provides a functions to generate a chessboard.
package chessboard

import "fmt"

// Get returns a chessboard of the given width and length.
func Get(width int, length int) (string, error) {
	if width > 1000 || length > 1000 {
		return "", fmt.Errorf("width and length must be less than or equal to 1000")
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
