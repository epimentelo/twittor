package main

import (
	"log"

	"github.com/epimentelo/twittor/bd"
	"github.com/epimentelo/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()

}
