package ticker

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go-tgbot/comm/global"
	"time"
)

func KeepAlive(bot *tgbotapi.BotAPI) {
	var (
		err     error
		message string
	)

	for {
		select {
		case <-time.After(60 * time.Minute):
			message = fmt.Sprintf("keep alive: %s", time.Now().Format("2006-01-02 15:04:05"))
			_, err = bot.Send(tgbotapi.NewMessage(global.Conf.App.KeepAliveChatID, message))
			if err != nil {
				err = errors.Wrapf(err, "KeepAlive err")
				logrus.Error(err.Error())
			}
		}
	}
}
