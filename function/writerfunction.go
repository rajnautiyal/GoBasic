package function

import (
	"log"
	"os"
)

/*
	type school struct {
		Name string
	}

func (s school) WriterFunc(w io.Writer) {

	t := school{
		Name: "raje",
	}
	fmt.Println(t.Name, s.Name)

}
*/
func WriteInFIile() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Println("there is some error in the file ")
	}
	defer f.Close()
	s := []byte("hello Rajendra")

	_, err = f.Write(s)
	if err != nil {
		log.Println("not able to write the file ")
	}
}
