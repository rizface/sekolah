package service

import "github.com/rizface/sekolah/app/model/response"

type GuruMapel interface {
	GetMapel(guruId interface{}) []response.Mapel
}
