package main

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	r.Use(func(c *gin.Context) {
		s := time.Now()

		c.Next()

		log.Info("incoming request",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("elapsed", time.Now().Sub(s)),
		)
	}, func(c *gin.Context) {
		c.Set("requestId", rand.Int())
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		h := gin.H{
			"message": "pong",
		}
		if requestId, exists := c.Get("requestId"); exists {
			h["requestId"] = requestId
		}
		c.JSON(http.StatusOK, h)
	})
	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
