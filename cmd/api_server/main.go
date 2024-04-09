package main

import (
	"fmt"
	"net/http"
	"openfeeder/pkg/rss"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	r := make(map[string]string)
	r["message"] = "ok"
	c.IndentedJSON(http.StatusOK, r)
}

// get list of feeds in database
func GetFeeds(c *gin.Context) {
	feeds := rss.FetchFeedsFromDb()
	c.IndentedJSON(http.StatusOK, feeds)
}

// get list of feed items using feed id
func GetFeedItems(c *gin.Context) {
	feedId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	feedItems := rss.FetchFeedItemsFromDb(feedId)
	c.IndentedJSON(http.StatusOK, feedItems)
}

// get feed item using feed id and item id
func GetFeedItemContent(c *gin.Context) {
	feedId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	feedItemId, _ := strconv.ParseInt(c.Param("item"), 10, 64)
	feedItem := rss.FetchFeedItemFromDb(feedId, feedItemId)
	c.IndentedJSON(http.StatusOK, feedItem)
}

func main() {
	fmt.Println("This is openfeeder api server")

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/api/healthcheck", HealthCheck)
	router.GET("/api/feeds", GetFeeds)
	router.GET("/api/feedItems/:id", GetFeedItems)
	router.GET("/api/feedItem/:id/:item", GetFeedItemContent)
	router.Run("localhost:9080")
}
