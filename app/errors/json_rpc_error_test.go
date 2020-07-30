package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapToJsonRpcError(t *testing.T) {
	err := testError{newBaseError(int(CertificateActivationCodeNotValidError), errors.New("lorem ipsum dolor sit amet"))}

	assert.EqualError(t, WrapToJsonRpcError(err), "certificate's activation code is not found: lorem ipsum dolor sit amet")

	err.baseError = newBaseError(3333, errors.New("lorem ipsum dolor sit amet"))

	assert.EqualError(t, WrapToJsonRpcError(err), "certificate transaction error: lorem ipsum dolor sit amet")

	err.baseError = newBaseError(1111111, errors.New("lorem ipsum dolor sit amet"))

	assert.EqualError(t, WrapToJsonRpcError(err), "undefined error: lorem ipsum dolor sit amet")

	err.baseError = newBaseError(int(CertificateActivationCodeNotValidError), nil)

	assert.EqualError(t, WrapToJsonRpcError(err), "certificate's activation code is not found")
}
