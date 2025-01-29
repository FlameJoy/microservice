package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"microsvc/common/utils"
	"os"
	"strconv"
	"time"
	"unicode"

	emailverifier "github.com/AfterShip/email-verifier"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

type User struct {
	ID         int       `json:"id"`
	Username   string    `json:"username" validate:"required,gte=4,lte=24"`
	Pswd       string    `json:"password" validate:"required,gte=8,lte=24"`
	PswdRepeat string    `json:"password_repeat" validate:"required,gte=8,lte=24"`
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
	minEntropyBits, _ = strconv.Atoi(os.Getenv("MIN_ENTROPY_BITS"))
	verifier          = emailverifier.NewVerifier()
	serveLoc          = os.Getenv("server_location")
)

func init() {
	verifier = verifier.EnableDomainSuggest()
	if serveLoc == "remote" {
		verifier.EnableSMTPCheck()
	}
	dispEmailsDomains := utils.DispEmailDomains()
	verifier = verifier.AddDisposableDomains(dispEmailsDomains)
}

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
			return errors.New("username must contain only letters and numbers")
		}
	}
	if len(u.Username) > 4 && len(u.Username) < 25 {
		return nil
	}
	return errors.New("username must be greater than 4 and less than 25 characters")
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
	if len(u.Pswd) < 8 && len(u.Pswd) > 24 {
		return errors.New("password must be greater than 7 and less than 25 characters")
	}

	if len(u.Pswd) != len(u.PswdRepeat) {
		return errors.New("different passwords")
	}

	for i := 0; i < len(u.Pswd); i++ {
		if u.Pswd[i] != u.PswdRepeat[i] {
			return errors.New("passwords aren't comparable")
		}
	}

	if err := passwordvalidator.Validate(u.Pswd, float64(minEntropyBits)); err != nil {
		return err
	}

	if len(u.Pswd) >= 8 && len(u.PswdRepeat) <= 24 {
		return nil
	}

	return errors.New("password must be greater than 7 and less than 25 characters")
}
