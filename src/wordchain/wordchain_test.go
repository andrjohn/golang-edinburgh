package main

import "testing"

func assertArraysMatch(t *testing.T, expect, actual []string) {
	if len(actual) != len(expect) {
		t.Fatalf("Length mismatch: exp %d, act %d", len(expect), len(actual))
	}
	for i, _ := range expect {
		if actual[i] != expect[i] {
			t.Errorf("Chain didn't match at pos %d"+
				"    expect: %s"+
				"    actual: %s\n",
				i, expect[i], actual[i])
		}
	}
}

func TestRotateOne(t *testing.T) {
	start := "got"
	end := "get"

	expect := []string{start, end}
	actual := findChain(start, end)

	assertArraysMatch(t, expect, actual)
}

func TestLoadDict(t *testing.T) {
	words := loadWords()

	if len(words[4]) < 1000 {
		t.Errorf("Expect more 4 letter words: Found %d\n", len(words[4]))
	}
}

func TestOffByOne(t *testing.T) {
	expect := true
	actual := offByOne("got", "get")

	if actual != expect {
		t.Errorf("Not off by one: got get")
	}
}

func TestNotOffByOne(t *testing.T) {
	expect := false
	actual := offByOne("aaa", "bbb")

	if actual != expect {
		t.Errorf("Not off by one: aaa bbb")
	}
}

func TestCandidates(t *testing.T) {
	start := "go"
	words := loadWords()

	expect := []string{"Co", "Ho", "Io", "Jo", "Mo", "Po", "do", "gs", "ho", "lo", "no", "so", "to", "yo"}
	actual := nextWords(words[len(start)], start)

	assertArraysMatch(t, expect, actual)

}

func TestRotateTwo(t *testing.T) {
	t.Skip()
	start := "got"
	end := "gel"

	expect := []string{start, "get", end}
	actual := findChain(start, end)

	assertArraysMatch(t, expect, actual)
}
