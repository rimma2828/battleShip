package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testError struct {
	baseError
}

func TestBaseError_Error(t *testing.T) {
	err := testError{newBaseError(int(CertificateActivationCodeNotValidError), errors.New("lorem ipsum dolor sit amet"))}

	assert.EqualError(t, err, "(2003) certificate's activation code is not found: lorem ipsum dolor sit amet")

	err.baseError = newBaseError(3333, errors.New("lorem ipsum dolor sit amet"))

	assert.EqualError(t, err, "(3333) certificate transaction error: lorem ipsum dolor sit amet")

	err.baseError = newBaseError(1111111, errors.New("lorem ipsum dolor sit amet"))

	assert.EqualError(t, err, "(1111111) undefined error: lorem ipsum dolor sit amet")

	err.baseError = newBaseError(int(CertificateActivationCodeNotValidError), nil)

	assert.EqualError(t, err, "(2003) certificate's activation code is not found")
}

func TestBaseError_ExtractJsonRpsErrorParams(t *testing.T) {
	err := testError{newBaseError(int(CertificateActivationCodeNotValidError), errors.New("lorem ipsum dolor sit amet"))}

	code, msg := err.ExtractJsonRpsErrorParams()

	assert.Equal(t, 2003, code)

	assert.Equal(t, "certificate's activation code is not found: lorem ipsum dolor sit amet", msg)
}
