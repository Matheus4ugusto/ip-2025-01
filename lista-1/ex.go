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

	for i, c := range casos {
		var result float64
		result += (c[1] / 100) * c[0]
		result += (c[2] / 100) * c[0] * 5
		result += (c[3] / 100) * c[0] * 10
		result += (c[4] / 100) * c[0] * 20
		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f \n", i+1, result)
	}
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

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Números inválidos")
		return
	}
	if num1 >= 10 || num2 >= 10 || num3 >= 10 || num1 < 0 || num2 < 0 || num3 < 0 {
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
	} else if num1 == 0 && num2 == 0 {
		fmt.Println(txt3)
		return
	} else if num1 == 0 && num2 == 0 && num3 == 0 {
		fmt.Println(txt3)
		return
	} else {
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
	fmt.Printf("Custo com desconto: R$%.2f \n", debit*0.9)
}

func Ex5() {
	scanner.Scan()
	i := scanner.Text()

	infosSlice := strings.Fields(i)

	numeroIdentificador := infosSlice[0]

	consumo, err := strconv.ParseFloat(infosSlice[1], 64)

	tipo := infosSlice[2]

	if err != nil {
		fmt.Println("Insira números válidos")
		return
	}

	var valor float64

	if tipo == "R" {
		valor = 5 + (0.05 * consumo)
	}
	if tipo == "C" {
		if consumo <= 80 {
			valor = 500
		} else {
			valor = 500 + (0.25 * (consumo - 80))
		}
	}
	if tipo == "I" {
		if consumo <= 100 {
			valor = 800
		} else {
			valor = 800 + (0.04 * (consumo - 100))
		}
	}

	fmt.Printf("CONTA = %s \n", numeroIdentificador)
	fmt.Printf("VALOR DA CONTA = %.2f \n", valor)
}

func Ex6() {
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())

	if err != nil {
		fmt.Println("Insira um número válido")
	}

	temperaturasFahrenheit := make([]float64, 0)
	temperaturasCelsius := make([]float64, 0)

	for range n {
		scanner.Scan()
		i, err1 := strconv.ParseFloat(scanner.Text(), 64)

		if err1 != nil {
			fmt.Println("Insira um número válido")
		}

		temperaturasFahrenheit = append(temperaturasFahrenheit, i)
	}

	for _, e := range temperaturasFahrenheit {
		temperaturasCelsius = append(temperaturasCelsius, 5*(e-32)/9)
	}

	for i, e := range temperaturasFahrenheit {
		fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS \n", e, temperaturasCelsius[i])
	}
}

func Ex7() {
	scanner.Scan()
	i := scanner.Text()
	t, err := strconv.ParseFloat(i, 64)

	scanner.Scan()
	i1 := scanner.Text()
	l, err1 := strconv.ParseFloat(i1, 64)

	if err != nil || err1 != nil {
		fmt.Println("Insira números válidos")
	}

	t = (5*t - 160) / 9

	l = l * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f \n", t)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f \n", l)
}

func Ex8() {
	PI := 3.14159
	SQRM := 100

	scanner.Scan()
	i := scanner.Text()

	r, err := strconv.ParseFloat(i, 64)

	scanner.Scan()
	i = scanner.Text()

	h, err1 := strconv.ParseFloat(i, 64)

	if err != nil || err1 != nil {
		fmt.Println("Insira números válidos")
	}

	area := 2 * PI * r * (r + h)

	price := area * float64(SQRM)

	fmt.Printf("O VALOR DO CUSTO E = %.2f \n", price)
}

func Ex9() {
	scanner.Scan()
	i := scanner.Text()
	A, err := strconv.ParseFloat(i, 64)

	scanner.Scan()
	i = scanner.Text()
	B, err1 := strconv.ParseFloat(i, 64)

	scanner.Scan()
	i = scanner.Text()
	C, err2 := strconv.ParseFloat(i, 64)

	if err != nil || err1 != nil || err2 != nil {
		fmt.Println("Insira números válidos")
	}

	D := (B * B) - (4 * A * C)

	fmt.Printf("O VALOR DE DELTA E = %.2f \n", D)
}

func Ex10() {
	scanner.Scan()
	i := scanner.Text()
	A11, err := strconv.ParseFloat(i, 64)

	scanner.Scan()
	i = scanner.Text()
	A21, err1 := strconv.ParseFloat(i, 64)

	scanner.Scan()
	i = scanner.Text()
	A12, err2 := strconv.ParseFloat(i, 64)

	scanner.Scan()
	i = scanner.Text()
	A22, err3 := strconv.ParseFloat(i, 64)

	if err != nil || err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Insira números válidos")
	}

	det := (A11 * A22) - (A21 * A12)

	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f \n", det)

}

