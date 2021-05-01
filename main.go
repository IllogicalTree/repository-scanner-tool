package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/go-git/go-git"
	"github.com/go-git/go-git/plumbing/object"
	"github.com/go-git/go-git/plumbing/transport/ssh"
	"github.com/go-git/go-git/storage/memory"
)

func main() {

	url, privateKeyFile := os.Args[1], os.Args[2]

	if len(os.Args) < 2 {
		fmt.Println("Usage: <url> <privateKeyFile>")
		os.Exit(1)
	}

	_, err := os.Stat(privateKeyFile)
	if err != nil {
		fmt.Println("Failed to read keyfile", privateKeyFile, err.Error())
		return
	}

	fmt.Println("Scanning repository", url)

	publicKeys, err := ssh.NewPublicKeysFromFile("git", privateKeyFile, "")

	if err != nil {
		fmt.Println("Failed to generate publickeys", err.Error())
		return
	}

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:      url,
		Auth:     publicKeys,
		Progress: os.Stdout,
	})

	if err != nil {
		fmt.Println("Failed to clone repository", err.Error())
	}

	ref, err := r.Head()
	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})

	outcomeCommits := make(map[string]int)
	outcomes := [30]string{
		"1.1.1.1", "1.1.1.2", "1.1.1.3", "1.1.1.4", "1.1.1.5",
		"1.1.2.1", "1.1.2.2", "1.1.2.3", "1.1.2.4", "1.1.2.5",
		"1.1.3.1", "1.1.3.2", "1.1.3.3", "1.1.3.4", "1.1.3.5",
		"1.2.1.1", "1.2.1.2", "1.2.1.3", "1.2.1.4", "1.2.1.5",
		"1.3.1.1", "1.3.1.2", "1.3.1.3", "1.3.1.4", "1.3.1.5",
		"1.3.2.1", "1.3.2.2", "1.3.2.3", "1.3.2.4", "1.3.2.5",
	}

	// Set number of commits for each outcome as 0;
	for _, outcome := range outcomes {
		outcomeCommits[outcome] = 0
	}

	// Configure regex rules to find commits with outcomes
	outcomeReg, _ := regexp.Compile("(?i)lo*\\s?\\d.\\d.\\d.\\d")
	outcomeNumReg, _ := regexp.Compile("\\d.\\d.\\d.\\d")

	cIter.ForEach(func(c *object.Commit) error {

		// If commit contains matches regex for learning outcome
		match := outcomeReg.MatchString(c.Message)
		if match {
			// Find the actual learning outcome number (without LO text)
			outcomeNo := outcomeNumReg.FindString(c.Message)
			// Increment number of outcome commits for outcome
			outcomeCommits[outcomeNo]++
		}
		return nil
	})

	// Output results
	for index, outcome := range outcomes {

		// Seperate by module
		if index%5 == 0 {
			fmt.Println("")
		}

		fmt.Println("Outcome", outcome, "has", outcomeCommits[outcome], "commits")
	}
}
