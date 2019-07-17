package commands

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func removeFile(fileName string) {
	if _, err := os.Stat(fileName); err == nil {
		errr := os.Remove(fileName)
		if errr != nil {
			fmt.Printf("ошибка %s", errr)
		}
	}
}

func TestGetWriter(t *testing.T) {
	writer, err := getWriter("123")
	writer.Close()
	assert.Nil(t, err)
	removeFile("123")
}

func TestGetReaderNoFile(t *testing.T) {
	_, err := getReader("123")
	nErr := fmt.Errorf("не удается открыть исходный файл %s", "123")
	assert.Equal(t, nErr, err)
}

func TestGetReaderWitjFile(t *testing.T) {
	writer, _ := getWriter("123")
	writer.Close()
	reader, err := getReader("123")
	assert.Nil(t, err)
	reader.Close()
	removeFile("123")
}

func TestCopy(t *testing.T) {
	reader := strings.NewReader("12345")
	writer := &bytes.Buffer{}
	err := copy(reader, writer, int64(5), func(progress int64) {
		assert.Equal(t, int64(5), progress)
	})
	assert.Nil(t, err)
	assert.Equal(t, "12345", writer.String())
}
