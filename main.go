package main

import (
	lista1 "IP-2025-01/lista-1"
	lista2 "IP-2025-01/lista-2"
	lista3 "IP-2025-01/lista-3"
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	listas := map[string]map[int]func(){
		"lista 1": {
			1:  lista1.Ex1,
			2:  lista1.Ex2,
			3:  lista1.Ex3,
			4:  lista1.Ex4,
			5:  lista1.Ex5,
			6:  lista1.Ex6,
			7:  lista1.Ex7,
			8:  lista1.Ex8,
			9:  lista1.Ex9,
			10: lista1.Ex10,
			11: lista1.Ex11,
			12: lista1.Ex12,
			13: lista1.Ex13,
			14: lista1.Ex14,
			15: lista1.Ex15,
			16: lista1.Ex16,
			17: lista1.Ex17,
			18: lista1.Ex18,
			19: lista1.Ex19,
			20: lista1.Ex20,
		},
		"lista 2": {
			1:  lista2.Ex1,
			2:  lista2.Ex2,
			3:  lista2.Ex3,
			4:  lista2.Ex4,
			5:  lista2.Ex5,
			6:  lista2.Ex6,
			7:  lista2.Ex7,
			8:  lista2.Ex8,
			9:  lista2.Ex9,
			10: lista2.Ex10,
			11: lista2.Ex11,
			12: lista2.Ex12,
			13: lista2.Ex13,
			14: lista2.Ex14,
			15: lista2.Ex15,
			16: lista2.Ex16,
			17: lista2.Ex17,
			18: lista2.Ex18,
			19: lista2.Ex19,
			20: lista2.Ex20,
			21: lista2.Ex21,
			22: lista2.Ex22,
			23: lista2.Ex23,
		},
		"lista 3": {
			1:  lista3.Ex1,
			2:  lista3.Ex2,
			3:  lista3.Ex3,
			4:  lista3.Ex4,
			5:  lista3.Ex5,
			6:  lista3.Ex6,
			7:  lista3.Ex7,
			8:  lista3.Ex8,
			9:  lista3.Ex9,
			10: lista3.Ex10,
			11: lista3.Ex11,
			12: lista3.Ex12,
			13: lista3.Ex13,
			14: lista3.Ex14,
			15: lista3.Ex15,
			16: lista3.Ex16,
			17: lista3.Ex17,
			18: lista3.Ex18,
			19: lista3.Ex19,
			20: lista3.Ex20,
			21: lista3.Ex21,
			22: lista3.Ex22,
			23: lista3.Ex23,
			24: lista3.Ex24,
			25: lista3.Ex25,
			26: lista3.Ex26,
			27: lista3.Ex27,
			28: lista3.Ex28,
			29: lista3.Ex29,
			30: lista3.Ex30,
			31: lista3.Ex31,
			32: lista3.Ex32,
			33: lista3.Ex33,
			34: lista3.Ex34,
			35: lista3.Ex35,
			36: lista3.Ex36,
			37: lista3.Ex37,
			38: lista3.Ex38,
			39: lista3.Ex39,
			40: lista3.Ex40,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Por favor, selecione uma lista dentre as presentes abaixo:")
		var listaNomes []string
		for l := range listas {
			listaNomes = append(listaNomes, l)
		}
		sort.Strings(listaNomes)
		for _, l := range listaNomes {
			fmt.Println(l)
		}
		fmt.Println("Digite 'sair' para encerrar o programa.")

		fmt.Println("----------")

		scanner.Scan()

		fmt.Println("----------")

		lista := scanner.Text()
		if lista == "sair" {
			break
		}
		if exercicios, found := listas[lista]; found {
			for {
				fmt.Println("Agora selecione um exercício dentre os presentes abaixo:")
				var exNums []int
				for ex := range exercicios {
					exNums = append(exNums, ex)
				}
				sort.Ints(exNums)
				for _, ex := range exNums {
					fmt.Println(ex)
				}
				fmt.Println("Digite 'voltar' para escolher outra lista.")

				fmt.Println("----------")

				scanner.Scan()
				exInput := scanner.Text()

				fmt.Println("----------")

				if exInput == "voltar" {
					break
				}

				
				exNum, err := strconv.Atoi(exInput)
				if err != nil {
					fmt.Println("Entrada inválida. Digite um número.")
					continue
				}
				if execFunc, exists := exercicios[exNum]; exists {
					execFunc()
					fmt.Println("----------")
				} else {
					fmt.Println("Exercício não encontrado.")
				}
			}
		} else {
			fmt.Println("A lista selecionada não é válida.")
		}
	}
}
