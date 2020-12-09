package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	conf := NewConfig()

	assert.NotNil(t, conf)
	assert.NotEmpty(t, conf)
}
func TestNewConfigFail(t *testing.T) {
	conf := &Config{}
	assert.Empty(t, conf)
}
func TestNewDatabase(t *testing.T) {
	conf := NewConfig()

	db, err := NewDatabase(conf)
	assert.NotNil(t, db)
	assert.NoError(t, err)
}
func TestNewDatabaseFail(t *testing.T) {
	conf := &Config{}
	db, err := NewDatabase(conf)
	assert.Error(t, err)
	assert.Nil(t, db)

}
