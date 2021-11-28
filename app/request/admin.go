package request

type Admin struct {
	NamaDepan    string `validate:"required"`
	NamaBelakang string
	Username     string `validate:"required"`
	Password     string `validate:"required"`
}
