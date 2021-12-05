package service

import "github.com/rizface/sekolah/app/model/response"

type WaliKelas interface {
	GetWaliKelas() ([]response.Walas,[]response.User,[]response.Kelas)
	PostKelas(userId string, kelasId string) string
	DeleteWaliKelas(userId string,kelasId string) string
}
