package main

import (
	"net/http"
	"openfeeder/pkg/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

// get list of feed items using feed id
func GetFeedItems(c *gin.Context) {
	feedId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	feedItems := db.SelectFeedItemsById(uint64(feedId))
	rssItems := make([]RssFeedItem, 0)
	for _, item := range feedItems {
		rssItems = append(rssItems, RssFeedItem{
			Id:          item.ID,
			Title:       item.Title,
			Link:        item.Link,
			PublishedAt: item.PubDate.String(),
		})
	}

	apiResponse := ApiResponse{
		Success: true,
		Data:    rssItems,
	}

	c.IndentedJSON(http.StatusOK, apiResponse)
}

// get feed item using feed id and item id
func GetFeedItemContent(c *gin.Context) {
	feedId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	feedItemId, _ := strconv.ParseInt(c.Param("item"), 10, 64)
	item := db.SelectFeedItem(uint64(feedId), uint64(feedItemId))

	rssItem := RssFeedItem{
		Id:          item.ID,
		Title:       item.Title,
		Description: *item.Description,
		Link:        item.Link,
		PublishedAt: item.PubDate.String(),
	}

	apiResponse := ApiResponse{
		Success: true,
		Data:    rssItem,
	}

	c.IndentedJSON(http.StatusOK, apiResponse)
}
