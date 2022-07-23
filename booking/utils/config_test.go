package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadConfig(t *testing.T) {
	var config ApiConfig
	err := LoadConfig("../", "ApiConfig", &config)

	require.NoError(t, err)
	require.NotEmpty(t, config)

	require.Equal(t, config.MaxPageSize, 100)
	require.Equal(t, config.PageSize, 10)
	require.Equal(t, config.RunMode, "debug")
	require.Equal(t, config.DBDriver, "postgres")
	require.Contains(t, config.DBSource, "postgres")
	require.Contains(t, config.ApiPort, "4000")

}
