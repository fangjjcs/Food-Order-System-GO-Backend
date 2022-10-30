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
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	Memo            string  `json:"memo"`
	FileString      string  `json:"fileString"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
	Opened          bool    `json:"opened"`
	Rating          float64 `json:"rating"`
	TotalVoter      int     `json:"totalVoter"`
	OrderCount      int     `json:"orderCount"`
	OrderTotalPrice int     `json:"orderTotalPrice"`
}

type OpenedMenu struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	Memo            string `json:"memo"`
	FileString      string `json:"fileString"`
	CreatedAt       string `json:"createdAt"`
	UpdatedAt       string `json:"updatedAt"`
	CloseAt         string `json:"closeAt"`
	Opened          bool   `json:"opened"`
	OrderCount      int    `json:"orderCount"`
	OrderTotalPrice int    `json:"orderTotalPrice"`
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
	UpdatedAt string `json:"updatedAt"`
	User      string `json:"user"`
	Count     string `json:"count"`
}

type Rating struct {
	ID         int     `json:"id"`
	Rating     float64 `json:"rating"`
	TotalVoter int     `json:"totalVoter"`
}
