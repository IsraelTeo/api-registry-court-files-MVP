package model

// Abogado
type Lawyer struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FullName  string `json:"full_name"`
	BarNumber string `json:"bar_number"`
	Email     string `json:"email"`
}
