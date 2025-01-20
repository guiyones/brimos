package personalfinances

import "time"

type Expenses struct {
	ID          string
	Description string
	Value       float64
	DueDate     time.Time
	PayDate     time.Time
	Type        string
	SubType     string
	Status      bool
}
