package main

import (
	"log"

	"github.com/kvillamizar05/Curso_go/bd"
	"github.com/kvillamizar05/Curso_go/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la base de datos.")
		return
	}

	handlers.Manejadores()
}
