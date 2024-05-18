package domain

import (
	"time"
)

type User struct {
	Id                  int
	Nip                 string
	Name                string
	Password            string
	IdentityCardScanImg string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           time.Time
}
