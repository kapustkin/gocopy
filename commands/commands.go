package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

// CopyFileToFile функция для копирования ИЗ В
// from - откуда копировать
// to 	- куда попировать
// buff - размер буфера
func CopyFileToFile(from string, to string) error {
	reader, err := getReader(from)
	if err != nil {
		return err
	}
	writer, err := getWriter(to)
	if err != nil {
		return err
	}

	fileLen, err := os.Stat(from)
	if err != nil {
		return err
	}

	progBar := pb.Full.Start64(fileLen.Size())
	progBar.SetWriter(os.Stdout)
	copy(reader, writer, fileLen.Size(), func(progress int64) {
		progBar.SetCurrent(progress)
	})
	progBar.Finish()
	return nil
}

func copy(reader io.Reader, writer io.Writer, size int64, callback func(progress int64)) error {

	step := size / int64(100)
	progress := int64(0)
	for progress < size {
		written, err := io.CopyN(writer, reader, step)
		if err == io.EOF {
			progress += written
			callback(progress)
			break
		}
		if err != nil {
			return err
		}
		progress += written
		callback(progress)
	}
	return nil
}

func getReader(from string) (io.Reader, error) {
	fromFile, ioErr := os.Open(from)
	var err error
	if ioErr != nil {
		err = fmt.Errorf("Не удается открыть исходный файл %s", from)
	}
	return fromFile, err
}

func getWriter(to string) (io.Writer, error) {
	toFile, ioErr := os.Create(to)
	var err error
	if ioErr != nil {
		err = fmt.Errorf("Не удается открыть целевой файл %s", to)
	}
	return toFile, err
}
