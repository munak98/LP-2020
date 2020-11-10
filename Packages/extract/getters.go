package extract

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"

	"github.com/montanaflynn/stats"
)

func getScore(recordLine []string, campo int) float64 {
	i, _ := strconv.ParseFloat(recordLine[campo], 64)
	return i
}

func getIntValue(recordLine []string, campo int) int {
	i, _ := strconv.Atoi(recordLine[campo])
	return i
}

func getScores(recordLine []string, scores [4][]float64) [4][]float64{ 
	scores[0] = append(scores[0], getScore(recordLine, 91)) // Ciencias da Natureza - campo 91
	scores[1] = append(scores[1], getScore(recordLine, 92)) // Ciencias Humanas - campo 92
	scores[2] = append(scores[2], getScore(recordLine, 93)) // Linguagens e Código - campo 93
	scores[3] = append(scores[3], getScore(recordLine, 94)) // Matemática - campo 94

	return scores
}

// Pega as notas de areas de conhecimento de cada UF
func getUFScores(recordLine []string, state *State) {
	state.Scores = getScores(recordLine, state.Scores)
	return
}

// Pega as notas de areas de conhecimento de Cada Raça
func getRaceScores(recordLine []string, state *State, raceType int) {
	state.Races[raceType].Scores = getScores(recordLine, state.Races[raceType].Scores)
	return
}

// Pega as medias das notas de cada area de conhecimento
func getMeanScores(scores [4][]float64) [4]float64 {
	meanScores := [4]float64{}

	for i := range scores {
		meanScores[i], _ = stats.Mean(scores[i])
	}

	return meanScores
}

// Pega os dados de cada Raça e poe na estrutura Estado
func getRacesData(
	recordLine []string,
	s *State,
) {

	// Tipo de cor/raça - campo 9
	raceType := getIntValue(recordLine, 9)

	// Tipo de escola - campo 17
	schoolType := getIntValue(recordLine, 17)

	switch raceType {
	case 0:
		s.Races[0].Total++ // Não informado
		getRaceScores(recordLine, s, raceType)
		getSchoolType(s, raceType, schoolType)

		break
	case 1:
		s.Races[1].Total++ // Branca
		getRaceScores(recordLine, s, raceType)
		getSchoolType(s, raceType, schoolType)

		break
	case 2:
		s.Races[2].Total++ // Preta
		getRaceScores(recordLine, s, raceType)
		getSchoolType(s, raceType, schoolType)

		break
	case 3:
		s.Races[3].Total++ // Parda
		getRaceScores(recordLine, s, raceType)
		getSchoolType(s, raceType, schoolType)

		break
	case 4:
		s.Races[4].Total++ // Amarela
		getRaceScores(recordLine, s, raceType)
		getSchoolType(s, raceType, schoolType)

		break
	case 5:
		s.Races[5].Total++ // Indigena
		getRaceScores(recordLine, s, raceType)
		getSchoolType(s, raceType, schoolType)

		break
	default:
		fmt.Println("Raça não reconhecida, possível E.T!")
		break
	}

	return
}

// pega os tipos de escola
func getSchoolType(s *State, raceType int, schoolType int) {

	switch schoolType {
	case 0: // Nao respondeu
		s.SchoolType[0]++
		s.Races[raceType].SchoolType[0]++
		break
	case 1: // Publica
		s.SchoolType[1]++
		s.Races[raceType].SchoolType[1]++
		break
	case 2: // Privada
		s.SchoolType[2]++
		s.Races[raceType].SchoolType[2]++
		break
	case 3: // Exterior
		s.SchoolType[3]++
		s.Races[raceType].SchoolType[3]++
		break
	default:
		fmt.Println("Algo errado, tipo de escola Invalido!")
		break
	}
}

func getData(
	reader *csv.Reader,
	states []State,
	scoresUF *[4][]float64,
	scoresPerRace *[6][4][]float64,
	done chan bool,
) {

	count := 0

	go func() {
		for /* i := begin; true ;i++ */ {
			recordLine, err := reader.Read()

			if err == io.EOF {
				break // chegou ao final do registro
			} else if err != nil { //checa por outros erros
				fmt.Println("An error encountered ::", err)
			}

			count++
			
			for i := range states {
				if states[i].Sigla == recordLine[5] {

					states[i].Total++

					// coleta as notas de cada area da UF
					getUFScores(recordLine, &states[i])

					// coleta dados de cada area por raça da UF
					getRacesData(recordLine, &states[i])
				}
			}
		}
	}()

	done <- true // recebe true no canal
}
