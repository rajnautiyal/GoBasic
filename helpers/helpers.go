package helpers

import "fmt"

type SomeType struct {
	Name string
	Age  string
}

func DeleteFromSlice() {
	xi := []int{1, 2, 3, 4, 5, 6}
	var len int = len(xi)
	xi = append(xi[0:4], xi[5:len]...)
}

func MultipleDiminesionSlice() {

	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mn := []string{"Miss", "Moneypenny", "I'm 008"}

	xi := [][]string{jb, mn}

	for _, v := range xi {

		for _, t := range v {
			fmt.Println("value of ", t)
		}
	}
}
