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
	t.SkipNow()
	admin := app.User{
		NamaDepan:    "muhammad",
		NamaBelakang: "al farizzi",
		Username:     "admin",
		Password:     helper.GeneratePassword("rahasia"),
		LevelId:      "b39a2d92-8867-48ba-8cff-8b3de7191019",
		GenderId:     "2069abf7-96ce-4a28-99e1-02c1e283fff2",
	}
	db := helper.Connection()
	db.Create(&admin)
}
