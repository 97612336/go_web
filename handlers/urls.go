package handlers

import (
	"go_web/handlers/web"
	"net/http"
)

func MyUrls() {
	http.HandleFunc("/",web.Index)
}
