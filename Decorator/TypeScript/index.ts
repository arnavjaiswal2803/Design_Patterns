interface IPizza {
    getPrice() : number;
};

class MargaritaPizza implements IPizza{
    getPrice() : number {
        return 10;
    }
};

class VegDelightPizza implements IPizza {
    getPrice() : number {
        return 15;
    }
};

class TomatoTopping implements IPizza {
    basePizza : IPizza;

    constructor(basePizza : IPizza) {
        this.basePizza = basePizza;
    }

    getPrice() {
        return this.basePizza.getPrice() + 10;
    }
};

class JalapenoTopping implements IPizza {
    basePizza : IPizza

    constructor(basePizza : IPizza) {
        this.basePizza = basePizza;
    }

    getPrice(): number {
        return this.basePizza.getPrice() + 30;
    }
};

const margarita = new MargaritaPizza();

const tomatoToppingMargarita = new TomatoTopping(margarita);
console.log(tomatoToppingMargarita.getPrice())

const jalapenoAndTomatoMargarita = new JalapenoTopping(tomatoToppingMargarita);
console.log(jalapenoAndTomatoMargarita.getPrice());

const VegDelight = new VegDelightPizza();

const tomatoToppingVegDelight = new TomatoTopping(VegDelight);
console.log(tomatoToppingVegDelight.getPrice())

const jalapenoAndTomatoVegDelight = new JalapenoTopping(tomatoToppingVegDelight);
console.log(jalapenoAndTomatoVegDelight.getPrice());