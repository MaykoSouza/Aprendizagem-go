package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct{
	ID int
	Title string
	Completed bool
}

var todos []Todo 
var nextID	 int = 1

func main () {
	for {
		printMenu()
		escolha := readInput()

		switch escolha {
		case "1":
			addTodo()
		case "2":
			listTodos()
		case "3":
			markCompleted()
		case "4":
			deleteTodo()
		case "0":
			fmt.Println("Saindo...")
			return
		default:
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}

func printMenu() {
	fmt.Println("\n---- TO DO LIST ----")
	fmt.Println("1. Adicionar tarefa")
	fmt.Println("2. Listar tarefas")
	fmt.Println("3. Marcar tarefa como concluída")
	fmt.Println("4. Deletar tarefa")
	fmt.Println("0. Sair")
	fmt.Println("Escoha uma opção:")
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input , _ := reader.ReadString('\n')
	return strings.TrimSpace(input) // Remove espaços em brancos no início e no final de uma string 
}

func addTodo() {
	fmt.Println("Adicione um título a tarefa")
	title := readInput()

	if title == "" {
		fmt.Println("Adicione um título válido")
	}

	newTodo := todo{
		ID: nextID,
		Title: title,
		Completed: false, // por padrão é false mas eu preferi deixar explícito
	}

	todos = append(todos, newTodo)

	nextID++
	fmt.Println("Tarefa adicionanda com sucesso!")
}

func listTodos() {
	fmt.Println("\n Lista de tarefas")

	if len(todos) == 0 { // len lista todos os dados encontrados na estutura (lenght)
		fmt.Printlnl("Você não possui tarefas!")
		return
	}

	for _, todo := range todos {

		status := "[]"

		if todo.Completed {
			status := "[x]"
		}

		fmt.Printf("%d. %s %s\n")
	}

}






