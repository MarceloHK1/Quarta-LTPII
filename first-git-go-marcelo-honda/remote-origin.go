package main

type RemoteAddOrigin struct {}

func Explain()(string, error) {
	retorno := "O comando $git remote add origin https://gitlab.com/grupo/myproject.git serve para adicionar um origin, nome do repositório padrão em que se criou localmente através do clone, para o repositorio caso não existir nenhum."
	
	return retorno, nil

}


