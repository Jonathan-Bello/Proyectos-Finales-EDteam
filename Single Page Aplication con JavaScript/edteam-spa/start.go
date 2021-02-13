package main

import (

	// Libretia para crear un servidor web
	"log"

	"github.com/labstack/echo"

	"github.com/labstack/echo/middleware"
	// "github.com/labstack/gommon/log"
)

// startServer lenvantara el servidor web
func startServer() {
	// e contiene el servidor
	e := echo.New()

	// Para logear los eventos del servidor
	e.Use(middleware.Logger())

	// Para si se llega a caer la aplicación se recupere
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// Aqui registramos todos los dominios que voy a poermitir que se conecten a mi servidor
		// AllowOrigins: []string{"midominio.com","midominio.com:9191"},

		// El * es para permitir todos los dominios, esto solo se hace para pruebas o practicas
		AllowOrigins: []string{"*"},
		// Los metodos https que queremos que permita nuestro servidor
		AllowMethods: []string{echo.GET, echo.POST, echo.DELETE, echo.PUT},
	}))

	apiRoute(e)
	socketRoute(e)

	log.Fatal(e.Start(":9393"))
}

func apiRoute(e *echo.Echo)  {
	// Creamos nuestros end-point
	e.POST("/api/v1/login",Login)
	e.POST("/api/v1/register",Register)
}

func socketRoute(e *echo.Echo)  {
	// Para conectar un web socket debemos de hacerlo por el metodo GET
	e.GET("/ws", WebSocket)
}
