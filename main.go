package main

import "github.com/rajnautiyal/myniceprogram/websitevalidator"

func main() {
	/*fmt.Println("hello raj")
	//log.Println("this is the programming for the oacjergs")
	var Wg sync.WaitGroup
	var myVar helpers.SomeType
	myVar.Name = "rajendra"
	myVar.Age = "12"

	//log.Println(myVar.Age, " now the name is ", myVar.Name)
	Wg.Wait()

	//helpers.MultipleDiminesionSlice()
	//helpers.MapLearning()
	//helpers.WordCountFrequency()
	//
	//helpers.StructWithSlice()
	//helpers.SliceString()
	go OSDetails(&Wg)
	fmt.Println("Goroutine : ", runtime.NumGoroutine())

	*/

	//checking the the Waitgroup to sync the threads

	/*
		var wg sync.WaitGroup
		wg.Add(1)

		helpers.WordCountFrequency()
		go helpers.OSDetails(&wg)
		// Wait for all goroutines to finish
		wg.Wait()
		fmt.Println("All workers have finished")
	*/

	//Race Codition
	//helpers.MutexRaceCodititon()

	//Atomic in Go

	//helpers.Atomic()

	//channel testing

	//channel.ChannelTest()

	//channel.ChannelUblockTest()

	//channel.RangeChannel()
	//channel.FainIn()
	//channel.Chan224()
	//sum := function.Sum(12, 13)
	//fmt.Println("this is the sum", sum)

	//function.TestMyInterface()
	//function.WriteInFIile()

	//function.TestRecursion()

	//fmt.Println(leedcodeproblm.NumDecodings("226"))

	//s := leedcodeproblm.Solution{}
	//result := s.NumDecodings("111111111111111111111111111111111111111111111")
	//fmt.Println(result)
	//result := leedcodeproblm.NumDecodings("223")

	//pointer.TestFun()
	//concurrency.TestRaceCondtion()

	//concurrency.TestAtomicCondtion()
	//words := []string{"abc", "aabc", "bc"}
	//leedcodeproblm.RedistributionCharacter(words)

	websitevalidator.WebValidator()
}
