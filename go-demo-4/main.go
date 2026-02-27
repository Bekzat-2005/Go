package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type account struct {
	login    string
	password string
	url      string
}

func (acc *account) outputData() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunnes[rand.IntN(len(letterRunnes))]
	}
	acc.password = string(res)
}

func newAccount(login, password, urlString string) (*account, error) {

	if login == "" {
		return nil, errors.New("INVALID_LOGIN")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	Newacc := &account{
		password: password,
		login:    login,
		url:      urlString,
	}

	if password == "" {
		Newacc.generatePassword(25)
	}
	return Newacc, nil
}

var letterRunnes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-!*")

func main() {

	str := []rune("Hello!!")
	for _, ch := range str {
		fmt.Println(ch, string(ch))
	}
	// fmt.Println(rand.IntN(10))
	// fmt.Println(generatePassword(12))
	login := promptData("Login: ")
	password := promptData("Password: ")
	url := promptData("URL: ")

	// myAccount := account{
	// 	login:    login,
	// 	password: password,
	// 	url:      url,
	// }

	myAccount, err := newAccount(login, password, url)
	if err != nil {
		fmt.Print("INNNVALID_URL")
		return
	}
	// myAccount.generatePassword(25)
	myAccount.outputData()
	fmt.Println(myAccount)

	// outputData(&myAccount)
}

func promptData(prompt string) string {
	fmt.Println(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
