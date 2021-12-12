package request

import "time"

type Spp struct {
	Nominal string `gorm:"required"`
	StudentId,EmployeeId string
	CreatedAt time.Time
}
