package domain

import "time"

type Book struct {
	Base
	Title string
	Author string
}

type Author struct {
	Base
	Name string
}

type Base struct {
	ID        uint `gorm:"primary_key" sql:"AUTO_INCREMENT"`
	CreatedAt time.Time `sql:"not null;type:date"`
	UpdatedAt time.Time `sql:"not null;type:date"`
}

type Books []Book