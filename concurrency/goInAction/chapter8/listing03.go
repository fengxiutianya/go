package main

import (
	"log"
)

func init() {
	log.SetPrefix("Trace: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	//println 写到标准日志记录器
	log.Println("message")

	//Patalln 在调用println之后会接着调用os.Exit()
	log.Panicln("fatal message")

	//panicln在调用pringln() 知乎会接着调用panic()
	log.Panicln("panic message")

}
