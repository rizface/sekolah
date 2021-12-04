package test

import (
	"github.com/rizface/sekolah/app/model/app"
	"github.com/rizface/sekolah/helper"
	"testing"
)

func TestFakerGender(t *testing.T) {
	t.SkipNow()
	gender := []app.Gender{
		app.Gender{
			Gender: "laki-laki",
		},
		app.Gender{
			Gender: "perempuan",
		},
	}
	db := helper.Connection()
	db.Create(&gender)
}

func TestFakerLevel(t *testing.T) {
	t.SkipNow()
	Level := []app.Level{{
		Level: "admin",
	}, app.Level{
		Level: "guru",
	}, app.Level{
		Level: "akuntansi",
	}, app.Level{Level: "murid"}}

	db := helper.Connection()
	db.Create(&Level)
}

func TestFakerAdmin(t *testing.T) {
	//t.SkipNow()
	admin := app.User{
		NamaDepan:    "muhammad",
		NamaBelakang: "al farizzi",
		Username:     "admin",
		Password:     helper.GeneratePassword("rahasia"),
		LevelId:      "8888308f-2d3f-465f-b4b3-d60877babf0a",
		GenderId:     "f3128026-1992-45f9-8acc-b323abaf8d74",
	}
	db := helper.Connection()
	db.Create(&admin)
}
