package guns

type Gun struct {
	Name  string
	Power int
}

func NewGun(name string, power int) Gun {
	return Gun{
		Name:  name,
		Power: power,
	}
}

func (g *Gun) GetName() string {
	return g.Name
}

func (g *Gun) SetName(name string) {
	g.Name = name
}

func (g *Gun) GetPower() int {
	return g.Power
}

func (g *Gun) SetPower(power int) {
	g.Power = power
}
