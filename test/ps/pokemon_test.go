package ps_test

import (
	"psapi/pkg/ps"
	"testing"
)

func TestComparePokemonId(t *testing.T) {
	p1 := &ps.Pokemon{
		Id: 1,
	}

	p2 := &ps.Pokemon{
		Id: 2,
	}

	if ps.ComparePokemonId(p1, p2) != -1 {
		t.Error("ComparePokemonId with p1:", p1.Id, "and p2:", p2.Id, "should return -1")
	}
	if ps.ComparePokemonId(p2, p1) != 1 {
		t.Error("ComparePokemonId with p2:", p2.Id, "and p1:", p1.Id, "should return -1")
	}
}
