package main

import (
	"fmt"
	"os/exec"
	"strings"
	"os"
	"bufio"
)

type GitStatus struct {
	LocalBranch string
	RemoteBranch string
	Untracked []string
	ModifiedStaged []string
	ModifiedUnstaged []string
	Added []string
	AddedThenModified []string
	DeletedStaged []string
	DeletedUnstaged []string
	ModifiedStagedModified []string

}

func getGitStatus() GitStatus {
	status := GitStatus{}
	cmd := exec.Command("git", "status", "--porcelain", "--branch")
	output, err := cmd.Output()

	lines := strings.Split(string(output), "\n")

	if err !=  nil {
		fmt.Println("error", err)
	}
	

	for _, line := range lines {
		if strings.HasPrefix(line, "##") {
			trimmedLine := strings.TrimPrefix(line, "## ")
			parts := strings.Split(trimmedLine, "...")
			status.LocalBranch = parts[0]
			if len(parts) > 1 {
				status.RemoteBranch = parts[1]
			} else {
				status.RemoteBranch= "no remote"
			}
		}
		if len(line) < 3 {
			continue
		}
		if strings.HasPrefix(line, " M") {
			filename := line[3:]
			status.ModifiedUnstaged = append(status.ModifiedUnstaged, filename)
		} else if strings.HasPrefix(line, "M ") {
			filename := line[3:]
			status.ModifiedStaged = append(status.ModifiedStaged, filename)
		} else if strings.HasPrefix(line, "A ") {
			filename := line[3:]
			status.Added = append(status.Added, filename)
		} else if strings.HasPrefix(line, "AM") {
			filename := line[3:]
			status.AddedThenModified = append(status.AddedThenModified, filename)
		} else if strings.HasPrefix(line, "D ") {
			filename := line[3:]
			status.DeletedStaged = append(status.DeletedStaged, filename)
		} else if strings.HasPrefix(line, " D") {
			filename := line[3:]
			status.DeletedUnstaged = append(status.DeletedUnstaged, filename)
		} else if strings.HasPrefix(line, "R "){
			filename := line[3:]
			status.Renamed = append(status.Renamed, filename)
		} else if strings.HasPrefix(line, "MM") {
			filename := line[3:]
			status.ModifiedStagedModified = append(status.ModifiedStagedModified, filename)
		} else if strings.HasPrefix(line, "??") {
			filename := line[3:]
			status.Unstaged = append(status.Unstaged, filename)
		}
	}
	return status
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
