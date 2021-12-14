package route

const(
	GURU = "/guru"
	GURU_MAPEL = "/guru/data-mapel"
	GURU_KELAS = "/guru/kelas"
	KELAS_SISWA = "/guru/kelas/{kelasId}"
	INPUT_NILAI = "/guru/kelas/{kelasId}/siswa/{siswaId}"
	GET_SISWA_NILAI = "/guru/nilai/kelas/{kelasId}/siswa/{siswaId}"
)
