package main

import "fmt"

type AbstractSpaceObject interface {
	GetCoordinates() Coordinates
}

type Coordinates struct {
	X, Y, Z int
}

type Planet struct {
	Coordinates
	Radius int
}

func (p *Planet) GetCoordinates() Coordinates {
	return p.Coordinates
}

type Galactic map[string]AbstractSpaceObject

type SpaceMap map[string]Galactic

func main() {
	println("Lesson-03 – array, slice, map")

	// var NN SpaceMap

	MM := make(SpaceMap)

	pp := &Planet{
		Coordinates: Coordinates{
			X: 100,
			Y: 200,
			Z: 300,
		},
		Radius: 20000,
	}

	fmt.Printf("%#v\n", pp)

	MM["MilkyWay"] = make(Galactic)

	galaxy, ok := MM["MilkyWay"]
	if ok {
		pl, ok := galaxy["Earth"]
		if ok {
			// что-то полезное
			fmt.Println(pl)
		} else {
			// такой планеты нет
			fmt.Println("No Planet")
		}
	} else {
		fmt.Println("No galaxy")
		// такой галактики нет
	}

	// m := make(map[string]int)

	// m["aa"] = 1
	// m["bb"] = 2

	// value, ok := m["aa"]
	// if ok {
	// 	// полезная работа
	// } else {
	// 	// такого ключа нет
	// }

	// fmt.Println(value)
}

func MyPrint(s ...string) {
	for _, v := range s {
		fmt.Println(v)
	}
}
