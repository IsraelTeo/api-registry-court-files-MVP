package model

import "time"

// Expediente Judicial
type JudicialFile struct {
	ID                 uint       `gorm:"primaryKey" json:"id"`
	FileNumber         string     `gorm:"size:100;not null" json:"file_number"`
	NotificationNumber string     `gorm:"size:100" json:"notification_number"`
	DigitizationNumber string     `gorm:"size:100" json:"digitization_number"`
	DocumentType       string     `gorm:"size:50" json:"document_type"`
	Headquarters       string     `gorm:"size:100" json:"headquarters"`
	Court              string     `gorm:"size:150" json:"court"`
	NotificationDate   *time.Time `json:"notification_date"`
	CreationDate       time.Time  `gorm:"autoCreateTime" json:"creation_date"`
	UpdateDate         time.Time  `gorm:"autoUpdateTime" json:"update_date"`
	CourtID            uint       `json:"court_id"`

	// relaciones
	Persons `gorm:"many2many:judicial_file_persons;constraint:OnDelete:CASCADE;" json:"persons"`
	Lawyers `gorm:"many2many:judicial_file_lawyers;constraint:OnDelete:CASCADE;" json:"lawyers"`
}

type JudicialFiles []JudicialFile
