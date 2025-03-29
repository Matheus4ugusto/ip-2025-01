package Lista01_3

import "fmt"


func printNumbersAndSum() {
	sum := 0
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Printf("A soma dos números de 1 a 100 é: %d\n", sum)
}

func main() {

	printNumbersAndSum()
}