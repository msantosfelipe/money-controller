package logger

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/msantosfelipe/money-controller/config"
	"github.com/sirupsen/logrus"
)

//Log -
func Log(message string, a ...interface{}) *logrus.Entry {
	m := fmt.Sprintf(message, a...)
	logger := logrus.WithFields(logrus.Fields{
		"log_version": "1.0",
		"product": map[string]interface{}{
			"name":        config.ProductName,
			"application": config.ApplicationName,
			"version":     config.ApplicationVersion,
		},
		"context": map[string]interface{}{
			"message": m,
		},
	})

	return logger
}

//GetLogger -
func GetLogger(c *gin.Context) *logrus.Entry {
	ctxLog, _ := c.MustGet("log").(*logrus.Entry)
	return ctxLog
}

// LogMiddleware Make base logger
func LogMiddleware(context *gin.Context) {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logger := logrus.WithFields(logrus.Fields{
		"event_origin": map[string]string{
			"system":         config.ApplicationName,
			"client_address": context.GetHeader("Origin"),
		},
		"path":           context.Request.URL.Path,
		"x-request-id":   uuid.New().String(),
		"date_time":      time.Now(),
		"request-app-id": context.GetHeader("request-app-id"),
		"version":        config.ApplicationVersion,
	})

	logger.Info("Request initialized.")
	context.Set("log", logger)
	context.Next()
}
