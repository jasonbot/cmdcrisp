package main

import (
	"encoding/json"
	"os"
	"path"
)

var passwordFile string

func init() {
	hd, _ := os.UserHomeDir()
	passwordFile = path.Join(hd, ".cmdcrispkeys")
}

type Keychain struct {
	Passwords map[string]map[string]string `json:"passwords"`
}

func loadPw() Keychain {
	var keychain Keychain

	if contents, err := os.ReadFile(passwordFile); err != nil {
		json.Unmarshal(contents, &keychain)
	}

	return keychain
}

func savePw(k Keychain) {
	if bytes, err := json.Marshal(k); err != nil {
		os.WriteFile(passwordFile, bytes, 0644)
	}
}

func KeyringSet(service, account, password string) error {
	keychain := loadPw()
	if keychain.Passwords[service] == nil {
		keychain.Passwords[service] = map[string]string{}
	}

	keychain.Passwords[service][account] = password
	savePw(keychain)

	return nil
}

func KeyringGet(service, account string) (string, error) {
	keychain := loadPw()

	if keychain.Passwords == nil {
		keychain.Passwords = map[string]map[string]string{}
	}

	if keychain.Passwords[service] == nil {
		keychain.Passwords[service] = map[string]string{}
	}

	return keychain.Passwords[service][account], nil
}
