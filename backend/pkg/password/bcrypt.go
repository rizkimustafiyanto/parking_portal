package password

import "golang.org/x/crypto/bcrypt"

func Hash(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(raw),
		bcrypt.DefaultCost,
	)

	return string(hash), err
}

func Compare(hash, raw string) error {
	return bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(raw),
	)
}