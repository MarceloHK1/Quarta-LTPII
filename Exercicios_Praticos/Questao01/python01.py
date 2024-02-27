from abc import ABC, abstractclassmethod

class Gopher(ABC):
    def __init__(self, nome):
        self.nome = nome

    @abstractclassmethod
    def falar(self):
        return f"{self.nome} falou oi!"