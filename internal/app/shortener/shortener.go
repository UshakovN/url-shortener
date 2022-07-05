package shortener

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

func GenerateShortUrl() string {
	runes := []rune(alphabet)
	rand.Seed(time.Now().UnixNano())
	url := make([]rune, 10)
	for i := range url {
		url[i] = runes[rand.Intn(len(alphabet))]
	}
	return string(url)
}
