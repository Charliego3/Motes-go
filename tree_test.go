package main

import "testing"

func TestFileTree(t *testing.T) {
	const path = "/Users/charlie/Library/Mobile Documents/com~apple~CloudDocs/notes"
	app := NewApp()
	tree, err := app.buildTree(path)
	if err != nil {
		panic(err)
	}

	t.Logf("%+v", tree)
}
