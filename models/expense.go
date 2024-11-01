// models/expense.go
package models

import "time"

type Expense struct {
	ID          int       `json:"id"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	PaidBy      int       `json:"paid_by"`
	GroupID     int       `json:"group_id"`
	CreatedAt   time.Time `json:"created_at"`
}
