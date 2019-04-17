package main

import (
	"log"
	"net/http"

	"github.com/goinaction/code/chapter9/listing17/handlers"
)

//服务端点 是指与服务宿主信息无关，用来分辨某个服务的地址，一般不包括宿主的一个路径。

func main() {
	handlers.Routes()
	log.Println("listener :Started : Listing on :4000")
	http.ListenAndServe(`:4000`, nil)
}
