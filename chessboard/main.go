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
	fmt.Printf("width: %d\n", width)

	var length int
	fmt.Print("Введите высоту: ")
	_, err = fmt.Scanf("%d", &length)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	fmt.Printf("Printing chessboard for	 %dx%d\n", width, length)

	board, err := chessboard.Get(width, length)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(board)
}
