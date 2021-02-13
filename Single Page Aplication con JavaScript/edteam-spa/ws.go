package main

import (
	"encoding/json"
	"log"

	"github.com/labstack/echo"
	"github.com/olahol/melody"
)

/* var mel *melody.Melody */
var mel = melody.New() // se actualiza var mel *melody.Melody

// MessageWS estructura para los mensajes del WebSocket
type MessageWS struct {
	Type string
	From string
	Data interface{} // => interface{} es un objeto general, una estructura sin campos especificos
}

// WebSocket funcion para detectar los estados de WS
func WebSocket(c echo.Context) error {
	// Para crear websocjets, incluso los creadores de go recomiendan
	// Usar paquetes de terceros por la complejidad del websocket
	// Aqui se ocupara una llamada melody

	// Con esto registramos la petición
	mel.HandleRequest(c.Response().Writer, c.Request())
	// Con esto reacionamos a la coneción
	mel.HandleConnect((hConnect))
	// Con esto reaccionamos a la desconección
	mel.HandleDisconnect(hDisconnect)
	// detectamos cuando se manda un mensaje
	mel.HandleMessage(hMessage)

	return nil
}

func hConnect(s *melody.Session) {
	if !valideToken(s) {
		return
	}
	nick := getNickFromURL(s)
	m := &MessageWS{
		Type: "Connect",
		From: nick,
		Data: "Conectado",
	}
	sendMessage(m)
}

func hDisconnect(s *melody.Session) {
	if !valideToken(s) {
		return
	}
	nick := getNickFromURL(s)
	m := &MessageWS{
		Type: "Disconnect",
		From: nick,
		Data: "Desconectado",
	}

	sendMessage(m)
}

func hMessage(s *melody.Session, msg []byte) {
	// Verificar quien envio el mensaje
	nick := getNickFromURL(s)

/* 	var data map[string]interface{}
	err := json.Unmarshal(msg, data) */
	var data = make(map[string]interface{}) //se corrige var data map[string]interface{}
	err := json.Unmarshal(msg, &data)       //se corrige err := json.Unmarshal(msg, data)

	if err != nil {
		log.Printf("No se pudo procesar el mensaje %v", err)
		return
	}

	if data["type"] == "ping" {
		mel.BroadcastFilter([]byte("pong"), func(ss *melody.Session) bool {
			return ss == s
		})
	}

	m := &MessageWS{
		Type: data["type"].(string),
		From: nick,
		Data: data["data"].(string),
	}

	sendMessage(m)
}

func valideToken(s *melody.Session) bool {
	t := s.Request.URL.Query().Get("token")
	return t == token
}

func getNickFromURL(s *melody.Session) string {
	return s.Request.URL.Query().Get("nick")
}


func sendMessage(m *MessageWS) {
	j, err := json.Marshal(m)
	if err != nil {
		log.Printf("No se pudo convertir mensaje a json %v", err)
		return
	}

	mel.Broadcast(j)
}
