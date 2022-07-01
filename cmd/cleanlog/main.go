package main

import (
	"github.com/Ubbo-Sathla/tools/cmd/cleanlog/instance"
	service2 "github.com/Ubbo-Sathla/tools/cmd/cleanlog/service"
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(ginprom.PromMiddleware(nil))
	r.GET("/metrics", ginprom.PromHandler(promhttp.Handler()))

	v1 := r.Group("/api")
	{
		v1.POST("/addService", service2.AddServiceConfig)
		v1.GET("/listService", service2.ListServiceConfig)

		v1.POST("/addInstance", instance.AddInstance)
		v1.GET("/listInstance", instance.ListInstance)
		v1.POST("/delInstance", instance.DeleteInstance)

	}
	r.Run(":8088") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
