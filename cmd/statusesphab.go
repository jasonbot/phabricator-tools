package main

import (
	"fmt"

	phabricatortools "github.com/jasonbot/phabricator-tools"
)

func main() {
	statuses, err := phabricatortools.GetStatuses()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		for _, status := range statuses {
			defaultString := "  "
			if status.Special == "default" {
				defaultString = "* "
			}
			fmt.Printf("%v%v\t%v\n", defaultString, status.Value, status.Name)
		}
	}
}
