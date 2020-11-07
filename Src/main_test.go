package main

import (
	"testing"

	"../Packages/extract"
)

const msgError = "Valor esperado %v, mas o valor encontrado foi %v"

func TestMeanScoresUF(t *testing.T) {
	t.Parallel()

	reader := extract.CsvReader()
	state := extract.NormalMeanScoresUF(reader, "DF")	// testa dados do DF

	totalParticipantsExpected := state.Total

	// total de particpantes somando os participantes de cada raça
	totalParticipantsFromRaces := 0

	totalPublicParticipants := state.SchoolType[1]

	// total de participantes de escola publica somando os participantes de cada raça
	totalPublicParticipantsFromRaces := 0

	for i := range state.Races {
		totalParticipantsFromRaces += state.Races[i].Total
		totalPublicParticipantsFromRaces += state.Races[i].SchoolType[1]
	}

	if totalParticipantsFromRaces != totalParticipantsExpected {
		t.Errorf(msgError, totalParticipantsExpected, totalParticipantsFromRaces)
	}

	if totalPublicParticipantsFromRaces != totalPublicParticipants {
		t.Errorf(msgError, totalPublicParticipants, totalPublicParticipantsFromRaces)
	}

}
