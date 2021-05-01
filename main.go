package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/go-git/go-git"
	"github.com/go-git/go-git/plumbing/object"
	"github.com/go-git/go-git/plumbing/transport/ssh"
	"github.com/go-git/go-git/storage/memory"
)

func getRepositories(filePath string) []string {

	file, err := os.Open(filePath)
	repositories := []string{}

	if err != nil {
		fmt.Println("Error", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		repositories = append(repositories, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error", err.Error())
		os.Exit(1)
	}

	return repositories
}

func main() {

	repositoryFile, privateKeyFile := os.Args[1], os.Args[2]

	if len(os.Args) < 2 {
		fmt.Println("Usage: <repositoryFile> <privateKeyFile>")
		os.Exit(1)
	}

	_, err := os.Stat(privateKeyFile)
	if err != nil {
		fmt.Println("Failed to read keyfile", privateKeyFile, err.Error())
		os.Exit(1)
	}

	publicKeys, err := ssh.NewPublicKeysFromFile("git", privateKeyFile, "")

	if err != nil {
		fmt.Println("Failed to generate publickeys", err.Error())
		os.Exit(1)
	}

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

	for _, url := range getRepositories(repositoryFile) {

		fmt.Println("Scanning for commits in", url)

		r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
			URL:      url,
			Auth:     publicKeys,
			Progress: os.Stdout,
		})

		if err != nil {
			fmt.Println("Failed to clone repository", err.Error())
			os.Exit(1)
		}

		ref, err := r.Head()
		cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})

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
	}

	// Output results
	for index, outcome := range outcomes {

		// Seperate by module
		if index%5 == 0 {
			fmt.Println("")
		}

		fmt.Println("Outcome", outcome, "has", outcomeCommits[outcome], "commits")
	}
}
