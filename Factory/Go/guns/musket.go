package guns

type Musket struct {
	Gun
}

func NewMusket() *Musket {
	return &Musket{
		Gun: NewGun("Musket", 1),
	}
}
