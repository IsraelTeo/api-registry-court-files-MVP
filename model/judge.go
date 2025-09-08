package model

// Juez
type Judge struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FullName  string `gorm:"size:200;not null" json:"full_name"`
	Specialty string `gorm:"size:100" json:"specialty"`
	CourtID   uint   `json:"court_id"`
}

type Judges []Judge
