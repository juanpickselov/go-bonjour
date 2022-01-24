package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var scriptToRun = `{Get-TimeZone; Write-Host 'blah'}`

	cmd := exec.Command("cmd.exe", "/C", "start", "PShell Example", "powershell.exe", "-NoExit", "-Command", "& "+scriptToRun+";Read-Host 'Please hit enter to exit'; exit")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
	if err := cmd.Process.Release(); err != nil {
		fmt.Println(errors.New("Error releasing process"))
	}
}
