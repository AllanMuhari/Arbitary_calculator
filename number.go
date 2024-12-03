package main

import (
	"fmt"
)

type BigNumber struct {
	Digits []int
	Base   int
}

// Initialize the BigNumber from a string
func NewBigNumber(input string, base int) *BigNumber {
	digits := make([]int, len(input))
	for i, digit := range input {
		digits[len(input)-1-i] = int(digit - '0') 
	}
	return &BigNumber{Digits: digits, Base: base}
}

// Converts BigNumber to a string
func (b *BigNumber) String() string {
	result := ""
	for i := len(b.Digits) - 1; i >= 0; i-- {
		result += fmt.Sprintf("%d", b.Digits[i])
	}
	return result
}


func Add(a, b *BigNumber) *BigNumber {
	maxLength := max(len(a.Digits), len(b.Digits))
	result := make([]int, maxLength+1)
	carry := 0

	for i := 0; i < maxLength || carry != 0; i++ {
		sum := carry
		if i < len(a.Digits) {
			sum += a.Digits[i]
		}
		if i < len(b.Digits) {
			sum += b.Digits[i]
		}
		result[i] = sum % a.Base
		carry = sum / a.Base
	}
	return &BigNumber{Digits: trimLeadingZeros(result), Base: a.Base}
}
func Subtract(a, b *BigNumber) *BigNumber {
	maxLength := max(len(a.Digits), len(b.Digits))
	result := make([]int, maxLength)
	borrow := 0

	for i := 0; i < maxLength; i++ {
		diff := borrow
		if i < len(a.Digits) {
			diff += a.Digits[i]
		}
		if i < len(b.Digits) {
			diff -= b.Digits[i]
		}
		if diff < 0 {
			diff += a.Base
			borrow = -1
		} else {
			borrow = 0
		}
		result[i] = diff
	}
	return &BigNumber{Digits: trimLeadingZeros(result), Base: a.Base}
}


func Multiply(a, b *BigNumber) *BigNumber {
	result := make([]int, len(a.Digits)+len(b.Digits))
	for i, digitA := range a.Digits {
		for j, digitB := range b.Digits {
			result[i+j] += digitA * digitB
		}
	}
	for i := 0; i < len(result)-1; i++ {
		result[i+1] += result[i] / a.Base
		result[i] %= a.Base
	}
	return &BigNumber{Digits: trimLeadingZeros(result), Base: a.Base}
}

func Divide(a, b *BigNumber) (*BigNumber, *BigNumber) {
	if len(b.Digits) == 1 && b.Digits[0] == 0 {
		panic("Division by zero")
	}

	quotient := make([]int, len(a.Digits))
	remainder := &BigNumber{Digits: make([]int, len(a.Digits)), Base: a.Base}
	copy(remainder.Digits, a.Digits)

	for i := len(a.Digits) - 1; i >= 0; i-- {
		remainder.Digits = append([]int{0}, remainder.Digits...)
		remainder.Digits[0] = a.Digits[i]

		count := 0
		for Compare(remainder, b) >= 0 {
			remainder = Subtract(remainder, b)
			count++
		}
		quotient[i] = count
	}

	return &BigNumber{Digits: trimLeadingZeros(quotient), Base: a.Base}, remainder
}


func Compare(a, b *BigNumber) int {
	if len(a.Digits) != len(b.Digits) {
		if len(a.Digits) > len(b.Digits) {
			return 1
		}
		return -1
	}
	for i := len(a.Digits) - 1; i >= 0; i-- {
		if a.Digits[i] != b.Digits[i] {
			if a.Digits[i] > b.Digits[i] {
				return 1
			}
			return -1
		}
	}
	return 0
}

// Find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Remove leading zero digits
func trimLeadingZeros(digits []int) []int {
	i := len(digits) - 1
	for i > 0 && digits[i] == 0 {
		i--
	}
	return digits[:i+1]
}
