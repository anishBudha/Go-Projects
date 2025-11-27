package main

import(
	"fmt"
	"os/exec"
	"strings"
)

func main () {
	cmd := exec.Command("git", "status")
	output, err := cmd.Output()

	lines := strings.Split(string(output), "\n")

	if err != nil {
		fmt.Println("error: ", err)
	}
	
	var branch string
	var untracked []string // slice of string
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
	fmt.Println("Branch:", branch)
	fmt.Println("Untracked files:")
	for _, file := range untracked {
		fmt.Println(file)
	}
}
