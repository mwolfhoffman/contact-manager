package main

type Contact struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string
	Phone string
}
