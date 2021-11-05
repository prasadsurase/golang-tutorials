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

	m := map[string]person{
		"bond_james":   p1,
		"flemming_ian": p2,
	}

	for i, p := range m {
		fmt.Printf("Record: %v\n", i)
		fmt.Println(p.firstname)
		fmt.Println(p.lastname)
		for i, v := range p.favFlavours {
			fmt.Println(i, v)
		}
	}
}
