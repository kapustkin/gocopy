package internal

import (
	"fmt"
	"io"
	"os"
)

// CopyFileToFile функция для копирования ИЗ В
// from - откуда копировать
// to 	- куда попировать
// buff - размер буфера
func CopyFileToFile(from string, to string, buff int) error {
	reader, err := getReader(from)
	if err != nil {
		return err
	}
	writer, err := getWriter(to)
	if err != nil {
		return err
	}
	buffer, err := getBuffer(buff)
	if err != nil {
		return err
	}
	localBuffer := make([]byte, buffer)
	io.CopyBuffer(writer, reader, localBuffer)

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

func getBuffer(buff int) (int, error) {
	var err error
	buffer := buff * 1024
	if buffer < 1024 {
		err = fmt.Errorf("Размер буфера не может быть меньше 1024 байт")
	} else if buffer > 1024*1024*8 {
		err = fmt.Errorf("Размер буфера не может быть больше 8 Мб")
	}
	return buffer, err
}
