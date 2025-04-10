package lista1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func calcularMedia(notas [3]float64) (media float64) {
	var soma float64
	for _, nota := range notas {

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

	if media >= 6 {
		fmt.Println("APROVADO")
	} else {
		fmt.Println("REPROVADO")
	}
}

func Ex2() {
	fmt.Println("Quantos casos deverão ser testados?")
	scanner.Scan()
	quantidadeCasos := scanner.Text()
	numeroCasos, err := strconv.Atoi(quantidadeCasos)
	if err != nil {
		fmt.Println("A quantidade de casos deve ser um número inteiro")
		fmt.Println(numeroCasos)
		return
	}

	casos := make([][5]float64, 0)

	for range numeroCasos {
		scanner.Scan()

		infos := scanner.Text()

		infosSlice := strings.Fields(infos)

		infosFloats := [5]float64{}

		for i := 0; i <= 4; i++ {
			number, err := strconv.ParseFloat(infosSlice[i], 64)
			if err != nil {
				fmt.Println("Números inválidos")
				return
			}
			infosFloats[i] = number
		}

		var sum float64

		for i, f := range infosFloats {
			if i != 0 {
				sum += f
			}
		}

		if sum != 100 {
			fmt.Println("A soma das porcentagens deve corresponder a 100%")
			return
		}

		casos = append(casos, infosFloats)
	}
	
	for i, c := range casos{
		var result float64
		result += (c[1] / 100) * c[0]
		result += (c[2] / 100) * c[0] * 5
		result += (c[3] / 100) * c[0] * 10
		result += (c[4] / 100) * c[0] * 20
		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f \n", i + 1, result)
	}
	return
}

func Ex3() {
	scanner.Scan()
	txt1 := scanner.Text()
	scanner.Scan()
	txt2 := scanner.Text()
	scanner.Scan()
	txt3 := scanner.Text()
	num1, err1 := strconv.Atoi(txt1)
	num2, err2 := strconv.Atoi(txt2)
	num3, err3 := strconv.Atoi(txt3)

	if err1 != nil || err2 != nil || err3 != nil{
		fmt.Println("Números inválidos")
		return
	} 
	if num1 >= 10 || num2 >= 10 || num3 >= 10 || num1 < 0 || num2 < 0 || num3 < 0{
		fmt.Println("Somente os número 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 são válidos")
		return
	}
	if num1 == 0 {
		result, err := strconv.Atoi(txt2 + txt3)
		if err != nil {
			fmt.Println("Insira números válidos")
		}
		fmt.Printf("%d, %f", result, math.Pow(float64(result), 2))
		return
	} else if num1 == 0 && num2 == 0{
		fmt.Println(txt3)
		return
	}else if num1 == 0 && num2 == 0 && num3 == 0{
		fmt.Println(txt3)
		return
	}else {
		fmt.Println(txt1 + txt2 + txt3)
		return
	}
}

func Ex4() {
	scanner.Scan()
	s := scanner.Text()
	minSalary, err1 := strconv.ParseFloat(s, 64)

	scanner.Scan()
	q := scanner.Text()
	quantity, err2 := strconv.ParseFloat(q, 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Insira números válidos")
	}

	kW := 0.7 * minSalary / 100

	debit := kW * quantity

	fmt.Printf("Custo por kW: R$%.2f \n", kW)
	fmt.Printf("Custo do consumo: R$%.2f \n", debit)
	fmt.Printf("Custo com desconto: R$%.2f \n", debit * 0.9)
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
