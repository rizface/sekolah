package service

import (
	request2 "github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
)

type Spp interface {
	GetStudents() []response.SppStudents
	GetDetailSpp(userId string) []response.DetailSpp
	PostSpp(request request2.Spp) string
}
