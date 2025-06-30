package uuider

import "github.com/google/uuid"

type UUIDer struct{}

func New() *UUIDer {
	return &UUIDer{}
}

func (u *UUIDer) Create() string {
	return uuid.New().String()
}
