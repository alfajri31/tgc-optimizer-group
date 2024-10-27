package controller

import (
	"net/http"
	"tgc-optimizer-group/services"
)

func ApiMessages(w http.ResponseWriter, r *http.Request) {
	services.ApiMessageService(w, r)
}

func ApiTest(w http.ResponseWriter, r *http.Request) {
	services.ApiMessageService(w, r)
}
