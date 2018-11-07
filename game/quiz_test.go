package game_test

import (
	"reflect"
	"testing"

	"github.com/vinixavier/goquiz/game"
)

func TestParseCSV(t *testing.T) {

	filepath := "../test/test.csv"
	var expected [][]string
	expected = append(expected, []string{"5+5", "10"})

	result := game.NewQuiz().ParseCSV(&filepath)

	if !reflect.DeepEqual(result, expected) {
		t.Error("For input: ", filepath,
			"expected: ", expected[0][0],
			"got: ", result)
	}

}
