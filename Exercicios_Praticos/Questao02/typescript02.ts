import { Gopher } from "../Questao01/typescript01";

class GopherCabeludo extends Gopher {
    constructor(nome: string, public corCabelo: string) {
        super(nome);
    }

    falar(): string {
        return `${this.nome} com cabelo {this.corCabelo} falou oi!`;
    }
}