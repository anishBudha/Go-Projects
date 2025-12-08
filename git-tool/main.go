package main

import (
	"fmt"
	"os/exec"
	"strings"
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
	} else {
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
		fmt.Println(output)
	}

	commitSuccess := "commit all files successfully"
	return commitSuccess
}

func gitPush() string {
	cmd := exec.Command("git", "push", "origin", "main")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println(output)
	}

	pushMessage := "Files pushed successfully."
	return pushMessage
}


func main () {
	
	branch, untracked := getGitStatus()
	fmt.Println("Branch:", branch)
	fmt.Println("Untracked files:")
	for _, file := range untracked {
		fmt.Println(file)
	}

	addMsg := gitAddAll()
	fmt.Println(addMsg)

	commitMsg := gitCommit()
	fmt.Println(commitMsg)

	pushMsg := gitPush()
	fmt.Println(pushMsg)

}
