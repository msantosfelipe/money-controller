package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/msantosfelipe/money-controller/commons/utils"
	"github.com/msantosfelipe/money-controller/config"
)

func CorsMiddleware() func(c *gin.Context) {
	allowedHeader := []string{
		"X-Application-Key",
		"Authorization",
		"X-Amz-Date",
		"X-Api-Key",
		"X-Amz-Security-Token",
		"X-Bifrost-Authorization",
	}

	config := cors.DefaultConfig()
	config.AddAllowHeaders(allowedHeader...)
	config.AddExposeHeaders("X-Request-Id")
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	return cors.New(config)
}

func OriginApplicationIDMiddleware(context *gin.Context) {
	appID := context.GetHeader("application-id")
	if appID == "" {
		err := fmt.Errorf("required header application-id is missing")
		context.AbortWithStatusJSON(http.StatusForbidden, utils.BuildResponseFromError(err))
		return
	}

	AllowedApplications := []string{
		config.ENV.APP_ID_1,
	}

	for _, app := range AllowedApplications {
		if app == appID {
			context.Set(config.ApplicationIdKey, appID)
			config.ApplicationIdKey = appID
			context.Next() // Application authorized
			return
		}
	}
	err := fmt.Errorf("application-id (%s) does not have permission to do this operation", appID)
	context.AbortWithStatusJSON(http.StatusUnauthorized, utils.BuildResponseFromError(err))
	return
}
