package repository

import (
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type absen struct{}

func NewAbsen() Absen {
	return &absen{}
}

func (a *absen) GetStudentAbsent(db *gorm.DB, siswaId string) ([]response.Absen, error) {
	var absents []response.Absen
	result := db.Select("absen_students.id, subjects.subject AS mapel, absen_students.status AS keterangan, TO_CHAR(absen_students.created_at :: DATE, 'dd Monthyyyy') AS tgl").Joins("INNER JOIN subjects ON subjects.id = absen_students.subject_id").Where(" absen_students.student_id = ?",siswaId).Find(&[]app.AbsenStudent{}).Scan(&absents)
	return absents,result.Error
}

func (a *absen) Absen(db *gorm.DB,request request.Absen) error {
	result := db.Create(&app.AbsenStudent{
		SubjectId: request.SubjectId,
		StudentId: request.Id,
		Status:    request.Key,
	})
	return result.Error
}

