package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var scriptToRun = `{Get-TimeZone; Write-Host 'blah'}`

	cmd := exec.Command("cmd.exe", "/C", "start", "", "powershell.exe", "-NoExit", "-Command", "& "+scriptToRun+";Read-Host 'Please hit enter to exit'; exit")
	// cmd := exec.Command("cmd.exe", "/C", "start", "", "powershell.exe", "-NoExit", "-Command", "& "+scriptToRun)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
