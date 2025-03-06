package interfaces

type IGun interface {
	GetName() (name string)
	SetName(name string)
	GetPower() (power int)
	SetPower(power int)
}
