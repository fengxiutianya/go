package main

import (
	"flag"
	"fmt"
	"os"
)

// 实际中应该用更好的变量名
var (
	h    bool
	H    bool
	v, V bool
	t, T bool
	q    *bool

	s    string
	p    string
	c    string
	g    string
	test *bool
)

func init() {

	flag.BoolVar(&h, "h", false, "this help")

	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&V, "V", false, "show version and configure options then exit")

	flag.BoolVar(&t, "t", false, "test configuration and exit")
	flag.BoolVar(&T, "T", false, "test configuration, dump it and exit")

	// 另一种绑定方式
	q = flag.Bool("q", false, "suppress non-error messages during configuration testing")

	// 注意 `signal`。默认是 -s string，有了 `signal` 之后，变为 -s signal
	flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")
	flag.StringVar(&p, "p", "/usr/local/nginx/", "set `prefix` path")
	flag.StringVar(&c, "c", "conf/nginx.conf", "set configuration `file`")
	flag.StringVar(&g, "g", "conf/nginx.conf", "set global `directives` out of configuration file")

	// 改变默认的 Usage
	flag.Usage = usage
}

func main() {
	test := flag.Bool("H", false, "this help")
	flag.Parse()
	var args = flag.Args()
	fmt.Println(args)
	if h {
		flag.Usage()
	}
	if *test {
		flag.Usage()
	}
	if v {
		fmt.Println("version 1.0")
	}
	if len(s) != 0 {
		fmt.Println("dd", s)
	}
	fmt.Println(s)
	// var arg = flag.Arg(1)
	// fmt.Println(arg)

}

func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`)
	flag.PrintDefaults()
}
