package filehandeler

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"
)

// Inter is just a shorthand for interface{}
type Inter interface{}

// YmlList returns a list of Yml
type YmlList []struct {
	fileName     string
	fileContents Inter
}

// ReadAllYamlIn reads all yml files in a directory and returns the files + filenames
func ReadAllYamlIn(path string) (YmlList, error) {
	toReturn := YmlList{}
	files, err := ListFiles(path)
	if err != nil {
		return toReturn, err
	}
	for _, file := range files {
		fmt.Println(file)
	}
	return toReturn, nil
}

// ReadYmlFile reads a yaml file and returns the contents of the yaml file as object
func ReadYmlFile(path string) (Inter, error) {
	fileData, err := OpenFile(path)
	if err != nil {
		return "", err
	}
	yamlObject, err := ReadYml(fileData)
	if err != nil {
		return "", err
	}
	return yamlObject, err
}

// ReadYml returns the contents of a yaml file into an object
func ReadYml(data string) (Inter, error) {
	var t Inter
	err := yaml.Unmarshal([]byte(data), &t)
	return t, err
}
