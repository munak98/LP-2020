package main

import (
  "fmt"
  "encoding/csv"
  "os"
  "golang.org/x/text/encoding/charmap"
  "github.com/munak98/LP-2020/Packages/extract"
  )

func main() {
    if len(os.Args) < 2 {
      fmt.Printf("No input file. Please supply a valid csv file \n");
      return
    }

    fmt.Printf("Reading: %s\n", fmt.Sprintf("../%s", os.Args[1]))

    csv_file, err := os.Open(fmt.Sprintf("../%s", os.Args[1]))
    if err != nil { //check for error in opening
      fmt.Println("An error encountered ::", err)
    }
    reader := csv.NewReader(charmap.ISO8859_1.NewDecoder().Reader(csv_file))
    reader.Comma = ';'

    extract.MeanScoresUF(reader, csv_file, "DF")

    return
  }
