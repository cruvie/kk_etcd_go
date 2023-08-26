#!/bin/bash

cd ..

(
  echo 'update_dependence'
  go get -d -u ./...
)
