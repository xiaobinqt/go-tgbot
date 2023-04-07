#!/bin/sh

# 获取 tgbot 进程的 PID
PID=$(ps -ef | grep tgbot | grep -v grep | awk '{print $2}')

if [ -z "$PID" ]; then
  echo "服务启动...."
  nohup /usr/bin/tgbot >/tmp/tgbot.log 2>&1 &
else
  echo "tgbot 进程的 PID 为 $PID"
  echo "正在停止 tgbot 进程..."
  kill -9 $PID
  echo "等待 5 秒钟，以确保进程已经停止"
  sleep 5
  echo "正在启动 tgbot 进程..."
  # 在此处添加启动 tgbot 的命令
  nohup /usr/bin/tgbot >/tmp/tgbot.log 2>&1 &
fi
