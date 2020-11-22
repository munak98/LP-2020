package extract

import "fmt"

//PrintYearData mostra os dados sobre um dado ano
func PrintYearData(year Year) {
	fmt.Printf("Dados nacionais ENEM %d", year.Year)
	fmt.Printf("\n*****************************************\n")
	PrintYearMeanScores(year)
	fmt.Printf("\n\n*****************************************\n")
	PrintYearRacesMeanScores(year)
	fmt.Printf("\n*****************************************\n")
	PrintYearSchoolsMeanScores(year)
	fmt.Printf("\n*****************************************\n")
	MostParticipantsUF(year.States)
	fmt.Printf("\n*****************************************\n")
	MenuStates(year.States)
}

//PrintYearMeanScores mostra as medias nacionais de um ano
func PrintYearMeanScores(year Year) {
	fmt.Printf("Médias nacionais do ano %d\n", year.Year)
	for i := range year.Medias {
		fmt.Printf("\n\t%s: %.2f", year.Subjects[i], year.Medias[i])
	}
}

//PrintYearRacesMeanScores mostra as medias nacionais de raça de um ano
func PrintYearRacesMeanScores(year Year) {
	fmt.Printf("Médias nacionais por raça\n")
	for j := range year.Subjects {
		fmt.Printf("\n%s: \n\t\t", year.Subjects[j])
		for i := range year.Races {
			fmt.Printf("%s: %.2f | ", year.Races[i].Name, year.Races[i].Medias[j])
		}
	}
}

//PrintYearSchoolsMeanScores mostra as medias nacionais de tipo de Escola de um ano
func PrintYearSchoolsMeanScores(year Year) {
	fmt.Printf("Médias nacionais para escola pública e privada\n\n")
	for i := range year.Subjects {
		fmt.Printf("%s:\n", year.Subjects[i])
		fmt.Printf("\tPública: %.2f | ", year.SchoolMeanScores[1][i])
		fmt.Printf("Privada: %.2f\n", year.SchoolMeanScores[2][i])
	}
}

//PrintUFData printa dados de uma UF
func PrintUFData(state State) {
	fmt.Printf("\n*****************************************\n")
	fmt.Printf("Dados de %s", state.Sigla)
	fmt.Printf("\nTotal de participantes: %d", state.Total)
	fmt.Printf("\n*****************************************\n")
	fmt.Printf("Numero de participantes de Escola Não informada: %d\n", state.SchoolType[0])
	fmt.Printf("Numero de participantes de Escola Publica: %d\n", state.SchoolType[1])
	fmt.Printf("Numero de participantes de Escola Privada: %d\n", state.SchoolType[2])
	fmt.Printf("Numero de participantes de Escola Exterior: %d\n", state.SchoolType[3])
	fmt.Printf("\n*****************************************\n")
	PrintUFMeanScores(state)
	fmt.Printf("\n*****************************************\n")
	PrintUFRacesMeanScores(state)
}

//PrintUFMeanScores printa dados acerca da UF
func PrintUFMeanScores(state State) {
	fmt.Println("Médias:")
	fmt.Printf("\tCiências da natureza: %.2f \n\tCiências humanas: %.2f \n\tLinguagens e códigos: %.2f\n\tMatemática: %.2f\n\n",
		state.Medias[0],
		state.Medias[1],
		state.Medias[2],
		state.Medias[3],
	)
}

//PrintUFRaceMeanScores printa dados estaduais acerca de uma raça
func PrintUFRaceMeanScores(race Race) {
	fmt.Printf("\n*****************************************\n")
	fmt.Printf("Dados de %s", race.Name)
	fmt.Printf("\nTotal de participantes: %d\n", race.Total)
	fmt.Println("Médias:")
	fmt.Printf("\tCiências da natureza: %.2f \n\tCiências humanas: %.2f \n\tLinguagens e códigos: %.2f\n\tMatemática: %.2f\n\n",
		race.Medias[0],
		race.Medias[1],
		race.Medias[2],
		race.Medias[3],
	)
	fmt.Printf("Número de participantes por tipo de Escola:\n")
	fmt.Printf("\n\tNão responder: %d \n\tPública: %d \n\tPrivada: %d\n\tExterior: %d\n\n",
		race.SchoolType[0],
		race.SchoolType[1],
		race.SchoolType[2],
		race.SchoolType[3],
	)
}

//PrintUFRacesMeanScores printa dados acerca de todas raças de um estado
func PrintUFRacesMeanScores(state State) {
	for i := range state.Races {
		fmt.Printf("\n*****************************************\n")
		fmt.Printf("Dados de %s", state.Races[i].Name)
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

