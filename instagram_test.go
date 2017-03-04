package instagram_test

import (
	"testing"

	. "github.com/hieven/go-instagram"
	"github.com/hieven/go-instagram/config"
	"github.com/stretchr/testify/assert"
)

func TestInstagramNew(t *testing.T) {
	assert := assert.New(t)

	config := &config.Config{
		Username: "even",
		Password: "password",
	}
	ig, err := New(config)

	assert.Nil(err)
	assert.EqualValues(config.Username, ig.Config.Username)
	assert.EqualValues(config.Password, ig.Config.Password)
}
