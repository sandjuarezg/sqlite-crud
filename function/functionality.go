package function

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func CleanConsole() {
	fmt.Println("Wait a second . . .")

	time.Sleep(4 * time.Second)
	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
