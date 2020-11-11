package extract

import "fmt"

//PrintUFMeanScores printa dados acerca da UF
func PrintUFMeanScores(state State) {
	fmt.Println("---------------------------------------")
	fmt.Printf("\nDados de %s", state.Sigla)
	fmt.Printf("\nTotal de participantes: %d\n", state.Total)
	fmt.Println("Médias:")

	fmt.Printf("\tCiências da natureza: %.2f \n\tCiências humanas: %.2f \n\tLinguagens e códigos: %.2f\n\tMatemática: %.2f\n\n",
		state.Medias[0],
		state.Medias[1],
		state.Medias[2],
		state.Medias[3],
	)

	fmt.Printf("Numero de participantes de Escola Publica: %d\n", state.SchoolType[1])
	fmt.Printf("Numero de participantes de Escola Privada: %d\n", state.SchoolType[2])
}

//PrintRacesMeanScores printa dados acerca de cada raça
func PrintRacesMeanScores(state State) {

	for i := range state.Races {
		fmt.Println("---------------------------------------")
		fmt.Printf("\nDados de %s", state.Races[i].Name)
		fmt.Printf("\nTotal de participantes: %d\n", state.Races[i].Total)
		fmt.Printf("Médias: \n\tCiências da natureza: %.2f \n\tCiências humanas: %.2f \n\tLinguagens e códigos: %.2f\n\tMatemática: %.2f\n\n",
			state.Races[i].Medias[0],
			state.Races[i].Medias[1],
			state.Races[i].Medias[2],
			state.Races[i].Medias[3],
		)
		fmt.Printf("Numero de participantes de Escola Publica: %d\n", state.Races[i].SchoolType[1])
		fmt.Printf("Numero de participantes de Escola Privada: %d\n", state.Races[i].SchoolType[2])
	}

}

//MostParticipantsUF printa o estado com maior numero de participantes de uma UF
func MostParticipantsUF(states []State) {
	var maior = State{}

	for i := 0; i < len(states)-1; i++ {
		if states[i].Total < states[i+1].Total {
			maior = states[i+1]
		}
	}

	fmt.Printf("\nO Estado com maior numero de participantes: %s\n", maior.Sigla)
	fmt.Printf("Numero de participantes: %v\n", maior.Total)
}
