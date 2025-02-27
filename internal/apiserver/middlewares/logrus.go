package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logrus() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Infof("Request: %s %s", c.Request.Method, c.Request.URL)
		c.Next()
	}
}
