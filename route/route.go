package route

const (
	LOGIN = "/"

	// kelola data admin (CRUD)
	ADMIN_DASHBOARD = "/admin/data-admin"
	TAMBAH_ADMIN = "/admin/data-admin/tambah"
	HAPUS_ADMIN = "/admin/data-admin/hapus/{userId}"
	UPDATE_ADMIN = "/admin/data-admin/update/{userId}"


	// kelola data guru (CRUD)
	GURU_DASHBOARD = "/admin/data-guru"
	TAMBAH_GURU = "/admin/data-guru/tambah"
	HAPUS_GURU = "/admin/data-guru/hapus/{userId}"
	UPDATE_GURU = "/admin/data-guru/update/{userId}"

	// kelola data guru (CRUD)
	SISWA_DASHBOARD = "/admin/data-siswa"
	TAMBAH_SISWA = "/admin/data-siswa/tambah"
	HAPUS_SISWA = "/admin/data-siswa/hapus/{userId}"
	UPDATE_SISWA = "/admin/data-siswa/update/{userId}"
)
