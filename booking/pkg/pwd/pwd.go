package pwd

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPwd(plainText string) (string, error) {
	hashByte, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Generate hash password error", err)
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashByte), err
}

func CheckedPwd(hashPwd, plainText string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(plainText))
}
