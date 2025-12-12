package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// permutation generator
func permute(words []string, start int, result *[][]string) {
	if start == len(words) {
		copyWords := make([]string, len(words))
		copy(copyWords, words)
		*result = append(*result, copyWords)
		return
	}

	for i := start; i < len(words); i++ {
		words[start], words[i] = words[i], words[start]
		permute(words, start+1, result)
		words[start], words[i] = words[i], words[start]
	}
}

// try mounting using a password
func tryPassword(dmgPath, password string) bool {
	cmd := exec.Command("hdiutil", "attach", "-passphrase", password, dmgPath)
	out, err := cmd.CombinedOutput()

	// success usually means "mounted successfully" and exit code 0
	if err == nil {
		fmt.Println("Mount output:", string(out))
		return true
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter path to the DMG file: ")
	dmgPath, _ := reader.ReadString('\n')
	dmgPath = strings.TrimSpace(dmgPath)

	words := make([]string, 20)
	for i := 0; i < 20; i++ {
		fmt.Printf("Enter Word %d: ", i+1)
		w, _ := reader.ReadString('\n')
		words[i] = strings.TrimSpace(w)
	}

	var combos [][]string
	permute(words, 0, &combos)

	fmt.Println("Trying combinations...")

	for _, combo := range combos {
		password := strings.Join(combo, "")
		fmt.Println("Trying:", password)

		if tryPassword(dmgPath, password) {
			fmt.Println("Success!")
			fmt.Println("Correct order:", combo)
			return
		}
	}
	fmt.Println("No combination succeeded.")
}
