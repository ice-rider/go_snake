package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"snake/core"
)

func main() {
	var width, height int = 15, 30
	game := core.NewGame(width, height)
	game.Start()

	fmt.Println("Игра Змейка")
	fmt.Println("Используйте W, A, S, D для управления")
	fmt.Println("Нажмите Enter после каждой команды")
	fmt.Println()
	fmt.Println(game.Field)

	// Запускаем горутину для обработки ввода
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := strings.ToLower(strings.TrimSpace(scanner.Text()))
			switch input {
			case "w":
				game.RotateSnake(core.Up)
			case "s":
				game.RotateSnake(core.Down)
			case "a":
				game.RotateSnake(core.Left)
			case "d":
				game.RotateSnake(core.Right)
			}
		}
	}()

	for {
		game.Next()
		time.Sleep(200 * time.Millisecond) // Увеличим задержку для лучшего контроля
		fmt.Print("\033[H\033[2J") // Очистка консоли
		fmt.Println("Игра Змейка")
		fmt.Println("Используйте W, A, S, D для управления")
		fmt.Println("Нажмите Enter после каждой команды")
		fmt.Println()
		fmt.Println(game.Field)
	}
}