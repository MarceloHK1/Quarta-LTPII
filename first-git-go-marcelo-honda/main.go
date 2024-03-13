package main

import "fmt"

var comandos = map[string]ComandosFodas {
	"origin": RemoteAddOrigin{},
	"ramificacao": Branch{},
	"pull": Pull{},
}

func main() {
	var comand string
	fmt.Scanf("%s", &comand)

	retorno, _ := comandos[comand].Explain()
	fmt.Printf(retorno)
}
