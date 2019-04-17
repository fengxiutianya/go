package search

import (
	"fmt"
	"log"
)

// 保存搜索的结果
type Result struct {
	Field string
	Content string
}

// 定义实现搜索的行为
type Matcher interface {
	Search(feed *Feed,searchTerm string)([]*Result,error)
}
// Match函数，为每个数据源启动􏷣􏷡􏷣􏷡􏷣􏷡􏷏􏷐􏺢􏺣􏷤􏱎goroutine来并发执行􏷥􏷈􏷉
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// 􏺙􏶝􏶞的􏶱配􏶲􏷈􏷉􏶡􏶢
	searchResults, err := matcher.Search(feed, searchTerm)
	if err!=nil{
		log.Print(err)
		return
	}

	// 将结果写入通道
	for _,result := range searchResults{
		results <- result
	}
}
// 从每个单独的goroutine接收数据后打印到终端
func Display(results chan *Result)  {
	// 通道会一直被阻塞，直到有数据被写入
	// 一旦通道被关闭，for循环就会中止
	for result := range results{
		fmt.Printf("%s\n%s\n\n",result.Field,result.Content)
	}
}