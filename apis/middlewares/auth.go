package middlewares

import (
	"github.com/gin-gonic/gin"
)

func (m *MiddleWare) Authenticate(c *gin.Context) {
	c.Next()
	//tokenValue := c.GetHeader("Authorization")
	//if tokenValue == "" {
	//	c.JSON(http.StatusUnauthorized, errors.NewCustomHttpError(errors.Unauthorized, "header_authorization_not_found"))
	//	c.Abort()
	//	return
	//}
	//c.Next()
}
