package main

import(
	"fmt"
	"os"
	"encoding/json"

)

type server struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Status bool `json:"status"`

}

func main() {
	leitorDoArquivo()
}

func leitorDoArquivo() {
	arquivo, err := os.ReadFile("server.json")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}
	
}