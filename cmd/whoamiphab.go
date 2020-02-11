package main

import (
	"fmt"

	phabricatortools "github.com/jasonbot/phabricator-tools"
)

func main() {
	user, err := phabricatortools.WhoAmI()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println(user.PHID)
	}
}
