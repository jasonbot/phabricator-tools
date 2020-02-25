package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sync"

	phabricatortools "github.com/jasonbot/phabricator-tools/v1"
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

func fetchRepo(repoToFetch repoInfo, updateOnly bool) {
	directoryName := "./r" + repoToFetch.Callsign
	if _, err := os.Stat(directoryName); os.IsNotExist(err) {
		if !updateOnly {
			cloneRepo(directoryName, repoToFetch)
		}
	} else {
		updateRepo(directoryName, repoToFetch)
	}
}

func main() {
	updateOnly := false
	flag.BoolVar(&updateOnly, "updateonly", false, "only update existing repos on disk")
	flag.Parse()

	repositories, err := phabricatortools.GetRepositories()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {

		repoPipeline := make(chan repoInfo)
		var done sync.WaitGroup

		for index := 1; index < 8; index++ {
			done.Add(1)
			go func() {
				defer done.Done()

				for item := range repoPipeline {
					fetchRepo(item, updateOnly)
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
