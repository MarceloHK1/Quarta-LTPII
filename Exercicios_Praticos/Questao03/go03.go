package main
import "fmt"

type Gopher struct {
	Nome string
	CorCabelo string
	Acessorio string
}

func (gopher *Gopher) Falar() string {
	return fmt.Sprintf("%s com cabelo %s e usando %s falou oi!", gopher.Nome, gopher.CorCabelo, gopher.Acessorio)
}