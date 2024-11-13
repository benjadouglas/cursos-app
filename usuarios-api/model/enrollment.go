package model

type Enrollment struct {
	Id_enrollment uint   `gorm:"primaryKey"`
	Id            string `gorm:"size:255"`
	Id_cursos     string `gorm:"size:255"`
}
