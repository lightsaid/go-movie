package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"lightsaid.com/go-movie/booking/utils"
)

func createRandomCinema(t *testing.T) Cinema {
	c := createRandomCity(t)
	arg := CreateCinemaParams{
		CityID:           c.ID,
		Name:             utils.RandomString(6),
		Long:             utils.RandomLonLat(),
		Lat:              utils.RandomLonLat(),
		TotalCinemaHalls: int32(utils.RandomInt(500, 2000)),
	}
	cn, err := testQueries.CreateCinema(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cn)

	require.Equal(t, arg.CityID, cn.CityID)
	require.Equal(t, arg.Name, cn.Name)
	require.Equal(t, arg.Long, cn.Long)
	require.Equal(t, arg.Lat, cn.Lat)
	require.Equal(t, arg.TotalCinemaHalls, cn.TotalCinemaHalls)

	return cn
}

func TestCreateCinema(t *testing.T) {
	_ = createRandomCinema(t)
}

func TestUpdateCinema(t *testing.T) {
	city := createRandomCity(t)
	c1 := createRandomCinema(t)
	arg := UpdateCinemaParams{
		ID:               c1.ID,
		CityID:           city.ID,
		Name:             utils.RandomString(6),
		Long:             utils.RandomLonLat(),
		Lat:              utils.RandomLonLat(),
		TotalCinemaHalls: int32(utils.RandomInt(500, 2000)),
	}
	c2, err := testQueries.UpdateCinema(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, c2)

	require.Equal(t, c1.ID, c2.ID)
	require.Equal(t, arg.CityID, c2.CityID)
	require.Equal(t, arg.Name, c2.Name)
	require.Equal(t, arg.TotalCinemaHalls, c2.TotalCinemaHalls)
	require.Equal(t, arg.Long, c2.Long)
	require.Equal(t, arg.Lat, c2.Lat)

}

func TestListCinema(t *testing.T) {
	rows := 10
	for i := 0; i < rows; i++ {
		_ = createRandomCinema(t)
	}
	arg := GetCinemaListParams{
		Limit:  5,
		Offset: 5,
	}
	list, err := testQueries.GetCinemaList(context.Background(), arg)
	require.NoError(t, err)
	require.True(t, len(list) > 0)
	for _, v := range list {
		require.NotEmpty(t, v)
		require.NotNil(t, v.ID)
		require.True(t, len(v.Name) > 0)
	}
}
