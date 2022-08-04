package object

type (
	StatusID = int64

	// Status status
	Status struct {
		// The internal ID of the status
		ID StatusID `json:"-"`

		// The account ID of the author of the status
		AccountID AccountID `json:"-" db:"account_id"`

		// The account of the author of the status
		*Account `json:"account" db:"account"`

		// The text of the status
		Content string `json:"content,omitempty"`

		// The time the status was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}

	DBStatus struct {
		AccountID AccountID `db:"account_id"`

		Content string

		CreateStatusAt DateTime `db:"create_s_at"`
	}
)
