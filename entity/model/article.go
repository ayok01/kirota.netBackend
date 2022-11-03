package model

import "time"

type Article struct {
	IllustrationImg string
	Title string
	CreateDate  time.Time
	Description string
}