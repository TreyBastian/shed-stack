package app

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_New_function(t *testing.T) {
	os.Setenv(APP_PORT, "8019032900")

	_, err := New()
	assert.Error(t, err, "should return error for invalid port")

	os.Setenv(APP_PORT, "8080")

	_, err = New()
	assert.NoError(t, err, "should not return error for valid port")
}
