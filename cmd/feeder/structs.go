package main

type RssFeed struct {
	Id          int64
	Title       string
	Url         string
	Description string
}

type RssFeedItem struct {
	Id          int64
	FeedId      int64
	Title       string
	Description string
	Link        string
	PubDate     string
}
