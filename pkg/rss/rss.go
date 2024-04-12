package rss

import (
	"fmt"
	"log"
	"time"

	"github.com/mmcdole/gofeed"
)

func parsePublishedTime(pubDate string) time.Time {
	layout := "Mon, 02 Jan 2006 15:04:05"
	t, err := time.Parse(layout, pubDate)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return t
	}
	return time.Now()
}

func ParseFeedItemsFromXml(feedUrl string) (*RssFeed, []*RssFeedItem) {
	log.Println("fetching feed items from rss url ", feedUrl)

	fp := gofeed.NewParser()
	root, err := fp.ParseURL(feedUrl)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	log.Printf("title: %s, link: %s, #items: %d\n", root.Title, root.Link, len(root.Items))

	feedHash := GetFeedHash(feedUrl)
	feed := RssFeed{
		Title:       root.Title,
		Url:         feedUrl,
		Description: root.Description,
		FeedHash:    feedHash,
	}

	var feedItems []*RssFeedItem
	for _, item := range root.Items {
		itemHash := GetFeedItemHash(item.Title, item.Description, item.Link)
		feedItem := &RssFeedItem{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			PubDate:     parsePublishedTime(item.Published),
			ItemHash:    itemHash,
		}
		feedItems = append(feedItems, feedItem)
	}
	return &feed, feedItems
}
