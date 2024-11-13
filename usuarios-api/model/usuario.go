package model

type Usuario struct {
	ID       uint   `gorm:"primaryKey"`
	Nombre   string `gorm:"size:255"`
	Email    string `gorm:"size:255;unique"`
	Password string `gorm:"size:255"` // Atributo que no se transfiere

}