func Ex11() {
	scanner.Scan()
	i := scanner.Text()
	n, err := strconv.Atoi(i)

	if err != nil {
		fmt.Println("Insira números válidos")
	}

	if n%3 == 0 && n%5 == 0 {
		fmt.Println("O NUMERO E DIVISIVEL")
	}
}

func Ex12() {
	scanner.Scan()
	i := scanner.Text()
	n, err := strconv.Atoi(i)

	if err != nil {
		fmt.Println("Insira números válidos")
	}

	var price float64

	if n < 3 {
		price = float64(5 * n)
	} else {
		price = float64(n / 3 * 10)
	}

	fmt.Printf("O VALOR A PAGAR E = %.2f \n", price)
}

func Ex13() {
	scanner.Scan()
	i := scanner.Text()
	n, err := strconv.ParseFloat(i, 64)

	var concept string

	if err != nil {
		fmt.Println("Insira números válidos")
	}

	if n < 6 {
		concept = "D"
	} else if n >= 6 && n < 7.5 {
		concept = "C"
	} else if n >= 7.5 && n < 9 {
		concept = "B"
	} else {
		concept = "A"
	}

	fmt.Printf("NOTA = %.1f CONCEITO = %s \n", n, concept)
}

func Ex14() {
	scanner.Scan()
	i := scanner.Text()
	iSlice := strings.Fields(i)
	h, err := strconv.ParseFloat(iSlice[0], 64)
	a, err1 := strconv.ParseFloat(iSlice[1], 64)

	if err != nil || err1 != nil {
		fmt.Println("Insira números válidos")
	}

	v := a * a * h * math.Sqrt(float64(3)) / 2

	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS \n", v)
}

func Ex15() {
	scanner.Scan()
	i := scanner.Text()
	n, err := strconv.Atoi(i)

	if err != nil {
		fmt.Println("Insira números válidos")
	}

	evens := make([]int, 0)

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			evens = append(evens, i)
		}
	}

	for _, e := range evens {
		r := e * e
		fmt.Printf("%d^%d = %d \n", e, e, r)
	}
}

func Ex16() {
	scanner.Scan()
	i := scanner.Text()
	n, err := strconv.ParseFloat(i, 64)

	if err != nil {
		fmt.Println("Insira números válidos")
	}

	var newSalary float64

	if n <= 300 {
		newSalary = 1.5 * n
	} else {
		newSalary = 1.3 * n
	}

	fmt.Printf("SALARIO COM REAJUSTE = %.2f \n", newSalary)
}

func Ex17() {
	scanner.Scan()
	i := scanner.Text()
	iSlice := strings.Fields(i)
	x, err := strconv.Atoi(iSlice[0])
	y, err1 := strconv.Atoi(iSlice[1])

	if err != nil || err1 != nil {
		fmt.Println("Insira números válidos")
	}

	if x%2 != 0 {
		fmt.Println("O PRIMEIRO NUMERO NAO E PAR")
		return
	}

	evens := make([]int, 0)

	evens = append(evens, x)

	for i := range y {
		evens = append(evens, evens[i]+2)
		fmt.Printf("%d ", evens[i])
	}

	fmt.Println()
}

func Ex18() {
	scanner.Scan()
	i := scanner.Text()
	iSlice := strings.Fields(i)
	a1, err := strconv.Atoi(iSlice[0])
	r, err1 := strconv.Atoi(iSlice[1])
	n, err2 := strconv.Atoi(iSlice[2])

	if err != nil || err1 != nil || err2 != nil {
		fmt.Println("Insira números válidos")
	}

	numbers := make([]int, n)
	var sum int

	for i := range n {
		numbers[i] = a1 + i*r
	}

	for _, e := range numbers {
		sum = sum + e
	}

	fmt.Println(sum)
}

func Ex19() {
	scanner.Scan()
	i := scanner.Text()

	n, err := strconv.Atoi(i)
	if err != nil || n <= 1 {
		fmt.Println("Numero invalido!")
		return
	}

	var sum float64 = 0.0
	for k := 1; k <= n; k++ {
		sum += 1.0 / float64(k)
	}

	fmt.Printf("%.6f\n", sum)
}

func Ex20() {
	scanner.Scan()
	i := scanner.Text()
	h, err := strconv.Atoi(i)

	scanner.Scan()
	i = scanner.Text()
	m, err1 := strconv.Atoi(i)

	scanner.Scan()
	i = scanner.Text()
	s, err2 := strconv.Atoi(i)

	if err != nil || err1 != nil || err2 != nil {
		fmt.Println("Insira número válidos")
		return
	}

	var seconds int
	seconds = s

	seconds += m * 60

	seconds += h * 60 * 60

	fmt.Printf("O TEMPO EM SEGUNDOS E = %d \n", seconds)
}
