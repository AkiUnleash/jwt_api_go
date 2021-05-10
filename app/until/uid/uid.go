package uid

import (
	"github.com/google/uuid"
)

func Generate() string {
	u, err := uuid.NewRandom()
	if err != nil {
		panic("No uid has been generated")
	}
	return u.String()
}
