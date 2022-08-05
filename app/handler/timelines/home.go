package timelines

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"
)

// Handle request for `GET /v1/timelines/home`
func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	follower := auth.AccountOf(r)

	var q TimelineQuery
	if err := q.setQuery(r.URL.Query()); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	f := h.app.Dao.Follow() // domain/repository の取得
	// フォローしているアカウントリストの取得
	accounts, err := f.GetFollowing(ctx, *follower, q.Limit)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	statuses := make([]*object.Status, 0)
	if len(accounts) > 0 {
		t := h.app.Dao.Timeline() // domain/repository の取得
		statuses, err = t.GetHome(ctx, accounts)
		if err != nil {
			httperror.InternalServerError(w, err)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(statuses); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
