package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var err error

func main() {

	//go Auth()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowOrigins:           nil,
		AllowOriginFunc:        nil,
		AllowMethods:           nil,
		AllowHeaders:           nil,
		AllowCredentials:       true,
		ExposeHeaders:          nil,
		MaxAge:                 0,
		AllowWildcard:          false,
		AllowBrowserExtensions: false,
		AllowWebSockets:        false,
		AllowFiles:             false,
	}))

	api := router.Group("/api")
	{
		api.GET("/login", AuthLogin)
		api.GET("/login/callback", CallBack)
	}
	// Simple group: v1
	v1 := router.Group("/v1")
	{

		v1.POST("/topics", v1GetTopics)

		v1.POST("/topic", v1CreateTopic)

		v1.POST("/clusters", v1Clusters)

		v1.GET("/clusters", v1GetClusters)

	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/topics", v2PostTopics)
		v2.GET("/topics", v2GetTopics)
		v2.GET("/groups", v2GetGroups)
		v2.GET("/querygroups", v2QueryAllGroups)
		v2.GET("/querygroupstatus", v2QueryGroupStatus)

		v2.POST("/allgroups", v2GetAllGroups)
		v2.POST("/groupstatus", v2GroupStatus)

		v2.POST("/addAlert", AddAlert)
		v2.POST("/deleteAlert", DeleteAlert)
		v2.POST("/sliceAlert", DoSilenceAlert)

		v2.POST("/listAlerts", ListAlerts)

		v2.GET("/todo", todo)

	}

	router.Run(":8081")
}

func todo(c *gin.Context) {
	c.String(http.StatusOK, "Todo")
}
