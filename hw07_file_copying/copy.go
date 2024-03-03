package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(from string, to string, offset int64, limit int64) error {
	src, err := os.Open(from)
	if err != nil {
		return err
	}
	defer src.Close()

	stat, err := src.Stat()
	if err != nil {
		return err
	}
	// Ожидаемое поведение согласно заданию
	if limit == 0 {
		limit = stat.Size()
	}

	dst, err := os.Create(to)
	if err != nil {
		return err
	}
	defer dst.Close()
	return copyWithOffset(src, dst, offset, limit)
}

func copyWithOffset(r io.ReaderAt, w io.Writer, offset int64, limit int64) error {
	r = io.NewSectionReader(r, offset, limit)
	_, err := io.Copy(w, r.(io.Reader))
	return err
}
