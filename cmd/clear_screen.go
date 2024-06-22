package main

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

func ClearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // For Windows
	} else {
		cmd = exec.Command("clear") // For Unix-like systems
	}

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("\n\tError clearing screen: %v\n\n", err)
	}
}
