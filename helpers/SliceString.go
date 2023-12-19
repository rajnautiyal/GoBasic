package helpers

import (
	"fmt"
	"strings"
)

func SliceString() {
	s := "rajendra nautiyal"
	sliceResult := strings.Split(s, " ")
	fmt.Println(sliceResult[0])
	// Converting the slice to an array

}
