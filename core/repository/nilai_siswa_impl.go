package repository

import (
	"fmt"
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
)

type nilaiSiswa struct{}

func NewNilaiSiswa() NilaiSiswa {
	return nilaiSiswa{}
}

func (n nilaiSiswa) Post(db *gorm.DB, request request.Nilai) (bool, error) {
	result := db.Create(&app.Grade{
		StudentId:  request.StudentId,
		TeacherId:  request.TeacherId,
		SemesterId: request.SemesterId,
		SubjectId:  request.SubjectId,
		Grade:      request.Grade,
	})
	return result.RowsAffected > 0, result.Error
}

func (n nilaiSiswa) GetAllGrade(db *gorm.DB, userId string) ([]response.Nilai, error) {
	var grade []response.Nilai
	result := db.Select("grades.id,subjects.subject,semesters.semester,grades.grade, TO_CHAR(grades.created_at :: DATE,'dd Monthyyyy') AS date").Order("grades.created_at DESC").Joins("INNER JOIN subjects ON subjects.id = grades.subject_id").Joins("INNER JOIN semesters ON semesters.id = grades.semester_id").Where("grades.student_id = ?",userId).Find(&[]app.Grade{}).Scan(&grade)
	return grade,result.Error
}

func (n nilaiSiswa) GetGradeBySemester(db *gorm.DB, userId string,semester string) ([]response.Nilai, error) {
	var sData app.Semester
	db.Where("semester = ?",semester).First(&sData)
	fmt.Println("userId :",userId,"semester :", sData.Semester)
	var grade []response.Nilai
	result := db.Select("grades.id,subjects.subject,semesters.semester,grades.grade, TO_CHAR(grades.created_at :: DATE,'dd Monthyyyy') AS date").Order("grades.created_at DESC").Joins("INNER JOIN subjects ON subjects.id = grades.subject_id").Joins("INNER JOIN semesters ON semesters.id = grades.semester_id").Where("grades.student_id = ? AND grades.semester_id = ?",userId,sData.Id).Find(&[]app.Grade{}).Scan(&grade)
	return grade,result.Error
}

func (n nilaiSiswa) GetGradeById(db *gorm.DB, gradeId string) (response.Nilai, error) {
	var grade response.Nilai
	result := db.Where("id = ?",gradeId).Select("id,semester_id AS semester,grade, subject_id AS subject").First(&app.Grade{}).Scan(&grade)
	return grade,result.Error
}

func (n nilaiSiswa) Delete(db *gorm.DB, gradeId string) (bool, error) {
	result := db.Where("id = ?", gradeId).Delete(&app.Grade{})
	return result.RowsAffected > 0 , result.Error
}

func (n nilaiSiswa) Update(db *gorm.DB, gradeId string, request request.Nilai) (bool, error) {
	result := db.Where("id = ?", gradeId).Updates(&app.Grade{
		SemesterId: request.SemesterId,
		SubjectId:  request.SubjectId,
		Grade:      request.Grade,
	})
	return result.RowsAffected > 0, result.Error
}


