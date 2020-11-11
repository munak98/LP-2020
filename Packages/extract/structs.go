package extract

//State - Estrutura de Estado (UF)
type State struct {
	Sigla      string
	Code       int
	Total      int
	Scores     [4][]float64 // 4 areas de conhecimento
	Medias     [4]float64
	Races      [6]Race // 6 raças ao total
	SchoolType [4]int
}

//Race - Estrutura de Raças
type Race struct {
	Name       string
	RaceType   int
	Total      int
	Scores     [4][]float64
	Medias     [4]float64
	SchoolType [4]int
}

//NewStates construtor de array de Estruturas de Estado (UF)
func NewStates() []State {

	siglas := []string{"AC", "AL", "AP", "AM", "BA", "CE", "DF", "ES", "GO", "MA", "MT", "MS", "MG",
		"PA", "PB", "PR", "PE", "PI", "RJ", "RN", "RS", "RO", "RR", "SC", "SP", "SE", "TO"}

	// gera array de estruturas de Estado (UFs)
	states := []State{}

	for i := range siglas {

		state := State{
			Races: [6]Race{
				Race{Name: "Raça Não Informada", RaceType: 0},
				Race{Name: "Raça Branca", RaceType: 1},
				Race{Name: "Raça Preta", RaceType: 2},
				Race{Name: "Raça Parda", RaceType: 3},
				Race{Name: "Raça Amarela", RaceType: 4},
				Race{Name: "Raça Indigena", RaceType: 5},
			},
		}

		states = append(states, state)
		states[i].Sigla = siglas[i]
		states[i].Total = 0
	}

	return states
}
