package service

import (
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
)

type AdminCrud interface{
	Get(level string) []response.User
	GetById(userId interface{}) response.User
	Post(request request.User, level string) string
	Delete(userId interface{}) string
	Update(userId interface{}, request request.User) string
}
