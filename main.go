package main

import (
	"flagger/parseData"
	"flagger/readData"
	"flagger/renderData"
	"flagger/writeData"
)

func main() {
	readData.DeclareFileLocation("./test/test.txt")
	writeData.DeclareFileLocation("./test/test.txt")
	//writeData.WriteRaw("")
	parseData.Parse()
	renderData.EditEntity("thisFloat", "i", "6")
	renderData.PushToFile()
}
