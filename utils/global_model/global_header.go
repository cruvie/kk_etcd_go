package global_model

import (
	"github.com/gin-gonic/gin"
)

const globalHeader = "globalHeader"

type RequestHeader struct {
	AuthorizationToken string
}

// SetRequestHeader set current request header to gin context
func SetRequestHeader(c *gin.Context, header RequestHeader) {
	c.Set(globalHeader, header)
}

// GetRequestHeader get current request header from gin context
func GetRequestHeader(c *gin.Context) RequestHeader {
	header, ok := c.Get(globalHeader)
	if !ok {
		return RequestHeader{}
	}
	return header.(RequestHeader)
}

// ClearRequestHeader clear current request header from gin context
func ClearRequestHeader(c *gin.Context) {
	delete(c.Keys, globalHeader)
}

// GetAuthorizationToken get AuthorizationToken from globalHeader
func GetAuthorizationToken(c *gin.Context) string {
	return GetRequestHeader(c).AuthorizationToken
}