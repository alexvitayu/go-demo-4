package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!")

// blueprint for struct
type Account struct {
	Url       string    `json:"url"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// method struct
func (acc *Account) OutputInfo() {
	color.Cyan(acc.Url)
	color.Blue(acc.Login)
	color.Green(acc.Password)
}

// method struct
func (acc *Account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

// constructor struct
func NewMyAccount(urlString, login, password string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}
	newAcc := &Account{
		Url:       urlString,
		Login:     login,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}
