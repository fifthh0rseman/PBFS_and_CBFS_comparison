package utils

import (
	"fmt"
	"time"
)

func SaveAndPrintElapsedTime(startTime time.Time, name string) time.Duration {
	elapsed := time.Since(startTime)
	fmt.Printf("%s took %s to execute.\n", name, elapsed)
	return elapsed
}
