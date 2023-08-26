#!/bin/bash

cd ..

(
  echo 'kk_etcd_go'
  swag fmt
  swag init -g ./main/main.go -o ./main/docs
)