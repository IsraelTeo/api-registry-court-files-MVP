package model

// Corte Judicial
type Court struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	Name         string  `gorm:"size:150;not null" json:"name"`
	Headquarters string  `gorm:"size:100" json:"headquarters"`
	Judges       []Judge `gorm:"foreignKey:CourtID" json:"judges,omitempty"`
}
