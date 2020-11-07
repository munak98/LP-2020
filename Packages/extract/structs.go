package extract

//State - Estrutura de Estado (UF)
type State struct {
	uf               string
	totalParticipants int
	medias           [4]float64
	races            [6]Race	// 6 raças ao total 
}

//Race - Estrutura de Raças 
type Race struct {
	name     string
	raceType int
	total    int
	medias   [4]float64
}

//NewState construtor de Estrutura do Estado (UF)
func NewState (uf string) State {

	// Cria a Estrutura do Estado (UF)
  state := State {
    races: [6]Race {
      Race {name: "Raça Não Informada", raceType: 0},
      Race {name: "Raça Branca", raceType: 1},
      Race {name: "Raça Preta", raceType: 2},
      Race {name: "Raça Parda", raceType: 3},
      Race {name: "Raça Amarela", raceType: 4},
      Race {name: "Indigena", raceType: 5},
    },
  }

  state.uf = uf
  state.totalParticipants = 0

	return state
}