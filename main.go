package main

import (
	"flag"
	"log"

	"github.com/kapustkin/gocopy/commands"
)

func main() {
	argFrom := flag.String("from", "", "Путь к исходному файлу")
	argTo := flag.String("to", "", "Путь к новому файлу")
	flag.Parse()
	err := commands.CopyFileToFile(*argFrom, *argTo)
	if err != nil {
		log.Fatalf("Ошибка при выполнении операции! %s", err)
	}
}
