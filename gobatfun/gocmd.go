package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("cmd.exe", "/K", "start", "", "powershell.exe", "-NoExit", "-Command", "& {Get-TimeZone; Write-Host 'blah'; Read-Host 'Please hit enter to exit'; exit}")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
}
