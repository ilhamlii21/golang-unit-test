package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Ilham")
	assert.Equal(t, "Hello Ilham", result, "Hasil Harusnya 'Hello Ilham'")
	fmt.Println("Test HelloWorldAssert dengan Assert Berhasil")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ilham")
	require.Equal(t, "Hello Ilham", result, "Hasil Harusnya 'Hello Ilham'")
	fmt.Println("Test HelloWorldAssert dengan Require Berhasil")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Ilham")

	if result != "Hello Ilham" {
		//if error
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
