package db

import (
	"log"
	"time"
)

type FeedItem struct {
	ID          uint64
	FeedId      uint64
	Title       string
	Link        string
	Description *string
	PubDate     time.Time
	CreatedAt   time.Time
	ItemHash    string
}

func (FeedItem) TableName() string {
	return "feed_item"
}

func FindFeedItemById(id uint64) *FeedItem {
	db := GetPGConn()

	var item FeedItem
	if err := db.First(&item, id).Error; err != nil {
		log.Println(err)
		return nil
	} else {
		return &item
	}
}

func SelectFeedItem(feedId uint64, itemId uint64) *FeedItem {
	db := GetPGConn()
	var item FeedItem
	if err := db.Where("feed_id = ? AND id = ?", feedId, itemId).First(&item).Error; err != nil {
		log.Println(err)
		return nil
	} else {
		return &item
	}
}

func SelectFeedItemsById(feedId uint64) []FeedItem {
	db := GetPGConn()
	var items []FeedItem
	db.Where("feed_id = ?", feedId).Find(&items)
	return items
}

func SelectFeedItemsByPublishedAtDesc(feedId uint64, limit int) []FeedItem {
	db := GetPGConn()
	var items []FeedItem
	db.Where("feed_id = ?", feedId).Order("pub_date desc").Limit(100).Find(&items)
	return items
}

func InsertFeedItems(items []FeedItem) []FeedItem {
	db := GetPGConn()

	if err := db.Create(&items).Error; err != nil {
		log.Println(err)
		return nil
	} else {
		return items
	}
}

func DeleteFeedItemsById(feedId uint64) bool {
	db := GetPGConn()
	if err := db.Where("feed_id = ?", feedId).Delete(&FeedItem{}).Error; err != nil {
		return false
	} else {
		return true
	}
}
