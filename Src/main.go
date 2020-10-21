package main

import (
	"fmt"
	"time"
	"../Packages/extract"
)

func main() {

	now := time.Now()
	// wait for all processes end to track the time passed
	defer func() {
		fmt.Println("\nTempo de execução:", time.Since(now))
	}()

	reader := extract.CsvReader()

	finished := make(chan bool)

	go extract.MeanScoresUF(reader, "DF", finished)

	// read channel
	<-finished

	return
}
