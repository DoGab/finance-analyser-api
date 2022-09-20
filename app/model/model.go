package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model           // Adds some metadata fields to the table
	ID         uuid.UUID `gorm:"type:uuid;primaryKey"` // Explicitly specify the type to be uuid
	Label      string    `json:"label"`
	ParentID   uuid.UUID `gorm:"type:uuid" json:"parent_id"`
	Entries    []Entry   `gorm:"foreignKey:ID" json:"entries"`
}

type Entry struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	ReferenceNr int       `json:"reference_nr"`
	Amount      float32   `json:"amount"`
	Type        string    `json:"type"`
	Saldo       float32   `json:"saldo"`
	ValutaDate  time.Time `json:"valuta_date"`
	ParentID    uuid.UUID `gorm:"type:uuid" json:"parent_id"`
}
