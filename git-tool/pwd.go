package main
import(
	"fmt"
	"os/exec"
)

func main () {
	cmd := exec.Command("pwd")
	output, err := cmd.Output()

	if err != nil {
		fmt.Print("Error:", err)
	}

	fmt.Println(string(output))
}
