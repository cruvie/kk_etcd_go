#!/bin/bash
clear
cd ..
(
  echo 'kk_etcd_go'
  swag fmt
  swag init -g ./main/main.go -o ./swagger --parseDependency true
)