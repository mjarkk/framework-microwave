package microw

import (
	"github.com/mjarkk/framework-microwave/pkg/checker"
	"github.com/mjarkk/framework-microwave/pkg/database"
	"github.com/mjarkk/framework-microwave/pkg/globalvars"
	"github.com/mjarkk/framework-microwave/pkg/gui"
	"github.com/mjarkk/framework-microwave/pkg/types"
)

// TypeInit are all the types that can be used by apps
type TypeInit *types.Init

// Init is the start of everthing
func Init(settings *types.Init) error {

	globalvars.SetSettings(settings)

	// Get env variables
	err := checker.Env()
	if err != nil {
		gui.CritErr(".env is incomplete missing key: " + err.Error())
	}

	database.Init()

	return nil
}
