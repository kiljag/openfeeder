package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	apiResponse := ApiResponse{
		Success: true,
		Message: "ok",
	}
	c.IndentedJSON(http.StatusOK, apiResponse)
}

func main() {
	fmt.Println("This is openfeeder api server")

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api/healthcheck", HealthCheck)
	router.GET("/api/feeds", GetFeeds)
	router.GET("/api/feedItems/:id", GetFeedItems)
	router.GET("/api/feedItem/:id/:item", GetFeedItemContent)

	router.POST("/api/addFeed", PostAddFeed)
	router.POST("/api/deleteFeed/:id", PostDeleteFeed)
	router.Run("localhost:9080")
}
