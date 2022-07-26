package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"lightsaid.com/go-movie/booking/pkg/pwd"
	"lightsaid.com/go-movie/booking/utils"
)

func createRandomUser(t *testing.T) User {
	hashPwd, err := pwd.GenerateHashPwd(utils.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Name:      utils.RandomString(6),
		AvatarUrl: utils.NullStr(fmt.Sprintf("http://%s.png", utils.RandomString(8))).String(),
		Password:  utils.NullStr(hashPwd).String(),
		Phone:     utils.NullStr(utils.RandomPhone()).String(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.True(t, user.ID > 0)
	require.Equal(t, arg.Name, user.Name)
	require.Equal(t, arg.AvatarUrl, user.AvatarUrl)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Phone, user.Phone)
	return user
}

func TestCreateUser(t *testing.T) {
	_ = createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user, user2)
}

func TestUpdateUser(t *testing.T) {
	user := createRandomUser(t)
	hashPwd, err := pwd.GenerateHashPwd(utils.RandomString(6))
	require.NoError(t, err)
	arg := UpdateUserParams{
		ID:        user.ID,
		Name:      utils.RandomString(6),
		AvatarUrl: utils.NullStr(fmt.Sprintf("http://%s.png", utils.RandomString(8))).String(),
		Password:  utils.NullStr(hashPwd).String(),
		Phone:     utils.NullStr(utils.RandomPhone()).String(),
	}
	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user.ID, user2.ID)
	require.Equal(t, arg.Name, user2.Name)
	require.Equal(t, arg.AvatarUrl, user2.AvatarUrl)
	require.Equal(t, arg.Password, user2.Password)
	require.Equal(t, arg.Phone, user2.Phone)
}
