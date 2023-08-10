package main

import (
	"fmt"
	"strings"
)

func main(){
	var names = [] string {"myskill", "id"}
	printMessages("hello", names)
	printMessages("hello lainnya", names)
}

func printMessages(message string, arr[] string){
	//Join untuk memecah dari bentuk array ke dalam bentuk string
	// " ", sebagai separator
	var nameString = strings.Join(arr, ",")
	fmt.Println(message, nameString)

}