package extract

import (
  "fmt"
  "encoding/csv"
  "io"
)

// MeanScoresUF pega os dados de medias das notas de uma UF do arquivo CSV
func MeanScoresUF(reader *csv.Reader, UF string, finished chan bool) State{

  // notas de areas de conhecimento da UF
  scoresUF := [4][]float64{}  

  // notas de areas de conhecimento de cada raça
  scoresPerRace := [6][4][]float64{}

  // gera nova estrutura de Estado (UF)
  state := NewState(UF)
  
  // leitura de linha a linha do registro 
  for /* i := 0; i < 1000000; i++ */ { 
    recordLine, err := reader.Read()

    if err == io.EOF {
      break   // chegou ao final do registro
    } else if err != nil{ //checa por outros erros
      fmt.Println("An error encountered ::", err)
    }

    // campo de UF - 5
    if recordLine[5] == UF {
      state.totalParticipants++;

      // coleta as notas de cada disciplina de toda UF
      scoresUF = getScores(recordLine, scoresUF)

      // coleta dados por raça da UF
      scoresPerRace = getRacesData(recordLine, &state, scoresPerRace)
    }
  }

  state.medias = getMeanScores(scoresUF)

  for i := range state.races {
    state.races[i].medias = getMeanScores(scoresPerRace[i])
  }

  printUFMeanScores(state)
  printRacesMeanScores(state)

  // escreve true no canal
  finished <- true

  return state
}

// MeanScoresUF pega os dados de medias das notas de uma UF do arquivo CSV
func MeanScoresUF2(reader *csv.Reader, UF string) State{

  // notas de areas de conhecimento da UF
  scoresUF := [4][]float64{}  

  // notas de areas de conhecimento de cada raça
  scoresPerRace := [6][4][]float64{}

  // gera nova estrutura de Estado (UF)
  state := NewState(UF)
  
  // leitura de linha a linha do registro 
  for /* i := 0; i < 1000000; i++ */ { 
    recordLine, err := reader.Read()

    if err == io.EOF {
      break   // chegou ao final do registro
    } else if err != nil{ //checa por outros erros
      fmt.Println("An error encountered ::", err)
    }

    // campo de UF - 5
    if recordLine[5] == UF {
      state.totalParticipants++;

      // coleta as notas de cada disciplina de toda UF
      scoresUF = getScores(recordLine, scoresUF)

      // coleta dados por raça da UF
      scoresPerRace = getRacesData(recordLine, &state, scoresPerRace)
    }
  }

  state.medias = getMeanScores(scoresUF)

  for i := range state.races {
    state.races[i].medias = getMeanScores(scoresPerRace[i])
  }

  printUFMeanScores(state)
  printRacesMeanScores(state)

  return state
}
