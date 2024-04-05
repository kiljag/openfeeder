package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
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
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var feed RssFeed
	for rows.Next() {
		rows.Scan(&feed.Id, &feed.Title, &feed.Url)
		feeds = append(feeds, feed)
	}
	return feeds
}

// fetch top 50 feed items
func GetFeedItemsFromDatabase(feedId int64) []RssFeedItem {

	db := GetDbConn()
	defer db.Close()

	var items []RssFeedItem
	query := fmt.Sprintf("SELECT id, title, description, link FROM feed_item WHERE feed_id = %d ORDER BY id DESC LIMIT 50", feedId)
	fmt.Println("query : ", query)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var item RssFeedItem
	for rows.Next() {
		rows.Scan(&item.Id, &item.Title, &item.Description, &item.Link)
		items = append(items, item)
	}
	return items
}

func GetFeedItemContentFromDatabase(feedId int64, feedItemId int64) RssFeedItem {
	db := GetDbConn()
	defer db.Close()

	query := fmt.Sprintf("SELECT id, title, description, link FROM feed_item WHERE feed_id = %d and id = %d", feedId, feedItemId)
	fmt.Println("query : ", query)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var item RssFeedItem
	for rows.Next() {
		rows.Scan(&item.Id, &item.Title, &item.Description, &item.Link)
	}
	return item
}
