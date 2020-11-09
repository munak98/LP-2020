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

// Pega as notas de areas de conhecimento
func getScores(recordLine []string, scores *[4][]float64) { 
	scores[0] = append(scores[0], getScore(recordLine, 91)) // Ciencias da Natureza - campo 91
	scores[1] = append(scores[1], getScore(recordLine, 92)) // Ciencias Humanas - campo 92
	scores[2] = append(scores[2], getScore(recordLine, 93)) // Linguagens e Código - campo 93
	scores[3] = append(scores[3], getScore(recordLine, 94)) // Matemática - campo 94

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
	scoresPerRace *[6][4][]float64,
) {

	// Tipo de cor/raça - campo 9
	raceType := getIntValue(recordLine, 9)

	// Tipo de escola - campo 17
	schoolType := getIntValue(recordLine, 17)

	switch raceType {
	case 0:
		s.Races[0].Total++ // Não informado
		getScores(recordLine, &scoresPerRace[0])
		getSchoolType(s, 0, schoolType)

		break
	case 1:
		s.Races[1].Total++ // Branca
		getScores(recordLine, &scoresPerRace[1])
		getSchoolType(s, 1, schoolType)

		break
	case 2:
		s.Races[2].Total++ // Preta
		getScores(recordLine, &scoresPerRace[2])
		getSchoolType(s, 2, schoolType)

		break
	case 3:
		s.Races[3].Total++ // Parda
		getScores(recordLine, &scoresPerRace[3])
		getSchoolType(s, 3, schoolType)

		break
	case 4:
		s.Races[4].Total++ // Amarela
		getScores(recordLine, &scoresPerRace[4])
		getSchoolType(s, 4, schoolType)

		break
	case 5:
		s.Races[5].Total++ // Indigena
		getScores(recordLine, &scoresPerRace[5])
		getSchoolType(s, 5, schoolType)

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
	s *State,
	scoresUF *[4][]float64,
	scoresPerRace *[6][4][]float64,
	count int, 
	done chan bool,
) {

	go func() {
		for /* i := begin; true ;i++ */ {
			recordLine, err := reader.Read()
			
			if err == io.EOF {
				break // chegou ao final do registro
			} else if err != nil { //checa por outros erros
				fmt.Println("An error encountered ::", err)
			}
				
			// campo de UF = 5
			if recordLine[5] == s.Sigla {
				s.Total++
				
				// coleta as notas de cada disciplina de toda UF
				getScores(recordLine, scoresUF)
				
				// coleta dados por raça da UF
				getRacesData(recordLine, s, scoresPerRace)
			}
			count++
		}
	}()

	done <- true // recebe true no canal
}
