package service

import (
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
)

type Mapel interface {
	Get() []response.Mapel
	GetById(mapelId string) response.Mapel
	GetPengampu(mapelId string) ([]response.User,[]response.User)
	Post(request request.Mapel) string
	PostPengampu(mapelId string, userId string) string
	Delete(mapelId string) string
	DeletePengampu(mapelId string, userId string) string
	Update(mapelId string,request request.Mapel) string
}
