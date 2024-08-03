package main

import "testing"

func TestSearches(t *testing.T) {
	root := &Node{Value: 100}
	bTree := BinaryTree{Root: root}

	bTree.Insert(50)
	bTree.Insert(75)
	bTree.Insert(160)

	findFifth := bTree.Search(50)
	if !findFifth {
		t.Errorf("Should return true for the value %d", 50)
	}

	findOneSixth := bTree.Search(160)
	if !findOneSixth {
		t.Errorf("Should return true for the value %d", 160)
	}

	find75 := bTree.Search(75)
	if !find75 {
		t.Errorf("Should return true for the value %d", 75)
	}

	notFound := bTree.Search(120)
	if notFound {
		t.Errorf("Should return false for the value %d", 120)
	}
}
