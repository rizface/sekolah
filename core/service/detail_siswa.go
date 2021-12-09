package service

import "github.com/rizface/sekolah/app/model/request"

type DetailSiswa interface{
	Update(userId string,request request.DetailSiswa) string
}
