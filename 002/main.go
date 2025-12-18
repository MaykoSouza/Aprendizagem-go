package main

import "fmt"

func adicao(num1, num2 float64) float64      { return num1 + num2 }
func subtracao(num1, num2 float64) float64   { return num1 - num2 }
func multiplicao(num1, num2 float64) float64 { return num1 * num2 }
func divisao(num1, num2 float64) (float64, string) {

	if num2 == 0 {
		return 0, "não é possível fazer divisões por zero"
	}
	return num1 / num2, ""
}

func main() {

	var opcao int
	var n1, n2 float64

	for {

		fmt.Println("\n--- CALCULADORA EM GO ---")
		fmt.Println("1. Adição")
		fmt.Println("2. Subtração")
		fmt.Println("3. Multiplicação")
		fmt.Println("4. Divisão")
		fmt.Println("0. Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scan(&opcao)

		if opcao == 0 {
			fmt.Println("Saindo...")
			break
		}

		if opcao >= 1 && opcao <= 4 {
			
			fmt.Println("Digite o primeiro número")
			fmt.Scan(&n1)
			fmt.Println("Digite o Segundo número")
			fmt.Scan(&n2)
		}
		switch opcao {
		case 1:
			fmt.Println("|---------------------------|")
			fmt.Println("     Resultado:", adicao(n1, n2))
			fmt.Println("|---------------------------|")
		case 2:
			fmt.Println("|---------------------------|")
			fmt.Println("     Resultado:", subtracao(n1, n2))
			fmt.Println("|---------------------------|")
		case 3:
			fmt.Println("|---------------------------|")
			fmt.Println("     Resultado:", multiplicao(n1, n2))
			fmt.Println("|---------------------------|")
		case 4:
			fmt.Println("|---------------------------|")
			res, erro := divisao(n1, n2)
			if erro != "" {
				fmt.Println(erro)
			} else {
				fmt.Println("      Resultado:", res)
			}
			fmt.Println("|---------------------------|")

		default:
			fmt.Println("Opção inválida! Tente novamente.")

		}
	}

}
