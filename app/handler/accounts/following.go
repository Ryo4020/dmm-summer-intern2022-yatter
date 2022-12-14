package accounts

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

// Handle request for `GET /v1/accounts/{username}/following`
func (h *handler) Following(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := chi.URLParam(r, "username")

	a := h.app.Dao.Account() // domain/repository の取得
	follower, err := a.FindByUsername(ctx, username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	var q FollowingQuery
	if err := q.setQuery(r.URL.Query()); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	f := h.app.Dao.Follow() // domain/repository の取得
	// フォローしているアカウントリストの取得
	accounts, err := f.GetFollowing(ctx, *follower, 0, 0, q.Limit)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(accounts); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
