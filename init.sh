#!/usr/bin/env zsh

appName=$1
if [ ! "$appName" ]; then
  appName="demo"
  prinf "no appName specific, will use demo as default\n"
fi

current=${PWD##*/}
cd .. && mv "$current" $appName && cd $appName

git init

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go mod init "$appName"

if [[ $(uname) == 'Darwin' ]]; then
  sed -i "" "s/{{app_name}}/$appName/g" ./docker-compose.yaml
fi

echo "require (
	github.com/stretchr/testify latest
	gorm.io/driver/mysql latest
	gorm.io/gorm latest
)" >> go.mod
go mod tidy
