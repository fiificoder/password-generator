package main

import (
	"fmt"
	"math/rand"
	"time"
)

func passwordGenerator() string {
	rand.Seed(time.Now().UnixNano())

	charset := "dscabcdefghijklmnofvjasevfvbufayvpqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-=_+"

	password := make([]byte, 10)
	for k := range password {
		password[k] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}

func main() {
	passcode := passwordGenerator(9)

	fmt.Println("Password:", passcode)

}
