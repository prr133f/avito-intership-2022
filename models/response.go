package models

type Error struct {
	Text string `json:"text"`
}

func (e Error) String() string {
	return e.Text
}

type Balance struct {
	Balance float64 `json:"balance"`
}

type ReservedBalance struct {
	ReservedBalance float64 `json:"reserved_balance"`
}

type MonthlyReport struct {
	ServiceName string  `json:"service_name"`
	Margin      float64 `json:"margin"`
}
