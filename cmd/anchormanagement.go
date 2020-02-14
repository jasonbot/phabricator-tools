package main

import (
	"fmt"

	phabricatortools "github.com/jasonbot/phabricator-tools"
)

func main() {
	_, err := phabricatortools.GetStatuses()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Anchor polish")
}
