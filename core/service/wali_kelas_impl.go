package service

import (
	"github.com/rizface/sekolah/app/model/response"
	"github.com/rizface/sekolah/core/repository"
	"github.com/rizface/sekolah/helper"
	"gorm.io/gorm"
	"sync"
)

type waliKelas struct {
	db *gorm.DB
	kelasRepo repository.Kelas
	walasRepo repository.WaliKelas
}

func NewWaliKelas(db *gorm.DB, kelasRepo repository.Kelas,walasRepo repository.WaliKelas) WaliKelas {
	return waliKelas{
		db:        db,
		kelasRepo: kelasRepo,
		walasRepo: walasRepo,
	}
}

func (w waliKelas) GetWaliKelas() ([]response.Walas, []response.User, []response.Kelas) {
	wg := sync.WaitGroup{}
	var kelas []response.Kelas
	var err error
	var withOutClass []response.User
	var withClass []response.Walas

	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		kelas,err = w.kelasRepo.GetWithoutWalas(w.db)
		helper.PanicIfError(err)
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		withClass,err = w.walasRepo.GetWalasWithClass(w.db)
		helper.PanicIfError(err)
	}(&wg)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		withOutClass,err = w.walasRepo.GetWalasWithoutClass(w.db)
		helper.PanicIfError(err)
	}(&wg)
	wg.Wait()
	return withClass,withOutClass,kelas
}

func (w waliKelas) PostKelas(userId string, kelasId string) string {
	success,err := w.walasRepo.PostWalas(w.db,userId,kelasId)
	helper.PanicIfError(err)
	if success {
		return "wali kelas berhasil ditambahkan"
	}
	return "wali kelas gagal ditambahkan"
}

func (w waliKelas) DeleteWaliKelas(userId string,kelasId string) string {
	success,err := w.walasRepo.DeleteWaliKelas(w.db,userId,kelasId)
	helper.PanicIfError(err)
	if success {
		return "wali kelas berhasil dihapus"
	}
	return "wali kelas gagal dihapus"
}
