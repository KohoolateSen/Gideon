package bot

import (
	N "Gideon/internal/bot/features/newsreader"
	"encoding/json"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type InlineResult = tgbotapi.InlineQueryResultArticle

func (tb *TelegramBot) HandleInlineQueries() {
	query := tb.update.InlineQuery

	if args, err := tb.InlineQueryArgValidator(query.Query); err == nil {
		category, sub_category := args[0], args[1]
		results := tb.GenerateQueriesInline(category, sub_category)
		json_results, _ := json.Marshal(results)

		params := tgbotapi.Params{
			"inline_query_id": tb.update.InlineQuery.ID,
			"results":         string(json_results),
		}

		resp, err := tb.MakeRequest("answerInlineQuery", params)
		fmt.Println(resp, "+", err)
	}
}

func (tb *TelegramBot) GenerateQueriesInline(category string, sub_category string) []interface{} {
	var results []interface{}
	category_lower := strings.ToLower(category)

	switch {
	case strings.Contains(category_lower, "news"):
		newsResults := tb.FeedNewsTitlesInline(sub_category)
		for _, res := range newsResults {
			results = append(results, res)
		}
	default:
		results = append(results, tgbotapi.NewInlineQueryResultArticle(
			category+"-1",
			"Echo: "+category,
			"You typed: "+category,
		))
	}

	return results
}

func (tb *TelegramBot) FeedNewsTitlesInline(subject string) []InlineResult {
	var results []InlineResult

	err := tb.newsreader.FetchNews(subject)
	if err != nil {
		logrus.Warnf("failed to fetch news ➜ %v", err)
	}

	for i, article := range tb.newsreader.Response.Articles {
		if i >= 50 {
			break
		}

		formatted_message := tb.newsreader.WrapArticleAsMessageByIndex(i)
		keyboard := tb.CustomSingleButtonKeyboard(N.KbBtnGoToWebPage, article.URL)

		result := tgbotapi.NewInlineQueryResultArticle(
			fmt.Sprintf("news-%d", i),
			article.Title,
			formatted_message,
		)

		result.ThumbURL = article.URLToImage
		result.ReplyMarkup = keyboard
		results = append(results, result)
	}

	return results
}
