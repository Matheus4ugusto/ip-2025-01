package lista01_1

import (
	"fmt"
)

func checkNumberSign(num int) {
	if num > 0 {
		fmt.Println("O número é positivo.")
	} else if num < 0 {
		fmt.Println("O número é negativo.")
	} else {
		fmt.Println("O número é nulo.")
	}
}

func main() {
	var num int

	fmt.Print("Digite um número: ")
	fmt.Scan(&num)
	checkNumberSign(num)
}