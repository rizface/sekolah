package request

type Nilai struct {
	StudentId  string  `validate:"required"`
	TeacherId  string  `validate:"required"`
	SemesterId string  `validate:"required"`
	SubjectId  string  `validate:"required"`
	Grade      float64 `validate:"required"`
}
