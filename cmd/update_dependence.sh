#!/bin/bash
clear
cd ..

(
  echo 'update_dependence'
  go get -d -u ./...
)
