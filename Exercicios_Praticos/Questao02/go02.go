type GopherCabeludo struct {
	Nome string
	CorCabelo string
}

func (nome *GopherCabeludo) Falar() string {
	return fmt.Sprintf("%s com cabelo %s falou oi!", nome.Nome, nome.CorCabelo)
}