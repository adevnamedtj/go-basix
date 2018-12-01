package main

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type user struct {
	FirstName         string
	LastName          string
	Email             string
	UserID            string
	PasswordHash      []byte
	PasswordHashAlt   []byte
	AddressLine1      string
	City              string
	State             string
	Zipcode           int
	LastSeen          time.Time
	JoinedOn          time.Time
	PlayTimeInMinutes float64
	Age               int
	GameHistory       map[string]bool
}

func hashPassword(s string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return hash
}

func validatePassword(passEntered, passUser []byte) bool {
	err := bcrypt.CompareHashAndPassword(passUser, passEntered)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
