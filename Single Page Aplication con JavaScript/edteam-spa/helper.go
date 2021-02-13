package main

const token = "TOKENMUYSEGUROQUENADIEPUEDEVIOLENTAR"

var users []*User

// User estructura de usuario
type User struct{
	Nick string
	Password string
}

// MessageResponse estructura para responder al cliente
type MessageResponse struct {
	Type string `json:"type"`
	From string `json:"from"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}


func addUser(u *User)  {
	// Es el equivalente a push
	users = append(users, u)
}

// login comprueba si el usuario existe
func login(u *User) bool {
	for _, v := range users {
		if v.Nick == u.Nick  && v.Password == u.Password{
			return true
		}
	}
	return false
}
