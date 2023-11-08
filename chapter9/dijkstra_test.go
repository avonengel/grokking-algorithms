package chapter9

import (
	"fmt"
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

type testCase struct {
	input  Graph
	wanted ShortestPath
}

func (c *testCase) run(t *testing.T) {
	// given graph input
	// when dijkstra is run
	got, err := c.input.ShortestPath("start", "finish")

	// then
	if err != nil {
		t.Fatal("unexpected error", err)
	}

	if !reflect.DeepEqual(got, c.wanted) {
		t.Errorf("got: %v, wanted: %v", got, c.wanted)
	}
}
func TestExercises(t *testing.T) {
	testcases := []testCase{
		{input: NewGraph([]Node{
			{
				Name: "start",
				neighbors: map[string]int{
					"UL": 5,
					"LL": 2,
				},
			},
			{
				Name: "LL",
				neighbors: map[string]int{
					"UL": 6,
					"LR": 7,
				},
			},
			{
				Name: "UL",
				neighbors: map[string]int{
					"UR": 4,
					"LR": 2,
				},
			},
			{
				Name: "UR",
				neighbors: map[string]int{
					"finish": 3,
					"LR":     6,
				},
			},
			{
				Name: "LR",
				neighbors: map[string]int{
					"finish": 1,
				},
			},
			{
				Name:      "finish",
				neighbors: nil,
			},
		}), wanted: ShortestPath{
			Path:   []string{"start", "UL", "LR", "finish"},
			Weight: 8,
		}},
		{
			input: Graph{
				Nodes: []Node{
					{
						Name: "start",
						neighbors: map[string]int{
							"UL": 10,
						},
					},
					{
						Name: "UL",
						neighbors: map[string]int{
							"UR": 20,
						},
					},
					{
						Name: "UR",
						neighbors: map[string]int{
							"finish": 30,
							"C":      1,
						},
					},
					{
						Name: "C",
						neighbors: map[string]int{
							"UL": 1,
						},
					},
					{
						Name:      "finish",
						neighbors: nil,
					},
				},
			},
			wanted: ShortestPath{
				Path:   []string{"start", "UL", "UR", "finish"},
				Weight: 60,
			},
		},
	}

	for i, testcase := range testcases {
		t.Run(fmt.Sprintf("graph-%d", i+1), testcase.run)
	}
}
