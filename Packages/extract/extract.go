package extract

import (
  "fmt"
  "encoding/csv"
  "io"
)

type state struct {
	uf string
	medias []float64
	races []int
}

// MeanScoresUF pega os dados de medias das notas de cada UF do arquivo CSV
func MeanScoresUF(reader *csv.Reader, UF string, finished chan bool) [4]float64{

  // Total de participantes da UF
  participantsUF := 0

  // Notas de areas de conhecimento da UF
  scoresUF := [4][]float64{}    

  races := [6]string{
    "Raça Não Informada",
    "Raça Branca",
    "Raça Preta",
    "Raça Parda",
    "Raça Amarela",
    "Raça Indigena",
  }

  // Numero de participantes de cada raça da UF
  participantsPerRace := [6]int{0,0,0,0,0,0}

  // Notas de areas de conhecimento de cada raça
  scoresPerRace := [6][4][]float64{}  

  // Medias da Notas de areas de conhecimento de cada raça
  meanScoresPerRace := [6][4]float64{}  

  // Read file line by line
  for /* i := 0; i < 1000000; i++ */ { 
    recordLine, err := reader.Read()

    if err == io.EOF {
      break   // reached end of the file
    } else if err != nil{ //check for other errors
      fmt.Println("An error encountered ::", err)
    }

    if recordLine[5] == UF {
      scoresUF = getScores(recordLine, scoresUF)

      participantsPerRace, scoresPerRace = 
      getRacesData(
        recordLine, 
        participantsPerRace, 
        scoresPerRace,
      )

      participantsUF++;
    }
  }

  fmt.Printf("\nTotal de participantes - %s: %d\n", UF, participantsUF)

  meanScoresUF := getMeanScores(scoresUF)
  printMeanScores(UF, meanScoresUF)

  for i := range races {
    fmt.Printf("Total de participantes - %s: %d\n", races[i],participantsPerRace[i])
    meanScoresPerRace[i] = getMeanScores(scoresPerRace[i])
    printMeanScores(races[i], meanScoresPerRace[i])
  }

  // set channel to true
  finished <- true

  return meanScoresUF
}
