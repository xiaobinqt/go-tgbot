package msg

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"go-tgbot/comm/image"
	"go-tgbot/comm/qweather"
	"go-tgbot/comm/tian"
	"go-tgbot/ecode"
	"go-tgbot/ticker"
	"os"
	"strings"
)

func HandleMsg(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message.Text == "程序员鼓励师" {
		Encourage(bot, update)
		return
	}

	msg := contextTextBypass(update.Message.Text, update.Message.Chat.ID)
	if msg == "" {
		return
	}

	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))
}

func Encourage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	imgURL, err := image.GetImage()
	if err != nil {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "鼓励师今天不在家，不要摸鱼，赶紧干活~"))
		return
	}

	savePath, err := image.SaveEncourageImg(imgURL)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "鼓励师今天不在家，BUG 虽好，但不要贪多哦~"))
		return
	}
	defer os.Remove(savePath)

	if err != nil {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "鼓励师今天不在家，么么哒~"))
		return
	}

	photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath(savePath))
	if _, err = bot.Send(photo); err != nil {
		bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "鼓励师出差了~"))
	}
}

func trimMsgContent(content string) string {
	content = strings.TrimLeft(content, " ")
	content = strings.TrimRight(content, " ")
	return content
}

func contextTextBypass(txt string, chatID int64) (retMsg string) {
	var (
		err error
	)
	if txt == "菜单" {
		return ecode.HelpMessage
	}

	if txt == "天气" {
		return "支持天气查询，如: 泾县天气。"
	}

	if txt == "菜谱" || txt == "做法" {
		return "支持菜谱查询，如: 红烧肉菜谱，红烧肉做法。"
	}

	// 天气处理
	if strings.HasSuffix(txt, "天气") {
		return handleWeatherMsg(txt)
	}

	// 毒鸡汤处理
	if txt == "毒鸡汤" {
		retMsg, err = tian.GetMessage(tian.C_dujitang)
		if err != nil {
			logrus.Error(err.Error())
			return ""
		}
		return retMsg
	}

	// 菜谱处理
	if strings.HasSuffix(txt, "菜谱") || strings.HasSuffix(txt, "做法") {
		return handleCookbookMsg(txt)
	}

	// 英语一句话
	if txt == "英语一句话" {
		retMsg, err = tian.GetMessage(tian.C_englishSentence)
		if err != nil {
			logrus.Error(err.Error())
			return ""
		}
		return retMsg
	}

	// 事件提醒
	if txt == "事件提醒" {
		return `
格式0：+s15:32//消息内容
格式0说明：今天 15:32 提醒我「消息内容」

格式1：+s15:32//消息内容//3//60
格式1说明：今天 15:32 提醒我「消息内容」,提醒 3 次每次间隔 60s

格式2: +st20221227 15:35//消息内容
格式2说明：20221227 日 15:35 提醒我「消息内容」。注意此格式的日期和时间中间的空格不能丢

格式3: +st20221227 15:35//消息内容//3//60
格式3说明：20221227 日 15:35 提醒我「消息内容」,提醒 3 次每次间隔 60s。注意此格式的日期和时间中间的空格不能丢
`
	}

	if ticker.IsScheduleNotice(txt) {
		return ticker.AddScheduleNotice(txt, chatID)
	}

	// 艾宾浩斯遗忘曲线复习提醒
	if ok, event := IsLeetCodeMsg(txt); ok {
		remindMsgs := ReminderMsgAssemble(event)
		returnMsg := make([]string, 0)
		for _, each := range remindMsgs {
			x := ticker.AddScheduleNotice(each, chatID)
			if x != "" {
				returnMsg = append(returnMsg, x)
			}
		}

		return strings.Join(returnMsg, "\n")
	}

	// todo 其他的一些

	return ""
}

func handleCookbookMsg(txt string) (cookbook string) {
	var (
		err error
	)
	originTxt := txt

	txt = strings.ReplaceAll(txt, "做法", "")
	txt = strings.ReplaceAll(txt, "菜谱", "")
	cookbook, err = tian.GetMessage(tian.C_caipu, txt)
	if err != nil && err != tian.ErrNotfoundCaiPu {
		logrus.Error(err.Error())
		return ""
	}

	if err == tian.ErrNotfoundCaiPu {
		return fmt.Sprintf("暂未找到%s，请重新输入关键字查询", originTxt)
	}

	return cookbook
}

func handleWeatherMsg(txt string) (formatWeatherMsg string) {
	var (
		err    error
		cityID string
	)
	originTxt := txt
	txt = strings.ReplaceAll(txt, "天气", "")
	cityID, err = qweather.GetLocationID(txt)
	if err != nil {
		err = errors.Wrapf(err, "handleWeatherMsg GetFormatWeatherMessage err")
		logrus.Error(err.Error())
		return fmt.Sprintf("(1)未查询到%s，请检查城市输入是否正确，当前只支持到区/县一级的城市查询，如：泾县天气，新市区天气。", originTxt)
	}

	formatWeatherMsg, err = qweather.GetQWeatherDetail(cityID, txt)
	if err != nil {
		err = errors.Wrapf(err, "handleWeatherMsg GetFormatWeatherMessage err")
		logrus.Error(err.Error())
		return fmt.Sprintf("(2)未查询到%s，请检查城市输入是否正确，当前只支持到区/县一级的城市查询，如：泾县天气，新市区天气。", originTxt)
	}

	return formatWeatherMsg
}
