package controller

import "net/http"

type Dashboard interface {
	Dashboard(w http.ResponseWriter, r *http.Request)
}

