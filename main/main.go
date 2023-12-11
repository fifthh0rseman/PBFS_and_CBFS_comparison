package main

import (
	"BFS/consecutive"
	"BFS/utils"
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(4)

	var i int
	fmt.Println("Consecutive - 1, parallel - 2")
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		fmt.Println("Error in scanf:", err)
		return
	}

	start := 0
	finish := 80496
	dim := 500

	if i == 1 {
		fmt.Println("Consecutive BFS is chosen.")
		startTimeCons := time.Now()
		consecutive.BFSCube(dim, start, finish)
		_ = utils.SaveAndPrintElapsedTime(startTimeCons, "Consecutive BFS")
	} else if i == 2 {
		fmt.Println("Parallel BFS is chosen.")
		startTimeCons := time.Now()
		//todo
		_ = utils.SaveAndPrintElapsedTime(startTimeCons, "Parallel BFS")
	} else {
		fmt.Println("Wrong input. Try again.")
	}
}
