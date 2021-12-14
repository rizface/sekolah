package controller

import (
	"github.com/rizface/sekolah/core/service"
	"github.com/rizface/sekolah/helper"
	"net/http"
)

type guru struct{
	kelasService service.Kelas
}

func NewGuru(kelasService service.Kelas) Dashboard {
	return &guru{
		kelasService: kelasService,
	}
}

func (g *guru) Dashboard(w http.ResponseWriter, r *http.Request) {
	kelas := g.kelasService.Get()
	tmp := helper.ViewGuru("view/guru/index.gohtml")
	err := tmp.Exec(w,r,"guru",map[string]interface{} {
		"classes": kelas,
	})
	helper.PanicIfError(err)
}


