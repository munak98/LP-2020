package extract

import "fmt"

//PrintYearData mostra todos os dados de um ano
func PrintYearData(year Year) {
	fmt.Printf("Dados nacionais Enem %d\n", year.Year)
	fmt.Printf("Total de participantes: %d", year.Total)
	PrintYearMeanScores(year)
	PrintYearRacesMeanScores(year)
	PrintYearSchoolsMeanScores(year)
	PrintYearMostParticipantsUF(year.States)

	MenuStates(year.States) // Chama o Menu de Estados
}

//PrintYearMeanScores mostra as medias nacionais de um ano
func PrintYearMeanScores(year Year) {
	fmt.Printf("\n*****************************************\n")
	fmt.Printf("\nMédias nacionais do ano %d\n", year.Year)
	for i := range year.Medias {
		fmt.Printf("\n\t%s: %.2f", year.Recurse.Subjects[i], year.Medias[i])
	}
}

//PrintYearRacesMeanScores mostra as medias nacionais de raça de um ano
func PrintYearRacesMeanScores(year Year) {
	fmt.Printf("\n*****************************************\n")
	fmt.Printf("\n\nMédias nacionais por raça\n")
	for j := range year.Recurse.Subjects {
		fmt.Printf("\n%s: \n\t\t", year.Recurse.Subjects[j])
		for i := range year.Races {
			fmt.Printf("%s: %.2f | ", year.Races[i].Name, year.Races[i].Medias[j])
		}
	}
}

//PrintYearSchoolsMeanScores mostra as medias nacionais de tipo de Escola de um ano
func PrintYearSchoolsMeanScores(year Year) {
	fmt.Printf("\n*****************************************\n")
	fmt.Printf("\n\nMédias nacionais para escola pública e privada\n\n")
	for i := range year.Recurse.Subjects {
		fmt.Printf("%s:\n", year.Recurse.Subjects[i])
		fmt.Printf("\tPública: %.2f | ", year.SchoolMeanScores[1][i])
		fmt.Printf("Privada: %.2f\n", year.SchoolMeanScores[2][i])
	}
}

//PrintYearMostParticipantsUF mostra o estado com maior numero de participantes de um ano
func PrintYearMostParticipantsUF(states []State) {
	var maior = State{}
	for i := 0; i < len(states)-1; i++ {
		if states[i].Total < states[i+1].Total {
			maior = states[i+1]
		}
	}
	fmt.Printf("\nO Estado com maior numero de participantes: %s\n", maior.Sigla)
	fmt.Printf("Numero de participantes: %v\n", maior.Total)
}

//PrintUFData printa dados de uma UF
func PrintUFData(state State) {
	fmt.Printf("\n*****************************************\n")
	fmt.Printf("Dados de %s", state.Sigla)
	fmt.Printf("\nTotal de participantes: %d\n", state.Total)
	PrintUFMeanScores(state)

	fmt.Printf("Numero de participantes de cada Raça:\n")
	for i := range state.Races {
		fmt.Printf("\t%s: %d\n", state.Races[i].Name, state.Races[i].Total)
	}
	fmt.Printf("\n*****************************************\n\n")
	fmt.Printf("Numero de participantes de cada Tipo de Escola:")
	for i := range state.Recurse.Schools {
		fmt.Printf("\t%s: %d\n", state.Recurse.Schools[i], state.SchoolType[i])
	}
}

//PrintUFMeanScores printa as media de uma UF
func PrintUFMeanScores(state State) {
	fmt.Println("Médias:")
	fmt.Printf("\t%s: %.2f \n\t%s: %.2f \n\t%s: %.2f\n\t%s: %.2f\n\n",
		state.Recurse.Schools[0], state.Medias[0],
		state.Recurse.Schools[1], state.Medias[1],
		state.Recurse.Schools[2], state.Medias[2],
		state.Recurse.Schools[3], state.Medias[3],
	)
}

