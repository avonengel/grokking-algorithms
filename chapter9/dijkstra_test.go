package chapter9

import (
	"reflect"
	"testing"
)

var pianoTradeGraph = NewGraph([]Node{
	Node{"book", map[string]int{
		"rare LP": 5,
		"poster":  0,
	}},
	Node{"rare LP", map[string]int{
		"bass guitar": 15,
		"drum set":    20,
	}},
	Node{"poster", map[string]int{
		"bass guitar": 30,
		"drum set":    35,
	}},
	Node{"bass guitar", map[string]int{
		"piano": 20,
	}},
	Node{"drum set", map[string]int{
		"piano": 10,
	}},
	Node{"piano", map[string]int{}},
})

func TestDijkstra(t *testing.T) {
	// given the pianoTradeGraph

	// when
	got, err := pianoTradeGraph.ShortestPath("book", "piano")

	// then
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	wanted := ShortestPath{
		Path:   []string{"book", "rare LP", "drum set", "piano"},
		Weight: 35,
	}
	if !reflect.DeepEqual(got, wanted) {
		t.Errorf("got: %v, wanted: %v", got, wanted)
	}
}
