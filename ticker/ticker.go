package ticker

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func Ticker(bot *tgbotapi.BotAPI) {
	go MasterTicker(bot)
	go ScheduleNoticeTicker(bot)
	go KeepAlive(bot)
}
