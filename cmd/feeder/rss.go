package main

import (
	"fmt"
	"log"

	gofeed "github.com/mmcdole/gofeed"
)

func GetRssFeedItems(feed RssFeed) []RssFeedItem {
	log.Println("fetching rss content for ", feed.Title, feed.Url)

	fp := gofeed.NewParser()
	root, _ := fp.ParseURL(feed.Url)
	fmt.Println("title: ", root.Title)
	fmt.Println("link: ", root.Link)
	fmt.Println("#items: ", len(root.Items))

	var feedItems []RssFeedItem
	for _, item := range root.Items {
		feedItem := &RssFeedItem{
			FeedId:      feed.Id,
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			PubDate:     item.Published,
		}
		feedItems = append(feedItems, *feedItem)
	}
	return feedItems
}
