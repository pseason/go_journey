package web

import (
	"encoding/json"
	"net/http"

	"springmars.com/http-mysql-viper/model"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	res := model.SuccessResponseData("hello http-mysql")
	bytes, _ := json.Marshal(res)
	w.Header().Set("content-type", "application/json;charset=UTF-8")
	w.Write(bytes)
}
