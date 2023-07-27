package logger

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	FileName string
}

func (l *Logger) Log(message string, args any) {
	logFile, err := os.OpenFile(l.FileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic("PathError: Wrong Path")
	}

	defer logFile.Close()

	wrt := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(wrt)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	if args != nil {
		log.Println(message, args)
	} else {
		log.Println(message)
	}
}
