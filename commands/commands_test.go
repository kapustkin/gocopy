package commands

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var fileName = "123"

func removeFile(fileName string) {
	if _, err := os.Stat(fileName); err == nil {
		errr := os.Remove(fileName)
		if errr != nil {
			fmt.Printf("Ошибка %s", errr)
		}
	}
}

func TestGetWriter(t *testing.T) {
	writer, err := getWriter(fileName)
	writer.Close()
	assert.Nil(t, err)
	removeFile(fileName)
}

func TestGetReaderNoFile(t *testing.T) {
	_, err := getReader(fileName)
	nErr := fmt.Errorf("Не удается открыть исходный файл %s", fileName)
	assert.Equal(t, nErr, err)
}

func TestGetReaderWitjFile(t *testing.T) {
	writer, err := getWriter(fileName)
	writer.Close()
	reader, err := getReader(fileName)
	assert.Nil(t, err)
	reader.Close()
	removeFile(fileName)
}

func TestCopy(t *testing.T) {
	reader := strings.NewReader("12345")
	writer := &bytes.Buffer{}
	copy(reader, writer, int64(5), func(progress int64) {
		assert.Equal(t, int64(5), progress)
	})
	assert.Equal(t, "12345", writer.String())
}
