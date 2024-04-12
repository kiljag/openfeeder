package rss

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func GetDbConn() *sql.DB {

	pgUser := os.Getenv("PG_USER")
	pgPass := os.Getenv("PG_PASS")
	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")

	connString := fmt.Sprintf("postgresql://%s:%s@%s:%s/staging?sslmode=disable", pgUser, pgPass, pgHost, pgPort)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func FetchFeedFromDb(feedUrl string) *RssFeed {
	db := GetDbConn()
	defer db.Close()
	query := "SELECT id, title, url FROM feed where url = %s"
	log.Println("query : ", query)
	rows, err := db.Query(query, feedUrl)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	var feed RssFeed
	for rows.Next() {
		rows.Scan(&feed.Id, &feed.Title, &feed.Url)
	}
	return &feed
}

func FetchFeedsFromDb() []*RssFeed {
	db := GetDbConn()
	defer db.Close()

	query := "SELECT id, title, url, description, feed_hash description FROM feed"
	log.Println("query : ", query)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var feeds []*RssFeed
	var feed RssFeed
	for rows.Next() {
		rows.Scan(&feed.Id, &feed.Title, &feed.Url, &feed.Description, &feed.FeedHash)
		feeds = append(feeds, &feed)
	}
	return feeds
}

func FetchFeedItemsFromDb(feedId int64) []*RssFeedItem {

	db := GetDbConn()
	defer db.Close()

	query := fmt.Sprintf("SELECT id, feed_id, title, link, pubdate, item_hash "+
		"FROM feed_item WHERE feed_id = %d ORDER BY id DESC LIMIT 50", feedId)
	log.Println("query : ", query)
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var items []*RssFeedItem
	var item RssFeedItem
	for rows.Next() {
		rows.Scan(&item.Id, &item.FeedId, &item.Title, &item.Link, &item.PubDate, &item.ItemHash)
		items = append(items, &item)
	}
	return items
}

func FetchFeedItemFromDb(feedId int64, itemId int64) *RssFeedItem {

	db := GetDbConn()
	defer db.Close()

	query := fmt.Sprintf("SELECT id, feed_id, title, description, link, pubdate, item_hash "+
		"FROM feed_item WHERE feed_id = %d and id = %d", feedId, itemId)
	log.Println("query : ", query)
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var item RssFeedItem
	for rows.Next() {
		rows.Scan(&item.Id, &item.FeedId, &item.Title, &item.Description, &item.Link, &item.PubDate, &item.ItemHash)
	}
	return &item
}

func PersistFeedToDb(feed *RssFeed) {

	db := GetDbConn()
	defer db.Close()

	result, err := db.Exec("INSERT INTO feed (title, description, url, feed_hash) VALUES ($1, $2, $3, $4)",
		feed.Title, feed.Description, feed.Url, feed.FeedHash)
	if err != nil {
		log.Println(err)
	}
	id, _ := result.LastInsertId()
	log.Println("Last Insert Id : ", id)
}

func PersistFeedsToDb(feeds []RssFeed) {

	db := GetDbConn()
	defer db.Close()

	for _, feed := range feeds {
		_, err := db.Exec("INSERT INTO feed (title, description, url, feed_hash) VALUES ($1, $2, $3, $4)",
			feed.Title, feed.Description, feed.Url, feed.FeedHash)

		if err != nil {
			log.Println(err)
		}
	}
}

func PersistFeedItemsToDb(items []*RssFeedItem) {

	db := GetDbConn()
	defer db.Close()

	for _, item := range items {
		_, err := db.Exec("INSERT INTO feed_item (feed_id, title, description, link, pubDate, item_hash)"+
			"VALUES ($1, $2, $3, $4, $5, $6)",
			item.FeedId, item.Title, item.Description, item.Link, item.PubDate, item.ItemHash)

		if err != nil {
			log.Println(err)
		}
	}
}
