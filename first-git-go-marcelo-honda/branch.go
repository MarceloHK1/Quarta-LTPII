package main

type Branch struct {}

func Explain()(string, error) {
	retorno := "O comando git branch serve para criar, listar, renomear e excluir ramificações com o nome da ramificação."

	return retorno, nil
}
