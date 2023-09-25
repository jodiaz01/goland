package models

import "gorm.io/gorm"

type Categoria struct {
	gorm.Model
	ID          int64  `gorm:"not null; unique_index"`
	Codigo      string `gorm:"not null; unique_index"`
	Categoria   string `gorm:"not null; unique_index"`
	Descripcion string
	Estado      bool `gorm:"default:false"`
}
