package data

import (
	"encoding/json"
	"io"
	"time"
)

type User struct {
	Name       string    `json:"name" validate:"required"`
	Pswd       string    `json:"pswd" validate:"required"`
	Email      string    `json:"email" validate:"required,email"`
	Age        uint8     `json:"age" validate:"gte=0,lte=130"`
	Role       string    `json:"role" validate:"required"`
	Gender     string    `json:"gender" validate:"oneof=male female prefer_not_to"`
	Created_At time.Time `json:"-"`
	Updated_At time.Time `json:"-"`
	Deleted_At time.Time `json:"-"`
}

func (u *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}
