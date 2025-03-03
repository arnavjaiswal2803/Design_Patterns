package pizzas

type TomatoToppingDecorator struct {
	BasePizza IPizza
}

func (p *TomatoToppingDecorator) GetPrice() int {
	return p.BasePizza.GetPrice() + 10
}
