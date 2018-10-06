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
	IsObject       bool     // Is this item an object with other items
	ObjectContents DBList   // If IsObject is true what is the content of this object
	IgnoreSettings bool     // Code needs to ignore Settings
	Settings       struct { // the object settings affect the data request parser
		IsArray bool // Is the item an array affect?
		Primary bool // This item can be used to search for linked item later
		Linked  bool // Is linked value (can't be eddited only when array can add array item)
	}
	IgnoreDataRequirements bool     // Code can ignore DataRequirements
	DataRequirements       struct { // Data requirements (this type )
		DataType     string   // Set data type inside var
		JSONType     string   // If the DataType is json set what is the json type (value in datatype and in here: json -> default, json:graphQL -> graphql, json:raw -> raw)
		file         string   // If dataType is file insert here the link to the file
		unique       bool     // Value needs to be unique to other items in array or document
		Required     bool     // Check if the value needs to be required
		MinLen       uint32   // Min lenght of value (default 0, max 4.294.967.295, min 0)
		MaxLen       uint32   // Max lenght of value (default 0, max 4.294.967.295, min 0)
		Regex        string   // Set matching regex (default empty)
		ReqUppercase bool     // Value needs at least 1 upper case
		ReqLowercase bool     // Value needs at least 1 lower case
		ReqSpecial   bool     // Value needs at least 1 special character like -, ;, *, etc..
		checkers     []string // A list of checkers defined by the user in order
	}
	IgnoreDataFilter bool     // Code can ignore DataFilter
	DataFilter       struct { // Data filters before saving (this type has items that do set data)
		Order         []string    // The order of executing the filters
		transformers  []string    // A list of custom tranformers defined by the user
		HasDefaultVal bool        // Value has default value
		DefaultVal    interface{} // The default value (will be changed to the datatype)
		HasHash       bool        // Has Hash filter
		Hash          string      // Transform value to hash (the contents of this value is the hashing algorithm)
	}
	IgnorePremissions bool     // Code can ignore Premissions
	Premissions       struct { // Set premisions for this item
		Read   []string
		Write  []string
		Delete []string
	}
}
