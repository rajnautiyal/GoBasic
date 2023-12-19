package helpers

import (
	"log"
)

func MapLearning() {
	am := map[string]int{
		"ram":  1,
		"raj":  2,
		"same": 3,
	}
	log.Println("this is ", am["ram"])
	am1 := make(map[string]int)
	am1["raj"] = 1
	am1["ram1"] = 3
	am1["ram2"] = 4
	am1["ram3"] = 5
	am1["ram4"] = 4
	log.Println("this is second way to create a map ", am1["ram1"])
	delete(am1, "raj")
	for k, v := range am1 {
		log.Println("key ", k, "value ", v)

	}

	//checking wether key exist in map or not if key exit it will give ok return else it will print
	if _, ok := am1["ram1"]; !ok {
		// Key "ram1" does not exist, so assigning value 4 to the key "ram4"
		am1["rama"] = 4
		log.Println("addin the rama in the list")
	}

}

func slice() {
	am := []string{"a", "b", "c"}
	log.Println("this is second way to create a map ", am[0])

	am1 := make([]string, 3, 4)
	am1[0] = "ram"
	log.Println("this is second way to create a map ", am[1])

}
