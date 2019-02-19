package main

import "strings"

type (
	Key string
)

func (k Key) Lowercase() Key {
	return Key(strings.ToLower(k.String()))
}

func (k Key) String() string {
	return string(k)
}
