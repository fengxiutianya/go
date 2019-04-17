package main
/**
切片之所以是切边，是因为穿件一个新的切片就是把底层数组切除一部分
**/
import "fmt"

func main(){
	//下面这俩个切片共享同一段底层数组，但是不同的是，共享不同的地址片段
	slice :=[]int{10,20,30,40,50}
	newSlice :=slice[1:3]
	newSlice[0] = 100
	//对于底层数组容量是k的切片slice[i:j]来说
	// 长度 j -i  容量 k - i
	fmt.Println(slice,newSlice)

	//切片增长，是通过append 切片增长时，容量有可能改变，也有可能不改变，这取决于备操作切片的可用容量
	slice1 :=[]int{10,20,30,40,50}
	newSlice2 :=slice1[1:3]
	newSlice2 = append(newSlice2,60)
	fmt.Println(newSlice2)
	//限制切片增加的最大容量
	source :=[]string{"apple","orange","plum","Banana","Grape"}
	slice2 :=source[2:3:4] //第三个值，限制切片的容量最大是原来切片的哪一个位置，所以现在容量的大小是 4 - 2 
	fmt.Println(slice2)

	//设置长度和容量一样的好处
	slice3 := source[2:3:3]
	slice3 = append(slice3,"kivi")  //此种情况下，会新创建一个slice3 ,不会再指向souce的底层数组，所以slice3的改版不影响source
	fmt.Println(slice3,source)

	//将一个切片里面的饿元素，追加到另外一个切片里面   在后一个元素后面加上 ... 就是追加
	//range 返回的是切片的副本
	fmt.Printf("%v\n",append(source,slice3...))

	for index,value :=range source{
		fmt.Printf("Index: %d Valjue: %s %d\n",index,value,&source[index])
	}
}