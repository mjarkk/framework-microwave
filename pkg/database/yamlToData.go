package database

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/mjarkk/framework-microwave/pkg/regex"

	"github.com/mjarkk/framework-microwave/pkg/types"
)

// ParseDBYmlS parses a list of yaml files to workable objects
func ParseDBYmlS(items []types.YmlList) ([]types.DBList, types.YmlError) {
	var fullObjects []types.DBList
	var errors []types.YmlError

	// Make process multiple threads
	var wg sync.WaitGroup
	wg.Add(len(items))

	for _, item := range items {
		go func(item types.YmlList) {
			defer wg.Done()
			toAdd, err := ParseYml(item.FileContents)
			if err != nil {
				errors = append(errors, types.YmlError{
					Err:  err,
					File: item.FileName,
				})
			}
			fullObjects = append(fullObjects, toAdd)
		}(item)
	}

	wg.Wait()

	// return first error when there are errors
	if len(errors) > 0 {
		return fullObjects, errors[0]
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
		out, err := loopOverDBData(content["data"])
		if err != nil {
			return toReturn, err
		}
		toReturn = out
	} else {
		return toReturn, errors.New("data object does not exsist")
	}

	if content["premissions"] != nil {
		// parse the premissions object
		newContent, err := loopOverPremissions(toReturn, content["premissions"], []string{})
		if err != nil {
			return toReturn, err
		}
		toReturn = newContent
	}

	if content["links"] != nil {
		// parse the links object
	}
	return toReturn, nil
}

// loopOverPremissions parses the premisions object in a DB yaml file
func loopOverPremissions(DBList types.DBList, contentRaw types.Inter, nDeep []string) (types.DBList, error) {
	var premissionsObj types.MapInter = contentRaw.(map[interface{}]interface{})
	for i, value := range premissionsObj {
		index := i.(string)
		iType := fmt.Sprintf("%T", value)

		if len(nDeep) > 0 {
			fmt.Println(nDeep, "=", index, "=", iType)
		}

		if iType == "map[interface {}]interface {}" {

			newDBList, err := loopOverPremissions(DBList, value, append(nDeep, index))
			if err != nil {
				return DBList, err
			}
			DBList = newDBList

		} else if iType == "[]interface {}" {

		} else if iType == "string" {

		} else {
			errMsg := ymlPath("premissions", nDeep, index) + " Hash a not supported data type `" + iType + "`, see: https://github.com/mjarkk/framework-microwave/blob/master/docs/databasefiles.md for more info"
			return DBList, errors.New(errMsg)
		}
	}
	return DBList, nil
}

// ymlPath generates a readable path to a item in a yaml file
func ymlPath(prefix string, path []string, itemName string) string {
	returnPath := prefix + "."
	if prefix == "" {
		returnPath = ""
	}
	returnPath = returnPath + strings.Join(path, ".")
	if itemName != "" {
		returnPath = returnPath + "." + itemName
	}
	return returnPath
}

// loopOverDBData parses the data object in a DB yaml file
func loopOverDBData(contentRaw types.Inter) (types.DBList, error) {
	// TODO add path to item error in yaml
	toReturn := make(types.DBList)
	var content types.MapInter = contentRaw.(map[interface{}]interface{})
	for i, value := range content {
		index := i.(string)
		iType := fmt.Sprintf("%T", value)

		toAdd := types.GenerateDBItem()

		if iType == "int" ||
			iType == "int8" || iType == "int16" || iType == "int32" || iType == "int64" ||
			iType == "uint" || iType == "uint8" || iType == "uint16" || iType == "uint32" || iType == "uint64" ||
			iType == "bool" || iType == "float32" || iType == "float64" || iType == "FloatType" {

			// detect invalid data type in yaml file
			return toReturn, errors.New(index + " Contains wrong data type '" + iType + "'")

		} else if regex.Match("HOST", index) ||
			regex.Match("edit", index) || regex.Match("view", index) ||
			regex.Match(".*HOST.*", index) || regex.Match("delete", index) {

			// detect wrong item name
			// these names might be later used by the framework or
			// might make errors with the premissions object
			return toReturn, errors.New("Data item can't be called: `" + index + "`")

		} else if iType == "string" {

			update, err := detectDBItem(toAdd, value.(string))
			if err != nil {
				return toReturn, err
			}
			toAdd = update

		} else if iType == "[]interface {}" {
			transform := value.([]interface{})[0]
			valueType := fmt.Sprintf("%T", transform)
			if valueType == "string" {

				update, err := detectDBItem(toAdd, transform.(string))
				if err != nil {
					return toReturn, err
				}
				update.Settings.IsArray = true
				update.IgnoreSettings = false
				toAdd = update
			} else {
				// is object in array
				toAdd.IsObject = true
				toAdd.Settings.IsArray = true
				toAdd.IgnoreSettings = false
				update, err := loopOverDBData(transform)
				if err != nil {
					return toReturn, err
				}
				toAdd.ObjectContents = update
			}
		} else {
			// is object in object
			toAdd.IsObject = true
			update, err := loopOverDBData(value)
			if err != nil {
				return toReturn, err
			}
			toAdd.ObjectContents = update
		}

		toReturn[index] = toAdd
	}
	return toReturn, nil
}
