package main

import "fmt"

func main() {
	var firstName string = "my"
	firstName = "ubah dari my"
	const lastName = "skill"
	// tidak akan bisa diubah
	// lastName ="ubah dari skill"
	fmt.Printf("halo %s %s!\n", firstName, lastName)

}
