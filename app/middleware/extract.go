package middleware

import (
	"net/http"
	"strings"
)

func Extract(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		url := request.URL.RequestURI()
		if strings.Contains(url, "data-admin") {
			request.Header.Add("X-ROLE", "admin")
			request.Header.Add("X-INDEX", "view/admin/admin/index.gohtml")
			request.Header.Add("X-INDEX-VIEW-NAME","data_admin_index")
			request.Header.Add("X-POST-PAGE","view/admin/admin/tambah.gohtml")
			request.Header.Add("X-POST-PAGE-VIEW-NAME", "tambah_data_admin")
			request.Header.Add("X-INDEX-PATH", "/admin/data-admin")
			request.Header.Add("X-UPDATE-PAGE", "view/admin/admin/update.gohtml")
			request.Header.Add("X-UPDATE-PAGE-VIEW-NAME","update_data_admin")
			request.Header.Add("X-UPDATE-REDIRECT","/admin/data-admin/update/")

		} else if strings.Contains(url, "data-guru") {
			request.Header.Add("X-ROLE", "guru")
			request.Header.Add("X-INDEX", "view/admin/guru/index.gohtml")
			request.Header.Add("X-INDEX-VIEW-NAME","data_guru_index")
			request.Header.Add("X-POST-PAGE","view/admin/guru/tambah.gohtml")
			request.Header.Add("X-POST-PAGE-VIEW-NAME", "tambah_data_guru")
			request.Header.Add("X-INDEX-PATH", "/admin/data-guru")
			request.Header.Add("X-UPDATE-PAGE", "view/admin/guru/update.gohtml")
			request.Header.Add("X-UPDATE-PAGE-VIEW-NAME","update_data_guru")
			request.Header.Add("X-UPDATE-REDIRECT","/admin/data-guru/update/")
		} else if strings.Contains(url, "data-siswa") {
			request.Header.Add("X-ROLE", "murid")
			request.Header.Add("X-INDEX", "view/admin/siswa/index.gohtml")
			request.Header.Add("X-INDEX-VIEW-NAME","data_siswa_index")
			request.Header.Add("X-POST-PAGE","view/admin/siswa/tambah.gohtml")
			request.Header.Add("X-POST-PAGE-VIEW-NAME", "tambah_data_siswa")
			request.Header.Add("X-INDEX-PATH", "/admin/data-siswa")
			request.Header.Add("X-UPDATE-PAGE", "view/admin/siswa/update.gohtml")
			request.Header.Add("X-UPDATE-PAGE-VIEW-NAME","update_data_siswa")
			request.Header.Add("X-UPDATE-REDIRECT","/admin/data-siswa/update/")
		}
		next.ServeHTTP(writer,request)
	})
}
