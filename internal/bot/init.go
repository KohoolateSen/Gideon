package bot

import (
	"Gideon/internal/bot/features/newsreader"
	c "Gideon/internal/config"
	"Gideon/internal/utils/httpproxy"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	*tgbotapi.BotAPI
	update     *tgbotapi.Update
	newsreader *newsreader.NewsReader
}

func InitBot() (TelegramBot, error) {
	bot_token := c.AppConfig.TelegramBot.APIToken
	api_endpoint := c.AppConfig.TelegramBot.APIEndPoint

	tgbot, err := tgbotapi.NewBotAPIWithClient(
		bot_token,
		api_endpoint,
		httpproxy.GlobalHTTPClient,
	)

	if err != nil {
		return TelegramBot{}, err
	}

	// Other features can be initiated here
	news_reader := newsreader.InitNewsReader()

	return TelegramBot{BotAPI: tgbot, newsreader: news_reader}, nil
}
