package pizzas

type MargaritaPizza struct {
}

func (p *MargaritaPizza) GetPrice() int {
	return 10
}

func NewMargaritaPizza() *MargaritaPizza {
	return &MargaritaPizza{}
}
