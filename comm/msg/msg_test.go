package msg

import (
	"testing"

	"github.com/json-iterator/go/extra"
	"github.com/sirupsen/logrus"
	conf2 "go-tgbot/comm/conf"
	"go-tgbot/comm/global"
)

func initAction(t *testing.T) (conf *conf2.Conf) {
	extra.RegisterFuzzyDecoders()
	logrus.SetLevel(logrus.DebugLevel)
	var (
		err error
	)
	conf, err = conf2.GetConf("../../config/prod.yaml")
	if err != nil {
		t.Logf("get conf err:%s ", err.Error())
		return
	}

	global.Conf = conf

	return conf
}
