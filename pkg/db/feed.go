package db

import (
	"log"
	"time"
)

type Feed struct {
	ID          uint64
	Title       string
	Url         string
	Description string
	CreatedAt   time.Time
	FeedHash    string
}

func (Feed) TableName() string {
	return "feed"
}

func SelectFeeds() []Feed {
	db := GetPGConn()

	var feeds []Feed
	db.Find(&feeds)
	return feeds
}

func SelectFeedById(id uint64) *Feed {
	db := GetPGConn()

	var feed Feed
	if err := db.First(&feed, id).Error; err != nil {
		return nil
	} else {
		return &feed
	}
}

func SelectFeedByHash(feedHash string) *Feed {

	db := GetPGConn()

	var feed Feed
	if err := db.First(&feed, "feed_hash = ?", feedHash).Error; err != nil {
		return nil
	} else {
		return &feed
	}
}

func InsertFeed(feed Feed) *Feed {
	db := GetPGConn()
	if err := db.Create(&feed).Error; err != nil {
		log.Println(err)
		return nil
	} else {
		return &feed
	}
}

func InsertFeeds(feeds []Feed) []Feed {
	db := GetPGConn()
	if err := db.Create(&feeds).Error; err != nil {
		log.Println(err)
		return nil
	} else {
		return feeds
	}
}
