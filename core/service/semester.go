package service

import "github.com/rizface/sekolah/app/model/response"

type Semester interface {
	GetSemester() []response.Semester
}