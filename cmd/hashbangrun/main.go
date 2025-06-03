// build +windows
package main

import (
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	toolName := os.Args[0]
	args := os.Args[1:]

	en := path.Base(toolName)
	pd := path.Dir(toolName)

	if !strings.HasSuffix(en, ".exe") {
		log.Fatalf("This is not supposed to happen -- Windows only")
	}

	splitName := en[:len(en)-4]

	scriptToRun := path.Join(pd, splitName)

	_, err := os.Stat(scriptToRun)
	if os.IsNotExist(err) {
		log.Fatalf("The corresponding script %s does not exist", scriptToRun)
	}

	cmd := exec.Command("bash", append([]string{scriptToRun}, args...)...)
	cmd.Run()
}
