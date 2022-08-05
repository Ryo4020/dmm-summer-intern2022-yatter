package timelines

import (
	"net/url"
	"strconv"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/query"
)

// Request query for `GET /v1/timelines/home` & `GET /v1/timelines/public`
type TimelineQuery struct {
	OnlyMedia string
	MaxID     object.AccountID
	SinceID   object.AccountID
	Limit     int
}

// Set request query to struct
func (q *TimelineQuery) setQuery(v url.Values) error {
	q.OnlyMedia = v.Get("only_media")

	m := query.GetQueryNumInStr(v, "max_id", "0")
	s := query.GetQueryNumInStr(v, "since_id", "0")
	l := query.GetQueryNumInStr(v, "limit", "40")

	var err error
	q.MaxID, err = strconv.ParseInt(m, 10, 64)
	if err != nil {
		return err
	}
	q.SinceID, err = strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	q.Limit, err = strconv.Atoi(l)
	if err != nil {
		return err
	}

	return nil
}
