package main

import (
	"fmt"

	"github.com/lightalloy/go-exercises/chessboard/chessboard"
)

func main() {
	var width int
	fmt.Print("Введите ширину: ")
	_, err := fmt.Scanf("%d", &width)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	if width < 0 || width > 1000 {
		fmt.Println("width must be > 0 and <= 1000")
		return
	}

	var length int
	fmt.Print("Введите высоту: ")
	_, err = fmt.Scanf("%d", &length)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	if length < 0 || length > 1000 {
		fmt.Println("length must be > 0 and <= 1000")
		return
	}

	fmt.Printf("Chessboard for %dx%d\n", width, length)

	board, err := chessboard.Get(width, length)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(board)
}
