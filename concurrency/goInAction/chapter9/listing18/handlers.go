package listing18

import (
	"encoding/json"
	"net/http"
)

//Routes 为网络服务设置路由
func Routes() {
	http.HandleFunc("/sendjson", SendJSON)
}

//SendJSON返回一个简单的Json文档
func SendJSON(rw http.ResponseWriter, r *http.Request) {
	u := struct {
		Name  string
		Email string
	}{
		Name:  "Bill",
		Email: "bill@ardantstudios.com",
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	json.NewEncoder(rw).Encode(&u)
}
