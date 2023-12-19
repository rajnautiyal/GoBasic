package helpers

type Engine struct {
	electric bool
}

type Vechical struct {
	engine Engine
	make   string
	model  string
	door   string
	color  string
}

func testMyEmbedFun() {
	var myFun Vechical
	myFun.engine.electric = true
	myFun.make = "TATA"
	myFun.model = "Nano"
	myFun.door = "2 door"
}
