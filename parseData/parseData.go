package parseData

import (
	"flagger/converter"
	"flagger/micro"
	"flagger/readData"
	"strings"
)

type Entity struct {
	Key      string
	TypeChar string
	Value    interface{}
}

var EntityList []Entity = make([]Entity, 0)

func Parse() {
	var dataString string = readData.Read()
	var dataList = strings.Split(dataString, "\n")
	var refinedDataList = make([]string, 0)
	for _, data := range dataList {
		if strings.TrimSpace(data) == "" {
			continue
		}
		refinedDataList = append(refinedDataList, data)
	}

	for _, refinedData := range refinedDataList {
		//fmt.Println(refinedData)
		var keyString string = strings.Split(refinedData, ".")[0]
		var typeString string = strings.Split(strings.Split(refinedData, ".")[1], " ")[0]
		var valueString string = strings.Split(refinedData, " ")[1]
		//fmt.Println(keyString, typeString, valueString)

		var tempEntity Entity
		tempEntity.Key = keyString

		switch typeString {
		case "i":
			tempEntity.Value = converter.StringToInt(valueString)
		case "f":
			tempEntity.Value = converter.StringToFloat(valueString)
		case "b":
			tempEntity.Value = converter.StringToBool(valueString)
		case "s":
			tempEntity.Value = micro.DecodeFromBase64(valueString)
		}
		tempEntity.TypeChar = typeString

		EntityList = append(EntityList, tempEntity)
	}
}

func GetEntityValue(key string) interface{} {
	for _, entry := range EntityList {
		if entry.Key == key {
			return entry.Value
		}
	}
	return nil
}
