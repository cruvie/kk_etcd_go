#!/bin/bash
go env -w GOPROXY=https://goproxy.cn,direct

OUTPUT_PATH="./program/kk_etcd_go"
MAIN_FILE="../main/main.go"

echo 'build for amd64 linux'
CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ${OUTPUT_PATH}-linux-amd64 ${MAIN_FILE}
echo 'build for arm64 linux'
CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o ${OUTPUT_PATH}-linux-arm64 ${MAIN_FILE}
echo 'build for amd64 mac'
CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o ${OUTPUT_PATH}-mac-amd64 ${MAIN_FILE}
echo 'build for arm64 mac'
CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -o ${OUTPUT_PATH}-mac-arm64 ${MAIN_FILE}
echo 'build for amd64 windows'
CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build -o ${OUTPUT_PATH}-windows-amd64 ${MAIN_FILE}
echo 'build for arm64 windows'
CGO_ENABLED=0 GOARCH=arm64 GOOS=windows go build -o ${OUTPUT_PATH}-windows-arm64 ${MAIN_FILE}
