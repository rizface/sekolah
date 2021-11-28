package service

import (
	"github.com/rizface/sekolah/app/model/response"
	"github.com/rizface/sekolah/app/request"
)

type Login interface {
	Login(login request.Login) response.Login
}
