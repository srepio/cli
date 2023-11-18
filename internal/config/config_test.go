package config

import (
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
