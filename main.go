package main

import "fmt"


func main(){
	var name string
	fmt.Println("이름을 입력하세요")
	fmt.Scanln(&name)
	fmt.Println(name,"님 환영합니다!")


}
