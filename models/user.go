package models

type UserRegister struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRegisterSuccess struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
