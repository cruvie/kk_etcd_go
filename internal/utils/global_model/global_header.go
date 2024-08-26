package global_model

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
)

const globalHeader = "globalHeader"

type RequestHeader struct {
	UserName           string //delete after todo use https://github.com/etcd-io/etcd/pull/16803
	Password           string //delete after
	AuthorizationToken string
}

// SetRequestHeader set current request header to gin context
func SetRequestHeader(stage *kk_stage.Stage, header RequestHeader) {
	stage.Set(globalHeader, header)
}

// GetRequestHeader get current request header from gin context
func GetRequestHeader(stage *kk_stage.Stage) RequestHeader {
	header, ok := stage.Get(globalHeader)
	if !ok {
		return RequestHeader{}
	}
	return header.(RequestHeader)
}

// GetAuthorizationToken get AuthorizationToken from globalHeader
func GetAuthorizationToken(stage *kk_stage.Stage) string {
	return GetRequestHeader(stage).AuthorizationToken
}
