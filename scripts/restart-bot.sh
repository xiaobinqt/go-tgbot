#!/bin/bash

RETRY_LIMIT=50
RETRY_COUNT=0

while true; do
  # 获取 tgbot 进程的 PID
  PID=$(ps -ef | grep tgbot | grep -v grep | awk '{print $2}')

  # 如果 tgbot 服务没有启动，则重试，最多重试 RETRY_LIMIT 次
  if [ -z "$PID" ]; then
    if [ $RETRY_COUNT -lt $RETRY_LIMIT ]; then
      echo "tgbot 服务未启动，正在第 $((RETRY_COUNT + 1)) 次重试..." >>/tmp/restart-bot.log
      sleep 5
      RETRY_COUNT=$((RETRY_COUNT + 1))
    else
      echo "已达到最大重试次数，tgbot 服务启动失败。" >>/tmp/restart-bot.log
      exit 1
    fi
  # 如果 tgbot 服务已经启动，则进行重启
  else
    echo "tgbot 进程的 PID 为 $PID" >>/tmp/restart-bot.log
    echo "正在停止 tgbot 进程..." >>/tmp/restart-bot.log
    kill -9 $PID
    echo "等待 5 秒钟，以确保进程已经停止" >>/tmp/restart-bot.log
    sleep 5
    echo "正在启动 tgbot 进程..." >>/tmp/restart-bot.log
    # 在此处添加启动 tgbot 的命令
    nohup /usr/bin/tgbot >/tmp/tgbot.log 2>&1 &

    # 检查 tgbot 是否重启成功，最多重试 RETRY_LIMIT 次
    COUNT=0
    while true; do
      PID=$(ps -ef | grep tgbot | grep -v grep | awk '{print $2}')
      if [ -n "$PID" ]; then
        echo "tgbot 重启成功！" >>/tmp/restart-bot.log
        exit 0
      elif [ $COUNT -lt $RETRY_LIMIT ]; then
        echo "tgbot 重启失败，正在第 $((COUNT + 1)) 次重试..." >>/tmp/restart-bot.log
        sleep 5
        COUNT=$((COUNT + 1))
      else
        echo "已达到最大重试次数，tgbot 服务启动失败。" >>/tmp/restart-bot.log
        exit 1
      fi
    done
  fi
done
