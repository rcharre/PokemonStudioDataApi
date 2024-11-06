package pkmn_test

import (
	"testing"

	"github.com/rcharre/psapi/pkg/pkmn"
)

func TestComparePokemonId(t *testing.T) {
	p1 := &pkmn.Pokemon{
		Id: 1,
	}

	p2 := &pkmn.Pokemon{
		Id: 2,
	}

	if pkmn.ComparePokemonId(p1, p2) != -1 {
		t.Error("ComparePokemonId with p1:", p1.Id, "and p2:", p2.Id, "should return -1")
	}
	if pkmn.ComparePokemonId(p2, p1) != 1 {
		t.Error("ComparePokemonId with p2:", p2.Id, "and p1:", p1.Id, "should return -1")
	}
}
