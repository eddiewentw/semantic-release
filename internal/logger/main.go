package logger

import "fmt"

func logWithLevel(message string, level string) {
	fmt.Print(logHeader)

	if level != "" {
		fmt.Print(" " + level)
	}

	fmt.Println(" " + message)
}

func DebugLog(message string, isDebug bool) {
	if isDebug == true {
		logWithLevel(message, colorDebug+"Debug"+colorReset+":")
	}
}

func Log(message string) {
	logWithLevel(message, "")
}

func Warning(message string) {
	logWithLevel(message, colorWarning+"Warning"+colorReset+":")
}

func Error(err error) {
	logWithLevel(err.Error(), colorError+"Error"+colorReset+":")
}
