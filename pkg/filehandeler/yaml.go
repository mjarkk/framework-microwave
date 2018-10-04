package filehandeler

import (
	"github.com/mjarkk/framework-microwave/pkg/types"
	yaml "gopkg.in/yaml.v2"
)

// ReadAllYamlIn reads all yml files in a directory and returns the files + filenames
func ReadAllYamlIn(path string) ([]types.YmlList, error) {
	toReturn := []types.YmlList{}
	files, err := ListFiles(path)
	if err != nil {
		return toReturn, err
	}
	for _, file := range files {
		ymlContent, err := ReadYmlFile(CDir(path, file))
		if err != nil {
			return toReturn, err
		}
		toReturn = append(toReturn, types.YmlList{
			FileName:     file,
			FileContents: ymlContent,
		})
	}
	return toReturn, nil
}

// ReadYmlFile reads a yaml file and returns the contents of the yaml file as object
func ReadYmlFile(path string) (types.Inter, error) {
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
func ReadYml(data string) (types.Inter, error) {
	var t types.Inter
	err := yaml.Unmarshal([]byte(data), &t)
	return t, err
}
