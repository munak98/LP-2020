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

//PrintUFRacesMeanScores printa dados acerca de cada raça por estado
func PrintUFRacesMeanScores(state State) {

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

//MostParticipantsUF mostra o estado com maior numero de participantes de um ano
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

//PrintYearMeanScores mostra as medias nacionais de um ano
func PrintYearMeanScores(year Year) {
	fmt.Printf("\nMédias nacionais do ano %d\n", year.Year)
	for i := range year.Medias {
		fmt.Printf("\n\t%s: %.2f", year.Subjects[i], year.Medias[i])
	}
}

//PrintYearRacesMeanScores mostra as medias nacionais de raça de um ano
func PrintYearRacesMeanScores(year Year) {
	fmt.Printf("\n\nMédias nacionais por raça\n")
	for j := range year.Subjects {
		fmt.Printf("\n%s: \n\t\t", year.Subjects[j])
		for i := range year.Races {
			fmt.Printf("%s: %.2f | ",year.Races[i].Name,  year.Races[i].Medias[j])
		}
	}
}

//PrintYearSchoolsMeanScores mostra as medias nacionais de tipo de Escola de um ano
func PrintYearSchoolsMeanScores(year Year) {
	fmt.Printf("\n\nMédias nacionais para escola pública e privada\n\n")
	for i := range year.Subjects {
		fmt.Printf("%s:\n", year.Subjects[i])
		fmt.Printf("\tPública: %.2f | ", year.SchoolMeanScores[1][i])
		fmt.Printf("Privada: %.2f\n", year.SchoolMeanScores[2][i])
	}
}