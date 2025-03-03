class Context {
    strategyObj : Strategy;
    
    constructor(strategy: Strategy) {
        this.strategyObj = strategy;
    }
    
    setStrategyObj(strategy : Strategy) {
        this.strategyObj = strategy;
    }
    
    do() {
        console.log("Doing something. Don't know how.");
        this.strategyObj.do(['h', 'e', 'q', 'i', 'o']);
    }
}

interface Strategy {
    do(data : string[]) : string[];
}
class ConcreteStrategyA implements Strategy {
    do(data : string[]) : string[] {
        console.log("sorting in normal order");
        data.sort();
        console.log(data);
        return data
    }
}

class ConcreteStrategyB implements Strategy {
    do(data : string[]) : string[] {
        console.log("sorting in reverse order");
        data.sort().reverse();
        console.log(data);
        return data;
    }
}


const context = new Context(new ConcreteStrategyA());
context.do();
context.setStrategyObj(new ConcreteStrategyB());
context.do();