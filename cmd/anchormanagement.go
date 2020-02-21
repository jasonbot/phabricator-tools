package main

import (
	"fmt"

	phabricatortools "github.com/jasonbot/phabricator-tools"
	"github.com/rivo/tview"
)

func main() {
	statuses, err := phabricatortools.GetStatusMap()
	if err != nil {
		panic(err)
	}

	tasks, err := phabricatortools.GetMyOpenTasks()
	if err != nil {
		panic(err)
	}

	table := tview.NewTable().SetBorders(true)

	for row, task := range tasks {
		table.SetCell(row, 0, tview.NewTableCell(statuses[task.Status.Value].Name))
		table.SetCell(row, 1, tview.NewTableCell(fmt.Sprintf("T%v", task.ID)))
		table.SetCell(row, 2, tview.NewTableCell(task.Name))
	}

	if err := tview.NewApplication().SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}
