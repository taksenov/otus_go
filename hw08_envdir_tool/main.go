// Package main -- HW08 otus: envdir tool.
package main

import (
	"fmt"
	"os"

	cerr "hw08_envdir_tool/app/customerrors"
)

func main() {
	if len(os.Args) < 2 {
		cerr.HandleErr(fmt.Errorf("not enough arguments! Must be 2 or more"), "os.Args length")
	}

	dirWithTreasures := os.Args[1]

	env, err := ReadDir(dirWithTreasures)
	if err != nil {
		cerr.HandleErr(err, "ReadDir")
	}

	returnCode := RunCmd(os.Args[2:], env)

	os.Exit(returnCode)
}
