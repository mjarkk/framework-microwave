package database

import (
	"errors"
	"fmt"
	"os"

	"github.com/globalsign/mgo"
	"github.com/mjarkk/framework-microwave/pkg/filehandeler"
	"github.com/mjarkk/framework-microwave/pkg/globalvars"
	"github.com/mjarkk/framework-microwave/pkg/gui"
	"github.com/mjarkk/framework-microwave/pkg/types"
)

// Init initializes the database
// check if it can connect
func Init() {
	// Conect to mongodb
	mongoURL := os.Getenv("MONGODB_URL")
	mongoDatabase := os.Getenv("MONGODB_DATABASE")

	session, err := mgo.Dial(mongoURL)
	if err != nil {
		gui.CritErr("Can't connect to mongodb, this has probebly something to do with the mongodb URL or mongodb service is not running\nerror: " + err.Error())
	}

	// database :=
	session.DB(mongoDatabase)

	// fmt.Println(globalvars.Settings.MigrationPath)
	YmlList, err := filehandeler.ReadAllYamlIn(globalvars.Settings.MigrationPath)
	if err != nil {
		gui.CritErr("can't read yaml database files in " + globalvars.Settings.MigrationPath + " error: " + err.Error())
	}
	parsedData, ymlErr := ParseDBYml(YmlList)
	if ymlErr.Err != nil {
		gui.CritErr("Error in yml file: " + ymlErr.File + ", error: " + ymlErr.Err.Error())
	}
	fmt.Println(parsedData)
}

// ParseDBYml parses a list of yaml files to workable objects
func ParseDBYml(items []types.YmlList) ([]types.DBList, types.YmlError) {
	var fullObjects []types.DBList
	for _, item := range items {
		toAdd, err := ParseYml(item.FileContents)
		if err != nil {
			return fullObjects, types.YmlError{
				Err:  err,
				File: item.FileName,
			}
		}
		fullObjects = append(fullObjects, toAdd)
	}
	return fullObjects, types.YmlError{}
}

// ParseYml parses a yaml object
func ParseYml(contentRaw types.Inter) (types.DBList, error) {
	var toReturn types.DBList
	var content types.MapInter
	content = contentRaw.(map[interface{}]interface{})

	if content["data"] != nil {
		// parse the data object
		out, err := loopItemData(content["data"])
		if err != nil {
			return toReturn, err
		}
		toReturn = out
	} else {
		return toReturn, errors.New("data object does not exsist")
	}

	if content["premissions"] != nil {
		// parse the premissions object
	}

	if content["links"] != nil {
		// parse the links object
	}
	return toReturn, nil
}

func loopItemData(contentRaw types.Inter) (types.DBList, error) {
	toReturn := make(types.DBList)
	var content types.MapInter = contentRaw.(map[interface{}]interface{})
	for i, value := range content {
		index := i.(string)
		iType := fmt.Sprintf("%T", value)

		toAdd := types.DBItem{}

		if iType == "int" ||
			iType == "int8" || iType == "int16" || iType == "int32" || iType == "int64" ||
			iType == "uint" || iType == "uint8" || iType == "uint16" || iType == "uint32" || iType == "uint64" ||
			iType == "bool" || iType == "float32" || iType == "float64" || iType == "FloatType" {
			// TODO add full path to yaml item
			return toReturn, errors.New(index + " Contains wrong data type '" + iType + "'")
		} else if iType == "string" {
			fmt.Println(index, "=", value)

		} else if iType == "[]interface {}" {
			transform := value.([]interface{})[0]
			valueType := fmt.Sprintf("%T", transform)
			if valueType == "string" {
				fmt.Println(index, "=", transform)
			} else {
				// is object in array in object
				fmt.Println(index, "= To parse object")
				// loopItemData(transform)
			}
		} else {
			// is object in object
			fmt.Println(index, "= To parse object")
			// loopItemData(transform)
		}

		toReturn[index] = toAdd
	}
	return toReturn, nil
}
