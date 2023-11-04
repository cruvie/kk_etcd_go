#!/bin/bash
#brew install golangci-lint
#https://golangci-lint.run/usage/linters/
#已包含go vet/staticcheck
clear
cd ..
  golangci-lint run
