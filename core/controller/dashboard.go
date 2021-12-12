package controller

import "net/http"

type Dashboard interface {
	DashboardAdmin(w http.ResponseWriter, r *http.Request)
}
