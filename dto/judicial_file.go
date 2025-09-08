package dto

import (
	"time"
)

// Expediente Judicial
type JudicialFile struct {
	ID                 uint       `json:"id"`
	FileNumber         string     `json:"file_number"`
	NotificationNumber string     `json:"notification_number"`
	DigitizationNumber string     `json:"digitization_number"`
	DocumentType       string     `json:"document_type"`
	NotificationDate   *time.Time `json:"notification_date"`
	CourtID            uint       `json:"court_id"`
	PersonsIDs         []uint     `json:"persons"`
	LawyersIDs         []uint     `json:"lawyers"`
}
