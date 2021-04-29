package main

import "fmt"

func main(){
	fmt.Println("Hello, World!")
	fmt.Println("Enter Name")
	fmt.Println("Enter your age")
	var name string
	var age int
	fmt.Scanln(&name)
	fmt.Scanln(&age)
	fmt.Printf("Name : %s  Age : %d",name,age)
}
