// Aqui se crearan los handred

package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// Register Metodo para registrar un usuario
func Register(c echo.Context) error {
	m := &MessageResponse{}
	u := &User{}
	err := c.Bind(u)

	if err != nil {
		log.Printf("La estructura recibida no es valida")
		m.Type = "error"
		m.Message = "La estructura no es valida"
		return c.JSON(http.StatusBadRequest, m)
	}

	addUser(u)
	m.Type = "ok"
	m.Message = "Registrado Correctamente"
	return c.JSON(http.StatusCreated, m)
}

// Login Metodo para validar la cuenta del usuario
func Login(c echo.Context) error {
	m := &MessageResponse{}
	u := &User{}

	err := c.Bind(u)
	if err != nil {
		log.Printf("La estructura recibida no es valida")
		m.Type = "error"
		m.Message = "La estructura no es valida"
		return c.JSON(http.StatusBadRequest, m)
	}

	// Valida primero el lado errado, como buena practica
	if !login(u) {
		m.Type = "error"
		m.Message = "usuario o contraseña no validos"
		return c.JSON(http.StatusUnauthorized, m)
	}

	m.Type = "ok"
	m.Message = "Bienvenido"
	// JWT => investiga sobre los TOKENS
	// Se va a simular que se cera un JWT
	m.Data = token

	return c.JSON(http.StatusOK, m)
}
