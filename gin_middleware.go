package logrus_starter

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"regexp"
	"time"
)

var printHealthFlag = true // Print health check api logs only once

// LogrusMiddleware Log print
func LogrusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		processTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUrl := c.Request.RequestURI
		statusCode := c.Writer.Status()
		matched, _ := regexp.MatchString(`^/(.*)?(health)$`, reqUrl)
		if matched {
			if !printHealthFlag {
				return
			}
			printHealthFlag = false
		}
		businessCode, exists := c.Get("bcode")
		if exists {
			statusCode = businessCode.(int)
		}
		log.Infof("| %s %3d %s | %13v |%s %s %s  %s ", getStatusColor(statusCode), statusCode, reset, processTime, getMethodColor(reqMethod), reqMethod, reset, reqUrl)
	}
}
