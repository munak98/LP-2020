package extract

import (
  "fmt"
  "encoding/csv"
  "io"
  "github.com/montanaflynn/stats"
)

type state struct {
	uf string
	medias []float64
	races []int
}

//MeanScoresUF gets means scores from a UF info from a csv file
func MeanScoresUF(reader *csv.Reader, UF string, finished chan bool) []float64{

  totalUFRecords := 0
  scores := [4][]float64{}
  races := []int{}

  //read file line by line
  for i := 0; i < 100000; i++{ 
    
    record, err := reader.Read()

    if err == io.EOF {
      break   // reached end of the file
    } else if err != nil{ //check for other errors
      fmt.Println("An error encountered ::", err)
    }

    if record[5] == UF {
      scores[0] = append(scores[0], getScore(record, 91))
      scores[1] = append(scores[1], getScore(record, 92))
      scores[2] = append(scores[2], getScore(record, 93))
      scores[3] = append(scores[3], getScore(record, 94))

      races = append(races, getIntValue(record, 9))

      totalUFRecords++;
    }
  }

  fmt.Printf("Total de registros analisados de %s: %d\n", UF, totalUFRecords)

  meanCN, _ := stats.Mean(scores[0])
  meanCH, _ := stats.Mean(scores[1])
  meanLC, _ := stats.Mean(scores[2])
  meanMT, _ := stats.Mean(scores[3])

  fmt.Printf("Médias de %s: \n\tCiências da natureza: %.2f \n\tCiências humanas: %.2f \n\tLinguagens e códigos: %.2f\n\tMatemática: %.2f\n", 
    UF, meanCN, meanCH, meanLC, meanMT)

  fmt.Printf("\tRaças: %v\n", races)

  // set channel to true
  finished <- true

  return []float64{meanCN, meanCH, meanLC, meanMT}
}
