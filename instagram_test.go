package instagram_test

import (
	"testing"

	. "github.com/hieven/go-instagram"
	"github.com/stretchr/testify/assert"
)

func TestInstagram(t *testing.T) {
	assert := assert.New(t)

	username := "even"
	password := "password"

	ig, err := Create(username, password)

	assert.Nil(err)
	assert.EqualValues(username, ig.Username)
	assert.EqualValues(password, ig.Password)
}
