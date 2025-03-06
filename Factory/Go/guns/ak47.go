package guns

type Ak47 struct {
	Gun
}

func NewAk47() *Ak47 {
	return &Ak47{
		Gun: NewGun("Ak47", 4),
	}
}
