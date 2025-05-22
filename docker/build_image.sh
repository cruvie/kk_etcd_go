clear
cd ..

export GOVERSION=1.23.9

#export PLATFORM=linux/amd64
#export EtcdPath=./docker/tarball/etcd-v3.6.0-linux-amd64/etcdutl
#export TAG=1.4.3-amd64

export PLATFORM=linux/arm64
export EtcdPath=./docker/tarball/etcd-v3.6.0-linux-arm64/etcdutl
export TAG=1.4.3-arm64

#since COPY doesn't support interpolation, so cp a temp file
cp ${EtcdPath} ./docker/tarball/etcdutl

# user echo to check cmd
docker buildx build --platform ${PLATFORM} \
            --build-arg GOVERSION=${GOVERSION} \
            -t cruvie/kk_etcd_go:${TAG} \
            -f docker/Dockerfile . \
            --push

rm -f ./docker/tarball/etcdutl
