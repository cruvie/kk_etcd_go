package global_model

import (
	"gitee.com/cruvie/kk_go_kit/kk_stage"
)

const globalHeader = "globalHeader"

type RequestHeader struct {
	AuthorizationToken string
}

// SetRequestHeader set current request header to gin context
func SetRequestHeader(stage *kk_stage.Stage, header RequestHeader) {
	stage.Set(globalHeader, header)
}

// getRequestHeader get current request header from gin context
func getRequestHeader(stage *kk_stage.Stage) RequestHeader {
	header, ok := stage.Get(globalHeader)
	if !ok {
		return RequestHeader{}
	}
	return header.(RequestHeader)
}

// GetAuthorizationToken get AuthorizationToken from globalHeader
func GetAuthorizationToken(stage *kk_stage.Stage) string {
	return getRequestHeader(stage).AuthorizationToken
}
