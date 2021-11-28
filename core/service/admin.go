package service

import (
	"github.com/rizface/sekolah/app/model/response"
	"github.com/rizface/sekolah/app/request"
)

type AdminCrudAdmin interface{
	Get() []response.Admin
	GetById(adminId interface{}) response.Admin
	Post(request request.Admin) string
	Delete(adminId interface{}) string
	Update(adminId interface{}, request request.Admin) string
}
