//Golang program to print log levels

package main

import "fmt"

type LogLevel int

const (
	LogTrace LogLevel = iota
	LogDebug
	LogInfo
	LogWarning
	LogError
)

var levelName = []string{"TRACE", "DEBUG", "INFO", "WARNING", "ERROR"}

func (l LogLevel) String() string {
	if l < LogTrace || l > LogError {
		return "INVALID LOG LEVEL"
	}
	return levelName[l]
}

func printLogLevel(level LogLevel) {
	fmt.Printf("Log Level: %d %s\n", level, level.String())

}

func main() {
	printLogLevel(LogTrace)
	printLogLevel(LogDebug)
	printLogLevel(LogInfo)
	printLogLevel(LogWarning)
	printLogLevel(LogError)
}
