package extract

import (
	"fmt"
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
func getScores(recordLine []string, scores [4][]float64) [4][]float64 {
	scores[0] = append(scores[0], getScore(recordLine, 91)) // Ciencias da Natureza - campo 91
	scores[1] = append(scores[1], getScore(recordLine, 92)) // Ciencias Humanas - campo 92
	scores[2] = append(scores[2], getScore(recordLine, 93)) // Linguagens e Código - campo 93
	scores[3] = append(scores[3], getScore(recordLine, 94)) // Matemática - campo 94

	return scores
}

// Pega as medias das notas de cada area de conhecimento
func getMeanScores(scores [4][]float64) ([4]float64) {
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
	scoresPerRace [6][4][]float64,
	) ([6][4][]float64) {

	// Tipo de cor/raça - campo 9
	switch getIntValue(recordLine, 9) {
	case 0:
		s.races[0].total++ // Não informado
		scoresPerRace[0] = getScores(recordLine, scoresPerRace[0])
		break
	case 1:
		s.races[1].total++ // Branca
		scoresPerRace[1] = getScores(recordLine, scoresPerRace[1])
		break
	case 2:
		s.races[2].total++ // Preta
		scoresPerRace[2] = getScores(recordLine, scoresPerRace[2])
		break
	case 3:
		s.races[3].total++ // Parda
		scoresPerRace[3] = getScores(recordLine, scoresPerRace[3])
		break
	case 4:
		s.races[4].total++ // Amarela
		scoresPerRace[4] = getScores(recordLine, scoresPerRace[4])
		break
	case 5:
		s.races[5].total++ // Indigena
		scoresPerRace[5] = getScores(recordLine, scoresPerRace[5])
		break
	default:
		fmt.Println("Raça não reconhecida, possível E.T!")
		break
	}

	return scoresPerRace
}
