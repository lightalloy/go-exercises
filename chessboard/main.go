package main

import (
	"fmt"

	"github.com/lightalloy/go-exercises/chessboard/chessboard"
)

func main() {
	var width int
	var length int
	var err error

	width, err = getValueFromUser("width")
	if err != nil {
		fmt.Println(err)
		return
	}

	length, err = getValueFromUser("length")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Chessboard for %dx%d:\n\n", width, length)

	board, err := chessboard.Get(width, length)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(board)
}

func getValueFromUser(dimension string) (int, error) {
	var value int

	fmt.Printf("%s: ", dimension)
	_, err := fmt.Scanf("%d", &value)
	if err != nil {
		return 0, err
	}

	if value < 0 || value > 1000 {
		return 0, fmt.Errorf("%s must be > 0 and <= 1000", dimension)
	}
	return value, nil
}
