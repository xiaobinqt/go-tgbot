# golang Telegram bot

⚠️本项目只是作者的一个玩具，大部分功能只为个人定制，并不通用。

## 功能演示

## 配置文件说明

```
app:
  debug: true
  chat_id: xxxx 
  keep_alive_chat_id: xxxx
  token: xxxxxxxx

keys:
  bot_name: 卫小兵
  weather_key: xxxxxxxxxxx # 高德天气 api key
  tianapi_key: xxxxxxxxxxxx # 天行 api key
  tianapi_key1: xxxxxxxxxxxxxx # 天行 api key
  qweather_key: xxxxxxxxxxxxxxxx # 和风天气 api key
  lover_ch_name: xxxxxx
  remind_msg: xxxx

# redis 配置
redis:
  ip: 39.xxxxx
  port: 6379
  passwd: xxxxxx
```

+ `app.debug` 是否是 debug 日志，会输出一些比较详细的日志
+ `app.chat_id` bot 定时发消息的 chat_id
+ `app.keep_alive_chat_id` 为了 bot 保活定时发消息 chat_id，如果的国外的服务器这个配置可以不要
+ `app.token` 机器人 token
+ `lover_ch_name` 和 `remind_msg` 私人订制的字符串:cry:

## 部署说明

clone 项目到本地，然后进入项目目录，将 `config/dev.yaml` 文件改成 `config/prod.yaml`， yaml 配置文件需要配置下，可以去对应的网站获取
apiKey。

执行如下命令：

```shell
go mod tidy # 下载依赖

go build -v -o tgbot  # 编译

nohup ./tgbot > /tmp/core.log & # 后台运行
```

也可以通过 Dockerfile 构建镜像运行容器。

## 功能列表

### 根据关键字回复

输入【英语一句话】关键字回复一句学习英语。

## 参考

+ [用 Go 寫 Telegram Bot](https://tonypepe.com/posts/telegram/go-tg-bot)

## 梯子

如果是国内的服务器，比如阿里云的，可以使用
clash，具体可以参考 [使用 Clash 为 Linux 提供网络代理服务](https://www.ahdark.blog/som/1643.shtml)
，我之前用的就是这个，但是感觉不太稳定，进程明明还在但是 bot 已经不收发消息了，其他人也有遇到同样的问题。

<div align="center"><img src="https://cdn.xiaobinqt.cn/xiaobinqt.io/20230410/3d5f57e8d752424a82bd4d89ff314982.png" width=  /></div>

后来我买了一个便宜的 vps，参考 [cheap-vps](https://renzhn.github.io/posts/cheap-vps/)，一年 10.78
刀，[840 MB KVM VPS (Easter 2023)](https://my.racknerd.com/cart.php?a=confproduct&i=0)

![](https://cdn.xiaobinqt.cn/xiaobinqt.io/20230410/9a3d07a8b57d461294db4e5e0e4b3a33.png)

## 联系方式

<div align="center"><img src="https://cdn.xiaobinqt.cn/xiaobinqt.io/20220319/d5616bfc809a45608437f9cc94b14044.jpg?imageView2/0/interlace/1/q/50|imageslim" width=320  /></div>

