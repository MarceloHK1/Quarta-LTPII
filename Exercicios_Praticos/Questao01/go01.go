package main

import "fmt"

type Gopher interface {
	Falar() string
}

type NomeGopher struct {
	Nome string
}

func (nome *NomeGopher) Falar() string {
	return fmt.Sprintf("%s falou oi!", nome.Nome)
}