package main

import (
	"fmt"
	"os/exec"
	"strings"
	"os"
	"bufio"
)

func getGitStatus() (string, []string){
	cmd := exec.Command("git", "status")
	output, err := cmd.Output()

	lines := strings.Split(string(output), "\n")

	if err !=  nil {
		fmt.Println("error", err)
	}

	var branch string
	var untracked []string // slice of strings
	collecting := false

	for _, line := range lines {
		if strings.HasPrefix(line, "On branch") {
			parts := strings.Split(line, " ")
			branch = parts[len(parts) - 1]
		}

		if strings.HasPrefix(line, "Untracked files:") {
			collecting = true
			continue
		}
		if collecting {
			trimmed := strings.TrimSpace(line)
			if trimmed == "" {
				collecting = false
				continue
			} 
			if strings.HasPrefix(trimmed, "(") {
				continue
			}
			untracked = append(untracked, trimmed)
		}
	}
	return branch, untracked
}

func gitAddAll() string {
	cmd := exec.Command("git", "add", ".")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println("error", err)
		fmt.Println(output)
	}
	addMessage := "Added all files successfully."
	return addMessage
}

func gitCommit() string {
	test := `"test"`
	cmd := exec.Command("git", "commit", "-m", test)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("error", err)
	} else {
		lines := strings.Split(string(output), "\n")
		fmt.Println(lines)
	}

	commitSuccess := "Commited all files successfully"
	return commitSuccess
}

func gitPush() string {
	cmd := exec.Command("git", "push", "origin", "main")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("error", err)
	} else {
		lines := strings.Split(string(output), "\n")
		fmt.Println(lines)
	}

	pushMessage := "Files pushed successfully."
	return pushMessage
}

func cont() string {
	fmt.Println("Would you like to continue? 1 to continue, 0 to abort")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}

func main () {
	
	branch, untracked := getGitStatus()
	fmt.Println("Branch:", branch)
	fmt.Println("length of untracked:", len(untracked))
	fmt.Println("Untracked files:")
	for _, file := range untracked {
		fmt.Println(file)
	}


	if untracked != nil {
		fmt.Println("Untracked files:")
		for _, file := range untracked {
			fmt.Println(file)
		}

		input := cont()

		if input == "1" {
			addMsg := gitAddAll()
			fmt.Println(addMsg)

			commitMsg := gitCommit()
			fmt.Println(commitMsg)

			input := cont()
			if input == "1" {
				pushMsg := gitPush()
				fmt.Println(pushMsg)
			}
		} else {
			fmt.Println("Aborting...")
		}
	} else {
		fmt.Println("No files to commit, everything is upto date.")
	}
}
