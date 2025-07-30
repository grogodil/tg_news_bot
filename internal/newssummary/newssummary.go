package newssummary

import (
	"fmt"
	"tg_news_bot/internal/db"

	"github.com/microcosm-cc/bluemonday"
	"github.com/mmcdole/gofeed"
)

func GetOneHabrGoSummary() (string, error) {
	rssURL := "https://habr.com/ru/rss/hub/go/articles/all/?fl=ru"
	// Инициализирцем новый парсер RSS-ленты
	fp := gofeed.NewParser()
	// Парсим RSS-ленту
	feed, err := fp.ParseURL(rssURL)

	if err != nil {
		return "", err
	}

	var article *gofeed.Item
	var summary string
	var i int64
	// Создаем объект политика отчистки режима StrictPolicy (для устранения тегов)
	policy := bluemonday.StrictPolicy()
	// будующий чистый текст без тегов
	var clean string

	// Перебераем найденные статьи и присваеваем значение summary
	for {
		article = feed.Items[i]
		summary = fmt.Sprintf("Заголовок: %s\nСсылка: %s\nОписание: \n%s", article.Title, article.Link, article.Description)

		// Если эту новость ранее не публиковали - то функция возвразает
		// summary и добавляет в ДБ ссылку на эту статью
		if !db.IsSent(article.Link) {
			db.MarkSent(article.Link)
			clean = policy.Sanitize(summary)
			return clean, nil
		}

		i++
	}
}
