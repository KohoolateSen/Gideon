package bot

import (
	N "Gideon/internal/bot/features/newsreader"
	"fmt"
)

func (tb *TelegramBot) HandleCommands() {
	args := tb.update.Message.CommandArguments()

	switch tb.update.Message.Command() {
	case "start":
		tb.StartCommand()

	case "help":
		tb.HelpCommand()

	case "news":
		tb.NewsCommand(args)
	}
}

func (tb *TelegramBot) StartCommand() {
	fmt.Println("Dummy Print")
}

func (tb *TelegramBot) HelpCommand() {
	fmt.Println("Dummy Print")
}

func (tb *TelegramBot) NewsCommand(subject string) {
	chat_id := tb.update.Message.Chat.ID

	tb.newsreader.Cpage = 0
	tb.newsreader.FetchNews(subject)

	if len(tb.newsreader.Response.Articles) > 0 {
		msg := tb.newsreader.WrapArticleAsMessage()
		tb.SendMessageWithMultiButtonKeyboard(chat_id, msg, N.NewsReadeCommandKeyboard)
	}
}
