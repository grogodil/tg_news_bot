package db

import (
	"log"
)

// IsSent проверяет, была ли уже отправлена статья с данным URL
func IsSent(url string) bool {
	var exists bool
	// Запрос, проверяет, есть ли хоть одно повторяющаяся запись в url_news,
	// возыращая bool значение
	query := `SELECT EXISTS(SELECT 1 FROM sent_articles WHERE url_news = ? LIMIT 1);`
	err := DB.QueryRow(query, url).Scan(&exists)

	if err != nil {
		log.Println("Ошибка поиска уникального URL")
		exists = false
	}

	return exists
}

// MarkSent отмечает статью как отправленную, записывая её URL
func MarkSent(url string) error {
	// Запрос, проверяет добавляет значение в url_news или
	// игнорирует добавление записей с дубликатами уникальных значений
	query := `INSERT OR IGNORE INTO sent_articles (url_news) VALUES (?);`
	_, err := DB.Exec(query, url)
	return err
}
