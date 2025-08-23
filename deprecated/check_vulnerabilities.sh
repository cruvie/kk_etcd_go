#!/bin/bash
clear
# go install golang.org/x/vuln/cmd/govulncheck@latest
cd ..

(
  echo 'check_vulnerabilities'
  govulncheck ./...
)