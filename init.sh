#!/usr/bin/env zsh

appName=$1
if [ ! "$appName" ]; then
  appName="demo"
  prinf "no appName specific, will use demo as default\n"
fi

rm -rf .git
git init

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go mod init "$appName"

if [[ $(uname) == 'Darwin' ]]; then
  sed -i "" "s/{{app_name}}/$appName/g" docker-compose.yaml
  sed -i "" "s/{{app_name}}/$appName/g" main.go
fi

echo "require (
  github.com/kataras/iris/v12 latest
  github.com/shurcooL/sanitized_anchor_name
	github.com/stretchr/testify latest
	golang.org/x/crypto
	github.com/go-sql-driver/mysql latest
	gorm.io/driver/mysql latest
	gorm.io/gorm latest
)" >> go.mod
go mod download
