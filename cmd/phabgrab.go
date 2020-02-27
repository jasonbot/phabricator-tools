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
	ShortName string
	Callsign  string
	URI       string
	Master    string
}

func cloneRepo(filePath string, repoToFetch repoInfo) {
	fmt.Printf("Cloning %v\n", repoToFetch.Callsign)

	cmd := exec.Command("git", "clone", repoToFetch.URI, filePath)
	err := cmd.Run()

	if err != nil {
		output, _ := cmd.CombinedOutput()
		if len(output) > 0 {
			fmt.Println(string(output))
			fmt.Println("---")
		}
		fmt.Printf("Error cloning %v: %v\n", repoToFetch.Callsign, err)
	}
}

func updateRepo(filePath string, repoToFetch repoInfo) {
	fmt.Printf("Updating %v\n", repoToFetch.Callsign)

	cmd := exec.Command("git", "pull", "origin", repoToFetch.Master)
	cmd.Dir = filePath
	err := cmd.Run()

	if err != nil {
		output, _ := cmd.CombinedOutput()
		if len(output) > 0 {
			fmt.Println(string(output))
			fmt.Println("---")
		}
		fmt.Printf("Error updating %v: %v\n", repoToFetch.Callsign, err)
	}
}

func fetchRepo(repoToFetch repoInfo, useShortname bool, updateOnly bool) {
	directoryName := "./r" + repoToFetch.Callsign
	if useShortname == true {
		directoryName = "./" + repoToFetch.ShortName
	}

	if _, err := os.Stat(directoryName); os.IsNotExist(err) {
		if !updateOnly {
			cloneRepo(directoryName, repoToFetch)
		} else {
			fmt.Printf("Skipping %v\n", repoToFetch.Callsign)
		}
	} else {
		updateRepo(directoryName, repoToFetch)
	}
}

func main() {
	updateOnly := false
	useShortname := false
	workers := uint(4)
	repoName := ""
	flag.UintVar(&workers, "workers", 4, "number of workers to run in parallel")
	flag.BoolVar(&updateOnly, "updateonly", false, "only update existing repos on disk")
	flag.BoolVar(&useShortname, "useshortname", false, "use short name instead of callsign for folder name and matching")
	flag.StringVar(&repoName, "repo", "", "the repo to clone (by default, will download all)")
	flag.Parse()

	repositories, err := phabricatortools.GetRepositories()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {

		repoPipeline := make(chan repoInfo)
		var done sync.WaitGroup

		for index := uint(0); index < workers; index++ {
			done.Add(1)
			go func() {
				defer done.Done()

				for item := range repoPipeline {
					fetchRepo(item, useShortname, updateOnly)
				}
			}()
		}

		for _, repo := range repositories {
			for _, URI := range repo.Attachments.URIs.URIs {
				uriTypeMatch := false

				if useShortname && URI.Fields.Builtin.Identifier == "shortname" {
					uriTypeMatch = true
				} else if !useShortname && URI.Fields.Builtin.Identifier == "callsign" {
					uriTypeMatch = true
				}

				if uriTypeMatch {
					matched := false

					if repoName == "" {
						matched = true
					} else {
						if useShortname == true && repoName == repo.Fields.ShortName {
							matched = true
						} else if useShortname == false && repoName == "r"+repo.Fields.Callsign {
							matched = true
						}
					}

					if matched == true {
						repoPipeline <- repoInfo{
							ShortName: repo.Fields.ShortName,
							Callsign:  repo.Fields.Callsign,
							URI:       URI.Fields.URI.Effective,
							Master:    repo.Fields.DefaultBranch,
						}
					}
				}
			}
		}
		close(repoPipeline)
		done.Wait()
	}
}
