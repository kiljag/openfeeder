package main

import (
	"log"
	"net/http"
	"openfeeder/pkg/db"
	"openfeeder/pkg/rss"

	"github.com/gin-gonic/gin"
)

// get list of feeds in database
func GetFeeds(c *gin.Context) {
	feeds := db.SelectFeeds()
	rssFeeds := make([]RssFeed, 0)
	for _, feed := range feeds {
		rssFeeds = append(rssFeeds, RssFeed{
			Id:          feed.ID,
			Title:       feed.Title,
			Description: feed.Description,
		})
	}

	apiResponse := ApiResponse{
		Success: true,
		Data:    rssFeeds,
	}

	c.IndentedJSON(http.StatusOK, apiResponse)
}

func addNewFeed(feedUrl string) (*db.Feed, []db.FeedItem) {
	feed, items := rss.ParseFeedItemsFromXml(feedUrl)
	dbFeed := &db.Feed{
		Title:       feed.Title,
		Description: feed.Description,
		Url:         feed.Url,
		FeedHash:    feed.FeedHash,
	}
	dbFeed = db.InsertFeed(*dbFeed)
	if dbFeed == nil {
		return &db.Feed{}, []db.FeedItem{}
	}

	dbItems := make([]db.FeedItem, 0)
	for _, item := range items {
		dbItems = append(dbItems, db.FeedItem{
			FeedId:      dbFeed.ID,
			Title:       item.Title,
			Link:        item.Link,
			Description: &item.Description,
			PubDate:     item.PubDate,
		})
	}
	dbItems = db.InsertFeedItems(dbItems)
	return dbFeed, dbItems
}

// add feed
func PostAddFeed(c *gin.Context) {
	var request AddFeedRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}

	log.Println("feed url : ", request.Url)
	feedHash := rss.GetFeedHash(request.Url)
	log.Println("feed hash : ", feedHash)

	dbFeed := db.SelectFeedByHash(feedHash)
	var dbItems []db.FeedItem
	if dbFeed != nil {
		log.Println("Feed exists in Database")
		dbItems = db.SelectFeedItemsByPublishedAtDesc(dbFeed.ID, 50)
	} else {
		log.Println("Adding new feed")
		dbFeed, dbItems = addNewFeed(request.Url)
	}

	rssFeed := RssFeed{
		Id:          dbFeed.ID,
		Title:       dbFeed.Title,
		Description: dbFeed.Description,
	}

	rssItems := make([]RssFeedItem, 0)
	for _, item := range dbItems {
		rssItems = append(rssItems, RssFeedItem{
			Id:          item.ID,
			Title:       item.Title,
			Link:        item.Link,
			PublishedAt: item.PubDate.String(),
		})
	}

	addFeedResponse := AddFeedResponse{
		Feed:  rssFeed,
		Items: rssItems,
	}

	apiResponse := ApiResponse{
		Success: true,
		Data:    addFeedResponse,
	}

	c.IndentedJSON(http.StatusOK, apiResponse)
}
