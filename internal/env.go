package internal

import (
	"errors"
	"os"
)

type Env struct {
	Session string
}

func CheckEnv() (*Env, error) {
	session, exists := os.LookupEnv("SESSION")
	if !exists {
		return nil, errors.New("SESSION environment variable not set")
	}
	env := &Env{
		Session: session,
	}
	return env, nil
}
