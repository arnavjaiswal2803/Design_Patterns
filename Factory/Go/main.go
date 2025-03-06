package main

import (
	"factory/factories"
	"factory/interfaces"
	"fmt"
)

func main() {
	ak47, _ := factories.GetGun(factories.AK47_GUN)
	musket, _ := factories.GetGun(factories.MUSKET_GUN)

	printGunDetails(ak47)
	printGunDetails(musket)
}

func printGunDetails(gun interfaces.IGun) {
	fmt.Println()
	fmt.Println(gun.GetName())
	fmt.Println(gun.GetPower())
}
