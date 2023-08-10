// bisa menginputkan min. 1 jenis parameter hingga paramter tak terbatas dengan parameter sejenis
// ciri khas namaVariabel ...tipeData

package main

import "fmt"

func main() {
	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	// sprintf tidak mencetakannya ke terminal
	var msg = fmt.Sprintf("Rata-rata : %.2f", avg)
	fmt.Println(msg)
}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers)) // length
	return avg
}
