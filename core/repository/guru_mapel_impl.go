package repository

import (
	"fmt"
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type guruMapel struct{}

func NewGuruMapel() GuruMapel {
	return &guruMapel{}
}

func (g *guruMapel) GetMapel(db *gorm.DB, guruId interface{}) ([]response.Mapel, error) {
	var subjects []response.Mapel
	fmt.Println(guruId)
	result := db.Select("subjects.id AS id,subjects.subject AS subject").Where("users.id = ?",guruId).Joins("INNER JOIN subjects ON subjects.id = teacher_subjects.subject_id INNER JOIN users ON teacher_subjects.user_id = users.id").Find(&[]app.TeacherSubject{}).Scan(&subjects)
	return subjects,result.Error
}