//PrintUFRacesMeanScores printa dados estaduais acerca de todas raças
func PrintUFRacesMeanScores(state State) {
	for i := range state.Races {
		fmt.Printf("\n*****************************************\n")
		fmt.Printf("\nDados de %s", state.Races[i].Name)
		fmt.Printf("\nTotal de participantes: %d\n", state.Races[i].Total)
		fmt.Printf("Médias: \n\t%s: %.2f \n\t%s: %.2f \n\t%s: %.2f\n\t%s: %.2f\n\n",
			state.Recurse.Schools[0], state.Races[i].Medias[0],
			state.Recurse.Schools[1], state.Races[i].Medias[1],
			state.Recurse.Schools[2], state.Races[i].Medias[2],
			state.Recurse.Schools[3], state.Races[i].Medias[3],
		)

		for j := range state.Recurse.Schools {
			fmt.Printf("Numero de participantes de %s: %d\n", state.Recurse.Schools[j], state.Races[i].SchoolType[j])
		}
	}
}

//PrintUFRaceMeanScores printa dados estaduais acerca de uma raça
func PrintUFRaceMeanScores(race Race) {
	fmt.Printf("\n*****************************************\n")
	fmt.Printf("\nDados de %s", race.Name)
	fmt.Printf("\nTotal de participantes: %d\n", race.Total)
	fmt.Printf("Médias: \n\t%s: %.2f \n\t%s: %.2f \n\t%s: %.2f\n\t%s: %.2f\n\n",
		race.Recurse.Schools[0], race.Medias[0],
		race.Recurse.Schools[1], race.Medias[1],
		race.Recurse.Schools[2], race.Medias[2],
		race.Recurse.Schools[3], race.Medias[3],
	)
	for i := range race.Recurse.Schools {
		fmt.Printf("Numero de participantes de %s: %d\n", race.Recurse.Schools[i], race.SchoolType[i])
	}
}

//PrintUFSchoolsMeanScores printa dados estaduais acerca de todos tipos de Escola
func PrintUFSchoolsMeanScores(state State) {
	fmt.Printf("\n*****************************************\n")
	for i := range state.Recurse.Schools {
		fmt.Printf("\nTotal de participantes da escola %s: %d\n", state.Recurse.Schools[i], state.SchoolType[i])
		fmt.Println("Médias:")
		fmt.Printf("\t%s: %.2f \n\t%s: %.2f \n\t%s: %.2f\n\t%s: %.2f\n\n",
			state.Recurse.Schools[0], state.SchoolMeanScores[0],
			state.Recurse.Schools[1], state.SchoolMeanScores[1],
			state.Recurse.Schools[2], state.SchoolMeanScores[2],
			state.Recurse.Schools[3], state.SchoolMeanScores[3],
		)
		for j := range state.Recurse.Schools {
			fmt.Printf("Numero de participantes de %s: %d\n", state.Recurse.Schools[j], state.SchoolType[j])
		}
	}
}

//PrintUFSchoolMeanScores printa dados estaduais de um tipo de Escola
func PrintUFSchoolMeanScores(state State, school int) {
	fmt.Printf("\n*****************************************\n")
	fmt.Printf("\nDados de %s", state.Sigla)
	fmt.Printf("\nTotal de participantes da escola %s: %d\n", state.Recurse.Schools[school], state.SchoolType[school])
	fmt.Println("Médias:")
	fmt.Printf("\t%s: %.2f \n\t%s: %.2f \n\t%s: %.2f\n\t%s: %.2f\n\n",
		state.Recurse.Schools[0], state.Medias[0],
		state.Recurse.Schools[1], state.Medias[1],
		state.Recurse.Schools[2], state.Medias[2],
		state.Recurse.Schools[3], state.Medias[3],
	)

	fmt.Printf("Numero de participantes de %s: %d\n", state.Recurse.Schools[school], state.SchoolType[school])
	
}
