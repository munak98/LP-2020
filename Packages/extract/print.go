package extract

import "fmt"

//PrintUFMeanScores printa dados acerca da UF
func PrintUFMeanScores(state State) {
	fmt.Println("---------------------------------------")
	fmt.Printf("\nDados de %s", state.Sigla)
	fmt.Printf("\nTotal de participantes: %d\n", state.Total)
	fmt.Printf("Número de participantes de Escola Publica: %d\n", state.SchoolType[1])
	fmt.Printf("Número de participantes de Escola Privada: %d\n", state.SchoolType[2])
	fmt.Println("Médias:")

	fmt.Printf("\tCiências da natureza: %.2f \n\tCiências humanas: %.2f \n\tLinguagens e códigos: %.2f\n\tMatemática: %.2f\n\n",
		state.Medias[0],
		state.Medias[1],
		state.Medias[2],
		state.Medias[3],
	)
}

//PrintRacesMeanScores printa dados acerca de cada raça
func PrintRacesMeanScores(state State, i int) {
		fmt.Printf("\nDados de %s", state.Races[i].Name)
		fmt.Printf("\nTotal de participantes: %d\n", state.Races[i].Total)
		fmt.Printf("Número de participantes de Escola Pública: %d\n", state.Races[i].SchoolType[1])
		fmt.Printf("Número de participantes de Escola Privada: %d\n", state.Races[i].SchoolType[2])
		fmt.Printf("Médias: \n\tCiências da natureza: %.2f \n\tCiências humanas: %.2f \n\tLinguagens e códigos: %.2f\n\tMatemática: %.2f\n\n",
			state.Races[i].Medias[0],
			state.Races[i].Medias[1],
			state.Races[i].Medias[2],
			state.Races[i].Medias[3],
		)
}

//MostParticipantsUF printa o estado com maior Número de participantes de uma UF
func MostParticipantsUF(states []State) {
	var maior = State{}

	for i := 0; i < len(states)-1; i++ {
		if states[i].Total < states[i+1].Total {
			maior = states[i+1]
		}
	}

	fmt.Printf("\nO Estado com maior número de participantes: %s\n", maior.Sigla)
	fmt.Printf("Número de participantes: %v\n", maior.Total)
}

func ShowGeneralRace(race RaceScores) {
	meanPreta := getMeanScores(race.Preta)
	meanBPA := getMeanScores(race.BrcPrdAmar)
	meanIndigena := getMeanScores(race.Indigena)
	fmt.Printf("\nMédias nacionais por raça\n")
	fmt.Printf("\tCiências da natureza:\n")
	fmt.Printf("\t\tBranca/Parda/Amarela: %.2f\n", meanBPA[0])
	fmt.Printf("\t\tPreta: %.2f\n", meanPreta[0])
	fmt.Printf("\t\tIndigena: %.2f\n", meanIndigena[0])

	fmt.Printf("\n\tCiências humanas:\n")
	fmt.Printf("\t\tBranca/Parda/Amarela: %.2f\n", meanBPA[1])
	fmt.Printf("\t\tPreta: %.2f\n", meanPreta[1])
	fmt.Printf("\t\tIndigena: %.2f\n", meanIndigena[1])

	fmt.Printf("\n\tLinguagens e códigos:\n")
	fmt.Printf("\t\tBranca/Parda/Amarela: %.2f\n", meanBPA[2])
	fmt.Printf("\t\tPreta: %.2f\n", meanPreta[2])
	fmt.Printf("\t\tIndigena: %.2f\n", meanIndigena[2])

	fmt.Printf("\n\tMatemática:\n")
	fmt.Printf("\t\tBranca/Parda/Amarela: %.2f\n", meanBPA[3])
	fmt.Printf("\t\tPreta: %.2f\n", meanPreta[3])
	fmt.Printf("\t\tIndigena: %.2f\n", meanIndigena[3])
}

func ShowGeneralSchools(school SchoolScores) {
	meanPublica := getMeanScores(school.Public)
	meanPrivada := getMeanScores(school.Private)
	fmt.Printf("\nMédias nacionais para escola pública e privada\n")
	fmt.Printf("\tCiências da natureza:\n")
	fmt.Printf("\t\tPública: %.2f\n", meanPublica[0])
	fmt.Printf("\t\tPrivada: %.2f\n", meanPrivada[0])

	fmt.Printf("\n\tCiências humanas:    \n")
	fmt.Printf("\t\tPública: %.2f\n", meanPublica[1])
	fmt.Printf("\t\tPrivada: %.2f\n", meanPrivada[1])

	fmt.Printf("\n\tLinguagens e códigos:\n")
	fmt.Printf("\t\tPública: %.2f\n", meanPublica[2])
	fmt.Printf("\t\tPrivada: %.2f\n", meanPrivada[2])

	fmt.Printf("\n\tMatemática:\n")
	fmt.Printf("\t\tPública: %.2f\n", meanPublica[3])
	fmt.Printf("\t\tPrivada: %.2f\n", meanPrivada[3])
}
