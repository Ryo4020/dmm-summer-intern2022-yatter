package timelines

import (
	"net/url"
)

// Request quey for `POST /v1/timelines/public`
type PublicQuery struct {
	OnlyMedia string
	MaxID     string
	SinceID   string
	Limit     string
}

// Set request query to struct
func (q *PublicQuery) setQuery(v url.Values) {
	q.OnlyMedia = v.Get("only_media")
	q.MaxID = v.Get("max_id")
	q.SinceID = v.Get("since_id")
	q.Limit = v.Get("limit")
}
