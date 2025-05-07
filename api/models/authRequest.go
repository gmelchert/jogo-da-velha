package models

import "fmt"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func errParamIsRequired(name, typ string) error {
	fmt.Println("Entrou no erro")
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

func (r *LoginRequest) ValidateLogin() error {
	if r == nil {
		return fmt.Errorf("body da requisição está malformado")
	}

	if r.Username == "" {
		return errParamIsRequired("username", "string")
	}

	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}

	return nil
}

func (r *SignUpRequest) ValidateSignUp() error {
	if r == nil {
		return fmt.Errorf("body da requisição está malformado")
	}

	if r.Username == "" {
		return errParamIsRequired("username", "string")
	}

	if r.Email == "" {
		return errParamIsRequired("email", "string")
	}

	if r.Password == "" {
		return errParamIsRequired("password", "string")
	}

	return nil
}
