package tgbot

import (
	"log"
	"tg_news_bot/internal/newssummary"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func handlersBot(bot *tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
	// Обработка входящих сообщений
	for update := range updates {
		if update.Message != nil && update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				startHandler(bot, update)
			case "help":
				helpHandler(bot, update)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда.")
				bot.Send(msg)
			}
		}
	}
}

func startHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я новостной бот. Используйте /help для справки."+
		" C этой минуты каждый час тебе будут проходить новость с Хабр")
	bot.Send(msg)

	summary, err := newssummary.GetOneHabrGoSummary()
	if err != nil {
		log.Println("Ошибка парсинга RSS-ленты")
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Упс! Что-то пошло не так...")
		bot.Send(msg)
	}

	// Отправка новости сразу после старта бота
	msg = tgbotapi.NewMessage(update.Message.Chat.ID, summary)
	bot.Send(msg)

	go func() {
		// Инициализируем тикер, который срабатыет 1 раз в час
		ticker := time.NewTicker(1 * time.Hour) // 1 раз в час
		defer ticker.Stop()

		for range ticker.C {
			summary, err := newssummary.GetOneHabrGoSummary()
			if err != nil {
				log.Println("Ошибка парсинга RSS-ленты")
				continue
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, summary)
			bot.Send(msg)
		}
	}()
}

func helpHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Доступные команды:\n/start — запуск\n/help — помощь")
	bot.Send(msg)
}
