package object

type (
	StatusID = int64

	// Status status
	Status struct {
		// The internal ID of the status
		ID StatusID `json:"-"`

		AccountID AccountID `json:"-" db:"account_id"`

		// Biography of user
		Content string `json:"content,omitempty"`

		// The time the status was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)
