package service

import (
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
)

type Kelas interface {
	Get() []response.Kelas
	GetById(kelasId string) response.Kelas
	GetStudent(kelasId string) ([]response.User,[]response.User)
	Post(request request.Kelas) string
	AddStudent(kelasId string, userId string) string
	Update(kelasId string, request request.Kelas) string
	Delete(kelasId string) string
	DeleteStudent(kelasId string, userId string) string
}
