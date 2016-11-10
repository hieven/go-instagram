package instagram

import (
	"os"
	"testing"

	"errors"

	"github.com/stretchr/testify/assert"
)

func TestInstagramCreate(t *testing.T) {
	account := os.Getenv("INSTAGRAM_ACCOUNT")
	password := os.Getenv("INSTAGRAM_PASSWORD")

	ig, err := Create(account, password)

	assert.Equal(t, account, ig.Username)
	assert.Nil(t, err)
}

func TestInstagramCreateFail(t *testing.T) {
	account := "even"
	password := "password"

	// incorrect account, password
	ig, err := Create(account, password)
	expectedErr := errors.New("The password you entered is incorrect. Please try again.")
	assert.Nil(t, ig)
	assert.Equal(t, expectedErr, err)

	// missing account
	ig, err = Create("", password)
	expectedErr = errors.New("Invalid Parameters")
	assert.Nil(t, ig)
	assert.Equal(t, expectedErr, err)

	// missing password
	ig, err = Create(account, "")
	expectedErr = errors.New("Invalid Parameters")
	assert.Nil(t, ig)
	assert.Equal(t, expectedErr, err)
}
