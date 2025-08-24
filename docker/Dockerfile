##
## Build
##
ARG GOVERSION
FROM golang:${GOVERSION} AS build


ENV GOPROXY=https://goproxy.cn,direct

COPY ./ /kk_etcd/kk_etcd_go

WORKDIR /kk_etcd/kk_etcd_go

RUN CGO_ENABLED=0 GOOS=linux go build -o /kk_etcd/output /kk_etcd/kk_etcd_go/main/main.go

##
## Deploy
##
FROM alpine:latest

WORKDIR /

COPY --from=build /kk_etcd/output /kk_etcd/output
COPY ./docker/tarball/etcdutl /usr/local/bin/etcdutl

EXPOSE 2333

CMD ["/kk_etcd/output"]


# 本地构建测试
# docker build -t kk_etcd_go_local -f Dockerfile .


# 多平台构建-发布
# docker login
# docker buildx build --platform linux/amd64,linux/arm64 -t cruvie/kk_etcd_go:1.4.4 -f Dockerfile . --push
# docker buildx build --platform linux/amd64 -t cruvie/kk_etcd_go:1.4.4-amd64 -f Dockerfile . --push
# docker buildx build --platform linux/arm64 -t cruvie/kk_etcd_go:1.4.4-arm64 -f Dockerfile . --push

# 测试
# docker pull cruvie/kk_etcd_go:1.4.4
# docker run --name kk_etcd_go -p 2333:2333 -v ./config/config.yml:/kk_etcd_go/config/config.yml cruvie/kk_etcd_go:1.4.4
