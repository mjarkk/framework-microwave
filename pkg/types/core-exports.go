package types

// Init is the input in main.go
type Init struct {
	MigrationPath string
}

// MapInter is a wrapper around map[interface{}]interface{}
type MapInter map[interface{}]interface{}

// YmlList defines what a item is inside a yaml list
type YmlList struct {
	FileName     string
	FileContents Inter
}

// YmlError Is the contents of a error inside a yml file
type YmlError struct {
	File string // path file that has an error
	Err  error  // the actual error
}

// Inter is just a shorthand for interface{}
type Inter interface{}

// DBList contains the database structure
type DBList map[interface{}]DBItem

// DBItem is one database item's type
type DBItem struct {
	IsObject               bool   // Is this item an object with other items
	ObjectContents         DBList // If IsObject is true what is the content of this object
	IgnoreSettings         bool   // Code needs to ignore Settings
	Settings               DBItemSettings
	IgnoreDataRequirements bool // Code can ignore DataRequirements
	DataRequirements       DBItemDataRequirements
	IgnoreDataFilter       bool // Code can ignore DataFilter
	DataFilter             DBItemDataFilter
	IgnorePremissions      bool // Code can ignore Premissions
	Premissions            DBItemPremissions
}

// DBItemSettings is the Settings section in DBItem
// the object settings affect the data request parser
type DBItemSettings struct {
	IsArray bool // Is the item an array affect?
	Primary bool // This item can be used to search for linked item later
	Linked  bool // Is linked value (can't be eddited only when array can add array item)
}

// DBItemDataRequirements is the DataRequirements section in DBItem
// Data requirements (this type )
type DBItemDataRequirements struct {
	DataType     string   // Set data type inside var
	JSONType     string   // If the DataType is json set what is the json type (value in datatype and in here: json -> default, json:graphQL -> graphql, json:raw -> raw)
	File         bool     // If dataType is file
	Unique       bool     // Value needs to be unique to other items in array or document
	Required     bool     // Check if the value needs to be required
	MinLen       uint32   // Min lenght of value (default 0, max 4.294.967.295, min 0)
	MaxLen       uint32   // Max lenght of value (default 0, max 4.294.967.295, min 0)
	Regex        string   // Set matching regex (default empty)
	ReqUppercase bool     // Value needs at least 1 upper case
	ReqLowercase bool     // Value needs at least 1 lower case
	ReqSpecial   bool     // Value needs at least 1 special character like -, ;, *, etc..
	Checkers     []string // A list of checkers defined by the user in order
}

// DBItemDataFilter is the DataFilter section in DBItem
// Data filters before saving (this type has items that do set data)
type DBItemDataFilter struct {
	Order         []string    // The order of executing the filters
	transformers  []string    // A list of custom tranformers defined by the user
	HasDefaultVal bool        // Value has default value
	DefaultVal    interface{} // The default value (will be changed to the datatype)
	HasHash       bool        // Has Hash filter
	Hash          string      // Transform value to hash (the contents of this value is the hashing algorithm)
}

// DBItemPremissions is the Premissions section in DBItem
// Set premisions for this item
type DBItemPremissions struct {
	Read   []string
	Write  []string
	Delete []string
}

// GenerateDBItem returns a DBItem with default values
func GenerateDBItem() DBItem {
	return DBItem{
		IsObject:       false,
		ObjectContents: DBList{},
		IgnoreSettings: true,
		Settings: DBItemSettings{
			IsArray: false,
			Primary: false,
			Linked:  false,
		},
		IgnoreDataRequirements: true,
		DataRequirements: DBItemDataRequirements{
			DataType:     "",
			JSONType:     "",
			File:         false,
			Unique:       false,
			Required:     false,
			MinLen:       0,
			MaxLen:       0,
			Regex:        "",
			ReqUppercase: false,
			ReqLowercase: false,
			ReqSpecial:   false,
			Checkers:     []string{},
		},
		IgnoreDataFilter: true,
		DataFilter: DBItemDataFilter{
			Order:         []string{},
			transformers:  []string{},
			HasDefaultVal: false,
			DefaultVal:    "",
			HasHash:       false,
			Hash:          "sha256",
		},
		IgnorePremissions: true,
		Premissions: DBItemPremissions{
			Read:   []string{},
			Write:  []string{},
			Delete: []string{},
		},
	}
}
