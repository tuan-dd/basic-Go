package main

import (
	"os"
	"strings"
	"testing"
)

func TestCard(t *testing.T) {
	d := newCard()

	if len(d) != 25 {
		t.Errorf("Expect deck length of 25, but got only %v", len(d))
	}

	if !strings.Contains(d[0], "sức mạnh") {
		t.Errorf("The first element %v don't have 'sức mạnh' ", d[0])
	}

	if !strings.Contains(d[len(d)-1], "khoai bé") {
		t.Errorf("The last element %v don't have 'khoai bé' ", d[len(d)-1])
	}
}

func TestSaveFileAndNewCardFromFile(t *testing.T) {
	os.Remove("_cardString")
	card := newCard()

	card.saveToFile("_cardString")

	loadFile := readTofFile("_cardString")

	if len(loadFile) != 25 {
		t.Errorf("Expect file card length of 25, but got only %v", len(loadFile))
	}

	os.Remove(("_cardString"))
}
