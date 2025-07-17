package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ilham")

	if result != "Hello Ilham" {
		panic("Result salah / bukan Hello Ilham")
	}
}

func TestHelloWorldLii(t *testing.T) {
	result := HelloWorld("Lii")

	if result != "Hello Lii" {
		panic("Result salah / bukan Hello Ilham")
	}
}
