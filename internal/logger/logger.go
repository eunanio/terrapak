package logger

import (
	"io"
	"log/slog"
	"os"
)

func NewLogger() *os.File {
	logFile, err := os.OpenFile("terrapak.log",os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666); if err != nil {
		panic(err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logger := slog.New(slog.NewJSONHandler(multiWriter,nil))
	slog.SetDefault(logger)
	return logFile
}