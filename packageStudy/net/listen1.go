package main

import (
	"io"
	"log"
	"net"
	"net/http"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
func main() {

	l, err := net.Listen("tcp", "127.0.0.1:9000")

	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// http.HandleFunc("/hello", HelloServer)
	// http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

	// 	// log.SetOutput(os.Stdin)
	// 	log.Println("test")
	// 	io.WriteString(w, "index")

	// })

	for {
		//wait for a connection
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		//handle connection in a goroutine
		//the loop then returns to accepting,
		//so that mutiple connections can be servered concurrently
		go func(c net.Conn) {
			addr := c.RemoteAddr()
			log.Println("test" + addr.String())
			c.Write([]byte("zhangke"))
			c.Close()
		}(conn)
	}
	// http.Serve(l, nil)
	// hello world, the web server
	// http.HandleFunc("/hello", HelloServer)

	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
