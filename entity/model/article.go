package model

import "time"

type Article struct {
	IllustrationImg string
	Title string
	Date  time.Time
	Description string
}