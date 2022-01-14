package main

import (
	"fmt"

	"github.com/sergyomsilva/api-golang/database"
	"github.com/sergyomsilva/api-golang/routes"
)

func main() {

	database.DatabaseConection()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleResquest()
}
