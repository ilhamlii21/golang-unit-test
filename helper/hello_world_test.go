package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ilham")
	}
}

func BenchmarkHelloWorldLii(b *testing) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Lii")
	}
}

func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Ilham",
			request:  "Ilham",
			expected: "Hello Ilham",
		},
		{
			name:     "Lii",
			request:  "Lii",
			expected: "Hello Lii",
		},
		{
			name:     "Dwi",
			request:  "Dwi",
			expected: "Hello Dwi",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})

	}
}

func TestSubTest(t *testing.T) {
	t.Run("Ilham", func(t *testing.T) {
		result := HelloWorld("Ilham")
		require.Equal(t, "Hello Ilham", result, "Hasil Harusnya 'Hello Ilham'")

	})
	t.Run("Lii", func(t *testing.T) {
		result := HelloWorld("Lii")
		require.Equal(t, "Hello Lii", result, "Hasil Harusnya 'Hello Lii'")

	})
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("==UNIT TEST DIMULAI==")
	m.Run()
	//after
	fmt.Println("==UNIT TEST SELESAI==")

	//func TestMain(m *testing.M) berfungsi sebagai entry point untuk eksekusi semua unit test dalam package
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Tidak bisa berjalan di Windows")
	}
	//t.Skip berfungsi untuk membatalkan test

	result := HelloWorld("Ilham")
	require.Equal(t, "Hello Ilham", result, "Hasil Harusnya 'Hello Ilham'")
}

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
