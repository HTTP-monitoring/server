package model

type URL struct {
	ID       int
	UserID   int
	URL      string
	Period   int
	Statuses []Status
}
