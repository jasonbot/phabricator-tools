package main

import (
	"fmt"

	phabricatortools "github.com/jasonbot/phabricator-tools"
)

func main() {
	tasks, err := phabricatortools.GetMyOpenTasks()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		statusmap, _ := phabricatortools.GetStatusMap()

		for _, task := range tasks {
			statusName := statusmap[task.Status.Value]
			fmt.Printf("T%v | %-15s | %v\n", task.ID, statusName.Name, task.Name)
		}
	}
}
