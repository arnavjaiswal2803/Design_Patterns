package main

import (
	"decorator/pizzas"
	"fmt"
)

func main() {
	margarita := pizzas.NewMargaritaPizza()

	tomatoToppingMargaritaPizza := &pizzas.TomatoToppingDecorator{
		BasePizza: margarita,
	}
	fmt.Println(tomatoToppingMargaritaPizza.GetPrice())

	tomatoAndJalapenoToppingMargaritaPizza := &pizzas.JalapenoToppingDecorator{
		BasePizza: tomatoToppingMargaritaPizza,
	}
	fmt.Println(tomatoAndJalapenoToppingMargaritaPizza.GetPrice())

	vegDelight := &pizzas.VegDelightPizza{}

	tomatoToppingVegDelightPizza := &pizzas.TomatoToppingDecorator{
		BasePizza: vegDelight,
	}
	fmt.Println(tomatoToppingVegDelightPizza.GetPrice())

	tomatoAndJalapenoToppingVegDelightPizza := &pizzas.JalapenoToppingDecorator{
		BasePizza: tomatoToppingVegDelightPizza,
	}
	fmt.Println(tomatoAndJalapenoToppingVegDelightPizza.GetPrice())
}
