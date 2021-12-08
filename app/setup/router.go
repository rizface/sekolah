package setup

import (
	"github.com/rizface/sekolah/app/middleware"
	"github.com/rizface/sekolah/route"
	"net/http"
)

func Login() {
	r := R.NewRoute().Subrouter()
	controller := LoginController()

	r.HandleFunc(route.LOGIN, controller.LoginPage).Methods(http.MethodGet)
	r.HandleFunc(route.LOGIN, controller.Login).Methods(http.MethodPost)
}

func Admin() {
	r := R.NewRoute().Subrouter()
	r.Use(middleware.Admin)

	// CRUD GURU
	crudAdmin := CrudAdminController()
	r.HandleFunc(route.ADMIN_DASHBOARD, crudAdmin.Get).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_ADMIN,crudAdmin.PostPage).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_ADMIN, crudAdmin.Post).Methods(http.MethodPost)
	r.HandleFunc(route.HAPUS_ADMIN,crudAdmin.Delete).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_ADMIN,crudAdmin.UpdatePage).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_ADMIN,crudAdmin.Update).Methods(http.MethodPost)

	// CRUD GURU
	r.HandleFunc(route.GURU_DASHBOARD, crudAdmin.Get).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_GURU,crudAdmin.PostPage).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_GURU,crudAdmin.Post).Methods(http.MethodPost)
	r.HandleFunc(route.HAPUS_GURU,crudAdmin.Delete).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_GURU,crudAdmin.UpdatePage).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_GURU,crudAdmin.Update).Methods(http.MethodPost)

	// CRUD MURID
	r.HandleFunc(route.SISWA_DASHBOARD, crudAdmin.Get).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_SISWA,crudAdmin.PostPage).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_SISWA,crudAdmin.Post).Methods(http.MethodPost)
	r.HandleFunc(route.HAPUS_SISWA,crudAdmin.Delete).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_SISWA,crudAdmin.UpdatePage).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_SISWA,crudAdmin.Update).Methods(http.MethodPost)

	// CRUD AKUNTANSI
	r.HandleFunc(route.AKUNTANSI_DASHBOARD, crudAdmin.Get).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_AKUNTANSI,crudAdmin.PostPage).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_AKUNTANSI,crudAdmin.Post).Methods(http.MethodPost)
	r.HandleFunc(route.HAPUS_AKUNTANSI,crudAdmin.Delete).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_AKUNTANSI,crudAdmin.UpdatePage).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_AKUNTANSI,crudAdmin.Update).Methods(http.MethodPost)

	// CRUD KELAS
	crudKelas := CrudKelasController()
	r.HandleFunc(route.KELAS_DASHBOARD,crudKelas.Get).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_KELAS, crudKelas.PostPage).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_KELAS, crudKelas.Post).Methods(http.MethodPost)
	r.HandleFunc(route.HAPUS_KELAS,crudKelas.Delete).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_KELAS,crudKelas.UpdatePage).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_KELAS,crudKelas.Update).Methods(http.MethodPost)

	crudAssign := CrudAssignStudentToClass()
	crudWalas := CrudWalas()
	r.HandleFunc(route.DATA_SISWA_KELAS,crudAssign.GetStudents).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_ANGGOTA_KELAS,crudAssign.AddStudents).Methods(http.MethodPost)
	r.HandleFunc(route.HAPUS_ANGGOTA_KELAS,crudAssign.DeleteStudents).Methods(http.MethodGet)
	r.HandleFunc(route.DATA_WALI_KELAS,crudWalas.GetWaliKelas).Methods(http.MethodGet)
	r.HandleFunc(route.DATA_WALI_KELAS,crudWalas.PostWaliKelas).Methods(http.MethodPost)
	r.HandleFunc(route.HAPUS_WALAS,crudWalas.DeleteWaliKelas).Methods(http.MethodGet)
	r.HandleFunc(route.DETAIL_SISWA_KELAS,crudAssign.DetailStudents).Methods(http.MethodGet)

	// CRUD Mata Pelajaran
	crudMapel := CrudMapel()
	r.HandleFunc(route.DATA_MATA_PELAJARAN, crudMapel.Get).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_MATA_PELAJARAN,crudMapel.PostPage).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_MATA_PELAJARAN,crudMapel.Post).Methods(http.MethodPost)
	r.HandleFunc(route.HAPUS_MATA_PELAJARAN,crudMapel.Delete).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_MATA_PELAJARAN,crudMapel.UpdatePage).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_MATA_PELAJARAN,crudMapel.Update).Methods(http.MethodPost)

	crudPengampu := CrudPengampu()
	r.HandleFunc(route.DATA_PENGAMPU, crudPengampu.GetPengampu).Methods(http.MethodGet)
	r.HandleFunc(route.HAPUS_PENGAMPU,crudPengampu.DeletePengampu).Methods(http.MethodGet)
	r.HandleFunc(route.TAMBAH_PENGAMPU,crudPengampu.PostPengampu).Methods(http.MethodPost)

	crudNilai := CrudNilaisiswa()
	r.HandleFunc(route.INPUT_NILAI_SISWA,crudNilai.PostForm).Methods(http.MethodGet)
	r.HandleFunc(route.INPUT_NILAI_SISWA,crudNilai.Post).Methods(http.MethodPost)
	r.HandleFunc(route.HAPUS_NILAI_SISWA,crudNilai.Delete).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_NILAI_SISWA,crudNilai.UpdateForm).Methods(http.MethodGet)
	r.HandleFunc(route.UPDATE_NILAI_SISWA,crudNilai.Update).Methods(http.MethodPost)
}