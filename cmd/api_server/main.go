package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// get list of feeds in database
func GetFeeds(c *gin.Context) {
	feeds := GetFeedsFromDatabase()
	c.IndentedJSON(http.StatusOK, feeds)
}

// get list of feed items using feed id
func GetFeedItems(c *gin.Context) {
	feedId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	feedItems := GetFeedItemsFromDatabase(feedId)
	c.IndentedJSON(http.StatusOK, feedItems)
}

func main() {
	fmt.Println("This is openfeeder api server")

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/feeds", GetFeeds)
	router.GET("/feeds/:id", GetFeedItems)
	router.Run("localhost:9080")
}
