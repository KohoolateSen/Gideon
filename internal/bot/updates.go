package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (tb *TelegramBot) StartUpdater() {
	updates, _ := tb.GetUpdates(tgbotapi.UpdateConfig{
		Offset:  0,
		Limit:   100,
		Timeout: 0,
	})

	var last_updateid int
	if len(updates) > 0 {
		last_updateid = updates[len(updates)-1].UpdateID
	}

	u := tgbotapi.NewUpdate(last_updateid + 1)
	u.Timeout = 60

	updateChan := tb.GetUpdatesChan(u)

	for update := range updateChan {
		tb.update = &update

		switch {
		case update.Message != nil && update.Message.IsCommand():
			tb.HandleCommands()

		case update.CallbackQuery != nil:
			tb.HandleCallBackQueries()

		case update.InlineQuery != nil:
			tb.HandleInlineQueries()

		}
	}
}
