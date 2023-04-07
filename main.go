package main

import (
	"embed"
	"flag"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/json-iterator/go/extra"
	"github.com/sirupsen/logrus"
	embed2 "go-tgbot/comm/embed"
	"go-tgbot/comm/global"
	"go-tgbot/conf"
	"go-tgbot/handle"
	"go-tgbot/ticker"
	"log"
	"time"
)

//go:embed config/prod.yaml
var embedWebFiles embed.FS

func main() {
	extra.RegisterFuzzyDecoders()
	flag.Parse()
	time.LoadLocation("Asia/Shanghai")

	var (
		c   *conf.AllConfig
		err error
	)

	embed2.SetWebFs(embedWebFiles)

	// 加载配置文件
	c, err = conf.ReadConfig()
	if err != nil {
		logrus.Fatal(err)
	}
	if c.App.Debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	bot, err := tgbotapi.NewBotAPI(c.App.Token)
	if err != nil {
		log.Fatal(err.Error())
	}

	if c.App.Debug {
		bot.Debug = true
	}

	global.Conf = c

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	ticker.Ticker(bot)

	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		go handle.Bypass(bot, update)
	}
}
