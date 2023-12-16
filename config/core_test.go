package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	// Test the ReadConfig function
	app, err := ReadConfig("../configs")
	assert.NoError(t, err)
	assert.NotNil(t, app)

	// Validate the content of the loaded configuration
	assert.Equal(t, ":80", app.Configs.Address)
}

func TestInitAndGetConfig(t *testing.T) {
	Init("../configs")
	app := GetConfig()

	// Validate the content of the loaded configuration
	assert.Equal(t, ":80", app.Configs.Address)
}

func TestUnmarshalCustomStruct(t *testing.T) {

	type appConfig struct {
		Config struct {
			Configs `mapstructure:"configs"`
		}

		Configs struct {
			Custom string `mapstructure:"custom"`
		}
	}

	appConf := appConfig{}
	Init("../configs")
	conf := GetConfig()

	err := Unmarshal(&appConf)
	assert.NoError(t, err)

	// Validate the content of the unmarshaled struct
	assert.Equal(t, "customConfig", appConf.Configs.Custom)
	assert.Equal(t, ":80", conf.Configs.Address)

}
