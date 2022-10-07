package models

import "gorm.io/gorm"

type Variant struct {
	gorm.Model

	Name        string `json:"name"`
	Description string `json:"description"`

	Users   []User   `gorm:"foreignKey:VariantID" json:"users"`
	Schemas []Schema `gorm:"many2many:schema_mappings" json:"schemas"`
}
