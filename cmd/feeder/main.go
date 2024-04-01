package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is feeder tool")
	feeds := GetFeedsFromDatabase()
	for _, feed := range feeds {
		fmt.Println("\n\n")
		fmt.Println(feed.Id, feed.Title, feed.Url)
		feedItems := GetRssFeedItems(feed)
		SaveFeedItemsToDatabase(feedItems)
	}
}
