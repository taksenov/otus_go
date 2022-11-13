/*
Copyright © 2022 taksenov@gmail.com
*/

// Package cmd -- comand line interface app.
package cmd

import (
	"errors"
	"io"
	"log"
	"os"

	"github.com/vardius/progress-go"
)

var (
	// ErrUnsupportedFile file is unsupported.
	ErrUnsupportedFile = errors.New("unsupported file")

	// ErrOffsetExceedsFileSize offset exceeds file size.
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

// Copy function for copy files.
func Copy(fromPath, toPath string, offset, limit int64) error {
	// Src
	file, err := os.OpenFile(fromPath, os.O_RDONLY, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	st, err := file.Stat()
	if err != nil {
		return err
	}
	if offset > st.Size() {
		return ErrOffsetExceedsFileSize
	}
	if limit == 0 {
		limit = st.Size()
	}
	if offset > 0 {
		if _, err := file.Seek(offset, io.SeekStart); err != nil {
			return err
		}
	}

	// Dst
	newFile, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	// Progress Bar
	bar := progress.New(0, limit)

	_, _ = bar.Start()
	defer func() {
		if _, err := bar.Stop(); err != nil {
			log.Printf("failed to finish progress: %v", err)
		}
	}()

	// Copy
	c := masterOfChunks(limit)
	chunk := c.size
	for {
		if c.count == 0 {
			chunk = c.tail
		}

		readData := readChunk(file, chunk)
		if readData != nil || c.count >= 0 {
			_, err := newFile.Write(readData)
			if err != nil {
				return err
			}

			c.count--

			_, _ = bar.Advance(chunk)
		} else {
			break
		}

		if c.count < 0 {
			break
		}
	}

	return nil
}

func readChunk(f *os.File, s int64) []byte {
	buf := make([]byte, s)

	n, err := f.Read(buf)
	if err == io.EOF {
		return nil
	}

	if err != nil {
		return nil
	}

	return buf[0:n]
}

type chunks struct {
	count int64
	size  int64
	tail  int64
}

func masterOfChunks(somethingWhole int64) chunks {
	const measure int64 = 10

	res := chunks{
		count: 0,
		size:  0,
		tail:  0,
	}

	if somethingWhole < measure {
		res.count = 1
		res.size = somethingWhole

		return res
	}

	q := somethingWhole / measure
	r := somethingWhole % measure

	res.count = measure
	res.size = q
	res.tail = r

	return res
}
