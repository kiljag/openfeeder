package rss

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func computeHash(content string) string {

	hasher := md5.New()
	hasher.Write([]byte(content))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func GetFeedHash(url string) string {
	return computeHash(url)
}

func GetFeedItemHash(title, description, link string) string {
	return computeHash(fmt.Sprintf("%s:%s:%s", title, description, link))
}
