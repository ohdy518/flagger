package renderData

import (
	"flagger/converter"
	"flagger/micro"
	"flagger/parseData"
	"flagger/writeData"
	"fmt"
)

var recognizedTypeChar = []string{"s", "i", "f", "b"}

func AppendEntity(key string, typeChar string, value string) {
	// Check if typeChar is recognized
	var isTypeCharRecognized bool = false
	for _, v := range recognizedTypeChar {
		if v == typeChar {
			isTypeCharRecognized = true
		}
	}
	if !isTypeCharRecognized {
		panic("Invalid type: " + typeChar)
	}
	// This function doesn't check if key name is valid, assuming the invoker has done the job.

	// If the type is string, convert the value to Base64.
	if typeChar == "s" {
		value = micro.EncodeToBase64(value)
	}

	// Put into a string
	var renderedString string = fmt.Sprintf("%s.%s %s", key, typeChar, value)

	// Send it to the writer
	writeData.AppendRaw(renderedString)
}

func EditEntity(key string, typeChar string, newValue string, optionalPanicOnAbsence ...bool) {
	var panicOnAbsence bool = true
	var isFound bool = false
	var entityIndex int
	var tempEntity parseData.Entity
	if len(optionalPanicOnAbsence) > 0 {
		panicOnAbsence = optionalPanicOnAbsence[0]
	}

	for index, entityKey := range parseData.EntityList {
		if entityKey.Key == key {
			fmt.Println(index, entityKey.Key, key)
			isFound = true
			entityIndex = index
			break
		}
	}

	if !isFound {
		if panicOnAbsence {
			panic("Invalid entity key: " + key)
		}
		AppendEntity(key, typeChar, newValue)
		return
	}
	tempEntity.Key = key
	switch typeChar {
	case "i":
		tempEntity.Value = converter.StringToInt(newValue)
	case "f":
		tempEntity.Value = converter.StringToFloat(newValue)
	case "b":
		tempEntity.Value = converter.StringToBool(newValue)
	case "s":
		tempEntity.Value = newValue
	}

	tempEntity.TypeChar = typeChar

	parseData.EntityList[entityIndex] = tempEntity

	// Rewrite the list to the file.
}

func PushToFile() {
	// Clear the file
	writeData.WriteRaw("")

	// Read the EntityList, and push it.
	for _, entity := range parseData.EntityList {
		var tempValue interface{}
		var stringTempValue string
		if entity.TypeChar == "s" {
			//tempValue = micro.EncodeToBase64(entity.Value.(string))
			tempValue = entity.Value
		} else {
			tempValue = entity.Value
		}
		// Change tempValue to string

		switch v := tempValue.(type) {
		case string:
			stringTempValue = v
		case float64:
			stringTempValue = converter.FloatToString(tempValue.(float64))
		case bool:
			stringTempValue = converter.BoolToString(tempValue.(bool))
		case int:
			stringTempValue = converter.IntToString(tempValue.(int))

		}
		AppendEntity(entity.Key, entity.TypeChar, stringTempValue)
	}
}
