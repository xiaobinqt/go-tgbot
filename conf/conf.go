package conf

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	embed2 "go-tgbot/comm/embed"
	"gopkg.in/yaml.v3"
	"io"
	"io/fs"
	"os"
)

// 说明：mapstructure 是viper的tag

// AllConfig .
type AllConfig struct {
	App       App       `json:"app" yaml:"app"`
	Keys      Keys      `json:"keys" yaml:"keys"`
	RedisConf RedisConf `json:"redis" yaml:"redis"`
}

// App .
type App struct {
	Debug           bool   `json:"debug" yaml:"debug" mapstructure:"debug"`
	ChatID          int64  `json:"chat_id" yaml:"chat_id" mapstructure:"chat_id"`
	KeepAliveChatID int64  `json:"keep_alive_chat_id" yaml:"keep_alive_chat_id" mapstructure:"keep_alive_chat_id"`
	Token           string `json:"token" yaml:"token" mapstructure:"token"`
}

type RedisConf struct {
	IP     string `json:"ip" yaml:"ip" mapstructure:"ip"`
	Port   string `json:"port" yaml:"port" mapstructure:"port"`
	Passwd string `json:"passwd" yaml:"passwd" mapstructure:"passwd"`
}

type Keys struct {
	BotName     string `json:"bot_name" yaml:"bot_name" mapstructure:"bot_name"`
	WeatherKey  string `json:"weather_key" yaml:"weather_key" mapstructure:"weather_key"`
	TianapiKey  string `json:"tianapi_key" yaml:"tianapi_key" mapstructure:"tianapi_key"`
	TianapiKey1 string `json:"tianapi_key1" yaml:"tianapi_key1" mapstructure:"tianapi_key1"`
	LoverChName string `json:"lover_ch_name" yaml:"lover_ch_name" mapstructure:"lover_ch_name"`
	QweatherKey string `json:"qweather_key" yaml:"qweather_key" mapstructure:"qweather_key"`
	RemindMsg   string `json:"remind_msg" yaml:"remind_msg" mapstructure:"remind_msg"`
}

func ReadConfig(path ...string) (c *AllConfig, err error) {
	var (
		yamlFile = make([]byte, 0)
		f        fs.File
	)

	if len(path) > 0 {
		f, err = os.Open(path[0])
	} else {
		f, err = embed2.GetWebFileSystem().Open("prod.yaml")
	}
	if err != nil {
		err = errors.Wrapf(err, "embed open config.yaml err")
		return nil, err
	}

	yamlFile, err = io.ReadAll(f)
	if err != nil {
		err = errors.Wrapf(err, "ReadFile error")
		logrus.Errorf(err.Error())
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		err = errors.Wrapf(err, "yaml.Unmarshal error")
		logrus.Errorf(err.Error())
		return nil, err
	}

	return c, nil
}
