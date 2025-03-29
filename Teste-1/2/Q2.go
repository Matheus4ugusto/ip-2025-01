package lista01_2

import (
	"fmt"
)

func isNumberInRange(num int) {
	if num >= 20 && num <= 90 {
		fmt.Println("O número está entre 20 e 90.")
	} else {
		fmt.Println("O número não está entre 20 e 90.")
	}
}

func main() {
	var num int

	fmt.Print("Digite um número para verificar se está entre 20 e 90: ")
	fmt.Scan(&num)
	isNumberInRange(num)

}