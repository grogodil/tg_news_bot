package tgbot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// Инициализируем телеграм-бота
func InitBot(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic("Error ")
	}

	bot.Debug = true

	// Настройка обновлений
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	// Получение канала обновлений
	updates := bot.GetUpdatesChan(updateConfig)
	handlersBot(bot, updates)
}
