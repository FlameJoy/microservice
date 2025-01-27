package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
	"unicode"

	emailverifier "github.com/AfterShip/email-verifier"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"name" validate:"required,gte=4,lte=24"`
	Pswd       string    `json:"pswd" validate:"required,gte=8,lte=24" gorm:"-"`
	PswdRepeat string    `json:"pswdRepeat" validate:"required,gte=8,lte=24" gorm:"-"`
	PswdHash   string    `json:"-"`
	Email      string    `json:"email" validate:"required,email"`
	VerHash    string    `json:"-"`
	TimeoutAt  time.Time `json:"-"`
	VerifiedAt time.Time `json:"-"`
	CreatedAt  time.Time `json:"-"`
	UpdatedAt  time.Time `json:"-"`
	DeletedAt  time.Time `json:"-"`
}

var (
	minEntropyBits, _ = strconv.Atoi(os.Getenv("minEntropyBits"))
	verifier          = emailverifier.NewVerifier()
	serveLoc          = os.Getenv("server_location")
)

func (u *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(u)
}

func (u *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(u)
}

func (u *User) ValidateUsername() error {
	for _, char := range u.Username {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) {
			return fmt.Errorf("username must contain only letters and numbers")
		}
	}
	if len(u.Username) > 4 && len(u.Username) < 25 {
		return nil
	}
	return fmt.Errorf("username must be greater than 4 and less than 25 characters")
}

func (u *User) ValidateEmail() error {
	result, err := verifier.Verify(u.Email)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("email verification failed: %v", err)
	}
	if !result.Syntax.Valid {
		return errors.New("email syntax is invalid")
	}
	if result.Disposable {
		return errors.New("sorry we don't accept disposable email addresses")
	}
	if result.Suggestion != "" {
		return errors.New("email address is not reachable, looking for: " + result.Suggestion + " instead of " + result.Email + "?")
	}
	if result.Reachable == "no" {
		return errors.New("email address is unreachable")
	}
	if !result.HasMxRecords {
		return errors.New("domain entered not properly setup to recieve emails, MX record not found")
	}
	return nil
}

func (u *User) ValidatePswd() error {
	// Check length
	if len(u.Pswd) >= 8 && len(u.PswdRepeat) <= 24 {
		return nil
	}

	// Compare pswds
	for i := 0; i < len(u.Pswd); i++ {
		if u.Pswd[i] != u.PswdRepeat[i] {
			return errors.New("passwords aren't comparable")
		}
	}
	// Validation
	err := passwordvalidator.Validate(u.Pswd, float64(minEntropyBits))
	if err != nil {
		return err
	}

	return errors.New("password must be greater than 7 and less than 25 characters")
}
