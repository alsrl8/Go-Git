package terminal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ClearTerminal() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func PrintAlert(str string) {
	fmt.Print(ColorRed + str + ColorReset)
}

func PrintNotice(str string) {
	fmt.Print(ColorCyan + str + ColorReset)
}
