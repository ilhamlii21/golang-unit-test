package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ilham")

	if result != "Hello Ilham" {
		//if erro
		// panic("Result salah / bukan Hello Ilham")

		// t.Fail()
		t.Error("Result harusnya 'hello ilham' ")
		//t.Error untuk memberikan keterangan error
	}

	fmt.Println("Test HelloWorld Berhasil")
}

func TestHelloWorldLii(t *testing.T) {
	result := HelloWorld("Lii")

	if result != "Hello Lii" {
		//if error
		// panic("Result salah / bukan Hello Ilham")
		// t.FailNow()
		t.Fatal("Hasil harusnya 'Hello Lii'")
		//t.Fatal memberikan keterangan error dan langsung keluar program
	}

	fmt.Println("Test HelloWorldLii Berhasil")
}
