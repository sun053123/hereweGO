package main

import "fmt"

/*
4 main feature of Object Oreinted Programming
	1. Abstraction
	2. Encapsulation
	3. Polymopism
	4. Inheritance <- ใน golang ไม่มี จะใช้ Composition
*/
type RegularHuman interface {
	chill()
}

type Ninja struct {
	human RegularHuman
}

func (Ninja) chill() {
	fmt.Println("Chilling as a Ninja ...")
}

func (Ninja) attack() {
	fmt.Println("Throwing Ninja Stars.")
}

type SeniorNinja struct {
	ninja Ninja
	human RegularHuman
}

func (sn SeniorNinja) attack() {
	sn.ninja.attack()
	fmt.Println("Swinging Ninja Swords")
	sn.human.chill()

}
