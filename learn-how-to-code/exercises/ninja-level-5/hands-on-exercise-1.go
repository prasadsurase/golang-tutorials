package main

import "fmt"

type person struct {
	firstname   string
	lastname    string
	favFlavours []string
}

func main() {
	p1 := person{
		firstname: "James", lastname: "Bond", favFlavours: []string{"martini", "carrot"},
	}

	p2 := person{
		firstname: "Ian", lastname: "Flemming", favFlavours: []string{"cigar"},
	}

	fmt.Println(p1.firstname)
	fmt.Println(p1.lastname)
	for i, v := range p1.favFlavours {
		fmt.Println(i, v)
	}

	fmt.Println(p2.firstname)
	fmt.Println(p2.lastname)
	for i, v := range p2.favFlavours {
		fmt.Println(i, v)
	}
}
