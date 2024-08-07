package readData

import (
	"flagger/micro"
	"os"
	"strings"
)

var fileLocation string = ""

func DeclareFileLocation(newLocation string) {
	fileLocation = newLocation
}

func Read() string {
	content, err := os.ReadFile(fileLocation)
	micro.CheckError(err)
	var stringContent string = string(content)

	// Resolve CRLF issues.
	stringContent = strings.ReplaceAll(stringContent, "\r", "")

	return stringContent
}
