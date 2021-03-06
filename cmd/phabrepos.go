package main

import (
	"fmt"

	phabricatortools "github.com/jasonbot/phabricator-tools/v1"
)

func main() {
	repositories, err := phabricatortools.GetRepositories()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		for _, repo := range repositories {
			if repo.Fields.Callsign != "" {
				fmt.Printf("r%v (%v)\n", repo.Fields.Callsign, repo.Fields.DefaultBranch)
			} else {
				fmt.Printf("%v (%v) [NO CALL SIGN]\n", repo.Fields.ShortName, repo.Fields.DefaultBranch)
			}
			for _, URI := range repo.Attachments.URIs.URIs {
				if URI.Fields.Builtin.Identifier != "" {
					fmt.Printf("    %v (%v)\n", URI.Fields.URI.Effective, URI.Fields.Builtin.Identifier)
				}
			}
		}
	}
}
