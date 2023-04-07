package handle

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go-tgbot/ecode"
)

func HandleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message.IsCommand() == false {
		return
	}

	text := update.Message.Text
	chatID := update.Message.Chat.ID
	replyMsg := tgbotapi.NewMessage(chatID, text)
	switch update.Message.Command() {
	case "help", "start":
		replyMsg.Text = ecode.HelpMessage
	default:
		replyMsg.Text = "No such command!"
	}
	_, _ = bot.Send(replyMsg)
}
