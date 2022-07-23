package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomCity(t *testing.T) City {
	arg := CreateCityParams{
		Name: "广州",
		Lat:  "0.0",
		Long: "0.0",
	}

	c, err := testQueries.CreateCity(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, c)
	require.Equal(t, c.Name, arg.Name)
	require.Equal(t, c.Lat, arg.Lat)
	require.Equal(t, c.Long, arg.Long)

	return c
}

func TestCity(t *testing.T) {
	_ = createRandomCity(t)
}

func TestGetCityList(t *testing.T) {
	// cities, err := testQueries.GetAllCity(context.Background())
	// require.NoError(t, err)
	// require.NotEmpty(t, cities)

	// for k, v := range cities {

	// }
}
