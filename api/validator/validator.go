package validator

import "fmt"

func errParamIsRequired(name, typ string) error {
	fmt.Println("Entrou no erro")
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}