package main

import (
	"fmt"
	"os"
	"strings"
)

func getPATH(listed bool) string {
	result := os.Getenv("PATH")
	if listed {
		result = strings.Join(strings.Split(result, ":"), "\n")
	}
	return result
}

func addToPATH(value string) string {
	result := os.Getenv("PATH")
	if value != "" {
		result += ":" + value
	}
	return result
}

func removeFromPath(value string) string {
	result := os.Getenv("PATH")

	// pathArray = strings.Split(result, ":")
	// ...

	return result
}

func main() {
	fmt.Println(addToPATH("testing..."))
}
