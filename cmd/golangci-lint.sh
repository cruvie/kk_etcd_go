clear
# export GOPROXY=https://goproxy.io,direct
# go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
cd ..
  golangci-lint version
  golangci-lint run