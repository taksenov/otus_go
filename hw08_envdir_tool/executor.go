// Package main -- HW08 otus: envdir tool.
package main

import (
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	ts := ""
	e := make([]string, 0)
	for k, v := range env {
		if v.NeedRemove {
			os.Unsetenv(k)
		} else {
			t := k + "=" + v.Value
			ts += " " + t
			e = append(e, t)
		}
	}

	commd := exec.Command(cmd[0], cmd[1:]...)
	commd.Env = append(os.Environ(), e...)
	commd.Stdin = os.Stdin
	commd.Stdout = os.Stdout
	commd.Stderr = os.Stderr

	err := commd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return
}
