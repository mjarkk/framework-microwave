package database

import (
	"errors"
	"strconv"
	"strings"

	"github.com/mjarkk/framework-microwave/pkg/regex"
	"github.com/mjarkk/framework-microwave/pkg/safety"
	"github.com/mjarkk/framework-microwave/pkg/types"
	funk "github.com/thoas/go-funk"
)

func detectDBItem(DBItem types.DBItem, input string) (types.DBItem, error) {
	filters := strings.Split(input, ":")
	IgnoreSettings := true
	IgnoreDataRequirements := true
	IgnoreDataFilter := true
	for _, f := range filters {

		if regex.Match(`(string)|(int)|(boolean)|(byteArr)`, f) {
			// Set DataType normal types
			DBItem.DataRequirements.DataType = f
			IgnoreDataRequirements = false
		} else if regex.Match(`json(:(raw)|(graphql))?`, f) {
			// Set DataType for json
			DBItem.DataRequirements.DataType = "json"
			IgnoreDataRequirements = false
		} else if regex.Match(`required`, f) {
			// Set required vield
			DBItem.DataRequirements.Required = true
			IgnoreDataRequirements = false
		} else if regex.Match(`linked`, f) {
			// Set linked field
			DBItem.Settings.Linked = true
			IgnoreSettings = false
		} else if regex.Match(`primary`, f) {
			// Set primary field
			DBItem.Settings.Primary = true
			IgnoreSettings = false
		} else if regex.Match(`unique`, f) {
			// Set unique vield
			DBItem.DataRequirements.Unique = true
			IgnoreDataRequirements = false
		} else if regex.Match(`default=.*`, f) {
			// Set default value
			DBItem.DataFilter.HasDefaultVal = true
			DBItem.DataFilter.DefaultVal = regex.FindMatch(f, `default=(.*)`, 1)
			DBItem.DataFilter.Order = append(DBItem.DataFilter.Order, "DefaultVal")
			IgnoreDataFilter = false
		} else if regex.Match(`min=\d+`, f) {
			// Set min lenght of the input
			value := regex.FindMatch(f, `min=(\d+)`, 1)
			intValue, err := strconv.Atoi(value)
			if err != nil {
				return DBItem, err
			}
			DBItem.DataRequirements.MinLen = uint32(intValue)
			IgnoreDataRequirements = false
		} else if regex.Match(`max=\d+`, f) {
			// Set max lenght of the input
			value := regex.FindMatch(f, `max=(\d+)`, 1)
			intValue, err := strconv.Atoi(value)
			if err != nil {
				return DBItem, err
			}
			DBItem.DataRequirements.MaxLen = uint32(intValue)
			IgnoreDataRequirements = false
		} else if regex.Match(f, `regex=/.*/`) {
			// Set regex filter
			DBItem.DataRequirements.Regex = regex.FindMatch(f, `regex=/(.*)/`, 1)
			IgnoreDataRequirements = false
		} else if regex.Match(f, `regex=/.*/.*`) {
			// Error wrong regex type
			return DBItem, errors.New("Regex can't contain ")
		} else if regex.Match(f, `reqUppercase`) {
			// Set reqUppercase requirement
			DBItem.DataRequirements.ReqUppercase = true
			IgnoreDataRequirements = false
		} else if regex.Match(f, `reqLowercase`) {
			// Set reqLowercase requirement
			DBItem.DataRequirements.ReqLowercase = true
			IgnoreDataRequirements = false
		} else if regex.Match(f, `reqSpecial`) {
			// Set reqSpecial requirement
			DBItem.DataRequirements.ReqSpecial = true
			IgnoreDataRequirements = false
		} else if regex.Match(f, `hashPassword`) {
			// Set hashPassword data filter
			DBItem.DataFilter.HasHash = true
			DBItem.DataFilter.Hash = "pbkdf2"
			DBItem.DataFilter.Order = append(DBItem.DataFilter.Order, "Hash")
			IgnoreDataFilter = false
		} else if regex.Match(f, `hash=.*`) {
			// Set hasing data filter
			hashType := regex.FindMatch(f, `hash=(.*)`, 1)
			if !funk.Contains(safety.ValidHashTypes, hashType) {
				return DBItem, errors.New("Type `" + hashType + "` doesn't exsist see https://github.com/mjarkk/framework-microwave/blob/master/pkg/safety/README.md for more info")
			}
			DBItem.DataFilter.HasHash = true
			DBItem.DataFilter.Hash = hashType
			DBItem.DataFilter.Order = append(DBItem.DataFilter.Order, "Hash")
			IgnoreDataFilter = false
		} else if regex.Match(f, `file`) {
			// Set content is file
			DBItem.DataRequirements.File = true
			IgnoreDataRequirements = false
		} else if regex.Match(f, `check=(\d|\w|,)+`) {
			// Set user defined function checks
			// TODO check if all function are specified by the user
			checks := strings.Split(regex.Replace(f, "", "check="), ",")
			DBItem.DataRequirements.Checkers = checks
		} else if regex.Match(f, `transformer=(\d|\w|,)+`) {
			// Set user defined function data transforms
			// TODO check if all function are specified by the user
			transforms := strings.Split(regex.Replace(f, "", "transformer="), ",")
			DBItem.DataFilter.Transformers = transforms
			DBItem.DataFilter.Order = append(DBItem.DataFilter.Order, "Transformers")
			IgnoreDataFilter = false
		} else {
			return DBItem, errors.New("Value flag: `" + f + "` doesn't match the api schema look at https://github.com/mjarkk/framework-microwave/blob/master/docs/databasefiles.md for more info")
		}

	}
	DBItem.IgnoreSettings = IgnoreSettings
	DBItem.IgnoreDataRequirements = IgnoreDataRequirements
	DBItem.IgnoreDataFilter = IgnoreDataFilter
	return DBItem, nil
}
