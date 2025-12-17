package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	cfg := New()

	require.Equal(t, "8080", cfg.Port)
}
