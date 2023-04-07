package handle

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message.IsCommand() == false {
		return
	}

	text := update.Message.Text
	chatID := update.Message.Chat.ID
	replyMsg := tgbotapi.NewMessage(chatID, text)
	switch update.Message.Command() {
	case "help":
		replyMsg.Text = `
天气查询，如：泾县天气。
菜谱查询，如: 红烧肉菜谱，红烧肉做法。
输入【程序员鼓励师】收到程序员鼓励师的回复。
输入【事件提醒】获取设置事件提醒的格式。
输入【毒鸡汤】关键字回复毒鸡汤。
输入【英语一句话】关键字回复一句学习英语。
`
	default:
		replyMsg.Text = "No such command!"
	}
	_, _ = bot.Send(replyMsg)
}
