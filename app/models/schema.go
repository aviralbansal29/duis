package models

import (
	"gorm.io/gorm"
)

// SchemaType defines type of Schema for UI
type SchemaType string

// UI component types
const (
	HomePage           SchemaType = "home_page"
	BottomBar          SchemaType = "bottom_bar"
	RecommendationPage SchemaType = "recommendation_page"
)

// Schema stores different schemas
type Schema struct {
	gorm.Model

	Name     string     `json:"name"`
	Type     SchemaType `json:"type"`
	Children []byte     `gorm:"type:jsonb" json:"children"`

	Variants []Variant `gorm:"many2many:schema_mappings" json:"variants"`
}
