package conf

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	cfg := cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "OPTION"},
		AllowHeaders:     []string{"Token", "Origin", "Content-Type"},
		AllowCredentials: true,
	}

	return cors.New(cfg)
}
