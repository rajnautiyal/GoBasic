// Ninga five
package helpers

import "fmt"

type Person struct {
	FirstName      string
	LastName       string
	FavortIceCream []string
}

func StructWithSlice() {
	person1 := Person{
		FirstName:      "raj",
		LastName:       "nautiyal",
		FavortIceCream: make([]string, 0, 3),
	}
	person1.FavortIceCream = append(person1.FavortIceCream, "Chocolate", "Vanilla", "Strawberry")

	person2 := Person{
		FirstName:      "raj",
		LastName:       "nautiyal",
		FavortIceCream: []string{"straw", "banana"},
	}

	for _, v := range person1.FavortIceCream {
		fmt.Println(v)
	}

	for _, v := range person2.FavortIceCream {
		fmt.Println(v)
	}

	personMap := map[string]Person{
		person1.LastName: person1,
		person2.LastName: person2,
	}

	fmt.Println(personMap)

	personMap1 := make(map[string]Person, 0)
	personMap1["key1"] = person1
	personMap1["key2"] = person1

}
