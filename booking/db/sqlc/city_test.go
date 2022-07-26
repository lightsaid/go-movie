package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"lightsaid.com/go-movie/booking/utils"
)

func createRandomCity(t *testing.T, args ...CreateCityParams) City {
	var arg CreateCityParams
	if len(args) > 0 {
		arg = args[0]
	} else {
		arg = CreateCityParams{
			Name: utils.RandomString(6),
			Long: utils.RandomLonLat(),
			Lat:  utils.RandomLonLat(),
		}
	}

	c, err := testQueries.CreateCity(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, c)
	require.Equal(t, c.Name, arg.Name)
	require.Equal(t, c.Lat, arg.Lat)
	require.Equal(t, c.Long, arg.Long)

	return c
}

// 添加一些正常数据
// func TestAddDevCity(t *testing.T) {
// 	args := []CreateCityParams{
// 		CreateCityParams{
// 			Name: "广州",
// 			Long: "113.264499",
// 			Lat:  "23.130061",
// 		},
// 		CreateCityParams{
// 			Name: "深圳",
// 			Long: "114.057939",
// 			Lat:  "22.543527",
// 		},
// 		CreateCityParams{
// 			Name: "杭州",
// 			Long: "120.210792",
// 			Lat:  "30.246026",
// 		},
// 		CreateCityParams{
// 			Name: "北京",
// 			Long: "116.322056",
// 			Lat:  "39.89491",
// 		},
// 		CreateCityParams{
// 			Name: "上海",
// 			Long: "121.473667",
// 			Lat:  "31.230525",
// 		},
// 	}
// 	for _, arg := range args {
// 		_ = createRandomCity(t, arg)
// 	}
// }

func TestCity(t *testing.T) {
	_ = createRandomCity(t)
}

func TestGetCityList(t *testing.T) {
	var rows = 10
	for i := 0; i < rows; i++ {
		_ = createRandomCity(t)
	}
	arg := GetCitiesParams{
		Limit:  5,
		Offset: 5,
	}
	cs, err := testQueries.GetCities(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, cs)

	for _, v := range cs {
		require.NotEmpty(t, v)
		require.NotEmpty(t, v.ID)
		require.NotEmpty(t, v.Name)
	}
}
