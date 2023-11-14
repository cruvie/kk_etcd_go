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

// getRequestHeader get current request header from gin context
func getRequestHeader(c *gin.Context) RequestHeader {
	header, ok := c.Get(globalHeader)
	if !ok {
		return RequestHeader{}
	}
	return header.(RequestHeader)
}

// GetAuthorizationToken get AuthorizationToken from globalHeader
func GetAuthorizationToken(c *gin.Context) string {
	return getRequestHeader(c).AuthorizationToken
}
