package repository

import (
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"gorm.io/gorm"
	"strconv"
)

type spp struct{}

func NewSpp() Spp {
	return &spp{}
}

func (s *spp) GetStudents(db *gorm.DB) ([]response.SppStudents, error) {
	var level app.Level
	var students []response.SppStudents
	db.Where("level = ?", "murid").First(&level)
	result := db.Select("users.id,users.nama_depan,users.nama_belakang,username, CASE WHEN monthly_fees.student_id IS NULL THEN 'Belum Bayar' ELSE 'Sudah Bayar' END AS status").Joins("LEFT JOIN monthly_fees ON monthly_fees.student_id = users.id").Where(" users.level_id = ? AND EXTRACT(month FROM monthly_fees.created_at) = EXTRACT(month FROM NOW()) OR monthly_fees.created_at IS NULL AND users.level_id = ?",level.Id,level.Id).Find(&[]app.User{}).Scan(&students)
	return students,result.Error
}

func (s *spp) GetByUserId(db *gorm.DB, userId string) (app.MonthlyFee,error) {
	var student app.MonthlyFee
	result := db.Where("student_id = ?", userId).Last(&student)
	return student,result.Error
}

func (s *spp) GetDetailSpp(db *gorm.DB, userId string) ([]response.DetailSpp, error) {
	var detail []response.DetailSpp
	result := db.Select("CONCAT(s.nama_depan,' ',s.nama_belakang) AS nama_siswa, CONCAT(o.nama_depan,' ',o.nama_belakang) AS nama_petugas, TO_CHAR(monthly_fees.created_at - INTERVAL '1 Month','dd Monthyyyy') AS tgl_bayar, TO_CHAR(monthly_fees.created_at,'dd Monthyyyy') AS tgl_lunas").Joins("INNER JOIN users s ON s.id = monthly_fees.student_id INNER JOIN users o  ON o.id = monthly_fees.officer_id").Where("monthly_fees.student_id = ?", userId).Find(&[]app.MonthlyFee{}).Scan(&detail)
	return detail,result.Error
}

func (s *spp) PostSpp(db *gorm.DB, request request.Spp) (bool, error) {
	nominal,_ := strconv.Atoi(request.Nominal)
	result := db.Create(&app.MonthlyFee{
		StudentId: request.StudentId,
		OfficerId: request.EmployeeId,
		Amount:    nominal,
		CreatedAt: request.CreatedAt,
	})
	return result.RowsAffected > 0, result.Error
}
