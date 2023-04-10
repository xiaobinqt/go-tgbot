# golang Telegram bot

⚠️本项目只是作者的一个玩具，大部分功能只为个人定制，并不通用。

## 功能演示

## 部署说明

clone 项目到本地，然后进入项目目录，将 `config/dev.yaml` 文件改成 `config/prod.yaml`， yaml 配置文件需要配置下，可以去对应的网站获取
apiKey。

执行如下命令：

```shell
go mod tidy # 下载依赖

go build -v -o tgbot  # 编译

nohup ./wxbot > core.log & # 后台运行
```

## 功能列表


### 根据关键字回复

基于 [天行](https://www.tianapi.com/) api 和 [和风天气](https://console.qweather.com/#/console?lang=zh) 查询接口开发。

比如在群里发送【泾县天气】机器人会回复泾县今日的天气情况。

现在支持的关键字查询如下：

```
天气查询，如：泾县天气。
菜谱查询，如: 红烧肉菜谱，红烧肉做法。
输入【程序员鼓励师】收到程序员鼓励师的回复。
输入【毒鸡汤】关键字回复毒鸡汤。
输入【英语一句话】关键字回复一句学习英语。
```

## 联系方式

<div align="center"><img src="https://cdn.xiaobinqt.cn/xiaobinqt.io/20220319/d5616bfc809a45608437f9cc94b14044.jpg?imageView2/0/interlace/1/q/50|imageslim" width=320  /></div>

## 参考

+ [用 Go 寫 Telegram Bot](https://tonypepe.com/posts/telegram/go-tg-bot)

## 梯子

如果是国内的服务器，比如阿里云的，可以使用 clash，具体可以参考 [使用 Clash 为 Linux 提供网络代理服务](https://www.ahdark.blog/som/1643.shtml)，我之前用的就是这个，但是感觉不太稳定，进程明明还在但是 bot 已经不收发消息了，其他人也有遇到同样的问题。

<div align="center"><img src="https://cdn.xiaobinqt.cn/xiaobinqt.io/20230410/3d5f57e8d752424a82bd4d89ff314982.png?imageView2/0/q/75|watermark/2/text/eGlhb2JpbnF0/font/dmlqYXlh/fontsize/1000/fill/IzVDNUI1Qg==/dissolve/52/gravity/SouthEast/dx/15/dy/15" width=  /></div>

后来我买了一个便宜的 vps，参考 [cheap-vps](https://renzhn.github.io/posts/cheap-vps/)，一年 10.78 刀，[840 MB KVM VPS (Easter 2023)](https://my.racknerd.com/cart.php?a=confproduct&i=0)

![](https://cdn.xiaobinqt.cn/xiaobinqt.io/20230410/9a3d07a8b57d461294db4e5e0e4b3a33.png)
