package main

import (
	"github.com/gin-gonic/gin"
	// "fmt"
	"github.com/Nyuuk/learn-api-golang/helpers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", pong)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

func pong(c *gin.Context) {
	config, err := helpers.LoadConfig("config.yaml")
	
	if err != nil {
		panic(err)
	}

	dt := map[string]interface{}{
		// "message": config.App.HelloMsg,
		// "name":    config.App.Name,
		// "port":    config.App.Port,
		"database": config.Database.Postgres,
		"config":  config,
	}
	helpers.Response(c, dt, helpers.ResponseHelper{Code: 200, Message: "success hello"})
}