package main

import (
	"embed"
	"flag"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/json-iterator/go/extra"
	"github.com/sirupsen/logrus"
	embed2 "go-tgbot/comm/embed"
	"go-tgbot/comm/global"
	"go-tgbot/conf"
	"go-tgbot/handle"
	"go-tgbot/ticker"
	"log"
	"os"
	"runtime"
	"time"
)

//go:embed config
var embedConfigFiles embed.FS

//go:embed static
var embedStaticsFiles embed.FS

var (
	gitCommit string
	buildAt   string
)

func main() {
	extra.RegisterFuzzyDecoders()

	flag.Parse()
	if len(os.Args) >= 2 && os.Args[1] == "version" {
		print(fmt.Sprintf("Git commit:	%s\nGo version:   %s\nBuilt:	%s\nOS/Arch:	%s/%s",
			gitCommit, runtime.Version(), buildAt, runtime.GOOS, runtime.GOARCH))
		return
	}

	time.LoadLocation("Asia/Shanghai")

	var (
		c   *conf.AllConfig
		err error
	)

	embed2.SetConfigFs(embedConfigFiles)
	embed2.SetStaticFsFs(embedStaticsFiles)

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
