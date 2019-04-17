package search

import (
	"encoding/json"
	"os"
)

const datafile  = "data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}



// 读取并反序列化源数据
func RetrieveFeeds() ([]*Feed,error){
	// 打开文件
	file,err :=os.Open(datafile)
	if err != nil {
		return nil,err
	}
	//当关闭函数时，关闭文件
	defer file.Close()

	// 将文件解码到一个Feed数组里
	// 数据里每一个元素指向一个Feed类型的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	// 这个函数不需要检查错误，调用者会做这件事
	return feeds,err
}
