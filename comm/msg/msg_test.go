package msg

import (
	"fmt"
	"testing"

	"github.com/json-iterator/go/extra"
	"github.com/sirupsen/logrus"
	"go-tgbot/comm/global"
	conf2 "go-tgbot/conf"
)

func initAction(t *testing.T) (conf *conf2.AllConfig) {
	extra.RegisterFuzzyDecoders()
	logrus.SetLevel(logrus.DebugLevel)
	var (
		err error
	)
	conf, err = conf2.ReadConfig("../../config/prod.yaml")
	if err != nil {
		t.Logf("get conf err:%s ", err.Error())
		return
	}

	global.Conf = conf

	return conf
}

func Test_ReminderMsgAssemble(t *testing.T) {
	x := ReminderMsgAssemble("leetcode 01")
	t.Log(x)
}

func TestIsLeetCodeMsg(t *testing.T) {
	fmt.Println(IsLeetCodeMsg("leetcode: leetcode 01"))
	//fmt.Println(IsLeetCodeMsg("lc,207，算法205"))
	//fmt.Println(IsLeetCodeMsg("leetcode,207,算法205"))
	//fmt.Println(IsLeetCodeMsg("leetcode:207"))
	//fmt.Println(IsLeetCodeMsg("leetcode:  207，哈哈哈哈"))
	//fmt.Println(IsLeetCodeMsg("leetcode,   207"))
	//fmt.Println(IsLeetCodeMsg("leetcode   207"))
	//fmt.Println(IsLeetCodeMsg("leetcodex   207"))
}

func Test1(t *testing.T) {
	if ok, event := IsLeetCodeMsg("leetcode: leetcode 01"); ok {
		remindMsgs := ReminderMsgAssemble(event)
		fmt.Println(remindMsgs)
	}
}
