package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name  string `json:"name"`
	Email string `gorm:"index" json:"email"`

	VariantID uint    `json:"variant_id"`
	Variant   Variant `gorm:"foreignKey:VariantID;references:ID"`
}
