package main

import "testing"

func Test_getCount(t *testing.T) {
	output := make(chan uint32)
	go getCount_Tester(output, 10)
	go getCount_Tester(output, 10)
	for i := 0; i < 20; i++ {
		value := <-output
		if i != int(value) {
			t.Errorf("getCount() = %v, want %v", value, i)
		}
	}
}

func getCount_Tester(output chan<- uint32, reads int) {
	for i := 0; i < reads; i++ {
		output <- getCount()
	}
}
