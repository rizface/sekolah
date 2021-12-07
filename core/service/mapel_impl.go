package service

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/rizface/sekolah/app/model/request"
	"github.com/rizface/sekolah/app/model/response"
	"github.com/rizface/sekolah/core/repository"
	"github.com/rizface/sekolah/helper"
	"gorm.io/gorm"
	"sync"
)

type mapel struct {
	db       *gorm.DB
	validate *validator.Validate
	repo     repository.Mapel
}

func NewMapel(db *gorm.DB, validate *validator.Validate, repo repository.Mapel) Mapel {
	return &mapel{
		db:       db,
		validate: validate,
		repo:     repo,
	}
}

func (m *mapel) Get() []response.Mapel {
	mapel, err := m.repo.Get(helper.Connection())
	helper.PanicIfError(err)
	return mapel
}

func (m *mapel) GetById(mapelId string) response.Mapel{
	mapel,err := m.repo.GetById(m.db,mapelId)
	if err != nil {
		if errors.Is(err,gorm.ErrRecordNotFound) {
			panic(errors.New("data tidak ditemukan"))
		}
		panic(errors.New("proses gagal"))
	}
	return mapel
}

func (m *mapel) GetPengampu(mapelId string) ([]response.User,[]response.User) {
	var pengampu []response.User
	var guru []response.User
	var err error
	wg := sync.WaitGroup{}

	wg. Add(2)
	go func(wg *sync.WaitGroup){
		defer wg.Done()
		pengampu,err = m.repo.GetPengampu(m.db,mapelId)
		helper.PanicIfError(err)
	}(&wg)
	go func(wg *sync.WaitGroup){
		defer wg.Done()
		guru,err = m.repo.GetTeacher(m.db,mapelId)
		helper.PanicIfError(err)
	}(&wg)
	wg.Wait()
	return pengampu,guru
}

func (m *mapel) Post(request request.Mapel) string {
	err := m.validate.Struct(request)
	helper.PanicIfError(err)
	success,err := m.repo.Post(m.db,request)
	helper.PanicIfError(err)
	if success {
		return fmt.Sprintf("mata pelajaran %s berhasil ditambahkan",request.Mapel)
	}
	return fmt.Sprintf("mata pelajaran %s gagal ditambahkan", request.Mapel)
}

func (m *mapel) PostPengampu(mapelId string, userId string) string {
	success,err := m.repo.PostPengampu(m.db,mapelId,userId)
	helper.PanicIfError(err)
	if success {
		return "data berhasil ditambahkan"
	}
	return "data gagal ditambahkan"
}

func (m *mapel) Delete(mapelId string) string {
	success,err := m.repo.Delete(m.db,mapelId)
	helper.PanicIfError(err)
	if success {
		return "mata pelajaran berhasil dihapus"
	}
	return "mata pelajaran gagal dihapus"
}

func (m *mapel) DeletePengampu(mapelId string, userId string) string {
	success,err := m.repo.DeletePengampu(m.db,mapelId,userId)
	helper.PanicIfError(err)
	if success {
		return "data berhasil dihapus"
	}
	return "data gagal dihapus"
}

func (m *mapel) Update(mapelId string,request request.Mapel) string {
	success,err := m.repo.Update(m.db,mapelId,request)
	helper.PanicIfError(err)
	if success {
		return "mata pelajaran berhasil diupdate"
	}
	return "mata pelajaran gagal diupdate"
}
