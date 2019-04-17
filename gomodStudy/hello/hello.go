package hello

//import "rsc.io/quote"

//import "rsc.io/quote"

import "example.com/hello/search"

func Hello() string {
	//return quote.HelloV2()
	return search.Run("hello world")
}
