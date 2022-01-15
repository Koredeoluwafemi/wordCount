package main

import (
	"testing"
)

func TestTextHighlighter(t *testing.T) {

	text := "my name is name and korede and my name and you and me and me and name korede is is you is korede come come den den den for us us we"
	sortArray := highlighter(text)

	firstResult := Pair{
		Key:   "and",
		Value: 6,
	}
	arrayLen := len(sortArray)
	if arrayLen > 0 {
		if sortArray[0] != firstResult {
			t.Errorf("expected {and 6}, but got %v", sortArray[0])
		}
	} else {
		t.Errorf("test failed")
	}

}
