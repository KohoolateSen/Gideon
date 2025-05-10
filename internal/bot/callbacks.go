package bot

import (
	N "Gideon/internal/bot/features/newsreader"
	utils "Gideon/internal/utils/scraper"
	"strings"
)

func (tb *TelegramBot) HandleCallBackQueries() {
	callback_str := strings.Split(tb.update.CallbackQuery.Data, ":")
	switch callback_str[0] {
	case "news":
		tb.HandleNewsCallBacks(callback_str[1])
	}

}

func (tb *TelegramBot) HandleNewsCallBacks(callback string) {
	articles := tb.newsreader.Response.Articles
	if len(articles) == 0 {
		return
	}

	chat_id := tb.update.CallbackQuery.Message.Chat.ID
	wrapped_msg := tb.newsreader.WrapArticleAsMessage

	switch callback {
	case N.KbBtnPrevCallback:
		tb.newsreader.Cpage = (tb.newsreader.Cpage - 1 + len(articles)) % len(articles)
		tb.SendMessageWithMultiButtonKeyboard(chat_id, wrapped_msg(), N.NewsReadeCommandKeyboard)

	case N.KbBtnReadCallback:
		url := articles[tb.newsreader.Cpage].URL
		content := utils.Scrape("p", url, 4091)
		tb.SendRawMessage(chat_id, content)

	case N.KbBtnNextCallback:
		tb.newsreader.Cpage = (tb.newsreader.Cpage + 1) % len(articles)
		tb.SendMessageWithMultiButtonKeyboard(chat_id, wrapped_msg(), N.NewsReadeCommandKeyboard)
	}
}
