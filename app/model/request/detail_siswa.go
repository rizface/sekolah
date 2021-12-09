package request

type DetailSiswa struct {
	Nisn string `validate:"required"`
	Nis string `validate:"required"`
}
