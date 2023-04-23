package main

import (
	"fmt"
	"math/rand"
	"time"
)

func passwordGenerator(length int) string {
	rand.Seed(time.Now().UnixNano())

	charset := "dscabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-=_+"

	password := make([]byte, length)
	for k, _ := range password {
		password[k] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}

func main() {
	passcode := passwordGenerator(9)

	fmt.Println("Password:", passcode)

}

/* basically I realized another simple sequence of doing it:

func main() {
	password := make([]byte, 10)

	rand.Seed(time.Now().UnixNano())
	rand.Read(password)

	fmt.Printf("password: %x", password)
}
*/
