package search

// 默认匹配器
type defaultMatcher struct {

}

// 默认匹配器注册到程序里
func init()  {
	var matcher defaultMatcher
	Register("default",matcher)
}

// search默认实现类
func (m defaultMatcher)Search(feed *Feed,searchTerm string)([]*Result,error) {
	return nil, nil
}