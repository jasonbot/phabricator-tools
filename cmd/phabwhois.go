package main

import (
	"flag"
	"fmt"

	phabricatortools "github.com/jasonbot/phabricator-tools"
)

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Println("Usage: whoisphab <user-PHID>")
		return
	}

	PHID := flag.Args()[0]

	user, err := phabricatortools.WhoIs(PHID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("%v (%v)\nRoles: %v\n", user.RealName, user.UserName, user.Roles)
	}
}
