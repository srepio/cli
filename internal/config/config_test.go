package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItReturnsAnErrorForInvalidConnections(t *testing.T) {
	c := &Config{
		CurrentConnection: "bongo",
	}

	err := c.validate()

	assert.ErrorIs(t, err, ErrInvalidConnection)
}

func TestItDoesntErrorForValidConnections(t *testing.T) {
	c := &Config{
		CurrentConnection: "bongo",
		Connections: []Api{
			{Name: "bongo"},
		},
	}

	err := c.validate()

	assert.Nil(t, err)
}

func TestItLoadsAFileWhenEnvVarIsSet(t *testing.T) {
	path, err := filepath.Abs("testdata/.srep-valid.yaml")
	assert.Nil(t, err)
	t.Setenv("SREP_CONFIG", path)

	c, err := GetConfig()

	assert.Nil(t, err)

	assert.Equal(t, "default", c.CurrentConnection)
}

func TestItFailsWhenNoConfigFileExists(t *testing.T) {
	t.Setenv("SREP_CONFIG", "/tmp/bongo.yaml")
	_, err := GetConfig()
	assert.ErrorIs(t, err, ErrNoConfig)
}

func TestItCreatesANewDefaultConfig(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("SREP_CONFIG", fmt.Sprintf("%s/.srep.yaml", dir))

	_, err := os.Stat(fmt.Sprintf("%s/.srep.yaml", dir))
	assert.Error(t, err)

	err = Initialise()
	assert.Nil(t, err)

	_, err = os.Stat(fmt.Sprintf("%s/.srep.yaml", dir))
	assert.Nil(t, err)
}
