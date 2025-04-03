package lista1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func calcularMedia(notas [3]float64) (media float64){
	var soma float64
	for _, nota := range notas{
		
		soma += nota
	}

	media = soma / 3
	return
}

func Ex1() {
	fmt.Println("Digite as três notas do aluno no modelo de numeração americado e separadas por um espaço")
	scanner.Scan()
	notas := scanner.Text()
	notasSlice := strings.Fields(notas)
	if len(notasSlice) != 3 {
		fmt.Println("Por favor, insira exatamente três notas separadas por espaço.")
		return
	}

	var notasFloat []float64
	for _, nota := range notasSlice {
		notaFloat, err := strconv.ParseFloat(nota, 64)
		if err != nil {
			fmt.Println("Erro ao converter as notas. Certifique-se de que são números válidos.")
			return
		}
		notasFloat = append(notasFloat, notaFloat)
	}

	media := calcularMedia([3]float64(notasFloat))

	fmt.Printf("MÉDIA = %.2f \n", media)

	if media >= 6{
		fmt.Println("APROVADO")
	}else {
		fmt.Println("REPROVADO")
	}
}

func Ex2() {
	fmt.Println("Exercício 2")
}

func Ex3() {
	fmt.Println("Exercício 3")
}

func Ex4() {
	fmt.Println("Exercício 4")
}

func Ex5() {
	fmt.Println("Exercício 5")
}

func Ex6() {
	fmt.Println("Exercício 6")
}

func Ex7() {
	fmt.Println("Exercício 7")
}

func Ex8() {
	fmt.Println("Exercício 8")
}

func Ex9() {
	fmt.Println("Exercício 9")
}

func Ex10() {
	fmt.Println("Exercício 10")
}

func Ex11() {
	fmt.Println("Exercício 11")
}

func Ex12() {
	fmt.Println("Exercício 12")
}

func Ex13() {
	fmt.Println("Exercício 13")
}

func Ex14() {
	fmt.Println("Exercício 14")
}

func Ex15() {
	fmt.Println("Exercício 15")
}

func Ex16() {
	fmt.Println("Exercício 16")
}

func Ex17() {
	fmt.Println("Exercício 17")
}

func Ex18() {
	fmt.Println("Exercício 18")
}

func Ex19() {
	fmt.Println("Exercício 19")
}

func Ex20() {
	fmt.Println("Exercício 20")
}
