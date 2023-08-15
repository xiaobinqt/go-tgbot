package msg

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/sirupsen/logrus"
	"strings"
)

func IsLeetCodeMsg(txt string) (bool, string) {
	if txt == "" {
		return false, ""
	}

	var (
		txtArr = make([]string, 0)
	)

	logrus.Infof("IsLeetCodeMsg: %s", txt)
	txt = strings.ToLower(txt)

	// 两个 //
	txtArr = strings.Split(txt, "//")
	if len(txtArr) < 2 {
		return false, ""
	}

	for index, each := range txtArr {
		each = strings.TrimLeft(each, " ")
		txtArr[index] = strings.TrimRight(each, " ")
	}

	if txtArr[0] != "leetcode" && txtArr[0] != "lc" {
		return false, ""
	}

	return true, txtArr[1]
}

// 每天的 10 点提醒
// +st20221227 15:35,消息内容
func ReminderMsgAssemble(txt string) (msgs []string) {
	msgs = make([]string, 0)

	format := func(time string) string {
		return fmt.Sprintf("%s // %s", time, txt)
	}

	// 设置提醒格式,分别隔 1 天、2 天、4 天、7 天、15 天、1 个月、3 个月、6 个月
	//carbon.Parse(fmt.Sprintf("%s 10:00", carbon.Now().ToDateString())).AddDays(1)
	msgs = append(msgs,
		format(fmt.Sprintf("+st%s 09:30", carbon.Now().AddDays(1).Format("Ymd"))),
		format(fmt.Sprintf("+st%s 09:30", carbon.Now().AddDays(2).Format("Ymd"))),
		format(fmt.Sprintf("+st%s 09:30", carbon.Now().AddDays(4).Format("Ymd"))),
		format(fmt.Sprintf("+st%s 09:30", carbon.Now().AddDays(7).Format("Ymd"))),
		format(fmt.Sprintf("+st%s 09:30", carbon.Now().AddMonths(1).Format("Ymd"))),
		format(fmt.Sprintf("+st%s 09:30", carbon.Now().AddMonths(3).Format("Ymd"))),
		format(fmt.Sprintf("+st%s 09:30", carbon.Now().AddMonths(6).Format("Ymd"))),
	)

	return msgs
}
