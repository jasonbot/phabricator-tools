package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	phabricatortools "github.com/jasonbot/phabricator-tools"
)

type repoInfo struct {
	Callsign string
	URI      string
	Master   string
}

func cloneRepo(filePath string, repoToFetch repoInfo) {
	fmt.Printf("Cloning %v\n", repoToFetch.Callsign)

	cmd := exec.Command("git", "clone", repoToFetch.URI, filePath)
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}

func updateRepo(filePath string, repoToFetch repoInfo) {
	fmt.Printf("Updating %v\n", repoToFetch.Callsign)

	cmd := exec.Command("git", "pull", "origin", repoToFetch.Master)
	cmd.Dir = filePath
	err := cmd.Run()

	if err != nil {
		fmt.Println(err)
	}
}

func fetchRepo(repoToFetch repoInfo) {
	directoryName := "./r" + repoToFetch.Callsign
	if _, err := os.Stat(directoryName); os.IsNotExist(err) {
		cloneRepo(directoryName, repoToFetch)
	} else {
		updateRepo(directoryName, repoToFetch)
	}
}

func main() {
	repositories, err := phabricatortools.GetRepositories()

	repoPipeline := make(chan repoInfo)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		var done sync.WaitGroup

		for index := 1; index < 8; index++ {
			done.Add(1)
			go func() {
				defer done.Done()

				for item := range repoPipeline {
					fetchRepo(item)
				}
			}()
		}

		for _, repo := range repositories {
			for _, URI := range repo.Attachments.URIs.URIs {
				if URI.Fields.Builtin.Identifier == "callsign" {
					repoPipeline <- repoInfo{Callsign: repo.Fields.Callsign, URI: URI.Fields.URI.Effective, Master: repo.Fields.DefaultBranch}
				}
			}
		}
		close(repoPipeline)
		done.Wait()
	}
}
