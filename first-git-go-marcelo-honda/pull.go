package main

type Pull struct {}

func Explain()(string, error) {
	retorno := "O comando git pull serve para baixar o conteúdo de repositórios remotos, fazendo uma atualização imediata no repositório local para que os conteúdos sejam iguais. Ele mistura o git fetch, que recupera as alterações do repositório remoto, e o git merge, que mescla as alterações do git fetch com o repositório local."	

	return retorno, nil	
}
