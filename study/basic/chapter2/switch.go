package main
import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Println("Go runs on")
	//在条件语句前面执行一条语句
	switch os := runtime.GOOS; os{
	case "darwin":fmt.Println("OS X.")
	case "linux":fmt.Println("Linux")
	default :fmt.Println("%s.",os)
	}
}