package models

type User struct {
	ID       uint   
	Username string 
	Password string // Hashé (en prod)
	Email    string 
}
