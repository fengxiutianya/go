package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

/**************************************
 *      Author : zhangke
 *      Date   : 2019-04-14 14:25
 *      email  : 398757724@qq.com
 *      Desc   : 生成markdown目录
 ***************************************/
type md_path struct {
	Header,Path string

	Paths   []string
}

// h2
const h2  =`{{"\n"}}## {{.Header}}{{$head := .Header}}{{$path := .Path}}
{{range $key,$value := .Paths}}{{$key}}{{"."}} [{{$value}}]({{$path}}/{{$value}}){{"\n"}}{{end}}`

// h3
const h3  =`### {{.Header}}
{{range $key,$value := .Paths}}{{$key}}{{"."}} [{{$value}}]({{$value}}){{"\n"}}{{end}}`


func main() {
	file, err := os.OpenFile("test.md", os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	t := template.Must(template.New("h2").Parse(h2))


	path := "/Users/zhangke/Documents/书籍"
	mds := getPath(path)
	for _,value := range mds {
		t.Execute(file,value)
	}
}
// 获取指定路径下面所有pdf文件，目前只会检测俩层
func getPath(path string)[]md_path  {

	files,err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	mds := make([]md_path,0)
	
	md := md_path{
		Header:path[strings.LastIndex(path,"/")+1:],
		Path:path,
		Paths:make([]string,0),
	}

	// 获取指定path下面所有文件路径
	for _, value := range files {

		// 对文件夹进行处理
		if  value.IsDir() &&
			!strings.HasPrefix(strings.TrimSpace(value.Name()),"."){

			for _ , tem := range getPath(path+"/"+value.Name()) {
				mds = append(mds,tem)
			}

		// 对文件进行处理
		}else if strings.HasSuffix(strings.TrimSpace(value.Name()),".pdf")  {
			s := strings.ReplaceAll(value.Name()," ","%20")
			md.Paths = append(md.Paths,s)
		}
	}
	mds = append(mds,md)
	return mds
}

func getFileNames(path string) *md_path {
	md := new(md_path)
	md.Header = path
	
	files, err := ioutil.ReadDir(path);
	if err != nil {
		log.Fatalln(err)
	}
	paths := make([]string,5)
	for _,value := range files {
		paths = append(paths, value.Name())

	}
	md.Paths = paths
	return md
}

func echo_markdwon(path string, content []string) {
	//_, err := os.Stat(path)    //os.Stat获取文件信息
	//if err != nil && !os.IsExist(err) {
	//	log.Fatal(err)
	//}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range content {
		file.WriteString(line + "\n")
	}

}
