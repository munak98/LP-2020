package extract

import (
  "fmt"
  "encoding/csv"
  "os"
  "io"
  "github.com/montanaflynn/stats"
)

//MeanScoresUF gets means scores from a UF info from a csv file
func MeanScoresUF(reader *csv.Reader, file *os.File, UF string, finished chan bool) []float64{
  scores := [4][]float64{}
  totalUFRecords := 0

  //read file line by line
  for i := 0; i < 100; i++{ 
    
    record, err := reader.Read()

    if err == io.EOF {
      break   // reached end of the file
    } else if err != nil{ //check for other errors
      fmt.Println("An error encountered ::", err)
    }

    if fromUF(record, UF){
      scores[0] = append(scores[0], getScoreCN(record))
      scores[1] = append(scores[1], getScoreCH(record))
      scores[2] = append(scores[2], getScoreLC(record))
      scores[3] = append(scores[3], getScoreMT(record))

      totalUFRecords++;
    }

  }

  fmt.Printf("Total de registros analisados de %s: %d\n", UF, totalUFRecords)

  meanCN, _ := stats.Mean(scores[0])
  meanCH, _ := stats.Mean(scores[1])
  meanLC, _ := stats.Mean(scores[2])
  meanMT, _ := stats.Mean(scores[3])

  fmt.Printf("%s médias: \n\tCiências da natureza: %.2f \n\tCiências humanas: %.2f \n\tLinguagens e códigos: %.2f\n\tMatemática: %.2f\n", UF, meanCN, meanCH, meanLC, meanMT)
	finished <- true

  return []float64{meanCN, meanCH, meanLC, meanMT}
}
