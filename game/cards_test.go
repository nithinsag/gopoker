package game

import "testing"

func TestCards(t *testing.T) {
	card := Card{Number: 8, Suite: Spade}

	cstring := card.String()
	if cstring != "Spade8" {
		t.Error("Expeced Spade8 got: ", cstring)
	}

	t.Log("got card for Spade 8", cstring)
}
