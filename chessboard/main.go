package main

import (
	"fmt"

	"github.com/lightalloy/go-exercises/chessboard/chessboard"
)

func main() {
	board, err := chessboard.Get(100, 101)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(board)
}
