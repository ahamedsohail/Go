package main

import "fmt"

func main(){
	fmt.Println("Hello, World!")
	fmt.Println("Enter Name")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("Your name is %s",name)
}
