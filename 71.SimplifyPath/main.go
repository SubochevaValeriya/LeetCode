package main

import (
	"strings"
)

func main() {

}

func simplifyPath(path string) string {
	path = strings.Trim(path, "/")
	absolutePath := strings.Split(path, "/")

	canonicalPath := []string{}
	for _, directory := range absolutePath {
		switch directory {
		case "..":
			if len(canonicalPath) != 0 {
				canonicalPath = canonicalPath[:len(canonicalPath)-1]
			}
		case ".":
			continue
		default:
			if string(directory[0]) == "/" {
				continue
			}

			canonicalPath = append(canonicalPath, directory)
		}
	}

	return "/" + strings.Join(canonicalPath, "/")
}
