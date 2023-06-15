package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

//声明一个Hero结构体切片类型
type HeroSlice []Hero

//实现Interface接口
func (hs HeroSlice) Len() int {
	return len(hs)
}
func (hs HeroSlice) Less(i, j int) bool { //Less方法就是决定你使用什么标准进行排序 按照Hero的年龄从小到大排序
	return hs[i].Age > hs[j].Age
}
func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

//接口最佳实践
func main() {
	//先定义一个数组/切片
	var intSlice = []int{0, -1, 10, 7, 90}
	//使用系统提供的方法
	/*引用类型：指针、slice切片、map、管道、interface*/
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//对结构体切片进行排序
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		//将hero append到heroes切片
		heroes = append(heroes, hero)
	}
	//查看排序前的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}
	//调用sort.Sort
	sort.Sort(heroes)
	for _, v := range heroes {
		fmt.Println(v)
	}
}
