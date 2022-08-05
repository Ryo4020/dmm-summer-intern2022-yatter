package accounts

import (
	"net/url"
	"strconv"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/query"
)

// Request quey for `GET /v1/accounts/{username}/following``
type FollowingQuery struct {
	Limit int
}

// Set request query to struct
func (q *FollowingQuery) setQuery(v url.Values) error {
	l := query.GetQueryNumInStr(v, "limit", "40")

	var err error
	q.Limit, err = strconv.Atoi(l)
	if err != nil {
		return err
	}

	return nil
}

// Request quey for `GET /v1/accounts/{username}/followers`
type FollowersQuery struct {
	MaxID   object.AccountID
	SinceID object.AccountID
	Limit   int
}

func (q *FollowersQuery) setQuery(v url.Values) error {
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

// Request quey for `GET /v1/accounts/relationships``
type RelationshipsQuery struct {
	Username string
}

func (q *RelationshipsQuery) setQuery(v url.Values) {
	q.Username = v.Get("username")
}
