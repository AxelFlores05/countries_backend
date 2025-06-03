package main

import (
	"countries_backend/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	
)

func main() {
	//Crea la nueva instancia de echo
	e := echo.New()

	// Rutas
	e.GET("/countries", handler.GetCountries)

	//CORS
	e.Use(middleware.CORS())

	// Lenvantamos el servidor
	e.Logger.Fatal(e.Start(":8080"))
}
