package movements

import "gorm.io/gorm"

type MovementType struct {
	gorm.Model
	Name         string `gorm:"unique; not null"`
	Abbreviation string `gorm:"unique; not null"`
}

type Movement struct {
	gorm.Model
	Name           string `gorm:"unique; not null"`
	Abbreviation   string `gorm:"unique; not null"`
	ImageUrl       string
	MovementTypeID uint
}
