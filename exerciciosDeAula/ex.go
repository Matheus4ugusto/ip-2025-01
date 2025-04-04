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
