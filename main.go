package main

import (
	"fmt"

	"github.com/guicalmeida/goCMS/database"
)

func main() {
	database.ConnectToDB()
	fmt.Println("iniciando servidor REST com Go")
}
