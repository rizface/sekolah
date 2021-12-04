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

	// kelola data akuntansi (CRUD)
	AKUNTANSI_DASHBOARD = "/admin/data-akuntansi"
	TAMBAH_AKUNTANSI = "/admin/data-akuntansi/tambah"
	HAPUS_AKUNTANSI = "/admin/data-akuntansi/hapus/{userId}"
	UPDATE_AKUNTANSI = "/admin/data-akuntansi/update/{userId}"

	// kelola data kelas (CRUD)
	KELAS_DASHBOARD = "/admin/data-kelas"
	TAMBAH_KELAS = "/admin/data-kelas/tambah"
	TAMBAH_ANGGOTA_KELAS = "/admin/data-kelas/tambah-anggota/kelas/{kelasId}/siswa/{userId}"
	HAPUS_ANGGOTA_KELAS = "/admin/data-kelas/hapus-anggota/kelas/{kelasId}/siswa/{userId}"
	HAPUS_KELAS = "/admin/data-kelas/hapus/{kelasId}"
	UPDATE_KELAS = "/admin/data-kelas/update/{kelasId}"
	DATA_SISWA_KELAS = "/admin/data-kelas/siswa/{kelasId}"

)
