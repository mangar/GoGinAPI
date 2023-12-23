package util

import "testing"

func TestRandomString(t *testing.T) {

	require.NotEmpty(t, RandomString(10))

}