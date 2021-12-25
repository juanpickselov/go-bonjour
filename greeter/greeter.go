package greeter

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("name not provided")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string {
		"Yo, %v. Welcome!",
		"Cool to see you, %v!",
		"Hola, %v. Well done!",
	}
	return formats[rand.Intn(len(formats))]
}
