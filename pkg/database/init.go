package database

import (
	"os"

	"github.com/globalsign/mgo"
	"github.com/mjarkk/framework-microwave/pkg/filehandeler"
	"github.com/mjarkk/framework-microwave/pkg/globalvars"
	"github.com/mjarkk/framework-microwave/pkg/gui"
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

	YmlList, err := filehandeler.ReadAllYamlIn(globalvars.Settings.MigrationPath)
	if err != nil {
		gui.CritErr("can't read yaml database files in " + globalvars.Settings.MigrationPath + " error: " + err.Error())
	}
	_, ymlErr := ParseDBYmlS(YmlList)
	if ymlErr.Err != nil {
		gui.CritErr("Error in yml file: " + ymlErr.File + ", error: " + ymlErr.Err.Error())
	}
	// spew.Dump(parsedData)
}
