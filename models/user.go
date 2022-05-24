package models

type Users struct {
	ID     string `gorm:"size:25;not null; uniqueIndex;primaryKey"`
	Name   string `gorm:"size:100;not null"`
	Email  string `gorm:"size:100;not null"`
	Phone  string `gorm:"size:100;not null"`
	Alamat string `gorm:"size:100;not null"`
}
