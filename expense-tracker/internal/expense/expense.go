package expense

import "time"

// Expense is the struct to handle a single expense tracked.
type Expense struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
	Id          int       `json:"id"`
	Amount      float64   `json:"amount"`
}

// ExpenseStorage is the struct used to handle the serialization and saving of
// expenses.
type ExpenseStorage struct {
	Expenses []Expense `json:"expenses"`

	// Metadata for the storage. This helps to keep next ID to use for the record
	// since I don't have any other way to keep track of the used and next IDs.
	Metadata struct {
		NextId int `json:"nextId"`
	} `json:"metadata"`
}
