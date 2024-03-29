package ticker

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"go-tgbot/comm/embed"
	conf2 "go-tgbot/conf"
	"log"
	"testing"
	"time"

	"github.com/json-iterator/go/extra"
	"github.com/sirupsen/logrus"
	"go-tgbot/comm/global"
)

func initAction(t *testing.T) {
	extra.RegisterFuzzyDecoders()
	logrus.SetLevel(logrus.DebugLevel)
	var (
		err error
	)

	conf, err := conf2.ReadConfig("../config/prod.yaml")
	if err != nil {
		t.Logf("get conf err:%s ", err.Error())
		return
	}

	global.Conf = conf
}

func TestSendLoveMessage(t *testing.T) {
	//var done = make(chan struct{})
	//go LoveTicker()
	//
	//fmt.Println("done...")
	//<-done
}

func TestNoticeMessage(t *testing.T) {
	count, interval, startTimestamp, message, err := parseNoticeMessage("+st20221227 15:35,消息内容")
	fmt.Println(count, interval, startTimestamp, message, err)
	fmt.Println("-------------------------------------------------------")
	count, interval, startTimestamp, message, err = parseNoticeMessage("+st20221227 15:35,消息内容,3,15")
	fmt.Println(count, interval, startTimestamp, message, err)
	fmt.Println("-------------------------------------------------------")
	count, interval, startTimestamp, message, err = parseNoticeMessage("+s15:35,消息内容")
	fmt.Println(count, interval, startTimestamp, message, err)
	fmt.Println("-------------------------------------------------------")
	count, interval, startTimestamp, message, err = parseNoticeMessage("+s15:35,消息内容,1")
	fmt.Println(count, interval, startTimestamp, message, err)
	fmt.Println("-------------------------------------------------------")
	count, interval, startTimestamp, message, err = parseNoticeMessage("+s15:35,消息内容,2,60")
	fmt.Println(count, interval, startTimestamp, message, err)
	fmt.Println("-------------------------------------------------------")
	count, interval, startTimestamp, message, err = parseNoticeMessage("+s16:11,记得喝水,2,60")
	fmt.Println(count, interval, startTimestamp, message, err)
	fmt.Println("-------------------------------------------------------")
}

func TestParseTime(t *testing.T) {
	count, interval, startTimestamp, message, err := parseNoticeMessage("+st20221227 15:35,消息内容,3,15")
	fmt.Println(count, interval, startTimestamp, message, err)
	if err != nil {
		fmt.Println("失败11111", err.Error())
		return
	}
	printMember(formatMember(count, interval, startTimestamp, message, 1111))

	fmt.Println("---------------------------------------------------------------------")
	count, interval, startTimestamp, message, err = parseNoticeMessage("+s15:32,消息内容2222,3,15")
	fmt.Println(count, interval, startTimestamp, message, err)
	if err != nil {
		fmt.Println("失败222222", err.Error())
		return
	}
	printMember(formatMember(count, interval, startTimestamp, message, 1111))

	fmt.Println("---------------------------------------------------------------------")
	count, interval, startTimestamp, message, err = parseNoticeMessage("+s15:32,消息内容3333333,1,45")
	fmt.Println(count, interval, startTimestamp, message, err)
	if err != nil {
		fmt.Println("失败33333", err.Error())
		return
	}
	printMember(formatMember(count, interval, startTimestamp, message, 1111))
}

func printMember(members []*redis.Z) {
	for _, each := range members {
		fmt.Println(each.Member, each.Score)
	}
}

func TestZSetRedis(t *testing.T) {
	initAction(t)
	err := set("+s19:32,记得买辣椒,2,60", 697611681)
	fmt.Println(err)
}

func TestZGetRedis(t *testing.T) {
	initAction(t)
	msg, err := get(time.Now().AddDate(0, 0, 15).Unix())
	if err != nil {
		fmt.Println("获取错误...", err.Error())
		return
	}

	fmt.Println(msg)
}

func TestZDelRedis(t *testing.T) {
	initAction(t)
	del(time.Now().Unix())
}

func Test_ELM(t *testing.T) {
	initAction(t)

	bot, err := tgbotapi.NewBotAPI(global.Conf.App.Token)
	if err != nil {
		log.Fatal(err.Error())
	}

	f, err := embed.GetStaticFileSystem().Open("elm.png")
	if err != nil {
		err = errors.Wrapf(err, "GetStaticFileSystem open elm.png")
		logrus.Error(err.Error())
	} else {
		reader := tgbotapi.FileReader{Name: "elm.png", Reader: f}
		_, _ = bot.Send(tgbotapi.NewPhoto(global.Conf.App.ChatID, reader))
	}
}
