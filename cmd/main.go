package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"ollama-bot/internal/telegram_bot/tg_app"
)

func main() {
	err := godotenv.Load("configs/.env")
	if err != nil {
		fmt.Println(err)
		godotenv.Load("../configs/.env")
	}
	tg_app.Run()
}
