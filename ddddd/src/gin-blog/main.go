package main

import (
	"fmt"
	"github.com/dearllu/ddddd/src/gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"message":"test"})
	})

	s := &http.Server{
		Addr: fmt.Sprintf(":%d",setting.HTTPPort),
		Handler: router,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1<<20,
	}

	s.ListenAndServe() //开始监听服务
}