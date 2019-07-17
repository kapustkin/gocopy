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
	err = copy(reader, writer, fileLen.Size(), func(progress int64) {
		progBar.SetCurrent(progress)
	})
	progBar.Finish()
	defer writer.Close()
	defer reader.Close()
	if err != nil {
		return fmt.Errorf("Ошибка при копировании файла %s", err)
	}
	return nil
}

func copy(reader io.Reader, writer io.Writer, size int64, callback func(progress int64)) error {
	step := size / int64(100)
	if step == 0 {
		step = size
	}
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

func getReader(from string) (io.ReadCloser, error) {
	fromFile, ioErr := os.Open(from)
	var err error
	if ioErr != nil {
		err = fmt.Errorf("Не удается открыть исходный файл %s", from)
	}
	return fromFile, err
}

func getWriter(to string) (io.WriteCloser, error) {
	toFile, ioErr := os.Create(to)
	var err error
	if ioErr != nil {
		err = fmt.Errorf("Не удается открыть целевой файл %s", to)
	}
	return toFile, err
}
