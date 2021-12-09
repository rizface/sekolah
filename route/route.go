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

	// kelola data suswa (CRUD)
	SISWA_DASHBOARD = "/admin/data-siswa"
	TAMBAH_SISWA = "/admin/data-siswa/tambah"
	HAPUS_SISWA = "/admin/data-siswa/hapus/{userId}"
	UPDATE_SISWA = "/admin/data-siswa/update/{userId}"

	// Detail
	UPDATE_DETAIL_SISWA = "/admin/data-siswa/update-detail/{userId}"
	UPDATE_DETAIL_PEGAWAI = "/admin/data-pegawai/detail/{userId}"

	// kelola data akuntansi (CRUD)
	AKUNTANSI_DASHBOARD = "/admin/data-akuntansi"
	TAMBAH_AKUNTANSI = "/admin/data-akuntansi/tambah"
	HAPUS_AKUNTANSI = "/admin/data-akuntansi/hapus/{userId}"
	UPDATE_AKUNTANSI = "/admin/data-akuntansi/update/{userId}"

	// kelola data kelas (CRUD)
	KELAS_DASHBOARD = "/admin/data-kelas"
	TAMBAH_KELAS = "/admin/data-kelas/tambah"
	HAPUS_KELAS = "/admin/data-kelas/hapus/{kelasId}"
	UPDATE_KELAS = "/admin/data-kelas/update/{kelasId}"

	// Anggota Kelas
	TAMBAH_ANGGOTA_KELAS = "/admin/data-kelas/tambah-anggota/kelas/{kelasId}/siswa/{userId}"
	HAPUS_ANGGOTA_KELAS = "/admin/data-kelas/hapus-anggota/kelas/{kelasId}/siswa/{userId}"
	DATA_SISWA_KELAS = "/admin/data-kelas/siswa/{kelasId}"
	DETAIL_SISWA_KELAS = "/admin/data-kelas/siswa/detail/{userId}"

	// Wali Kelas
	DATA_WALI_KELAS = "/admin/data-kelas/wali-kelas"
	HAPUS_WALAS = "/admin/data-kelas/wali-kelas/hapus/walas/{userId}/kelas/{kelasId}"

	// Mata Pelajaran
	DATA_MATA_PELAJARAN = "/admin/data-mata-pelajaran"
	TAMBAH_MATA_PELAJARAN = "/admin/data-mata-pelajaran/tambah"
	HAPUS_MATA_PELAJARAN = "/admin/data-mata-pelajaran/hapus/{mapelId}"
	UPDATE_MATA_PELAJARAN = "/admin/data-mata-pelajaran/update/{mapelId}"

	// Pengampu
	DATA_PENGAMPU = "/admin/data-mata-pelajaran/{mapelId}/guru"
	HAPUS_PENGAMPU = "/admin/data-mata-pelajaran/hapus/mapel/{mapelId}/guru/{userId}"
	TAMBAH_PENGAMPU = "/admin/data-mata-pelajaran/tambah/mapel/{mapelId}/guru/{userId}"

	// Nilai Siswa
	INPUT_NILAI_SISWA = "/admin/data-siswa/nilai/siswa/{userId}"
	HAPUS_NILAI_SISWA = "/admin/data-siswa/nilai/hapus/{gradeId}"
	UPDATE_NILAI_SISWA = "/admin/data-siswa/nilai/update/{gradeId}"
)
