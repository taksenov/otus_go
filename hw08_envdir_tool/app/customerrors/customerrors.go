// Package customerrors -- error handlers.
package customerrors

import (
	"bytes"
	"log"
)

// HandleErr custom error handler.
func HandleErr(err error, fn string) {
	var buf bytes.Buffer
	var logger = log.New(&buf, "[ERROR]: ", log.Lmsgprefix)

	logger.Print("func: ", fn, "; message: ", err)
	log.Fatal(&buf)
}

