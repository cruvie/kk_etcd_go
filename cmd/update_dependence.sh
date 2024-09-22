#!/bin/bash
clear
cd ..

(
  echo 'update_dependence'
  go get -u ./...
  go mod tidy
  go mod vendor
)
