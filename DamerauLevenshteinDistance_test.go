package DamerauLevenshteinDistance

import (
	"testing"
	"sort"
)

func TestEmptySource(t *testing.T) {
	target := "test"
	dist := Distance("", target)
	if dist != len(target) {
		t.Error(dist, " not equal ", len(target))
	}
}

func TestEmptyTarget(t *testing.T) {
	source := "test"
	dist := Distance(source, "")
	if dist != len(source) {
		t.Error(dist, " not equal ", len(source))
	}
}

func TestEqual(t *testing.T) {
	source := "I like go."
	target := "I like go."
	dist := Distance(source, target)
	if dist != 0 {
		t.Error(dist, " not equal ", 0)
	}
}

func TestOneSubstitution(t *testing.T) {
	source := "I like go."
	target := "I like og."
	dist := Distance(source, target)
	if dist != 1 {
		t.Error(dist, " not equal ", 1)
	}
}

func TestOneInsertion(t *testing.T) {
	source := "I like go."
	target := "I like go.."
	dist := Distance(source, target)
	if dist != 1 {
		t.Error(dist, " not equal ", 1)
	}
}

func TestOneDeletion(t *testing.T) {
	source := "I like go."
	target := "I like go"
	dist := Distance(source, target)
	if dist != 1 {
		t.Error(dist, " not equal ", 1)
	}
}

func TestStrings(t *testing.T) {
	good := []string {
		"cat",
		"cat",
		"cat",
		"cat",
		"I love go.",
		"I love go.",
		"I love go.",
		"I love go.",
	}

	bad := []string {
		"cta",
		"ca",
		"tac",
		"dog",
		"I love og.",
		"Ilove go.",
		"I lvoe .go",
		"zzzzzzzzzz",
	}

	distances := []int {
		1,
		1,
		2,
		3,
		1,
		1,
		3,
		len(good[len(good)-1]),
	}

	for i := 0; i < len(good); i++ {
		dist := Distance(good[i], bad[i])
		if dist != distances[i] {
			t.Error(dist, " not equal ", distances[i], good[i], bad[i])
		}

		dist = Distance(bad[i], good[i])
		if dist != distances[i] {
			t.Error(dist, " not equal ", distances[i], bad[i], good[i])
		}
	}
}

func TestSort(t *testing.T) {
	reference := "cat"

	list := DLStrings {
		{ Reference: reference, Value: "cta" },
		{ Reference: reference, Value: "dog" },
		{ Reference: reference, Value: "cat" },
	}

	sort.Sort(list)
	if list[0].Value != "cat" {
		t.Error(list[0].Value, " was not ", reference)
	}
	if list[1].Value != "cta" {
		t.Error(list[1].Value, " was not ", "cta")
	}
	if list[2].Value != "dog" {
		t.Error(list[2].Value, " was not ", "dog")
	}
}
