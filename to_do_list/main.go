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
			toggleTodo()
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
		return
	}

	newTodo := Todo{
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
		fmt.Println("Você não possui tarefas!")
		return
	}

	for _, todo := range todos {

		status := "[]"

		if todo.Completed {
			status = "[x]"
		}

		fmt.Printf("%d. %s %s\n", todo.ID, status, todo.Title)
	}

}

func toggleTodo() {

	listTodos()
	fmt.Println("Digite o ID da tarefa para alterar o status:")
	var id int 

	_, err := fmt.Scanln(&id)

	bufio.NewReader(os.Stdin) .ReadString('\n')

	if err != nil {
		fmt.Println("Digite um número válido")
		return
	}

	found := false

	for i, _ := range todos {
		if todos[i].ID == id {
			todos[i].Completed = !todos[i].Completed // inverte o boolean
			fmt.Printf("Tarefa '%s' Atualizada!\n", todos[i].Title)
			found = true
			break
		}
	}

	if !found {

		fmt.Println("ID não encontrado")

	}

}




