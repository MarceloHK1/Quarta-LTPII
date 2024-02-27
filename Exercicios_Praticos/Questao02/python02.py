from Questao01 import Gopher


class GopherCabeludo(Gopher):
    def __init__(self, nome, cor_cabelo):
        super().__init__(nome)
        self.cor_cabelo = cor_cabelo

    def falar(self):
        return f"{self.nome} com cabelo {self.cor_cabelo} falou oi!"