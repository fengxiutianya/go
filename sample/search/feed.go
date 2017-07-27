package search

import (
	"encoding/json"
	"os"
)
const dataFile = "data/data.json"

// Feed 包含我们需要处理的数据源的信息
type Feed struct{
	Name string  `json:"site"`
	URI  string  `json:"link"`
	Type string  `json:"type"`
}

//RetrieveFeeds 读取并反序列化源数据文件
func RetrieveFeeds()([]*Feed,error){
	//打开文件
	file,err := os.Open(dataFile)
	if err != nil{
		return nil,err
	}

	//当函数返回时 关闭文件  defer 关键字会安排随后的函数调用在韩硕返回时才执行，并且保证一定会执行，
	//即使程序遇到错误
	defer file.Close()

	//将文件解码到一个切片里面
	//这个切片的每一项是一个指向一个Feed类型值的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	//这个函数不需要检查错误，调用者会做这件事
	return feeds,err
}