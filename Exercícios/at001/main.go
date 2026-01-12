package main

import(
	
	"fmt"
	"os"
	"encoding/json"
)

type Server struct {
	Name string  `json:"name"`
	IP string  `json:"ip"`
	Status bool `json:"status"`
}

func main(){

	servers := Server{
		Name: "servidor01",
		IP: "192.168.0.1",
		Status: true,
	}

data, err := json.MarshalIndent(servers,"","    ")
if err != nil{
	fmt.Println("Falha ao importar JSON:", err)
	return
}
err = os.WriteFile("servidores.json", data, 0644)
if err != nil{
	fmt.Println("Erro ao tentar salvar o arquivo", err)
	return
}
fmt.Println("Arquivo salvo com sucesso")
}