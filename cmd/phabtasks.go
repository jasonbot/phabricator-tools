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
		for _, task := range tasks {
			fmt.Printf("T%v %-15s Name: %v\n", task.ID, task.Status.Value, task.Name)
		}
	}
}
