package libraries

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinJwt(c *gin.Context) {
	noAuth := []string{"/api/v1/customers/test"}
	path := c.Request.RequestURI
	//log.Print("url path ",path)
	for _, val := range noAuth {
		if val == path {
			c.Next()
			return
		}
	}
	header := c.GetHeader("Authorization")
	if header == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	token := SplitToken(header)
	valid, uid := ValidateToken(token)
	if !valid {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Invalid Token",
		})
		return
	}
	c.Set("user", uid)
	c.Next()
}
