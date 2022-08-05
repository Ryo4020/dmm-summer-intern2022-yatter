package accounts

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/handler/auth"
	"yatter-backend-go/app/handler/httperror"
)

// Handle request for `GET /v1/accounts/relationships`
func (h *handler) Relationships(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	follower := auth.AccountOf(r)

	var q RelationshipsQuery
	q.setQuery(r.URL.Query())

	a := h.app.Dao.Account()
	followee, err := a.FindByUsername(ctx, q.Username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	f := h.app.Dao.Follow() // domain/repository の取得
	relation, err := f.FindRelation(ctx, *follower, *followee)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(relation); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
