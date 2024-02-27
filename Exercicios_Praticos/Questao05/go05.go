package main

import "fmt"
import "time"

type GopherVoador interface {
	Voar(segundos int)
}

type GopherAtirador interface {
	Atirar(dano int)
}
