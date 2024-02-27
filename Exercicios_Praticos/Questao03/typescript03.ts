class Gopher {
    constructor(
        public nome: string,
        public corCabelo: string,
        public acessorio: string
    ){}

    falar(): string {
        return `${this.nome} com cabelo ${this.corCabelo} e usando ${this.acessorio} falou oi!`;
    }
}