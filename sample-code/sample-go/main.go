package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	//"fmt"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	//r := gin.New()
	//r.Use(gin.Recovery())
	r := gin.Default()

	/*
		r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
			// your custom format
					// your custom format
			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
					param.ClientIP,
					param.TimeStamp.Format(time.RFC3339Nano),
					param.Method,
					param.Path,
					param.Request.Proto,
					param.StatusCode,
					param.Latency,
					param.Request.UserAgent(),
					param.ErrorMessage,
			)
		}))
		//r.Use(gin.Recovery())
	*/
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		latency := c.Query("latency")
		size := c.Query("size")
		if latency != "" {
			i, err := strconv.Atoi(latency)
			if err == nil {
				time.Sleep(time.Duration(i) * time.Millisecond)
			}
		}
		if size != "" {
			i, err := strconv.Atoi(size)
			if err == nil {
				c.String(http.StatusOK, RandStringRunes(i))
			}
		} else {
			c.String(http.StatusOK, "pong")
		}
	})

	r.GET("/meetings", func(c *gin.Context) {
		c.String(http.StatusOK, c.Request.URL.Path+" ok ")
	})

	r.GET("/entries", func(c *gin.Context) {
		c.String(http.StatusOK, c.Request.URL.Path+" ok ")
	})

	r.GET("/horses", func(c *gin.Context) {
		c.String(http.StatusOK, c.Request.URL.Path+" ok ")
	})

	r.GET("/racingprogramme", func(c *gin.Context) {
		c.String(http.StatusOK, c.Request.URL.Path+" ok ")
	})

	return r
}

func main() {

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8000")
}
