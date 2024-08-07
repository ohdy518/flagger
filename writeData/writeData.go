package writeData

import (
	"bufio"
	"flagger/micro"
	"os"
	"strings"
)

var fileLocation string = ""

func DeclareFileLocation(newLocation string) {
	fileLocation = newLocation
}

func WriteRaw(rawString string) {
	micro.PanicIfEmpty(fileLocation)
	var byteData = []byte(rawString)
	var err = os.WriteFile(fileLocation, byteData, 0644)
	micro.CheckError(err)
}

func AppendRaw(rawString string, appendWithNewLine ...bool) {
	// Checks if appendWithNewLine is given.
	var useNewLine bool = true
	if len(appendWithNewLine) > 0 {
		useNewLine = appendWithNewLine[0]
	}

	// Appends \n at the beginning of the rawString if useNewLine is true.
	if useNewLine {
		var sb strings.Builder
		sb.WriteString("\n")
		sb.WriteString(rawString)
		rawString = sb.String()
	}

	micro.PanicIfEmpty(fileLocation)
	file, err := os.OpenFile(fileLocation, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	micro.CheckError(err)
	defer func() { err := file.Close(); micro.CheckError(err) }()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(rawString)
	micro.CheckError(err)
	err = writer.Flush()
	micro.CheckError(err)
}
