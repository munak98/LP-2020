package extract

import (
  "fmt"
  "encoding/csv"
  "io"
)

//UFData pega os dados de Medias das notas de uma UF do arquivo CSV
func UFData(reader *csv.Reader, UF string, finished chan<- bool) State{

  // notas de areas de conhecimento da UF
  scoresUF := [4][]float64{}  

  // notas de areas de conhecimento de cada raça
  scoresPerRace := [6][4][]float64{}

  // gera nova estrutura de Estado (UF)
  state := NewState(UF)

  totalRecords := totalRecords(reader)

  worker1 := make(chan bool)
  worker2 := make(chan bool)
  
  getData(reader, &state, &scoresUF, &scoresPerRace, 0, totalRecords / 2,  worker1)
  getData(reader, &state, &scoresUF, &scoresPerRace, totalRecords / 2, totalRecords, worker2)

  // recebe 
  <-worker1
  <-worker2

  fmt.Println("Numero de registros analisados:", totalRecords)

  state.Medias = getMeanScores(scoresUF)

  for i := range state.Races {
    state.Races[i].Medias = getMeanScores(scoresPerRace[i])
  }

  printUFMeanScores(state)
  printRacesMeanScores(state)

  // escreve true no canal
  finished <- true

  return state
}

//UFDataNormal pega os dados de Medias das notas de uma UF do arquivo CSV - Sem Goroutine!
func UFDataNormal(reader *csv.Reader, UF string) State{

  // notas de areas de conhecimento da UF
  scoresUF := [4][]float64{}  

  // notas de areas de conhecimento de cada raça
  scoresPerRace := [6][4][]float64{}

  // gera nova estrutura de Estado (UF)
  state := NewState(UF)

  count := 0
  
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
      state.Total++;

      // coleta as notas de cada disciplina de toda UF
      getScores(recordLine, &scoresUF)

      // coleta dados por raça da UF
      getRacesData(recordLine, &state, &scoresPerRace)
    }
  }

  fmt.Println("Numero de registros analisados:", count)

  state.Medias = getMeanScores(scoresUF)

  for i := range state.Races {
    state.Races[i].Medias = getMeanScores(scoresPerRace[i])
  }

  printUFMeanScores(state)
  printRacesMeanScores(state)

  return state
}