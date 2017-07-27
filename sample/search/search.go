package search
import (
	"log"
	"sync"
)
//注册用于搜索的匹配器映射  首字母大写表示公开变量 开头小写表示不公开变量 公开变量可以在导入包之后直接调用
var matchers = make(map[string]Matcher)

//run 执行搜索逻辑
func Run(searchTerm string){
	//获取搜索的数据源列表
	feeds,err:=RetrieveFeeds()

	if err !=nil{
		log.Fatal(err)
	}

	//创建一个无缓冲的通道，接受匹后的结果
	results :=make(chan *Result)

	//构造一个waitGroup，以便处理所有的数据源
	var waitGroup sync.waitGroup

	//设置需要等待处理
	//每个数据源的goroutine 的数量
	waitGroup.Add(len(feeds))

	//为每个数据源启动一个goroutine来查找结果
	for _, feed := range feeds{
		//读取一个匹配器用于查找
		matcher,exists :=matchers[feed.Type]
		if !exists{
			matcher =matchers["default"]
		}

		//启动一个goruntine来执行搜索
		go func(matcher Matcher,feed *Feed){
			Match(matcher,feed,searchTerm,results)
			waitGroup.Done()
		}(matcher,feed)
	}

	//启动一个goroutine 来监控是否所有的工作都做完了
	go func(){
		//等待所有任务完成
		waitGroup.Wait()

		//用关闭公道的方式，通知Display函数
		//可以退出程序了 
		close(results)
	}()

	//启动函数，显示返回的结果，并且在最后一个结果显示完成后返回
	Display(results)
}

//Register 调用是，会注册一个匹配器，提供给后面的程序使用
func Register(feedType string,matcher Matcher){
	if _,exists := matchers[feedType]; exists{
		log.Fatalln(feedType,"Matcher already registered")
	}

	log.Println("Register",feedType,"matcher")
	matchers[feedType] = matcher
}