package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // add this
)

const CONN_STRING = "postgresql://postgres:mysecretpassword@localhost/staging?sslmode=disable"

func GetDbConn() *sql.DB {
	db, err := sql.Open("postgres", CONN_STRING)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetFeedsFromDatabase() []RssFeed {

	db := GetDbConn()
	defer db.Close()

	var feeds []RssFeed
	rows, err := db.Query("SELECT id, title, url FROM feed")
	defer rows.Close()

	if err != nil {
		log.Fatalln(err)
	}

	var feed RssFeed
	for rows.Next() {
		rows.Scan(&feed.Id, &feed.Title, &feed.Url)
		feeds = append(feeds, feed)
	}
	return feeds
}

func SaveFeedItemsToDatabase(items []RssFeedItem) {

	db := GetDbConn()
	defer db.Close()

	for _, item := range items {
		fmt.Println("feed id : ", item.FeedId)
		fmt.Println("title : ", item.Title)
		fmt.Println("link : ", item.Link)
		_, err := db.Exec("INSERT INTO feed_item (feed_id, title, description, link, pubDate) VALUES ($1, $2, $3, $4, $5)",
			item.FeedId, item.Title, item.Description, item.Link, item.PubDate)

		if err != nil {
			log.Fatalln(err)
		}
	}
}
