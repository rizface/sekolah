package service

import (
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
)

type Kelas interface {
	Get() []response.Kelas
	GetById(kelasId string) response.Kelas
	Post(request request.Kelas) string
	Update(kelasId string, request request.Kelas) string
	Delete(kelasId string) string
}
