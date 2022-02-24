package models

import (
	"database/sql"
)

// Models is the wrapper for database
type Models struct {
	DB DBModel
}

// NewModels returns models with db pool
func NewModels(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

// User is the type for users
type User struct {
	ID         int
	Username   string
	EmployeeID string
}

// Create Menu Request
type Menu struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Memo       string `json:"memo"`
	FileString string `json:"fileString"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Opened     string `json:"opened"`
}

// Add Order Request
type Order struct {
	ID        int    `json:"id"`
	MenuID    int    `json:"menuId"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Item      string `json:"item"`
	Sugar     string `json:"sugar"`
	Ice       string `json:"ice"`
	Price     string `json:"price"`
	UserMemo  string `json:"memo"`
	UpdatedAt string `json:"updated_at"`
	User      string `json:"user"`
	Count     string `json:"count"`
}
