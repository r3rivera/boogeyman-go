package main

import (
	"fmt"
)


func main(){

	banner("Hello", 25)

}

func banner(text string, width int){
	padding:= (width - len(text))/2
	fmt.Println(padding)
}