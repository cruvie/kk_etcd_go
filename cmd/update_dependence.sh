#!/bin/bash
clear
cd ..

(
  echo 'update_dependence'
  go get gitee.com/cruvie/kk_go_kit@master
  go get -u ./...
  go mod tidy
#  go mod vendor
)
