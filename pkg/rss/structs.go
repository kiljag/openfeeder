package rss

import "time"

type RssFeed struct {
	Id          int64
	Title       string
	Url         string
	Description string
	FeedHash    string
}

type RssFeedItem struct {
	Id          int64
	FeedId      int64
	Title       string
	Description string
	Link        string
	PubDate     string
	Timestamp   time.Time
	IsViewed    bool
	ItemHash    string
}
