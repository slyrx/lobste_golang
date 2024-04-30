package models

import "time"

const ActiveStoriesPerPage = 25

type Story struct {
	ID            int64
	Title         string
	URL           string
	Description   string
	Score         int
	CommentsCount int
	Upvotes       int
	Downvotes     int
	Submitter     User
	SubmittedAt   time.Time
}
