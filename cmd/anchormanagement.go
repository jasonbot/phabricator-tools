package main

import (
	phabricatortools "github.com/jasonbot/phabricator-tools"
	"github.com/rivo/tview"
)

func main() {
	_, err := phabricatortools.GetStatuses()
	if err != nil {
		panic(err)
	}

	box := tview.NewBox().SetBorder(true).SetTitle("Anchor Management")
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}
