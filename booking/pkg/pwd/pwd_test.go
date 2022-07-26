package pwd

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"lightsaid.com/go-movie/booking/utils"
)

func TestPassword(t *testing.T) {
	pwd := utils.RandomString(6)

	hashPwd, err := GenerateHashPwd(pwd)
	require.NoError(t, err)
	require.NotEmpty(t, hashPwd)

	err = CheckedPwd(hashPwd, pwd)
	require.NoError(t, err)

	wrongPwd := utils.RandomString(7)
	err = CheckedPwd(hashPwd, wrongPwd)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
