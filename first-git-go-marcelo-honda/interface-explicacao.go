package main

type explicacao interface {
	Explain()(string, error);
}
