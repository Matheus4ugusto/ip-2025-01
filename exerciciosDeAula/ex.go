package exerciciosdeaula

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func CalculateSum() {
	scanner.Scan()
	notas := scanner.Text()
	notasSlice := strings.Fields(notas)
	notasSliceFloats := make([]float64, 0)
	var sum float64
	for _, n := range notasSlice{
		number, err := strconv.ParseFloat(n, 64)
		if err != nil {
			fmt.Println("Insira números válidos")
			return
		}
		notasSliceFloats = append(notasSliceFloats, number)
	}

	for _, n := range notasSliceFloats{
		sum += n
	}

	fmt.Println(sum)
	return
}

func Search(){
	fmt.Println("Insira o conjunto amostral de elementos")

	scanner.Scan()
	s := scanner.Text()

	fmt.Println("Insira o elemento a ser buscado")

	scanner.Scan()
	b := scanner.Text()

	bInt, err := strconv.Atoi(b)

	if err != nil {
		fmt.Println("Erro ao converter as notas. Certifique-se de que são números válidos.")
		return
	}

	sSplit := strings.Fields(s)

	var nums []int
	for _, n := range sSplit {
		num, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println("Erro ao converter as notas. Certifique-se de que são números válidos.")
			return
		}
		nums = append(nums, num)
	}

	for i, e := range nums {
		if e == bInt {
			fmt.Printf("O número %d é o elemento %d \n", bInt, i)
			return
		}
	}

	fmt.Print(-1)
	return
}