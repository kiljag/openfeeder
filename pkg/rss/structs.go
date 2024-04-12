package rss

import "time"

type RssFeed struct {
	Id          uint64
	Title       string
	Url         string
	Description string
	FeedHash    string
}

type RssFeedItem struct {
	Id          uint64
	FeedId      uint64
	Title       string
	Description string
	Link        string
	PubDate     time.Time
	IsViewed    bool
	ItemHash    string
}
