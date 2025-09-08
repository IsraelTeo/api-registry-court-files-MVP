package model

// Persona que puede estar relacionada con un expediente judicial
type Person struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

type Persons []Person
