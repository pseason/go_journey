package web

import (
	"encoding/json"
	"net/http"
)

var Handlers map[string]func(w http.ResponseWriter, r *http.Request)

func init() {
	Handlers = make(map[string]func(w http.ResponseWriter, r *http.Request))
	// hello
	Handlers["/hello"] = Hello
	// studnet
	Handlers["/student/find"] = FindStudentById
	Handlers["/student/findAll"] = FindAllStudent
	Handlers["/student/insert"] = InsertStudent
	Handlers["/student/update"] = UpdateStudent
	Handlers["/student/delete"] = DeleteStudent
}

//获取URL的GET参数
func GetUrlArg(r *http.Request, name string) string {
	values := r.URL.Query()
	return values.Get(name)
}

// 获取POST Body 参数
func GetBodyArg(r *http.Request, inter interface{}) {
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(inter)
}

// writer json
func WriteJson(w http.ResponseWriter, inter interface{}) {
	bytes, _ := json.Marshal(inter)
	w.Header().Set("content-type", "text/plain;charset=UTF-8")
	w.Write(bytes)
}
