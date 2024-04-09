package rss

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/mmcdole/gofeed"
)

func computeHash(content string) string {

	hasher := md5.New()
	hasher.Write([]byte(content))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func ParseFeedItemsFromXml(feedUrl string) (RssFeed, []RssFeedItem) {
	log.Println("fetching feed items from rss url ", feedUrl)

	fp := gofeed.NewParser()
	root, _ := fp.ParseURL(feedUrl)

	log.Printf("title: %s, link: %s, #items: %d\n", root.Title, root.Link, len(root.Items))

	feedHash := computeHash(fmt.Sprintf("%s:%s", feedUrl, root.Title))
	feed := RssFeed{
		Title:       root.Title,
		Url:         feedUrl,
		Description: root.Description,
		FeedHash:    feedHash,
	}

	var feedItems []RssFeedItem
	for _, item := range root.Items {
		itemHash := computeHash(fmt.Sprintf("%s:%s:%s", item.Title, item.Link, item.Published))
		feedItem := &RssFeedItem{
			Title:       item.Title,
			Description: item.Description,
			Link:        item.Link,
			PubDate:     item.Published,
			ItemHash:    itemHash,
		}
		feedItems = append(feedItems, *feedItem)
	}
	return feed, feedItems
}
