package types

import (
	"github.com/google/uuid"
)

type Funghi struct {
	ID             uuid.UUID `json:"id" gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	Name           string    `json:"name"`
	ScientificName string    `json:"scientific_name"`
	Edible         bool      `json:"edible"`
}

// func (f *Funghi) BeforeCreate(tx *gorm.DB) (err error) {
// 	f.ID = uuid.New()
// 	return
// }
