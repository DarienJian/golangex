package main

import "fmt"

func main(){
	var age int
	age = 30
	fmt.Printf("年紀是：%d\n",age)

	var name = "Darien"
	fmt.Printf("數值類型為：%T,數值為：%s\n",name,name)

	count := 123
	fmt.Println(count)

	var a,b,c int
	a = 1
	b=2
	c =3

	fmt.Println(a,b,c)

	var m,n int = 10,20
	fmt.Println(m,n)

	var (
		pet = "狗"
		size = "大"
		sex = "公"
	)
	fmt.Printf("動物種類：%s, 體型：%s, 性別：%s",pet,size,sex)
}
