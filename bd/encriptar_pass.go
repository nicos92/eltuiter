package bd

import "golang.org/x/crypto/bcrypt"

func EncriptarPass(pass string) (string, error) {
	costo := 6
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	if err != nil {
		return err.Error(), err
	}

	return string(bytes), nil
}
