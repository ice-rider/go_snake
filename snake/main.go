package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"snake/graphics"
)

func main() {
	if os.Getenv("CODESPACES") == "true" {
		// В Codespaces используем стандартный режим без TUI
		log.Println("Запуск в Codespaces, TUI недоступен")
		return
	}

	p := graphics.NewProgram(15, 30)
	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}