package factories

import (
	"errors"
	"factory/guns"
	"factory/interfaces"
)

type GunType string

const (
	AK47_GUN   GunType = "Ak47"
	MUSKET_GUN GunType = "Musket"
)

func GetGun(gunType GunType) (interfaces.IGun, error) {
	if gunType == AK47_GUN {
		return guns.NewAk47(), nil
	}
	if gunType == MUSKET_GUN {
		return guns.NewMusket(), nil
	}
	return nil, errors.New("invalid gun type")
}
