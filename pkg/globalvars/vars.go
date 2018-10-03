package globalvars

import (
	"github.com/mjarkk/framework-microwave/pkg/types"
	"github.com/oleiade/reflections"
)

// Settings can be accessed from everyware
var Settings = types.Init{}

// SetSettings can set new global settings
func SetSettings(newSettings *types.Init) {
	Settings = *newSettings
}

// GetSettings returns the settings object
func GetSettings() types.Init {
	return Settings
}

// SetOneSetting sets one sets one item
func SetOneSetting(index string, val interface{}) error {
	return reflections.SetField(&Settings, index, val)
}
