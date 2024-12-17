package model

type Usuario struct {
	ID       uint   `gorm:"primaryKey"`
	Nombre   string `gorm:"size:255"`
	Email    string `gorm:"size:255"`
	Password string `gorm:"size:255"`
	Admin    bool   `gorm:"default:false"`
}
