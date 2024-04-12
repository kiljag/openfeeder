package main

type ApiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type RssFeed struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type RssFeedItem struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Link        string `json:"link"`
	PublishedAt string `json:"publishedAt"`
}

type AddFeedRequest struct {
	Url string `json:"url"`
}

type AddFeedResponse struct {
	Feed  RssFeed       `json:"feed"`
	Items []RssFeedItem `json:"items"`
}
