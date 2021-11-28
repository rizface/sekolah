package helper

import "github.com/martinlindhe/notify"

func Notif(title string,msg string) {
	notify.Alert("Sekolah", title,msg,"")
}
