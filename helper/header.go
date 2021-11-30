package helper

import "net/http"

func GetHeader(r *http.Request, headerName string) string {
	return r.Header.Get(headerName)
}
