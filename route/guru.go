package route

const(
	GURU = "/guru"
	GURU_MAPEL = "/guru/data-mapel"
	GURU_KELAS = "/guru/kelas"
	KELAS_SISWA = "/guru/kelas/{kelasId}"
	INPUT_NILAI = "/guru/kelas/{kelasId}/siswa/{siswaId}"
	GET_SISWA_NILAI = "/guru/nilai/kelas/{kelasId}/siswa/{siswaId}"
	HAPUS_SISWA_NILAI = "/guru/nilai/hapus/{nilaiId}"
	UPDATE_SISWA_NILAI = "/guru/nilai/update/{nilaiId}"
	DETAIL_SISWA = "/guru/siswa/{siswaId}"
	ABSEN_SISWA = "/guru/absen/kelas/{kelasId}"
)
