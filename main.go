package microw

import (
	"os"

	"github.com/globalsign/mgo"
	"github.com/joho/godotenv"
	"github.com/mjarkk/framework-microwave/pkg/checker"
	"github.com/mjarkk/framework-microwave/pkg/gui"
)

// Init is the start of everthing
func Init() error {

	// Get env variables
	err := godotenv.Load()
	if err != nil {
		gui.CritErr("No env file found")
	}
	err = checker.Env()
	if err != nil {
		gui.CritErr(".env is incomplete missing key: " + err.Error())
	}

	// Conect to mongodb
	mongoURL := os.Getenv("MONGODB_URL")
	mongoDatabase := os.Getenv("MONGODB_DATABASE")

	session, err := mgo.Dial(mongoURL)
	if err != nil {
		gui.CritErr("Can't connect to mongodb, error: " + err.Error())
	}

	database := session.DB(mongoDatabase)

	database.C("test")

	return nil
}
