package main

import "sync"

var mutex sync.Mutex
var count uint32 = 0

func getCount() uint32 {
	mutex.Lock()
	defer mutex.Unlock()
	var currentCount = count
	count = count + 1
	return currentCount
}
