package path

import (
	"fmt"
	"os"
	"strings"
)

func filter(in []string, predicate func(string) bool) []string {
	out := make([]string, 0, cap(in))

	for _, value := range in {
		if predicate(value) {
			out = append(out, value)
		}
	}

	return out
}

func getPATH(path string, listed bool) string {
	if listed {
		path = strings.Join(strings.Split(path, ":"), "\n")
	}
	return path
}

func addToPATH(path string, value string) string {
	if value != "" {
		path += ":" + value
	}
	return path
}

func removeFromPath(path string, value string) string {
	compare := func(comparable string) bool {
		return comparable == value
	}

	pathArray := strings.Split(path, ":")

	resultArray := make

	pathArray = strings.Split(result, ":")

	for _, pathVar = range pathArray {

	}

	// ...

	return result
}

func main() {
	path := os.Getenv("PATH")

	fmt.Println(addToPATH("testing..."))
}
