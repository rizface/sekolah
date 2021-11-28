package helper

import "golang.org/x/crypto/bcrypt"

func GeneratePassword(plain string) string {
	result,err := bcrypt.GenerateFromPassword([]byte(plain),bcrypt.DefaultCost)
	PanicIfError(err)
	return string(result)
}

func ComparePassword(plain string, hashed string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed),[]byte(plain))
	return err
}