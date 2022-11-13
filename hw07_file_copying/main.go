// Package main -- HW07 otus.
package main

import (
	"fmt"
	"os"

	"github.com/taksenov/otus_go/hw07_file_copying/cmd"
)

func main() {
	root := cmd.Root()

	if err := root.Execute(); err != nil {
		fmt.Println("[ERROR]:", err)
		os.Exit(1)
	}
}
