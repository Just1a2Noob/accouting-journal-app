package api

import (
	"time"

	"github.com/google/uuid"
)

type UserDB struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Username string    `json:"username,omitempty"`
	Password string    `json:"password,omitempty"`
	Role     string    `json:"role,omitempty"`
	Section  string    `json:"section,omitempty"`
	Email    string    `json:"email,omitempty"`
}

type ExpenseType struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
}

type IncomeType struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
}

type ExpensesDB struct {
	ID               uuid.UUID `json:"id,omitempty"`
	ExpenseID        uuid.UUID `json:"expense_id,omitempty"`
	Total_expense    int64     `json:"total_expense,omitempty"`
	Transaction_date time.Time `json:"transaction_date,omitempty"`
	Description      string    `json:"description,omitempty"`
	Created_at       time.Time `json:"created_at,omitempty"`
	Updated_at       time.Time `json:"updated_at,omitempty"`
	UserID           uuid.UUID `json:"user_id,omitempty"`
}

type IncomesDB struct {
	ID               uuid.UUID `json:"id,omitempty"`
	IncomeDB         uuid.UUID `json:"income_db,omitempty"`
	Total_income     int64     `json:"total_expense,omitempty"`
	Transaction_date time.Time `json:"transaction_date,omitempty"`
	Description      string    `json:"description,omitempty"`
	Created_at       time.Time `json:"created_at,omitempty"`
	Updated_at       time.Time `json:"updated_at,omitempty"`
	UserID           uuid.UUID `json:"user_id,omitempty"`
}

type UserResponse struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Section  string `json:"section,omitempty"`
	Roles    string `json:"roles,omitempty"`
}
