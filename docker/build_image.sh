clear
cd ..

docker buildx build --platform linux/amd64 -t cruvie/kk_etcd_go:1.4.3-amd64 -f docker/Dockerfile . --push
#docker buildx build --platform linux/arm64 -t cruvie/kk_etcd_go:1.4.3-arm64 -f docker/Dockerfile . --push