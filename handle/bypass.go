package handle

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go-tgbot/comm/msg"
)

func Bypass(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message.IsCommand() {
		HandleCommand(bot, update)
	}

	// 非 command
	msg.HandleMsg(bot, update)
}
