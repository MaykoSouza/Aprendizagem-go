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
var servidores []server

func main() {
	leitorDoArquivo()
}

func leitorDoArquivo() {
	arquivo, err := os.ReadFile("servidores.json")
	if err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		return
	}

	err = json.Unmarshal(arquivo, &servidores)
	if err != nil {
		fmt.Println("Erro ao fazer o unmarshal:", err)
		return
	}
	fmt.Printf("Servidores lidos: %v\n", servidores)
}