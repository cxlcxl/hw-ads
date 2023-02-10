#!/bin/bash

if [[ $1 = "web" ]]
then
  echo "---- 打包 $1 服务中 ----"
  CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./web/ads ./web/main.go
  echo "---- 打包完成 ----"
elif [[ $1 = "cron" ]]
then
  echo "---- 打包 $1 服务中 ----"
  CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./cronJobs/cron ./cronJobs/cron.go
  echo "---- 打包完成 ----"
elif [[ $1 = "all" ]]
then
  echo "---- 打包 web 服务中 ----"
  CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./web/ads ./web/main.go
  echo "---- web 打包完成 ----"
  echo ""
  echo "---- 打包 cron 服务中 ----"
  CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./cronJobs/cron ./cronJobs/cron.go
  echo "---- cron 打包完成 ----"

  echo "---- 全部完成 ----"
else
  echo "error：请传入要打包的模块，当前包含模块「web、cron、all」"
fi

