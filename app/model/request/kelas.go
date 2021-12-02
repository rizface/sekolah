package request

type Kelas struct{
	Tingkat string `validate:"required"`
	Kelas string `validate:"required"`
}
