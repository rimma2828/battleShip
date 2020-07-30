package errors

import (
	"testing"

	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestGetErrorMessagePrefix(t *testing.T) {
	assert.Equal(t, "certificate is not found", getErrorMessagePrefix(2002))

	assert.Equal(t, "certificate error", getErrorMessagePrefix(2333))

	assert.Equal(t, "undefined error", getErrorMessagePrefix(11111111))
}

func TestCodes(t *testing.T) {
	data, _ := json.MarshalIndent(errorMessagePrefixes, "", "\t")
	fmt.Println(string(data))
}
