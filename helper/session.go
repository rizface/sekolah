package helper

import (
	"github.com/gorilla/sessions"
	"github.com/rizface/sekolah/app/model/response"
	"net/http"
)

var Store = sessions.NewCookieStore([]byte("user-data"))

func CreateSession(w http.ResponseWriter, r *http.Request,data response.Login) {
	session,_ := Store.Get(r,"user-data")
	session.Values["user_id"] = data.Id
	session.Values["nama_depan"] = data.NamaDepan
	session.Values["nama_belakang"] = data.NamaBelakang
	session.Values["username"] = data.Username
	session.Values["level"] = data.Level
	session.Save(r,w)
}
