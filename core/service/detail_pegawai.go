package service

import "github.com/rizface/sekolah/app/model/request"

type DetailPegawai interface {
	Post(userId string, request request.DetailPegawai) string
}
