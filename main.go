package main

import (
	"log"
	"os"
	"tg_news_bot/internal/tgbot"
	"tg_news_bot/internal/db"
	
	"github.com/joho/godotenv"
)

func main() {
	// Загружаем переменные окружения
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found")
	}

	defer db.DB.Close()

	// Получаем путь к БД из переменной окружения
	dbFile := os.Getenv("TG_DBFILE")

	// Если путь к файлу БД отсутствует - присваеваем переменной нужное нам значение
	if dbFile == "" {
		dbFile = "tgStorage.db"
	}

	// Инициализируем БД
	var err error
	if err = db.Init(dbFile); err != nil {
		log.Fatalf("Ошибка инициализации БД: %v", err)
	}

	// Присваеваем token соответствующее значение
	// переменной окружения
	token := os.Getenv("TELEGRAM_BOT_TOKEN")

	// Инициализируем бота
	tgbot.InitBot(token)
}
