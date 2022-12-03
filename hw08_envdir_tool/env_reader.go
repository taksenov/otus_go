// Package main -- HW08 otus: envdir tool.
package main

import (
	"io"
	"os"

	"hw08_envdir_tool/app/filesystem"
)

// Environment customers environment structure.
type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	res := make(Environment, 0)

	for _, file := range files {
		needRemove := false

		val, err := filesystem.ReadFileFirstLine(dir + "/" + file.Name())
		if err == io.EOF {
			needRemove = true
		} else if err != nil {
			continue
		}

		res[file.Name()] = EnvValue{Value: val, NeedRemove: needRemove}
	}

	return res, nil
}
