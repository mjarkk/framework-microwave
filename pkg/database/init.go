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

	// switch content := content.(type) {
	// case map[interface{}]interface{}:
	// 	for index := range content {
	// 		if index == "data" {
	// 			// handel the data object
	// 			if lastState == "" {
	// 				data, err := loopItemData(content)
	// 				if err != nil {
	// 					return data, err
	// 				}
	// 				toReturn = data
	// 			} else {
	// 				return toReturn, errors.New("The order of yml file content is wrong it needs to be: `data`, `premissions`, `links`")
	// 			}
	// 		} else if index == "premissions" {
	// 			// handel the premissions object
	// 		} else if index == "links" {
	// 			// handel the links object
	// 		}
	// 		lastState = index.(string)
	// 	}
	// }
	return toReturn, nil
}

func loopItemData(contentRaw types.Inter) (types.DBList, error) {
	var toReturn types.DBList
	var content types.MapInter = contentRaw.(map[interface{}]interface{})
	for index := range content {

		fmt.Println(index)
	}
	return toReturn, nil
}
