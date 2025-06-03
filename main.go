package main

import (
	"countries_backend/handler"
	"github.com/labstack/echo/v4"
	
)

func main() {
	//Crea la nueva instancia de echo
	e := echo.New()

	// Rutas
	e.GET("/countries", handler.GetCountries)

	// Lenvantamos el servidor
	e.Logger.Fatal(e.Start(":8080"))
}
