package search

import (
	"log"
	"sync"
)

// 注册用于搜索匹配器的映射
var matchers = make(map[string]Matcher)

// run执行搜索逻辑
func Run(searchTerm string){
    //获取需要搜索的数据源
    feeds,err := RetrieveFeeds()
    if err != nil{
    	log.Fatal(err)
	}

    // 创建一个无缓冲的通道，接收匹配后的结果
    results := make(chan *Result)

    // 构造一个waitGroup，以便处理所有的结果
	var waitGroup sync.WaitGroup

    // 设置需要等待处理的goroutine数量
    waitGroup.Add(len(feeds))

    //为每个数据源启动一个goroutine来查找结果
	 for _, feed := range feeds {
		 // 􏷌􏷍获取一个匹配器用于查找
		 matcher, exists := matchers[feed.Type]

		 if !exists {
			 matcher = matchers["default"]
		  }

		 // 启动􏷤􏱎一个goroutine用于查找 􏷥􏷈􏷉􏶡􏶢
		 go func(matcher Matcher, feed *Feed) {
			 Match(matcher, feed, searchTerm, results)
			 waitGroup.Done()
		 }(matcher, feed)
     }

	 // 启动一个goroutine来监控是否所有的工作都做完了
	 go func() {
	 	// 等待所有的任务完成
	 	waitGroup.Wait()
	 	// 关闭通道，通知Display函数可以退出了
	 	close(results)
	 }()
	 // 启动函数，显示返回结果，并且在最后一个结果显示完成后返回
	 Display(results)
}
// Register 调用时，􏰞􏶯􏶰一个􏶱配􏶲，提供􏹧􏷘􏺹的程序使用
func Register(feedType string,matcher Matcher)  {

	if _,exists := matchers[feedType];exists{
		log.Fatalln(feedType,"Matcher already register")
	}
	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}