package model

type Enrollment struct {
	Id_enrollment uint   `gorm:"primaryKey"`
	Id            string `json:"id" gorm:"size:255"`
	Id_cursos     string `json:"curso_id" gorm:"size:255"`
}
