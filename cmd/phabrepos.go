package main

import (
	"fmt"

	phabricatortools "github.com/jasonbot/phabricator-tools"
)

func main() {
	repositories, err := phabricatortools.GetRepositories()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		for _, repo := range repositories {
			fmt.Printf("r%v\n", repo.Fields.Callsign)
			for _, URI := range repo.Attachments.URIs.URIs {
				fmt.Printf("    %v (%v)\n", URI.Fields.URI.Effective, URI.Fields.Builtin.Identifier)
			}
		}
	}
}
