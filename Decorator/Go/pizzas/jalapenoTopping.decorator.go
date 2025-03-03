package pizzas

type JalapenoToppingDecorator struct {
	BasePizza IPizza
}

func (p *JalapenoToppingDecorator) GetPrice() int {
	return p.BasePizza.GetPrice() + 30
}
