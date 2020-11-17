package extract

//Year - Estrutura de ano do Enem
type Year struct {
	Year             int
	Total            int
	States           []State
	Races            [6]Race         // 6 raças ao total
	Medias           [4]float64      // 4 medias
	SchoolTypes      [4]int          // Numero de registros dos 4 tipos de Escola
	SchoolScores     [4][4][]float64 // 4 tipos de Escola, 4 areas, varias notas
	SchoolMeanScores [4][4]float64   // 4 tipos de Escola, 4 medias

	Recurse     Recurse
	CsvFilePath string
	Workers     int
}

//State - Estrutura de Estado (UF)
type State struct {
	Sigla            string
	Total            int
	Scores           [4][]float64
	Races            [6]Race
	Medias           [4]float64
	SchoolType       [4]int
	SchoolScores     [4][4][]float64
	SchoolMeanScores [4][4]float64
	Recurse          Recurse
}

//Race - Estrutura de Raça
type Race struct {
	Name             string
	RaceType         int
	Total            int
	Scores           [4][]float64
	Medias           [4]float64
	SchoolType       [4]int
	SchoolScores     [4][4][]float64
	SchoolMeanScores [4][4]float64
	Recurse          Recurse
}

//Recurse tem os nomes de recursos
type Recurse struct {
	Subjects [4]string // 4 areas de conhecimento
	Schools  [4]string // 4 tipos de Escola
}

//NewYears construtor de array de Anos do Enem
func NewYears() []Year {
	years := []Year{}
	csvFilePath17 := "../microdados_enem_2017/DADOS/MICRODADOS_ENEM_2017.csv"
	csvFilePath18 := "../microdados_enem_2018/DADOS/MICRODADOS_ENEM_2018.csv"
	csvFilePath19 := "../microdados_enem_2019/DADOS/MICRODADOS_ENEM_2019.csv"
	states17 := NewStates()
	states18 := NewStates()
	states19 := NewStates()
	races17 := NewRaces()
	races18 := NewRaces()
	races19 := NewRaces()
	recurse := NewRecurse()

	year17 := Year{
		Year:        2017,
		States:      states17,
		CsvFilePath: csvFilePath17,
		Total:       6731342, // total de registros
		Workers:     2,       // numero de processos, tem que ser um divisor do total de registros
		Races:       races17,
		Recurse:     recurse,
	}
	year18 := Year{
		Year:        2018,
		States:      states18,
		CsvFilePath: csvFilePath18,
		Total:       5513748,
		Workers:     12,
		Races:       races18,
		Recurse:     recurse,
	}
	year19 := Year{
		Year:        2019,
		States:      states19,
		CsvFilePath: csvFilePath19,
		Total:       5095271,
		Workers:     29,
		Races:       races19,
		Recurse:     recurse,
	}
	years = append(years, year17)
	years = append(years, year18)
	years = append(years, year19)
	return years
}

//NewStates construtor de array de Estruturas de Estado (UF)
func NewStates() []State {
	siglas := []string{"AC", "AL", "AP", "AM", "BA", "CE", "DF", "ES", "GO", "MA", "MT", "MS", "MG",
		"PA", "PB", "PR", "PE", "PI", "RJ", "RN", "RS", "RO", "RR", "SC", "SP", "SE", "TO"}
	states := []State{}
	races := NewRaces()
	recurse := NewRecurse()
	for i := range siglas {
		state := State{
			Races:    races,
			Recurse:  recurse,
		}
		states = append(states, state)
		states[i].Sigla = siglas[i]
		states[i].Total = 0
	}
	return states
}

//NewRaces construtor de array de Raças
func NewRaces() [6]Race {
	recurse := NewRecurse()
	races := [6]Race{
		Race{Name: "Raça Não Informada", RaceType: 0, Recurse: recurse},
		Race{Name: "Raça Branca", RaceType: 1, Recurse: recurse},
		Race{Name: "Raça Preta", RaceType: 2, Recurse: recurse},
		Race{Name: "Raça Parda", RaceType: 3, Recurse: recurse},
		Race{Name: "Raça Amarela", RaceType: 4, Recurse: recurse},
		Race{Name: "Raça Indigena", RaceType: 5, Recurse: recurse},
	}
	return races
}

//NewRecurse construtor de nomes de recursos
func NewRecurse() Recurse {
	subjects := [4]string{
		"Ciências da natureza",
		"Ciências humanas",
		"Linguagens e códigos",
		"Matemática",
	}
	schools := [4]string{
		"Não Respondeu",
		"Pública",
		"Privada",
		"Exterior",
	}
	recurse := Recurse{
		Subjects:     subjects,
		Schools:      schools,
	}

	return recurse
}
