# brew install just

set shell := ["zsh", "-ic"]

default: fmt

which-shell:
	echo ${SHELL}
	pwd

fmt:
	go tool gofumpt -l -w .

lint:
	go tool golangci-lint version
	go tool golangci-lint run

deadcode:
	go tool deadcode ./...

govulncheck:
	go tool govulncheck ./...

fix:
	go fix ./...


update-dep:
	go get -u ./...
	go mod tidy


update_self_dependence_in_one_mod:
  # export GOPROXY=https://goproxy.io,direct
  #  go env | grep  GOPROXY
  #  export https_proxy=http://127.0.0.1:7890 http_proxy=http://127.0.0.1:7890 all_proxy=socks5://127.0.0.1:7890
  #  go get github.com/cruvie/kk_etcd_go@master
  go get  gitee.com/cruvie/kk_go_kit@master
  go mod tidy
  go mod vendor


proto-gen:
    buf generate


test:
	go test ./...
