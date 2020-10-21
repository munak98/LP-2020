package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"../Packages/extract"
	"golang.org/x/text/encoding/charmap"
)

func worker(finished chan bool) {
	fmt.Println("Worker: Started")
	time.Sleep(time.Second)
	fmt.Println("Worker: Finished")
	finished <- true
}

func main() {

	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	if len(os.Args) < 2 {
		fmt.Printf("No input file. Please supply a valid csv file \n")
		return
	}

	fmt.Printf("Reading: %s\n", fmt.Sprintf("../%s", os.Args[1]))

	csvFile, err := os.Open(fmt.Sprintf("../%s", os.Args[1]))

	if err != nil { //check for error in opening
		fmt.Println("An error encountered ::", err)
	}

	reader := csv.NewReader(charmap.ISO8859_1.NewDecoder().Reader(csvFile))
  reader.Comma = ';'
  
  finished := make(chan bool)
  
  go extract.MeanScoresUF(reader, csvFile, "DF", finished)

	<- finished

	return
}
