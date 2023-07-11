package models

type Error struct {
	Text string `json:"text"`
}

func (e Error) String() string {
	return e.Text
}

type GetBalance struct {
	Balance float64 `json:"balance"`
}

type AddBalance struct {
	NewBalance float64 `json:"new_balance"`
}

type ReserveBalance struct {
	ReservedBalance float64 `json:"reserved_balance"`
}
