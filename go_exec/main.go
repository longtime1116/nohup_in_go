package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run go_exec/main.go <true|false> <true|false>")
		os.Exit(1)
	}
	nohupGo := os.Args[1] == "true"
	nohupSh := os.Args[2] == "true"
	var scriptPath string
	if nohupSh {
		scriptPath = "./scripts/nohup.sh"
	} else {
		scriptPath = "./scripts/hup.sh"
	}

	var cmd *exec.Cmd
	if nohupGo {
		cmd = exec.Command("nohup", "/bin/sh", scriptPath, "&")
	} else {
		cmd = exec.Command("/bin/sh", scriptPath)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Printf("Failed to start the script: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Started the script: %s with go nohup: %t, sh nohup: %t\n", scriptPath, nohupGo, nohupSh)
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("main.go running...")
	}
}
