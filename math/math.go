package math

import (
	"errors"
	"math"
)

// This function takes two parameters, a and b, and returns the result of
// adding them together.
// Example: Tambah(1, 2) returns 3.
func Tambah(a int, b int) int {
	return a + b

}

// Kurang subtracts the integer b from integer a and returns the result.
// If a is less than b, an error is returned indicating that a must be greater than b.
// If the subtraction would cause an overflow, an error is returned.
func Kurang(a int, b int) (int, error) {
    if a < b {
        return 0, errors.New("integer a harus lebih besar dari b")
    } else if a > math.MaxInt - b {
        return 0, errors.New("pengurangan akan menyebabkan overflow")
    } else {
        return a - b, nil
    }
}

// Kali returns the result of a multiplied by b.
func Kali(a int, b int) int {
	return a * b
}

// Bagi returns the result of a divided by b.
func Bagi(a int, b int) int {
	return a / b
}

// Sisa returns the remainder of a divided by b.
func Sisa(a int, b int) int {
	return a % b
}

// Pangkat returns the result of a raised to the power of b.
func Pangkat(a int, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}
	return result
}

// Akar returns the root of a and b.
func Akar(a int) float64 {
	return math.Sqrt(float64(a))
}

// Mod returns the modulus of a and b.
func Mod(a int, b int) int {
	return a % b
}

// Menentukan nilai minimum antara dua bilangan
func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// Menghitung nilai faktorial dari suatu bilangan
// Faktorial 0 dan 1 adalah 1
// Faktorial n adalah n * faktorial(n-1)
func faktorial(a int) int {
	// 
    if a == 0 || a == 1 {
        return 1
    }
    return a * faktorial(a-1)
}

// Permutasi returns the number of permutations of n items taken r at a time.
func Permutasi(n int, r int) int {
	return faktorial(n) / faktorial(n-r)
}