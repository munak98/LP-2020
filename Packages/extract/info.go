package extract

import "fmt"

/* Info Nacional */

//MostParticipantsYear mostra o ano com maior numero de participantes  
func MostParticipantsYear(years []Year) {
	var most = Year{}
	for i := 0; i < len(years)-1; i++ {
		if years[i].Total < years[i+1].Total {
			most = years[i+1]
		}
	}
	fmt.Printf("O Ano com maior numero de participantes: %d\n", most.Year)
	fmt.Printf("Numero de participantes: %v", most.Total)
}

//MostParticipantsUF mostra o estado com maior numero de participantes de um ano
func MostParticipantsUF(states []State) {
	var maior = State{}
	for i := 0; i < len(states)-1; i++ {
		if states[i].Total < states[i+1].Total {
			maior = states[i+1]
		}
	}
	fmt.Printf("O Estado com maior numero de participantes: %s\n", maior.Sigla)
	fmt.Printf("Numero de participantes: %v", maior.Total)
}
