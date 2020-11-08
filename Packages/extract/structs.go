package extract

//State - Estrutura de Estado (UF)
type State struct {
	UF         string
	Total      int
	Medias     [4]float64
	Races      [6]Race // 6 raças ao total
	SchoolType [4]int
}

//Race - Estrutura de Raças
type Race struct {
	Name       string
	RaceType   int
	Total      int
	Medias     [4]float64
	SchoolType [4]int
}

//NewState construtor de Estrutura do Estado (UF)
func NewState(UF string) State {

	// Cria a Estrutura do Estado (UF)
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

	state.UF = UF
	state.Total = 0

	return state
}
