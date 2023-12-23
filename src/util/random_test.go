package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {

	require.NotEmpty(t, RandomString(10))

}