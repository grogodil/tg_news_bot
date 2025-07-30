package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

// Создаем констатнту tgStorAge, которая определяет создаие таблицы
// для храннения уникальных URL новостей
const tgStorage = `
	CREATE TABLE IF NOT EXISTS sent_articles (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	url_news TEXT NOT NULL UNIQUE
	);`

var DB *sql.DB

func Init(dbFile string) error {
	var err error

	// Открываем БД
	DB, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("ошибка открытия БД: %w", err)
	}

	// Проверяем существование БД файла
	// Если БД не существует - создаем ее
	_, err = os.Stat(dbFile)
	if os.IsNotExist(err) {
		if _, err = DB.Exec(tgStorage); err != nil {
			return fmt.Errorf("ошибка создания БД: %w", err)
		}
	}

	return nil
}
