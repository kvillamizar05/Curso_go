package bd

import "golang.org/x/crypto/bcrypt"

/*EncriptarPassword sirve para hacer la modificacion del password enviado por el front*/
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
