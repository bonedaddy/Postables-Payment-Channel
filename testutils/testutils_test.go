package testutils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewBlockchain(t *testing.T) {
	auth, backend := NewBlockchain(t)
	require.NoError(t, backend.Close())
	fmt.Println(auth.From.String())
}
